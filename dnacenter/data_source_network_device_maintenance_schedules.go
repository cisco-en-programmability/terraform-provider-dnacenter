package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceMaintenanceSchedules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- This data source retrieves a list of scheduled maintenance windows for network devices based on filter parameters.
Each maintenance window is composed of a start schedule and end schedule, both of which have unique
identifiers(*startId* and *endId*). These identifiers can be used to fetch the status of the start schedule and end
schedule using the *GET /dna/intent/api/v1/activities/{id}* API. Completed maintenance schedules are automatically
removed from the system after two weeks.
`,

		ReadContext: dataSourceNetworkDeviceMaintenanceSchedulesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Min: 1, Max: 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_ids": &schema.Schema{
				Description: `networkDeviceIds query parameter. List of network device ids.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `A brief narrative describing the maintenance schedule.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the schedule maintenance window
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"maintenance_schedule": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_id": &schema.Schema{
										Description: `Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /dna/intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_time": &schema.Schema{
										Description: `End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"recurrence": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interval": &schema.Schema{
													Description: `Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"recurrence_end_time": &schema.Schema{
													Description: `The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},

									"start_id": &schema.Schema{
										Description: `Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /dna/intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `The status of the maintenance schedule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"network_device_ids": &schema.Schema{
							Description: `List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceMaintenanceSchedulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceIDs, okNetworkDeviceIDs := d.GetOk("network_device_ids")
	vStatus, okStatus := d.GetOk("status")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveScheduledMaintenanceWindowsForNetworkDevices")
		queryParams1 := dnacentersdkgo.RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams{}

		if okNetworkDeviceIDs {
			queryParams1.NetworkDeviceIDs = vNetworkDeviceIDs.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Devices.RetrieveScheduledMaintenanceWindowsForNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveScheduledMaintenanceWindowsForNetworkDevices", err,
				"Failure at RetrieveScheduledMaintenanceWindowsForNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveScheduledMaintenanceWindowsForNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItems(items *[]dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["maintenance_schedule"] = flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItemsMaintenanceSchedule(item.MaintenanceSchedule)
		respItem["network_device_ids"] = item.NetworkDeviceIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItemsMaintenanceSchedule(item *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["start_id"] = item.StartID
	respItem["end_id"] = item.EndID
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["recurrence"] = flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItemsMaintenanceScheduleRecurrence(item.Recurrence)
	respItem["status"] = item.Status

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesItemsMaintenanceScheduleRecurrence(item *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceScheduleRecurrence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["recurrence_end_time"] = item.RecurrenceEndTime

	return []map[string]interface{}{
		respItem,
	}

}
