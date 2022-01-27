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

func resourceTag() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Tag.

- Updates a tag specified by id

- Creates tag with specified tag attributes

- Deletes a tag specified by id
`,

		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"values": &schema.Schema{
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
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"values": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTagCreateTag(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.Tag.GetTagByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTagRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		queryParams1.Name = vvName
		response1, err := searchTagGetTag(m, queryParams1)

		if err != nil || response1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTagRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Tag.CreateTag(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTag", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTag", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateTag", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TaskID %s", taskId)
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
			log.Printf("[DEBUG] Error => %s", response2.Response.FailureReason)
			diags = append(diags, diagError(
				"Failure when executing CreateTag", err))
			return diags
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceTagRead(ctx, d, m)
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTag")
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		if okName {
			queryParams1.Name = vName
		}

		response1, err := searchTagGetTag(m, queryParams1)

		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetTag", err,
			// 	"Failure at GetTag, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		response2, restyResp2, err := client.Tag.GetTagByID(response1.ID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}

		vItem2 := flattenTagGetTagByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagByID response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTagByID")
		vvID := vID

		response2, restyResp2, err := client.Tag.GetTagByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTagGetTagByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vName := resourceMap["name"]
	var vvID string
	if vID != "" {
		vvID = vID
		getResp, _, err := client.Tag.GetTagByID(vvID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}
	}
	if vName != "" {
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		queryParams1.Name = vName
		response1, err := searchTagGetTag(m, queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTag, unexpected response", ""))
			return diags
		}
		vvID = response1.ID
	}
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }

	//Set value vvName = getResp
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Id used for update operation %s", vvID)
		request1 := expandRequestTagUpdateTag(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Tag.UpdateTag(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTag", err, restyResp1.String(),
					"Failure at UpdateTag, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTag", err,
				"Failure at UpdateTag, unexpected response", ""))
			return diags
		}
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateTag", err))
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
				log.Printf("[DEBUG] Error => %s", response2.Response.FailureReason)
				diags = append(diags, diagError(
					"Failure when executing UpdateTag", err))
				return diags
			}
		}
	}

	return resourceTagRead(ctx, d, m)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vName := resourceMap["name"]
	var vvID string
	if vID != "" {
		vvID = vID
		getResp, _, err := client.Tag.GetTagByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			d.SetId("")
			return diags
		}
	} else if vName != "" {
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		queryParams1.Name = vName
		getResp, err := searchTagGetTag(m, queryParams1)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			d.SetId("")
			return diags
		}
		vvID = getResp.ID
	}
	response1, restyResp1, err := client.Tag.DeleteTag(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTag", err, restyResp1.String(),
				"Failure at DeleteTag, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTag", err,
			"Failure at DeleteTag, unexpected response", ""))
		return diags
	}
	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteTag", err))
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
			log.Printf("[DEBUG] Error => %s", response2.Response.FailureReason)
			diags = append(diags, diagError(
				"Failure when executing DeleteTag", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTagCreateTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTag {
	request := dnacentersdkgo.RequestTagCreateTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_tag")))) {
		request.SystemTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_rules")))) {
		request.DynamicRules = expandRequestTagCreateTagDynamicRulesArray(ctx, key+".dynamic_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagCreateTagDynamicRules {
	request := []dnacentersdkgo.RequestTagCreateTagDynamicRules{}
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
		i := expandRequestTagCreateTagDynamicRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRules {
	request := dnacentersdkgo.RequestTagCreateTagDynamicRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestTagCreateTagDynamicRulesRules(ctx, key+".rules.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRulesRules {
	request := dnacentersdkgo.RequestTagCreateTagDynamicRulesRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".values")))) {
		request.Values = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestTagCreateTagDynamicRulesRulesItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRulesItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems {
	request := []dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems{}
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
		i := expandRequestTagCreateTagDynamicRulesRulesItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRulesItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems {
	var request dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTag {
	request := dnacentersdkgo.RequestTagUpdateTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_tag")))) {
		request.SystemTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_rules")))) {
		request.DynamicRules = expandRequestTagUpdateTagDynamicRulesArray(ctx, key+".dynamic_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagUpdateTagDynamicRules {
	request := []dnacentersdkgo.RequestTagUpdateTagDynamicRules{}
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
		i := expandRequestTagUpdateTagDynamicRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRules {
	request := dnacentersdkgo.RequestTagUpdateTagDynamicRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestTagUpdateTagDynamicRulesRules(ctx, key+".rules.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRulesRules {
	request := dnacentersdkgo.RequestTagUpdateTagDynamicRulesRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".values")))) {
		request.Values = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestTagUpdateTagDynamicRulesRulesItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRulesItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems {
	request := []dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems{}
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
		i := expandRequestTagUpdateTagDynamicRulesRulesItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRulesItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems {
	var request dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchTagGetTag(m interface{}, queryParams dnacentersdkgo.GetTagQueryParams) (*dnacentersdkgo.ResponseTagGetTagResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseTagGetTagResponse
	var ite *dnacentersdkgo.ResponseTagGetTag
	ite, _, err = client.Tag.GetTag(&queryParams)
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
			var getItem *dnacentersdkgo.ResponseTagGetTagResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
