package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceWirelessProvisionDeviceCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Provision wireless devices
`,

		ReadContext: dataSourceWirelessProvisionDeviceCreateRead,
		Schema: map[string]*schema.Schema{
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
			"payload": &schema.Schema{
				Description: `Array of RequestWirelessProvision`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_name": &schema.Schema{
							Description: `Controller Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_interfaces": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_gateway": &schema.Schema{
										Description: `Interface Gateway
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_ipaddress": &schema.Schema{
										Description: `Interface IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface Name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_netmask_in_cid_r": &schema.Schema{
										Description: `Interface Netmask In CIDR
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"lag_or_port_number": &schema.Schema{
										Description: `Lag Or Port Number
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vlan_id": &schema.Schema{
										Description: `VLAN ID
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"managed_aplocations": &schema.Schema{
							Description: `List of managed AP locations (Site Hierarchies)
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"site": &schema.Schema{
							Description: `Full Site Hierarchy where device has to be assigned
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessProvisionDeviceCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Provision")
		request1 := expandRequestWirelessProvisionDeviceCreateProvision(ctx, "", d)

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
	return diags
}

func expandRequestWirelessProvisionDeviceCreateProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessProvision {
	request := dnacentersdkgo.RequestWirelessProvision{}
	if v := expandRequestWirelessProvisionDeviceCreateProvisionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
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
	for item_no, _ := range objs {
		i := expandRequestWirelessProvisionDeviceCreateProvisionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
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
	for item_no, _ := range objs {
		i := expandRequestWirelessProvisionDeviceCreateProvisionItemDynamicInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
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
