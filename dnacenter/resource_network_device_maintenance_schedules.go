package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceMaintenanceSchedules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Devices.

- API to create maintenance schedule for network devices. The state of network device can be queried using API *GET
/dna/intent/api/v1/networkDevices*. The *managementState* attribute of the network device will be updated to
*UNDER_MAINTENANCE* when the maintenance window starts.
`,

		CreateContext: resourceNetworkDeviceMaintenanceSchedulesCreate,
		ReadContext:   resourceNetworkDeviceMaintenanceSchedulesRead,
		UpdateContext: resourceNetworkDeviceMaintenanceSchedulesUpdate,
		DeleteContext: resourceNetworkDeviceMaintenanceSchedulesDelete,
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

func resourceNetworkDeviceMaintenanceSchedulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevices(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vvNetworkDeviceIDs := interfaceToString(resourceItem["network_device_ids"])

	queryParamImport := dnacentersdkgo.RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams{}
	queryParamImport.NetworkDeviceIDs = vvNetworkDeviceIDs
	item2, err := searchDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices(m, queryParamImport, vvNetworkDeviceIDs)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["network_device_ids"] = vvNetworkDeviceIDs
		d.SetId(joinResourceID(resourceMap))
		return resourceNetworkDeviceMaintenanceSchedulesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Devices.CreateMaintenanceScheduleForNetworkDevices(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateMaintenanceScheduleForNetworkDevices", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateMaintenanceScheduleForNetworkDevices", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateMaintenanceScheduleForNetworkDevices", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
				"Failure when executing CreateMaintenanceScheduleForNetworkDevices", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams{}
	queryParamValidate.NetworkDeviceIDs = vvNetworkDeviceIDs
	item3, err := searchDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices(m, queryParamValidate, vvNetworkDeviceIDs)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateMaintenanceScheduleForNetworkDevices", err,
			"Failure at CreateMaintenanceScheduleForNetworkDevices, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["network_device_ids"] = vvNetworkDeviceIDs
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkDeviceMaintenanceSchedulesRead(ctx, d, m)
}

func resourceNetworkDeviceMaintenanceSchedulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["network_device_ids"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveScheduledMaintenanceWindowsForNetworkDevices")
		queryParams1 := dnacentersdkgo.RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams{}

		response1, restyResp1, err := client.Devices.RetrieveScheduledMaintenanceWindowsForNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveScheduledMaintenanceWindowsForNetworkDevices search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItem(item *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["description"] = item.Description
	respItem["maintenanceSchedule"] = flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItemMaintenanceSchedule(item.MaintenanceSchedule)
	respItem["networkDeviceIds"] = item.NetworkDeviceIDs
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItemMaintenanceSchedule(item *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["startId"] = item.StartID
	respItem["endId"] = item.EndID
	respItem["startTime"] = item.StartTime
	respItem["endTime"] = item.EndTime
	respItem["recurrence"] = flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItemRecurrence(item.Recurrence)
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesByIDItemRecurrence(item *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceScheduleRecurrence) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["recurrenceEndTime"] = item.RecurrenceEndTime
	return []map[string]interface{}{
		respItem,
	}
}

func resourceNetworkDeviceMaintenanceSchedulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDeviceMaintenanceSchedulesRead(ctx, d, m)
}

func resourceNetworkDeviceMaintenanceSchedulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete NetworkDeviceMaintenanceSchedules on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevices {
	request := dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".maintenance_schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".maintenance_schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".maintenance_schedule")))) {
		request.MaintenanceSchedule = expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule(ctx, key+".maintenance_schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_ids")))) {
		request.NetworkDeviceIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule {
	request := dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".recurrence")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".recurrence")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".recurrence")))) {
		request.Recurrence = expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence(ctx, key+".recurrence.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceMaintenanceSchedulesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence {
	request := dnacentersdkgo.RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence{}
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

func searchDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices(m interface{}, queryParams dnacentersdkgo.RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams, vID string) (*dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse
	var ite *dnacentersdkgo.ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices
	if vID != "" {
		queryParams.Offset = strconv.Itoa(1)
		nResponse, _, err := client.Devices.RetrieveScheduledMaintenanceWindowsForNetworkDevices(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			offset, err := strconv.Atoi(queryParams.Offset)
			if err != nil {
				return foundItem, err
			}
			offset += maxPageSize
			queryParams.Limit = strconv.Itoa(maxPageSize)
			queryParams.Offset += strconv.Itoa(offset)
			nResponse, _, err = client.Devices.RetrieveScheduledMaintenanceWindowsForNetworkDevices(&queryParams)
		}
		return nil, err
	} else if queryParams.NetworkDeviceIDs != "" {
		ite, _, err = client.Devices.RetrieveScheduledMaintenanceWindowsForNetworkDevices(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			for _, deviceId := range item.NetworkDeviceIDs {
				if deviceId == queryParams.NetworkDeviceIDs {
					foundItem = &item
					return foundItem, err
				}
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
