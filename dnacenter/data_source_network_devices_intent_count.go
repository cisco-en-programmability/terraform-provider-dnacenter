package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesIntentCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- API to fetch the count of network devices using basic filters. Use the */dna/intent/api/v1/networkDevices/query/count*
API if you need advanced filtering.
`,

		ReadContext: dataSourceNetworkDevicesIntentCountRead,
		Schema: map[string]*schema.Schema{
			"family": &schema.Schema{
				Description: `family query parameter. Product family of the network device. For example, Switches, Routers, etc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Network device Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_address": &schema.Schema{
				Description: `managementAddress query parameter. Management address of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_state": &schema.Schema{
				Description: `managementState query parameter. The status of the network device's manageability. Available values : MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachability_status": &schema.Schema{
				Description: `reachabilityStatus query parameter. Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Description: `role query parameter. Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Serial number of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_device": &schema.Schema{
				Description: `stackDevice query parameter. Flag indicating if the device is a stack device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesIntentCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vManagementAddress, okManagementAddress := d.GetOk("management_address")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vFamily, okFamily := d.GetOk("family")
	vStackDevice, okStackDevice := d.GetOk("stack_device")
	vRole, okRole := d.GetOk("role")
	vStatus, okStatus := d.GetOk("status")
	vReachabilityStatus, okReachabilityStatus := d.GetOk("reachability_status")
	vManagementState, okManagementState := d.GetOk("management_state")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountTheNumberOfNetworkDevices")
		queryParams1 := dnacentersdkgo.CountTheNumberOfNetworkDevicesQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okManagementAddress {
			queryParams1.ManagementAddress = vManagementAddress.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okStackDevice {
			queryParams1.StackDevice = vStackDevice.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okReachabilityStatus {
			queryParams1.ReachabilityStatus = vReachabilityStatus.(string)
		}
		if okManagementState {
			queryParams1.ManagementState = vManagementState.(string)
		}

		response1, restyResp1, err := client.Devices.CountTheNumberOfNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountTheNumberOfNetworkDevices", err,
				"Failure at CountTheNumberOfNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesCountTheNumberOfNetworkDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountTheNumberOfNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesCountTheNumberOfNetworkDevicesItem(item *dnacentersdkgo.ResponseDevicesCountTheNumberOfNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
