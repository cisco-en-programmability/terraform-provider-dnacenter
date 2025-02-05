package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceMaintenanceSchedulesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- API to retrieve the maintenance schedule information for the given id.
`,

		ReadContext: dataSourceNetworkDeviceMaintenanceSchedulesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Unique identifier for the maintenance schedule
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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
										Description: `Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.
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
										Description: `Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.
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

func dataSourceNetworkDeviceMaintenanceSchedulesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheMaintenanceScheduleInformation")
		vvID := vID.(string)

		response1, restyResp1, err := client.Devices.RetrievesTheMaintenanceScheduleInformation(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheMaintenanceScheduleInformation", err,
				"Failure at RetrievesTheMaintenanceScheduleInformation, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesRetrievesTheMaintenanceScheduleInformationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheMaintenanceScheduleInformation response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrievesTheMaintenanceScheduleInformationItem(item *dnacentersdkgo.ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["description"] = item.Description
	respItem["maintenance_schedule"] = flattenDevicesRetrievesTheMaintenanceScheduleInformationItemMaintenanceSchedule(item.MaintenanceSchedule)
	respItem["network_device_ids"] = item.NetworkDeviceIDs
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesRetrievesTheMaintenanceScheduleInformationItemMaintenanceSchedule(item *dnacentersdkgo.ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["start_id"] = item.StartID
	respItem["end_id"] = item.EndID
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["recurrence"] = flattenDevicesRetrievesTheMaintenanceScheduleInformationItemMaintenanceScheduleRecurrence(item.Recurrence)
	respItem["status"] = item.Status

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesRetrievesTheMaintenanceScheduleInformationItemMaintenanceScheduleRecurrence(item *dnacentersdkgo.ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceScheduleRecurrence) []map[string]interface{} {
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
