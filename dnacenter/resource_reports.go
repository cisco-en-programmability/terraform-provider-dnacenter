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

func resourceReports() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Reports.

- Create/Schedule a report configuration. Use "Get view details for a given view group & view" API to get the metadata
required to configure a report.

- Delete a scheduled report configuration. Deletes the report executions also.
`,

		CreateContext: resourceReportsCreate,
		ReadContext:   resourceReportsRead,
		UpdateContext: resourceReportsUpdate,
		DeleteContext: resourceReportsDelete,
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

						"data_category": &schema.Schema{
							Description: `data category of the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"deliveries": &schema.Schema{
							Description: `Array of available delivery channels
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"execution_count": &schema.Schema{
							Description: `Total number of report executions
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"executions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Report execution pipeline end time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"errors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"execution_id": &schema.Schema{
										Description: `Report execution Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"process_status": &schema.Schema{
										Description: `Report execution status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"request_status": &schema.Schema{
										Description: `Report execution acceptance status from scheduler
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Description: `Report execution pipeline start time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"warnings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `report name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"report_id": &schema.Schema{
							Description: `report Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"report_was_executed": &schema.Schema{
							Description: `true if atleast one execution has started
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
							Description: `array of tags for report
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"view": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `view description
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"field_groups": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"format": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default": &schema.Schema{
													Description: `true, if the format type is considered default
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `view name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"view_id": &schema.Schema{
										Description: `view Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"view_info": &schema.Schema{
										Description: `view filters info
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"view_group_id": &schema.Schema{
							Description: `viewGroupId of the viewgroup for the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"view_group_version": &schema.Schema{
							Description: `version of viewgroup for the report
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deliveries": &schema.Schema{
							Description: `Array of available delivery channels
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Description: `report name
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"report_id": &schema.Schema{
							Description: `reportId path parameter. reportId of report
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Description: `array of tags for report
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"view": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"field_groups": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"format": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `view name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"view_id": &schema.Schema{
										Description: `view Id
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"view_group_id": &schema.Schema{
							Description: `viewGroupId of the viewgroup for the report
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"view_group_version": &schema.Schema{
							Description: `version of viewgroup for the report
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

func resourceReportsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestReportsCreateOrScheduleAReport(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vReportID, okReportID := resourceItem["report_id"]
	vvReportID := interfaceToString(vReportID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okReportID && vvReportID != "" {
		getResponse2, _, err := client.Reports.GetAScheduledReport(vvReportID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["report_id"] = vvReportID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceReportsRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, err := searchReportsGetListOfScheduledReports(m, nil, vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["report_id"] = vvReportID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceReportsRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Reports.CreateOrScheduleAReport(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateOrScheduleAReport", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateOrScheduleAReport", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["report_id"] = resp1.ReportID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceReportsRead(ctx, d, m)
}

func resourceReportsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vReportID := resourceMap["report_id"]
	vName := resourceMap["name"]

	if vReportID != "" {
		log.Printf("[DEBUG] Selected method 2: GetAScheduledReport")
		vvReportID := vReportID

		response1, restyResp1, err := client.Reports.GetAScheduledReport(vvReportID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetAScheduledReport", err,
			// 	"Failure at GetAScheduledReport, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenReportsGetAScheduledReportItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAScheduledReport response",
				err))
			return diags
		}
		return diags

	}

	if vName != "" {
		response1, err := searchReportsGetListOfScheduledReports(m, nil, vName)
		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetAScheduledReport", err,
			// 	"Failure at GetAScheduledReport, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		response2, restyResp2, err := client.Reports.GetAScheduledReport(response1.ReportID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAScheduledReport", err,
				"Failure at GetAScheduledReport, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenReportsGetAScheduledReportItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAScheduledReport response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceReportsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceReportsRead(ctx, d, m)
}

func resourceReportsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vReportID := resourceMap["report_id"]
	vName := resourceMap["name"]
	// REVIEW: Add getAllItems and search function to get missing params

	if vReportID != "" {
		getResp, _, err := client.Reports.GetAScheduledReport(vReportID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if vName != "" {
		getResp, err := searchReportsGetListOfScheduledReports(m, nil, vName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		vReportID = getResp.ReportID
	}

	response1, restyResp1, err := client.Reports.DeleteAScheduledReport(vReportID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAScheduledReport", err, restyResp1.String(),
				"Failure at DeleteAScheduledReport, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAScheduledReport", err,
			"Failure at DeleteAScheduledReport, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestReportsCreateOrScheduleAReport(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReport {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReport{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deliveries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deliveries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deliveries")))) {
		request.Deliveries = expandRequestReportsCreateOrScheduleAReportDeliveriesArray(ctx, key+".deliveries", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schedule")))) {
		request.Schedule = expandRequestReportsCreateOrScheduleAReportSchedule(ctx, key+".schedule", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view")))) {
		request.View = expandRequestReportsCreateOrScheduleAReportView(ctx, key+".view.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_group_id")))) {
		request.ViewGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_group_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_group_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_group_version")))) {
		request.ViewGroupVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportDeliveriesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries{}
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
		i := expandRequestReportsCreateOrScheduleAReportDeliveries(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportDeliveries(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportSchedule {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportSchedule
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportView(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportView {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportView{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_groups")))) {
		request.FieldGroups = expandRequestReportsCreateOrScheduleAReportViewFieldGroupsArray(ctx, key+".field_groups", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestReportsCreateOrScheduleAReportViewFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".format")))) {
		request.Format = expandRequestReportsCreateOrScheduleAReportViewFormat(ctx, key+".format.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_id")))) {
		request.ViewID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFieldGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_group_display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_group_display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_group_display_name")))) {
		request.FieldGroupDisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_group_name")))) {
		request.FieldGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fields")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fields")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fields")))) {
		request.Fields = expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFieldsArray(ctx, key+".fields", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFieldsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFields(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFields(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestReportsCreateOrScheduleAReportViewFiltersValue(ctx, key+".value", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFiltersValue {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFiltersValue
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFormat(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFormat {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFormat{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".format_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".format_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".format_type")))) {
		request.FormatType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchReportsGetListOfScheduledReports(m interface{}, queryParams *dnacentersdkgo.GetListOfScheduledReportsQueryParams, vName string) (*dnacentersdkgo.ResponseItemReportsGetListOfScheduledReports, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemReportsGetListOfScheduledReports
	var ite *dnacentersdkgo.ResponseReportsGetListOfScheduledReports
	ite, _, err = client.Reports.GetListOfScheduledReports(nil)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == vName {
			var getItem *dnacentersdkgo.ResponseItemReportsGetListOfScheduledReports
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
