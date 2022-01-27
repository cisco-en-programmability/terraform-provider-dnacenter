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

func resourcePnpWorkflow() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Onboarding (PnP).

- Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database

- Deletes a workflow specified by id

- Updates an existing workflow
`,

		CreateContext: resourcePnpWorkflowCreate,
		ReadContext:   resourcePnpWorkflowRead,
		UpdateContext: resourcePnpWorkflowUpdate,
		DeleteContext: resourcePnpWorkflowDelete,
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

						"id": &schema.Schema{
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"add_to_inventory": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"added_on": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"curr_task_idx": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"exec_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"image_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lastupdate_on": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"curr_work_item_idx": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"task_seq_no": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"time_taken": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"end_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"output_str": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"state": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPnpWorkflowAddAWorkflow(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpWorkflowRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}
		queryParams1.Name = append(queryParams1.Name, vvName)
		getResponse2, err := searchDeviceOnboardingPnpGetWorkflows(m, queryParams1)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpWorkflowRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddAWorkflow(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddAWorkflow", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddAWorkflow", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourcePnpWorkflowRead(ctx, d, m)
}

func resourcePnpWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	if okName && vName != "" {
		log.Printf("[DEBUG] Selected method 1: GetWorkflows")
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}

		if okName {
			queryParams1.Name = append(queryParams1.Name, vName)
		}

		response1, err := searchDeviceOnboardingPnpGetWorkflows(m, queryParams1)

		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetWorkflows", err,
			// 	"Failure at GetWorkflows, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetWorkflowByID(response1.TypeID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflowByID", err,
				"Failure at GetWorkflowByID, unexpected response", ""))
			return diags
		}

		vItem1 := flattenDeviceOnboardingPnpGetWorkflowByIDItem(response2)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWorkflows search response",
				err))
			return diags
		}
		return diags

	}
	if okID && vID != "" {
		log.Printf("[DEBUG] Selected method 2: GetWorkflowByID")
		vvID := vID

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetWorkflowByID", err,
			// 	"Failure at GetWorkflowByID, unexpected response", ""))
			// return diags
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourcePnpWorkflowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vName := resourceMap["name"]

	vvID := ""
	if vName != "" {
		log.Printf("[DEBUG] Selected method 1: GetWorkflows")
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}

		queryParams1.Name = append(queryParams1.Name, vName)

		response1, err := searchDeviceOnboardingPnpGetWorkflows(m, queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflows", err,
				"Failure at GetWorkflows, unexpected response", ""))
			return diags
		}
		vvID = response1.TypeID
	} else if vID != "" {
		log.Printf("[DEBUG] Selected method 2: GetWorkflowByID")
		vvID = vID
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
	}

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestPnpWorkflowUpdateWorkflow(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil && request1.TypeID == "" {
			request1.TypeID = vvID
		}
		response1, restyResp1, err := client.DeviceOnboardingPnp.UpdateWorkflow(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateWorkflow", err, restyResp1.String(),
					"Failure at UpdateWorkflow, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateWorkflow", err,
				"Failure at UpdateWorkflow, unexpected response", ""))
			return diags
		}
	}

	return resourcePnpWorkflowRead(ctx, d, m)
}

func resourcePnpWorkflowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	var vvID string

	vID := resourceMap["id"]
	vName := resourceMap["name"]

	if vID != "" {
		vvID = vID
		getResp, _, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	} else if vName != "" {
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}
		queryParams1.Name = append(queryParams1.Name, vName)

		response1, err := searchDeviceOnboardingPnpGetWorkflows(m, queryParams1)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}
		vvID = response1.TypeID
	}

	response1, restyResp1, err := client.DeviceOnboardingPnp.DeleteWorkflowByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteWorkflowByID", err, restyResp1.String(),
				"Failure at DeleteWorkflowByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteWorkflowByID", err,
			"Failure at DeleteWorkflowByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPnpWorkflowAddAWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpWorkflowAddAWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks{}
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
		i := expandRequestPnpWorkflowAddAWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpWorkflowAddAWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpWorkflowAddAWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpWorkflowUpdateWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks{}
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
		i := expandRequestPnpWorkflowUpdateWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDeviceOnboardingPnpGetWorkflows(m interface{}, queryParams dnacentersdkgo.GetWorkflowsQueryParams) (*dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetWorkflows, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetWorkflows
	var ite *dnacentersdkgo.ResponseDeviceOnboardingPnpGetWorkflows
	ite, _, err = client.DeviceOnboardingPnp.GetWorkflows(&queryParams)
	if err != nil {
		return foundItem, err
	}

	if ite == nil {
		return foundItem, err
	}

	items := ite

	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name[0] {
			var getItem *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetWorkflows
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
