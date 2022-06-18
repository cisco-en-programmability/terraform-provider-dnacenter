package dnacenter

import (
	"context"
	"time"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessProvisionDeviceCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Provision wireless devices
`,

		CreateContext: resourceWirelessProvisionDeviceCreateCreate,
		ReadContext:   resourceWirelessProvisionDeviceCreateRead,
		DeleteContext: resourceWirelessProvisionDeviceCreateDelete,
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
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_url": &schema.Schema{
							Description: `Execution Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"provisioning_tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failed": &schema.Schema{
										Description: `Failed`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"success": &schema.Schema{
										Description: `Success`,
										Type:        schema.TypeList,
										Computed:    true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestWirelessProvision`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_name": &schema.Schema{
										Description: `Controller Name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"dynamic_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface_gateway": &schema.Schema{
													Description: `Interface Gateway
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"interface_ipaddress": &schema.Schema{
													Description: `Interface IP Address
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"interface_name": &schema.Schema{
													Description: `Interface Name
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"interface_netmask_in_cid_r": &schema.Schema{
													Description: `Interface Netmask In CIDR
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
												},
												"lag_or_port_number": &schema.Schema{
													Description: `Lag Or Port Number
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
												},
												"vlan_id": &schema.Schema{
													Description: `VLAN ID
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
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
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"site": &schema.Schema{
										Description: `Full Site Hierarchy where device has to be assigned
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceWirelessProvisionDeviceCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessProvisionDeviceCreateProvision(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Wireless.Provision(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing Provision", err,
			"Failure at Provision, unexpected response", ""))
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
				"Failure when executing Provision", err,
				"Failure at Provision execution", bapiError))
			return diags
		}
	}

	vItem1 := flattenWirelessProvisionItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting Provision response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceWirelessProvisionDeviceCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessProvisionDeviceCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessProvisionDeviceCreateProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessProvision {
	request := dnacentersdkgo.RequestWirelessProvision{}
	if v := expandRequestWirelessProvisionDeviceCreateProvisionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionDeviceCreateProvisionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvision {
	request := []dnacentersdkgo.RequestItemWirelessProvision{}
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
		i := expandRequestWirelessProvisionDeviceCreateProvisionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceCreateProvisionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvision {
	request := dnacentersdkgo.RequestItemWirelessProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_aplocations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_aplocations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".managed_aplocations")))) {
		request.ManagedApLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_interfaces")))) {
		request.DynamicInterfaces = expandRequestWirelessProvisionDeviceCreateProvisionItemDynamicInterfacesArray(ctx, key+".dynamic_interfaces", d)
	}
	return &request
}

func expandRequestWirelessProvisionDeviceCreateProvisionItemDynamicInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces {
	request := []dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces{}
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
		i := expandRequestWirelessProvisionDeviceCreateProvisionItemDynamicInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceCreateProvisionItemDynamicInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces {
	request := dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces{}
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

func flattenWirelessProvisionItem(item *dnacentersdkgo.ResponseWirelessProvision) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_url"] = item.ExecutionURL
	respItem["provisioning_tasks"] = flattenWirelessProvisionItemProvisioningTasks(item.ProvisioningTasks)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessProvisionItemProvisioningTasks(item *dnacentersdkgo.ResponseWirelessProvisionProvisioningTasks) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = item.Success
	respItem["failed"] = item.Failed

	return []map[string]interface{}{
		respItem,
	}

}
