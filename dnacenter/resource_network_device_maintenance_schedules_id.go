package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceMaintenanceSchedulesID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Devices.

- API to update the maintenance schedule for the network devices. The *maintenanceSchedule* can be updated only if the
*status* value is *UPCOMING* or *IN_PROGRESS*. User can exit *IN_PROGRESS* maintenance window by setting the *endTime*
to -1. This will update the endTime to the current time and exit the maintenance window immediately. When exiting the
maintenance window, only the endTime will be updated while other parameters remain read-only.

- API to delete maintenance schedule by id. Deletion is allowed if the maintenance window is in the *UPCOMING*,
*COMPLETED*, or *FAILED* state. Deletion of maintenance schedule is not allowed if the maintenance window is currently
*IN_PROGRESS*. To delete the maintenance schedule while it is *IN_PROGRESS*, first exit the current maintenance window
using *PUT /dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}* API, and then proceed to delete the maintenance
schedule.
`,

		CreateContext: resourceNetworkDeviceMaintenanceSchedulesIDCreate,
		ReadContext:   resourceNetworkDeviceMaintenanceSchedulesIDRead,
		UpdateContext: resourceNetworkDeviceMaintenanceSchedulesIDUpdate,
		DeleteContext: resourceNetworkDeviceMaintenanceSchedulesIDDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `A brief narrative describing the maintenance schedule.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Unique identifier for the maintenance schedule
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"maintenance_schedule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"recurrence": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interval": &schema.Schema{
													Description: `Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"recurrence_end_time": &schema.Schema{
													Description: `The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"start_time": &schema.Schema{
										Description: `Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"network_device_ids": &schema.Schema{
							Description: `List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
`,
							Type:     schema.TypeList,
							Optional: true,
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

func resourceNetworkDeviceMaintenanceSchedulesIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDeviceMaintenanceSchedulesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheMaintenanceScheduleInformation")
		vvID := vID

		response1, restyResp1, err := client.Devices.RetrievesTheMaintenanceScheduleInformation(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
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

		return diags

	}
	return diags
}

func resourceNetworkDeviceMaintenanceSchedulesIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformation(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Devices.UpdatesTheMaintenanceScheduleInformation(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesTheMaintenanceScheduleInformation", err, restyResp1.String(),
					"Failure at UpdatesTheMaintenanceScheduleInformation, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesTheMaintenanceScheduleInformation", err,
				"Failure at UpdatesTheMaintenanceScheduleInformation, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesTheMaintenanceScheduleInformation", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdatesTheMaintenanceScheduleInformation", err1))
				return diags
			}
		}

	}

	return resourceNetworkDeviceMaintenanceSchedulesIDRead(ctx, d, m)
}

func resourceNetworkDeviceMaintenanceSchedulesIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Devices.DeleteMaintenanceSchedule(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteMaintenanceSchedule", err, restyResp1.String(),
				"Failure at DeleteMaintenanceSchedule, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMaintenanceSchedule", err,
			"Failure at DeleteMaintenanceSchedule, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteMaintenanceSchedule", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteMaintenanceSchedule", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformation {
	request := dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".maintenance_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".maintenance_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".maintenance_schedule")))) {
		request.MaintenanceSchedule = expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule(ctx, key+".maintenance_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_ids")))) {
		request.NetworkDeviceIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule {
	request := dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".recurrence")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".recurrence")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".recurrence")))) {
		request.Recurrence = expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence(ctx, key+".recurrence.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceMaintenanceSchedulesIDUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence {
	request := dnacentersdkgo.RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".recurrence_end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".recurrence_end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".recurrence_end_time")))) {
		request.RecurrenceEndTime = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
