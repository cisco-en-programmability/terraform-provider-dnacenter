package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func templateParam() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instruction_text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_param": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"param_array": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parameter_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"selection": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"selection_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"selection_values": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceTemplate() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,
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
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"author": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"composite": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"failure_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_template_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"project_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rollback_template_content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"software_variant": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"composite": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"product_family": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_series": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     templateParam(),
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     templateParam(),
						},
					},
				},
			},
		},
	}
}

func constructUpdateTemplate(templateID string, template map[string]interface{}) *dnac.UpdateTemplateRequest {
	projectID := template["project_id"].(string)
	name := template["name"].(string)

	softwareType := template["software_type"].(string)
	updateTemplateRequest := dnac.UpdateTemplateRequest{Name: name, SoftwareType: softwareType}

	updateTemplateRequest.ID = templateID
	updateTemplateRequest.ProjectID = projectID
	if v, ok := template["author"]; ok && v != nil {
		updateTemplateRequest.Author = v.(string)
	}
	if v, ok := template["composite"]; ok && v != nil {
		updateTemplateRequest.Composite = v.(bool)
	}
	if v, ok := template["create_time"]; ok && v != nil {
		updateTemplateRequest.CreateTime = v.(int)
	}
	if v, ok := template["description"]; ok && v != nil {
		updateTemplateRequest.Description = v.(string)
	}
	if v, ok := template["failure_policy"]; ok && v != nil {
		updateTemplateRequest.FailurePolicy = v.(string)
	}
	if v, ok := template["last_update_time"]; ok && v != nil {
		updateTemplateRequest.LastUpdateTime = v.(int)
	}
	if v, ok := template["parent_template_id"]; ok && v != nil {
		updateTemplateRequest.ParentTemplateID = v.(string)
	}
	if v, ok := template["project_name"]; ok && v != nil {
		updateTemplateRequest.ProjectName = v.(string)
	}
	if v, ok := template["rollback_template_content"]; ok && v != nil {
		updateTemplateRequest.RollbackTemplateContent = v.(string)
	}
	if v, ok := template["software_variant"]; ok && v != nil {
		updateTemplateRequest.SoftwareVariant = v.(string)
	}
	if v, ok := template["software_version"]; ok && v != nil {
		updateTemplateRequest.SoftwareVersion = v.(string)
	}
	if v, ok := template["template_content"]; ok && v != nil {
		updateTemplateRequest.TemplateContent = v.(string)
	}
	if v, ok := template["version"]; ok && v != nil {
		updateTemplateRequest.Version = v.(string)
	}

	deviceTypes := template["device_types"].([]interface{})
	for _, deviceType := range deviceTypes {
		dT := deviceType.(map[string]interface{})
		deviceTypeObject := dnac.UpdateTemplateRequestDeviceTypes{}
		if v, ok := dT["product_family"]; ok && v != nil {
			deviceTypeObject.ProductFamily = v.(string)
		}
		if v, ok := dT["product_series"]; ok && v != nil {
			deviceTypeObject.ProductSeries = v.(string)
		}
		if v, ok := dT["product_type"]; ok && v != nil {
			deviceTypeObject.ProductType = v.(string)
		}
		updateTemplateRequest.DeviceTypes = append(updateTemplateRequest.DeviceTypes, deviceTypeObject)
	}
	if v, ok := template["containing_templates"]; ok && v != nil {
		containingTemplates := v.([]interface{})
		for _, containingTemplate := range containingTemplates {
			cT := containingTemplate.(map[string]interface{})
			containingTemplateObject := dnac.UpdateTemplateRequestContainingTemplates{}
			if v, ok := cT["composite"]; ok && v != nil {
				containingTemplateObject.Composite = v.(bool)
			}
			if v, ok := cT["id"]; ok && v != nil {
				containingTemplateObject.ID = v.(string)
			}
			if v, ok := cT["name"]; ok && v != nil {
				containingTemplateObject.Name = v.(string)
			}
			if v, ok := cT["version"]; ok && v != nil {
				containingTemplateObject.Version = v.(string)
			}
			updateTemplateRequest.ContainingTemplates = append(updateTemplateRequest.ContainingTemplates, containingTemplateObject)
		}
	}
	if v, ok := template["rollback_template_params"]; ok && v != nil {
		rollbackTemplateParams := v.([]interface{})
		for _, rollbackTemplateParam := range rollbackTemplateParams {
			tP := rollbackTemplateParam.(map[string]interface{})
			tPParams := dnac.UpdateTemplateRequestRollbackTemplateParams{}
			if v, ok := tP["binding"]; ok && v != nil {
				tPParams.Binding = v.(string)
			}
			if v, ok := tP["data_type"]; ok && v != nil {
				tPParams.DataType = v.(string)
			}
			if v, ok := tP["default_value"]; ok && v != nil {
				tPParams.DefaultValue = v.(string)
			}
			if v, ok := tP["description"]; ok && v != nil {
				tPParams.Description = v.(string)
			}
			if v, ok := tP["display_name"]; ok && v != nil {
				tPParams.DisplayName = v.(string)
			}
			if v, ok := tP["group"]; ok && v != nil {
				tPParams.Group = v.(string)
			}
			if v, ok := tP["id"]; ok && v != nil {
				tPParams.ID = v.(string)
			}
			if v, ok := tP["instruction_text"]; ok && v != nil {
				tPParams.InstructionText = v.(string)
			}
			if v, ok := tP["key"]; ok && v != nil {
				tPParams.Key = v.(string)
			}
			if v, ok := tP["not_param"]; ok && v != nil {
				tPParams.NotParam = v.(bool)
			}
			if v, ok := tP["order"]; ok && v != nil {
				tPParams.Order = v.(int)
			}
			if v, ok := tP["param_array"]; ok && v != nil {
				tPParams.ParamArray = v.(bool)
			}
			if v, ok := tP["parameter_name"]; ok && v != nil {
				tPParams.ParameterName = v.(string)
			}
			if v, ok := tP["provider"]; ok && v != nil {
				tPParams.Provider = v.(string)
			}
			if v, ok := tP["required"]; ok && v != nil {
				tPParams.Required = v.(bool)
			}
			updateTemplateRequest.RollbackTemplateParams = append(updateTemplateRequest.RollbackTemplateParams, tPParams)
		}
	}
	if v, ok := template["template_params"]; ok && v != nil {
		templateParams := v.([]interface{})
		for _, templateParam := range templateParams {
			tP := templateParam.(map[string]interface{})
			tPParams := dnac.UpdateTemplateRequestTemplateParams{}
			if v, ok := tP["binding"]; ok && v != nil {
				tPParams.Binding = v.(string)
			}
			if v, ok := tP["data_type"]; ok && v != nil {
				tPParams.DataType = v.(string)
			}
			if v, ok := tP["default_value"]; ok && v != nil {
				tPParams.DefaultValue = v.(string)
			}
			if v, ok := tP["description"]; ok && v != nil {
				tPParams.Description = v.(string)
			}
			if v, ok := tP["display_name"]; ok && v != nil {
				tPParams.DisplayName = v.(string)
			}
			if v, ok := tP["group"]; ok && v != nil {
				tPParams.Group = v.(string)
			}
			if v, ok := tP["id"]; ok && v != nil {
				tPParams.ID = v.(string)
			}
			if v, ok := tP["instruction_text"]; ok && v != nil {
				tPParams.InstructionText = v.(string)
			}
			if v, ok := tP["key"]; ok && v != nil {
				tPParams.Key = v.(string)
			}
			if v, ok := tP["not_param"]; ok && v != nil {
				tPParams.NotParam = v.(bool)
			}
			if v, ok := tP["order"]; ok && v != nil {
				tPParams.Order = v.(int)
			}
			if v, ok := tP["param_array"]; ok && v != nil {
				tPParams.ParamArray = v.(bool)
			}
			if v, ok := tP["parameter_name"]; ok && v != nil {
				tPParams.ParameterName = v.(string)
			}
			if v, ok := tP["provider"]; ok && v != nil {
				tPParams.Provider = v.(string)
			}
			if v, ok := tP["required"]; ok && v != nil {
				tPParams.Required = v.(bool)
			}
			updateTemplateRequest.TemplateParams = append(updateTemplateRequest.TemplateParams, tPParams)
		}
	}
	return &updateTemplateRequest
}

func constructCreateTemplate(template map[string]interface{}) *dnac.CreateTemplateRequest {
	name := template["name"].(string)
	softwareType := template["software_type"].(string)
	createTemplateRequest := dnac.CreateTemplateRequest{Name: name, SoftwareType: softwareType}

	if v, ok := template["author"]; ok && v != nil {
		createTemplateRequest.Author = v.(string)
	}
	if v, ok := template["composite"]; ok && v != nil {
		createTemplateRequest.Composite = v.(bool)
	}
	if v, ok := template["create_time"]; ok && v != nil {
		createTemplateRequest.CreateTime = v.(int)
	}
	if v, ok := template["description"]; ok && v != nil {
		createTemplateRequest.Description = v.(string)
	}
	if v, ok := template["failure_policy"]; ok && v != nil {
		createTemplateRequest.FailurePolicy = v.(string)
	}
	if v, ok := template["id"]; ok && v != nil {
		createTemplateRequest.ID = v.(string)
	}
	if v, ok := template["last_update_time"]; ok && v != nil {
		createTemplateRequest.LastUpdateTime = v.(int)
	}
	if v, ok := template["parent_template_id"]; ok && v != nil {
		createTemplateRequest.ParentTemplateID = v.(string)
	}
	if v, ok := template["project_name"]; ok && v != nil {
		createTemplateRequest.ProjectName = v.(string)
	}
	if v, ok := template["rollback_template_content"]; ok && v != nil {
		createTemplateRequest.RollbackTemplateContent = v.(string)
	}
	if v, ok := template["software_variant"]; ok && v != nil {
		createTemplateRequest.SoftwareVariant = v.(string)
	}
	if v, ok := template["software_version"]; ok && v != nil {
		createTemplateRequest.SoftwareVersion = v.(string)
	}
	if v, ok := template["template_content"]; ok && v != nil {
		createTemplateRequest.TemplateContent = v.(string)
	}
	if v, ok := template["version"]; ok && v != nil {
		createTemplateRequest.Version = v.(string)
	}

	deviceTypes := template["device_types"].([]interface{})
	for _, deviceType := range deviceTypes {
		dT := deviceType.(map[string]interface{})
		deviceTypeObject := dnac.CreateTemplateRequestDeviceTypes{}
		if v, ok := dT["product_family"]; ok && v != nil {
			deviceTypeObject.ProductFamily = v.(string)
		}
		if v, ok := dT["product_series"]; ok && v != nil {
			deviceTypeObject.ProductSeries = v.(string)
		}
		if v, ok := dT["product_type"]; ok && v != nil {
			deviceTypeObject.ProductType = v.(string)
		}
		createTemplateRequest.DeviceTypes = append(createTemplateRequest.DeviceTypes, deviceTypeObject)
	}
	if v, ok := template["containing_templates"]; ok && v != nil {
		containingTemplates := v.([]interface{})
		for _, containingTemplate := range containingTemplates {
			cT := containingTemplate.(map[string]interface{})
			containingTemplateObject := dnac.CreateTemplateRequestContainingTemplates{}
			if v, ok := cT["composite"]; ok && v != nil {
				containingTemplateObject.Composite = v.(bool)
			}
			if v, ok := cT["id"]; ok && v != nil {
				containingTemplateObject.ID = v.(string)
			}
			if v, ok := cT["name"]; ok && v != nil {
				containingTemplateObject.Name = v.(string)
			}
			if v, ok := cT["version"]; ok && v != nil {
				containingTemplateObject.Version = v.(string)
			}
			createTemplateRequest.ContainingTemplates = append(createTemplateRequest.ContainingTemplates, containingTemplateObject)
		}
	}
	if v, ok := template["rollback_template_params"]; ok && v != nil {
		rollbackTemplateParams := v.([]interface{})
		for _, rollbackTemplateParam := range rollbackTemplateParams {
			tP := rollbackTemplateParam.(map[string]interface{})
			tPParams := dnac.CreateTemplateRequestRollbackTemplateParams{}
			if v, ok := tP["binding"]; ok && v != nil {
				tPParams.Binding = v.(string)
			}
			if v, ok := tP["data_type"]; ok && v != nil {
				tPParams.DataType = v.(string)
			}
			if v, ok := tP["default_value"]; ok && v != nil {
				tPParams.DefaultValue = v.(string)
			}
			if v, ok := tP["description"]; ok && v != nil {
				tPParams.Description = v.(string)
			}
			if v, ok := tP["display_name"]; ok && v != nil {
				tPParams.DisplayName = v.(string)
			}
			if v, ok := tP["group"]; ok && v != nil {
				tPParams.Group = v.(string)
			}
			if v, ok := tP["id"]; ok && v != nil {
				tPParams.ID = v.(string)
			}
			if v, ok := tP["instruction_text"]; ok && v != nil {
				tPParams.InstructionText = v.(string)
			}
			if v, ok := tP["key"]; ok && v != nil {
				tPParams.Key = v.(string)
			}
			if v, ok := tP["not_param"]; ok && v != nil {
				tPParams.NotParam = v.(bool)
			}
			if v, ok := tP["order"]; ok && v != nil {
				tPParams.Order = v.(int)
			}
			if v, ok := tP["param_array"]; ok && v != nil {
				tPParams.ParamArray = v.(bool)
			}
			if v, ok := tP["parameter_name"]; ok && v != nil {
				tPParams.ParameterName = v.(string)
			}
			if v, ok := tP["provider"]; ok && v != nil {
				tPParams.Provider = v.(string)
			}
			if v, ok := tP["required"]; ok && v != nil {
				tPParams.Required = v.(bool)
			}
			createTemplateRequest.RollbackTemplateParams = append(createTemplateRequest.RollbackTemplateParams, tPParams)
		}
	}
	if v, ok := template["template_params"]; ok && v != nil {
		templateParams := v.([]interface{})
		for _, templateParam := range templateParams {
			tP := templateParam.(map[string]interface{})
			tPParams := dnac.CreateTemplateRequestTemplateParams{}
			if v, ok := tP["binding"]; ok && v != nil {
				tPParams.Binding = v.(string)
			}
			if v, ok := tP["data_type"]; ok && v != nil {
				tPParams.DataType = v.(string)
			}
			if v, ok := tP["default_value"]; ok && v != nil {
				tPParams.DefaultValue = v.(string)
			}
			if v, ok := tP["description"]; ok && v != nil {
				tPParams.Description = v.(string)
			}
			if v, ok := tP["display_name"]; ok && v != nil {
				tPParams.DisplayName = v.(string)
			}
			if v, ok := tP["group"]; ok && v != nil {
				tPParams.Group = v.(string)
			}
			if v, ok := tP["id"]; ok && v != nil {
				tPParams.ID = v.(string)
			}
			if v, ok := tP["instruction_text"]; ok && v != nil {
				tPParams.InstructionText = v.(string)
			}
			if v, ok := tP["key"]; ok && v != nil {
				tPParams.Key = v.(string)
			}
			if v, ok := tP["not_param"]; ok && v != nil {
				tPParams.NotParam = v.(bool)
			}
			if v, ok := tP["order"]; ok && v != nil {
				tPParams.Order = v.(int)
			}
			if v, ok := tP["param_array"]; ok && v != nil {
				tPParams.ParamArray = v.(bool)
			}
			if v, ok := tP["parameter_name"]; ok && v != nil {
				tPParams.ParameterName = v.(string)
			}
			if v, ok := tP["provider"]; ok && v != nil {
				tPParams.Provider = v.(string)
			}
			if v, ok := tP["required"]; ok && v != nil {
				tPParams.Required = v.(bool)
			}
			createTemplateRequest.TemplateParams = append(createTemplateRequest.TemplateParams, tPParams)
		}
	}
	return &createTemplateRequest
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	template := item.(map[string]interface{})

	projectID := template["project_id"].(string)
	name := template["name"].(string)

	// REVIEW: Use 'get project templates' to check resource, and remove
	searchResponse, _, err := client.ConfigurationTemplates.GetsTheTemplatesAvailable(&dnac.GetsTheTemplatesAvailableQueryParams{
		ProjectID: projectID,
	})
	if err == nil && searchResponse != nil {
		for _, templateAvailable := range *searchResponse {
			if templateAvailable.Name == name {

				updateTemplateRequest := constructUpdateTemplate(templateAvailable.TemplateID, template)
				_, _, err = client.ConfigurationTemplates.UpdateTemplate(updateTemplateRequest)
				if err != nil {
					return diag.FromErr(err)
				}

				// Update resource id
				d.SetId(templateAvailable.TemplateID)
				// Update resource data
				resourceTemplateRead(ctx, d, m)
				return diags
			}
		}
	}
	createTemplateRequest := constructCreateTemplate(template)
	response, _, err := client.ConfigurationTemplates.CreateTemplate(projectID, createTemplateRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(10 * time.Second)

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
			Summary:  "Unable to create template",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// REVIEW: Create template version, so it is now available
	client.ConfigurationTemplates.VersionTemplate(&dnac.VersionTemplateRequest{
		TemplateID: taskResponse.Response.Data,
	})

	// Update resource id
	d.SetId(taskResponse.Response.Data)
	// Update resource on Terraform
	resourceTemplateRead(ctx, d, m)
	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of template template.id)
	templateID := d.Id()

	response, _, err := client.ConfigurationTemplates.GetTemplateDetails(templateID, &dnac.GetTemplateDetailsQueryParams{})
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	templateRead := flattenTemplateReadItem(response)
	if err := d.Set("item", templateRead); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of template template.id)
	templateID := d.Id()

	_, _, err := client.ConfigurationTemplates.GetTemplateDetails(templateID, &dnac.GetTemplateDetailsQueryParams{})
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		template := item.(map[string]interface{})

		updateTemplateRequest := constructUpdateTemplate(templateID, template)
		response, _, err := client.ConfigurationTemplates.UpdateTemplate(updateTemplateRequest)
		if err != nil {
			return diag.FromErr(err)
		}

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
				Summary:  "Unable to update template",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}

		// REVIEW: Update template version, so it is now available
		client.ConfigurationTemplates.VersionTemplate(&dnac.VersionTemplateRequest{
			TemplateID: templateID,
		})

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	templateID := d.Id()

	_, _, err := client.ConfigurationTemplates.GetTemplateDetails(templateID, &dnac.GetTemplateDetailsQueryParams{})
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	// Call function to delete template project resource
	deleteResponse, _, err := client.ConfigurationTemplates.DeleteTemplate(templateID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call function to check task
	taskID := deleteResponse.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResponse != nil && taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete template",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	return diags
}
