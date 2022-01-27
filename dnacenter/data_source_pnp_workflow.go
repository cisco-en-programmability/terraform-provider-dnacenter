package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpWorkflow() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the list of workflows based on filter criteria. If a limit is not specified, it will default to return 50
workflows. Pagination and sorting are also supported by this endpoint

- Returns a workflow specified by id
`,

		ReadContext: dataSourcePnpWorkflowRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Limits number of results
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Workflow Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Index of first result
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sort": &schema.Schema{
				Description: `sort query parameter. Comma seperated lost of fields to sort on
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Workflow Type
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"add_to_inventory": &schema.Schema{
							Description: `Add To Inventory`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"added_on": &schema.Schema{
							Description: `Added On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"config_id": &schema.Schema{
							Description: `Config Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"curr_task_idx": &schema.Schema{
							Description: `Curr Task Idx`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"exec_time": &schema.Schema{
							Description: `Exec Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"image_id": &schema.Schema{
							Description: `Image Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"lastupdate_on": &schema.Schema{
							Description: `Lastupdate On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"state": &schema.Schema{
							Description: `State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"curr_work_item_idx": &schema.Schema{
										Description: `Curr Work Item Idx`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"end_time": &schema.Schema{
										Description: `End Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"state": &schema.Schema{
										Description: `State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"task_seq_no": &schema.Schema{
										Description: `Task Seq No`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"time_taken": &schema.Schema{
										Description: `Time Taken`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Description: `Command`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"output_str": &schema.Schema{
													Description: `Output Str`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"time_taken": &schema.Schema{
													Description: `Time Taken`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"use_state": &schema.Schema{
							Description: `Use State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"add_to_inventory": &schema.Schema{
							Description: `Add To Inventory`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"added_on": &schema.Schema{
							Description: `Added On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"config_id": &schema.Schema{
							Description: `Config Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"curr_task_idx": &schema.Schema{
							Description: `Curr Task Idx`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"exec_time": &schema.Schema{
							Description: `Exec Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"image_id": &schema.Schema{
							Description: `Image Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"lastupdate_on": &schema.Schema{
							Description: `Lastupdate On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"state": &schema.Schema{
							Description: `State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"curr_work_item_idx": &schema.Schema{
										Description: `Curr Work Item Idx`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"end_time": &schema.Schema{
										Description: `End Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"state": &schema.Schema{
										Description: `State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"task_seq_no": &schema.Schema{
										Description: `Task Seq No`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"time_taken": &schema.Schema{
										Description: `Time Taken`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Description: `Command`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"output_str": &schema.Schema{
													Description: `Output Str`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"time_taken": &schema.Schema{
													Description: `Time Taken`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"use_state": &schema.Schema{
							Description: `Use State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSort, okSort := d.GetOk("sort")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vType, okType := d.GetOk("type")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okType, okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetWorkflows")
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okSort {
			queryParams1.Sort = interfaceToSliceString(vSort)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okType {
			queryParams1.Type = interfaceToSliceString(vType)
		}
		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetWorkflows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflows", err,
				"Failure at GetWorkflows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetWorkflowsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWorkflows response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetWorkflowByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflowByID", err,
				"Failure at GetWorkflowByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceOnboardingPnpGetWorkflowByIDItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWorkflowByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetWorkflowsItems(items *dnacentersdkgo.ResponseDeviceOnboardingPnpGetWorkflows) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type_id"] = item.TypeID
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["description"] = item.Description
		respItem["lastupdate_on"] = item.LastupdateOn
		respItem["image_id"] = item.ImageID
		respItem["curr_task_idx"] = item.CurrTaskIDx
		respItem["added_on"] = item.AddedOn
		respItem["tasks"] = flattenDeviceOnboardingPnpGetWorkflowsItemsTasks(item.Tasks)
		respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
		respItem["instance_type"] = item.InstanceType
		respItem["end_time"] = item.EndTime
		respItem["exec_time"] = item.ExecTime
		respItem["start_time"] = item.StartTime
		respItem["use_state"] = item.UseState
		respItem["config_id"] = item.ConfigID
		respItem["name"] = item.Name
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetWorkflowsItemsTasks(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetWorkflowsTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetWorkflowsItemsTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetWorkflowsItemsTasksWorkItemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetWorkflowsTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetWorkflowByIDItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetWorkflowByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpGetWorkflowByIDItemTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpGetWorkflowByIDItemTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetWorkflowByIDTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetWorkflowByIDItemTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetWorkflowByIDItemTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}
