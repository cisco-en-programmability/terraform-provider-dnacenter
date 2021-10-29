package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{

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
				MaxItems: 1,
				Required: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_tag": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"rules": &schema.Schema{
										Type: schema.TypeList,
										// MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Default:  "",
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Default:  "",
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Default:  "",
												},
												"values": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"items": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
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
					},
				},
			},
		},
	}
}

func constructUpdateTag(tagID string, tag map[string]interface{}) *dnac.UpdateTagRequest {
	tagRequest := dnac.UpdateTagRequest{
		ID:          tagID,
		SystemTag:   tag["system_tag"].(bool),
		Description: tag["description"].(string),
		Name:        tag["name"].(string),
	}
	dataInstanceTenantID, okInstanceTenantID := tag["instance_tenant_id"]
	if okInstanceTenantID {
		tagRequest.InstanceTenantID = dataInstanceTenantID.(string)
	}

	if v, ok := tag["dynamic_rules"]; ok && v != nil {
		dynamicRules := tag["dynamic_rules"].([]interface{})

		if len(dynamicRules) > 0 {
			for _, dynamicRule := range dynamicRules {
				dR := dynamicRule.(map[string]interface{})

				tdr := dnac.UpdateTagRequestDynamicRules{
					MemberType: dR["member_type"].(string),
				}
				if drRules, ok := dR["rules"]; ok {
					ru := drRules.([]interface{})[0]
					rules := ru.(map[string]interface{})

					if rulesName, ok := rules["name"]; ok {
						tdr.Rules.Name = rulesName.(string)
					}
					if rulesOp, ok := rules["operation"]; ok {
						tdr.Rules.Operation = rulesOp.(string)
					}
					if rulesValue, ok := rules["value"]; ok {
						tdr.Rules.Value = rulesValue.(string)
					}
					if rulesItems, ok := rules["items"]; ok {
						items := rulesItems.([]interface{})
						if len(items) > 0 {
							for _, item := range items {
								tdr.Rules.Items = append(tdr.Rules.Items, item.(string))
							}
						}
					}
					if rulesValues, ok := rules["values"]; ok {
						values := rulesValues.([]interface{})
						if len(values) > 0 {
							for _, value := range values {
								tdr.Rules.Values = append(tdr.Rules.Values, value.(string))
							}
						}
					}
				}
				tagRequest.DynamicRules = append(tagRequest.DynamicRules, tdr)
			}
		}
	}
	return &tagRequest
}

func constructCreateTag(tag map[string]interface{}) *dnac.CreateTagRequest {
	tagRequest := dnac.CreateTagRequest{
		SystemTag:   tag["system_tag"].(bool),
		Description: tag["description"].(string),
		Name:        tag["name"].(string),
	}
	dataInstanceTenantID, okInstanceTenantID := tag["instance_tenant_id"]
	if okInstanceTenantID {
		tagRequest.InstanceTenantID = dataInstanceTenantID.(string)
	}

	if v, ok := tag["dynamic_rules"]; ok && v != nil {
		dynamicRules := tag["dynamic_rules"].([]interface{})

		if len(dynamicRules) > 0 {
			for _, dynamicRule := range dynamicRules {
				dR := dynamicRule.(map[string]interface{})

				tdr := dnac.CreateTagRequestDynamicRules{
					MemberType: dR["member_type"].(string),
				}
				if drRules, ok := dR["rules"]; ok {
					ru := drRules.([]interface{})[0]
					rules := ru.(map[string]interface{})

					if rulesName, ok := rules["name"]; ok {
						tdr.Rules.Name = rulesName.(string)
					}
					if rulesOp, ok := rules["operation"]; ok {
						tdr.Rules.Operation = rulesOp.(string)
					}
					if rulesValue, ok := rules["value"]; ok {
						tdr.Rules.Value = rulesValue.(string)
					}
					if rulesItems, ok := rules["items"]; ok {
						items := rulesItems.([]interface{})
						if len(items) > 0 {
							for _, item := range items {
								tdr.Rules.Items = append(tdr.Rules.Items, item.(string))
							}
						}
					}
					if rulesValues, ok := rules["values"]; ok {
						values := rulesValues.([]interface{})
						if len(values) > 0 {
							for _, value := range values {
								tdr.Rules.Values = append(tdr.Rules.Values, value.(string))
							}
						}
					}
				}
				tagRequest.DynamicRules = append(tagRequest.DynamicRules, tdr)
			}
		}
	}
	return &tagRequest
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	tag := item.(map[string]interface{})

	// Check if element already exists
	tagName := tag["name"].(string)
	tagQueryParams := &dnac.GetTagQueryParams{
		Name: tagName,
	}
	searchResponse, _, err := client.Tag.GetTag(tagQueryParams)
	if err == nil && searchResponse != nil && len(searchResponse.Response) > 0 {
		searchTagResponse := searchResponse.Response[0]

		updateTagRequest := constructUpdateTag(searchTagResponse.ID, tag)
		_, _, err := client.Tag.UpdateTag(updateTagRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Update resource id
		d.SetId(searchTagResponse.ID)
		// Update resource data (on Terraform and DNAC)
		resourceTagRead(ctx, d, m)
		return diags
	}

	// Construct payload from resource schema (item)
	tagRequest := constructCreateTag(tag)

	// Call function to create tag resource
	response, _, err := client.Tag.CreateTag(tagRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if task was completed successfully
	if taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create tag",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	d.SetId(taskResponse.Response.Data)
	// Update resource on Terraform
	resourceTagRead(ctx, d, m)
	return diags
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of tag.id)
	tagID := d.Id()

	// Call function to read tag.id
	response, _, err := client.Tag.GetTagByID(tagID)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	tagRead := flattenTagReadItem(response)
	if err := d.Set("item", tagRead); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagID := d.Id()
	_, _, err := client.Tag.GetTagByID(tagID)
	if err != nil {
		d.SetId("")
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		tag := item.(map[string]interface{})

		// Construct payload from resource schema (item)
		tagRequest := constructUpdateTag(tagID, tag)

		// Call function to update tag resource
		response, _, err := client.Tag.UpdateTag(tagRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Call function to check task
		taskID := response.Response.TaskID
		taskResponse, _, err := client.Task.GetTaskByID(taskID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Check if task was completed successfully
		if taskResponse.Response.IsError {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update tag",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceTagRead(ctx, d, m)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagID := d.Id()

	// Call function to delete resource
	response, _, err := client.Tag.DeleteTag(tagID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResponse != nil && taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete tag",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	return diags
}
