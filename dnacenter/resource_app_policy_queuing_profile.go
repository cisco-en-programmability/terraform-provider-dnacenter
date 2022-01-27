package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppPolicyQueuingProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing custom application queuing profile

- Create new custom application queuing profile

- Delete existing custom application policy queuing profile by id
`,

		CreateContext: resourceAppPolicyQueuingProfileCreate,
		ReadContext:   resourceAppPolicyQueuingProfileRead,
		UpdateContext: resourceAppPolicyQueuingProfileUpdate,
		DeleteContext: resourceAppPolicyQueuingProfileDelete,
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

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"clause": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"interface_speed_bandwidth_clauses": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"display_name": &schema.Schema{
																Description: `Display name
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"id": &schema.Schema{
																Description: `Id
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"instance_created_on": &schema.Schema{
																Description: `Instance created on
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"instance_updated_on": &schema.Schema{
																Description: `Instance updated on
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"instance_version": &schema.Schema{
																Description: `Instance version
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"traffic_class": &schema.Schema{
																Description: `Traffic Class
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

									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tc_dscp_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"contract_classifier": &schema.Schema{
							Description: `Contract classifier
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"create_time": &schema.Schema{
							Description: `Create time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"custom_provisions": &schema.Schema{
							Description: `Custom provisions
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Free test description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"gen_id": &schema.Schema{
							Description: `Gen id
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of Queueing profile
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_created_on": &schema.Schema{
							Description: `Instance created on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Description: `Instance updated on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"internal": &schema.Schema{
							Description: `Internal
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_deleted": &schema.Schema{
							Description: `Is deleted
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_seeded": &schema.Schema{
							Description: `Is seeded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_stale": &schema.Schema{
							Description: `Is stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ise_reserved": &schema.Schema{
							Description: `Is reserved
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Queueing profile name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"provisioning_state": &schema.Schema{
							Description: `Provisioning State
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pushed": &schema.Schema{
							Description: `Pushed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"qualifier": &schema.Schema{
							Description: `Qualifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"resource_version": &schema.Schema{
							Description: `Resource version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"target_id_list": &schema.Schema{
							Description: `Target id list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile`,
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"clause": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interface_speed_bandwidth_clauses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"traffic_class": &schema.Schema{
																Description: `Traffic Class
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
									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"tc_dscp_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Free test description
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `Id of Queueing profile
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Queueing profile name
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceAppPolicyQueuingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileArray(ctx, "parameters", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vName := resourceItem["name"]

	vvName := interfaceToString(vName)

	queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}
	queryParams1.Name = vvName
	item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
	if err != nil || item != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceAppPolicyQueuingProfileRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationPolicyQueuingProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfile", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfile", err))
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
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicyQueuingProfile")

		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}
		queryParams1.Name = vName
		item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
		if err != nil || item == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetApplicationPolicyQueuingProfile", err,
			// 	"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*item))

		vItem1 := flattenApplicationPolicyGetApplicationPolicyQueuingProfileItem(item)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyQueuingProfile search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAppPolicyQueuingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	//selectedMethod := 1
	//var vvID string

	queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}
	queryParams1.Name = vName
	item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetApplicationPolicyQueuingProfile", err,
			"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vName)
		request1 := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileArray(ctx, "parameters", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		// Add ID to update
		if request1 != nil && len(*request1) > 0 && item != nil {
			req := *request1
			req[0].ID = item.ID
			request1 = &req
		}
		response1, restyResp1, err := client.ApplicationPolicy.UpdateApplicationPolicyQueuingProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err, restyResp1.String(),
					"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err,
				"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err))
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
				diags = append(diags, diagError(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err))
				return diags
			}
		}
	}

	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	var vvID string
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}
	queryParams1.Name = vName
	item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetApplicationPolicyQueuingProfile", err,
			"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}

	vvID = item.ID
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationPolicyQueuingProfile(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationPolicyQueuingProfile", err, restyResp1.String(),
				"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationPolicyQueuingProfile", err,
			"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}
	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteApplicationPolicyQueuingProfile", err))
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
			diags = append(diags, diagError(
				"Failure when executing DeleteApplicationPolicyQueuingProfile", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileArray(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileArray(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchApplicationPolicyGetApplicationPolicyQueuingProfile(m interface{}, queryParams dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams) (*dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfile
	ite, _, err = client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
