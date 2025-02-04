package dnacenter

import (
	"context"
	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessProvisionDeviceUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Wireless.

- Updates wireless provisioning
`,

		CreateContext: resourceWirelessProvisionDeviceUpdateCreate,
		ReadContext:   resourceWirelessProvisionDeviceUpdateRead,
		DeleteContext: resourceWirelessProvisionDeviceUpdateDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Description: `Message
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestWirelessProvisionUpdate`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_name": &schema.Schema{
										Description: `Controller Name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"dynamic_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface_gateway": &schema.Schema{
													Description: `Interface Gateway. Required for AireOS.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"interface_ipaddress": &schema.Schema{
													Description: `Interface IP Address. Required for AireOS.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"interface_name": &schema.Schema{
													Description: `Interface Name. Required for AireOS and EWLC.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"interface_netmask_in_cid_r": &schema.Schema{
													Description: `Interface Netmask In CIDR. Required for AireOS.
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"lag_or_port_number": &schema.Schema{
													Description: `Lag Or Port Number. Required for AireOS.
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"vlan_id": &schema.Schema{
													Description: `VLAN ID. Required for AireOS and EWLC.
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"managed_aplocations": &schema.Schema{
										Description: `List of managed AP locations (Site Hierarchies)
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceWirelessProvisionDeviceUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vPersistbapioutput := resourceItem["persistbapioutput"]

	request1 := expandRequestWirelessProvisionDeviceUpdateProvisionUpdate(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.ProvisionUpdateHeaderParams{}

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.ProvisionUpdate(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ProvisionUpdate", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ProvisionUpdate", err,
				"Failure at ProvisionUpdate execution", bapiError))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessProvisionUpdateItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ProvisionUpdate response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessProvisionDeviceUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessProvisionDeviceUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessProvisionUpdate {
	request := dnacentersdkgo.RequestWirelessProvisionUpdate{}
	if v := expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionUpdate {
	request := []dnacentersdkgo.RequestItemWirelessProvisionUpdate{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionUpdate {
	request := dnacentersdkgo.RequestItemWirelessProvisionUpdate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_aplocations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_aplocations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".managed_aplocations")))) {
		request.ManagedApLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_interfaces")))) {
		request.DynamicInterfaces = expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemDynamicInterfacesArray(ctx, key+".dynamic_interfaces", d)
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemDynamicInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces {
	request := []dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemDynamicInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateItemDynamicInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces {
	request := dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_ipaddress")))) {
		request.InterfaceIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_netmask_in_cid_r")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) {
		request.InterfaceNetmaskInCIDR = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_gateway")))) {
		request.InterfaceGateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lag_or_port_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lag_or_port_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lag_or_port_number")))) {
		request.LagOrPortNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	return &request
}

func flattenWirelessProvisionUpdateItem(item *dnacentersdkgo.ResponseWirelessProvisionUpdate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
