package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceMaintenanceSchedulesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieve the total count of all scheduled maintenance windows for network devices.
`,

		ReadContext: dataSourceNetworkDeviceMaintenanceSchedulesCountRead,
		Schema: map[string]*schema.Schema{
			"network_device_ids": &schema.Schema{
				Description: `networkDeviceIds query parameter. List of network device ids.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
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
							Description: `Count of scheduled maintenance windows
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

func dataSourceNetworkDeviceMaintenanceSchedulesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceIDs, okNetworkDeviceIDs := d.GetOk("network_device_ids")
	vStatus, okStatus := d.GetOk("status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheTotalNumberOfScheduledMaintenanceWindows")
		queryParams1 := dnacentersdkgo.RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams{}

		if okNetworkDeviceIDs {
			queryParams1.NetworkDeviceIDs = vNetworkDeviceIDs.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}

		response1, restyResp1, err := client.Devices.RetrieveTheTotalNumberOfScheduledMaintenanceWindows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheTotalNumberOfScheduledMaintenanceWindows", err,
				"Failure at RetrieveTheTotalNumberOfScheduledMaintenanceWindows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheTotalNumberOfScheduledMaintenanceWindows response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsItem(item *dnacentersdkgo.ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
