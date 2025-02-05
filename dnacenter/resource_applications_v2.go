package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationsV2() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Edit the attributes of an existing application

- Create new custom application/s

- Delete existing custom application by id
`,

		CreateContext: resourceApplicationsV2Create,
		ReadContext:   resourceApplicationsV2Read,
		UpdateContext: resourceApplicationsV2Update,
		DeleteContext: resourceApplicationsV2Delete,
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

						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Id of Application
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `Type of identify source. NBAR: build in Application, APIC_EM: custom Application
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"indicative_network_identity": &schema.Schema{
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
									"lower_port": &schema.Schema{
										Description: `Lower port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ports": &schema.Schema{
										Description: `Ports
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"upper_port": &schema.Schema{
										Description: `Upper port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"instance_id": &schema.Schema{
							Description: `Instance id
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
						"name": &schema.Schema{
							Description: `Application name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Description: `Namespace, valid value scalablegroup:application
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_protocol": &schema.Schema{
										Description: `App protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_sub_type": &schema.Schema{
										Description: `Application sub type, LEARNED: discovered application, NONE: nbar and custom application
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_type": &schema.Schema{
										Description: `Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"category_id": &schema.Schema{
										Description: `Category id
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
									"dscp": &schema.Schema{
										Description: `Dscp
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"engine_id": &schema.Schema{
										Description: `Engine id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"help_string": &schema.Schema{
										Description: `Help string
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
									"long_description": &schema.Schema{
										Description: `Long description
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Application name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"popularity": &schema.Schema{
										Description: `Popularity
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"rank": &schema.Schema{
										Description: `Rank, any value between 1 to 65535
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"selector_id": &schema.Schema{
										Description: `Selector id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_name": &schema.Schema{
										Description: `Server name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"traffic_class": &schema.Schema{
										Description: `Traffic class
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": &schema.Schema{
										Description: `Url
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"network_identity": &schema.Schema{
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
									"ipv4_subnet": &schema.Schema{
										Description: `Ipv4 subnet
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ipv6_subnet": &schema.Schema{
										Description: `Ipv6 subnet
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"lower_port": &schema.Schema{
										Description: `Lower port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ports": &schema.Schema{
										Description: `Ports
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"upper_port": &schema.Schema{
										Description: `Upper port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"parent_scalable_group": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id_ref": &schema.Schema{
										Description: `Id reference to parent application set
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"qualifier": &schema.Schema{
							Description: `Qualifier, valid value application
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable_group_external_handle": &schema.Schema{
							Description: `Scalable group external handle, should be equal to Application name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable_group_type": &schema.Schema{
							Description: `Scalable group type, valid value APPLICATION
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type, valid value scalablegroup
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplications`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Application id
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"indicative_network_identity": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ipv4_subnet": &schema.Schema{
													Description: `Ipv4 subnet
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ipv6_subnet": &schema.Schema{
													Description: `Ipv6 subnet
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"lower_port": &schema.Schema{
													Description: `The minimum port when used as a port range. For single port number, ports attribute should be used.
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"ports": &schema.Schema{
													Description: `Ports
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"upper_port": &schema.Schema{
													Description: `The maximum port when used as a port range. For single port number, ports attribute should be used.
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Application name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": &schema.Schema{
										Description: `Namespace
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_applications": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"app_protocol": &schema.Schema{
													Description: `App protocol, in case of _servername should not be set, in case of _url should be set to TCP
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"application_sub_type": &schema.Schema{
													Description: `Application sub type, LEARNED: discovered application, NONE: nbar and custom application
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"application_type": &schema.Schema{
													Description: `Application type
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"category_id": &schema.Schema{
													Description: `Category id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dscp": &schema.Schema{
													Description: `Dscp, valid only in case of _server-ip custom application type
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"engine_id": &schema.Schema{
													Description: `Engine id, should be set to 6

ERROR: Different types for param engineId schema.TypeInt schema.TypeString`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"help_string": &schema.Schema{
													Description: `Help string
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ignore_conflict": &schema.Schema{
													Description: `Ignore conflict, true or false
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"long_description": &schema.Schema{
													Description: `Long description
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Application name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"popularity": &schema.Schema{
													Description: `Popularity
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"rank": &schema.Schema{
													Description: `Rank, should be set to 1
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"selector_id": &schema.Schema{
													Description: `Selector id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"server_name": &schema.Schema{
													Description: `Server name, should be set only in case of _servername
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic class
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"type": &schema.Schema{
													Description: `Custom application type
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"url": &schema.Schema{
													Description: `Url, should be set only in case of _url
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"network_identity": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ipv4_subnet": &schema.Schema{
													Description: `Ipv4 subnet
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ipv6_subnet": &schema.Schema{
													Description: `Ipv6 subnet
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"lower_port": &schema.Schema{
													Description: `Lower port
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"ports": &schema.Schema{
													Description: `Ports
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"upper_port": &schema.Schema{
													Description: `Upper port
`,
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"parent_scalable_group": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id reference to parent application set
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"qualifier": &schema.Schema{
										Description: `Qualifier, valid value application
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scalable_group_external_handle": &schema.Schema{
										Description: `Scalable group external handle, should be equal to Application name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scalable_group_type": &schema.Schema{
										Description: `Scalable group type, valid value APPLICATION
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `Type, valid value scalablegroup
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceApplicationsV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestApplicationsV2CreateApplicationsV2(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := dnacentersdkgo.GetApplicationsV2QueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchApplicationPolicyGetApplicationsV2(m, queryParamImport, "")
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceApplicationsV2Read(ctx, d, m)
	}
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationsV2(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplications", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplications", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplications", err))
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
				"Failure when executing CreateApplications", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetApplicationsV2QueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchApplicationPolicyGetApplicationsV2(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApplications", err,
			"Failure at CreateApplications, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item2.ID
	resourceMap["name"] = vvName

	d.SetId(joinResourceID(resourceMap))
	return resourceApplicationsV2Read(ctx, d, m)
}

func resourceApplicationsV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplications")
		queryParams1 := dnacentersdkgo.GetApplicationsV2QueryParams{}
		queryParams1.Name = vvName
		item1, err := searchApplicationPolicyGetApplicationsV2(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2Response{
			*item1,
		}
		vItem1 := flattenApplicationPolicyGetApplicationsV2Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplications search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceApplicationsV2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestApplicationsV2EditApplicationsV2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.ApplicationPolicy.EditApplicationsV2(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing EditApplications", err, restyResp1.String(),
					"Failure at EditApplications, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing EditApplications", err,
				"Failure at EditApplications, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing EditApplications", err))
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
					"Failure when executing EditApplications", err1))
				return diags
			}
		}

	}

	return resourceApplicationsV2Read(ctx, d, m)
}

func resourceApplicationsV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplication(&dnacentersdkgo.DeleteApplicationQueryParams{
		ID: vvID,
	})
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplication", err, restyResp1.String(),
				"Failure at DeleteApplication, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplication", err,
			"Failure at DeleteApplication, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteApplication", err))
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
				"Failure when executing DeleteApplication", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestApplicationsV2CreateApplicationsV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationsV2 {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationsV2{}
	if v := expandRequestApplicationsV2CreateApplicationsV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2 {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2CreateApplicationsV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2Item(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2 {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_applications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_applications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_applications")))) {
		request.NetworkApplications = expandRequestApplicationsV2CreateApplicationsV2ItemNetworkApplicationsArray(ctx, key+".network_applications", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_scalable_group")))) {
		request.ParentScalableGroup = expandRequestApplicationsV2CreateApplicationsV2ItemParentScalableGroup(ctx, key+".parent_scalable_group.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identity")))) {
		request.NetworkIDentity = expandRequestApplicationsV2CreateApplicationsV2ItemNetworkIDentityArray(ctx, key+".network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".indicative_network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".indicative_network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".indicative_network_identity")))) {
		request.IndicativeNetworkIDentity = expandRequestApplicationsV2CreateApplicationsV2ItemIndicativeNetworkIDentityArray(ctx, key+".indicative_network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_type")))) {
		request.ScalableGroupType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemNetworkApplicationsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2CreateApplicationsV2ItemNetworkApplications(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemNetworkApplications(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".help_string")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".help_string")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".help_string")))) {
		request.HelpString = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".app_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".app_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".app_protocol")))) {
		request.AppProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".category_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".category_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".category_id")))) {
		request.CategoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ignore_conflict")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ignore_conflict")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ignore_conflict")))) {
		request.IgnoreConflict = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".engine_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".engine_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".engine_id")))) {
		request.EngineID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemParentScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2CreateApplicationsV2ItemNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_subnet")))) {
		request.IPv4Subnet = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemIndicativeNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2CreateApplicationsV2ItemIndicativeNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2CreateApplicationsV2ItemIndicativeNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_subnet")))) {
		request.IPv4Subnet = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyEditApplicationsV2 {
	request := dnacentersdkgo.RequestApplicationPolicyEditApplicationsV2{}
	if v := expandRequestApplicationsV2EditApplicationsV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2 {
	request := []dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2EditApplicationsV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2Item(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2 {
	request := dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_version")))) {
		request.InstanceVersion = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".indicative_network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".indicative_network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".indicative_network_identity")))) {
		request.IndicativeNetworkIDentity = expandRequestApplicationsV2EditApplicationsV2ItemIndicativeNetworkIDentityArray(ctx, key+".indicative_network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".namespace")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".namespace")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".namespace")))) {
		request.Namespace = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_applications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_applications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_applications")))) {
		request.NetworkApplications = expandRequestApplicationsV2EditApplicationsV2ItemNetworkApplicationsArray(ctx, key+".network_applications", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identity")))) {
		request.NetworkIDentity = expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityArray(ctx, key+".network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_scalable_group")))) {
		request.ParentScalableGroup = expandRequestApplicationsV2EditApplicationsV2ItemParentScalableGroup(ctx, key+".parent_scalable_group.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qualifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qualifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qualifier")))) {
		request.Qualifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_external_handle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_external_handle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_external_handle")))) {
		request.ScalableGroupExternalHandle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_type")))) {
		request.ScalableGroupType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemIndicativeNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity {
	request := []dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2EditApplicationsV2ItemIndicativeNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemIndicativeNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity {
	request := dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkApplicationsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkApplications {
	request := []dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkApplications{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2EditApplicationsV2ItemNetworkApplications(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkApplications(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkApplications {
	request := dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkApplications{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_sub_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_sub_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_sub_type")))) {
		request.ApplicationSubType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".category_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".category_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".category_id")))) {
		request.CategoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".engine_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".engine_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".engine_id")))) {
		request.EngineID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".help_string")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".help_string")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".help_string")))) {
		request.HelpString = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".long_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".long_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".long_description")))) {
		request.LongDescription = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".popularity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".popularity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".popularity")))) {
		request.Popularity = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selector_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selector_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selector_id")))) {
		request.SelectorID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".app_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".app_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".app_protocol")))) {
		request.AppProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ignore_conflict")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ignore_conflict")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ignore_conflict")))) {
		request.IgnoreConflict = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity {
	request := []dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity {
	request := dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_subnet")))) {
		request.IPv4Subnet = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityIPv6SubnetArray(ctx, key+".ipv6_subnet", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityIPv6SubnetArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet {
	request := []dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityIPv6Subnet(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemNetworkIDentityIPv6Subnet(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet {
	var request dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsV2EditApplicationsV2ItemParentScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup {
	request := dnacentersdkgo.RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplicationsV2(m interface{}, queryParams dnacentersdkgo.GetApplicationsV2QueryParams, vID string) (*dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2Response, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2Response
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.ApplicationPolicy.GetApplicationsV2(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.ApplicationPolicy.GetApplicationsV2(&queryParams)
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.ApplicationPolicy.GetApplicationsV2(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
