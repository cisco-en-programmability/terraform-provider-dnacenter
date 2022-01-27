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

// dataSourceAction
func dataSourceConfigurationTemplateImportTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- Imports the templates provided in the DTO by project Name
`,

		ReadContext: dataSourceConfigurationTemplateImportTemplateRead,
		Schema: map[string]*schema.Schema{
			"do_version": &schema.Schema{
				Description: `doVersion query parameter. If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"project_name": &schema.Schema{
				Description: `projectName path parameter. Project name to create template under the project
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"payload": &schema.Schema{
				Description: `Array of RequestConfigurationTemplatesImportsTheTemplatesProvided`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"author": &schema.Schema{
							Description: `Author of template
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time of template
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"custom_params_order": &schema.Schema{
							Description: `Custom Params Order
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"description": &schema.Schema{
							Description: `Description of template
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"product_family": &schema.Schema{
										Description: `Device family
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"product_series": &schema.Schema{
										Description: `Device series
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"product_type": &schema.Schema{
										Description: `Device type
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"failure_policy": &schema.Schema{
							Description: `Define failure policy if template provisioning fails
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"language": &schema.Schema{
							Description: `Template language (JINJA or VELOCITY)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of template
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latest_version_time": &schema.Schema{
							Description: `Latest versioned template time
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"parent_template_id": &schema.Schema{
							Description: `Parent templateID
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"project_id": &schema.Schema{
							Description: `Project UUID
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"project_name": &schema.Schema{
							Description: `Project name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rollback_template_content": &schema.Schema{
							Description: `Rollback template content
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type: schema.TypeList,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"software_type": &schema.Schema{
							Description: `Applicable device software type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"software_variant": &schema.Schema{
							Description: `Applicable device software variant
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"software_version": &schema.Schema{
							Description: `Applicable device software version
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"template_content": &schema.Schema{
							Description: `Template content
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type: schema.TypeList,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"validation_errors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional: true,
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional: true,
									},
									"template_id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"template_version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Current version of template
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

func dataSourceConfigurationTemplateImportTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProjectName := d.Get("project_name")
	vDoVersion, okDoVersion := d.GetOk("do_version")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportsTheTemplatesProvided")
		vvProjectName := vProjectName.(string)
		request1 := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvided(ctx, "", d)
		queryParams1 := dnacentersdkgo.ImportsTheTemplatesProvidedQueryParams{}

		if okDoVersion {
			queryParams1.DoVersion = vDoVersion.(bool)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.ImportsTheTemplatesProvided(vvProjectName, request1, &queryParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportsTheTemplatesProvided", err,
				"Failure at ImportsTheTemplatesProvided, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesImportsTheTemplatesProvidedItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportsTheTemplatesProvided response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvided(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesImportsTheTemplatesProvided {
	request := dnacentersdkgo.RequestConfigurationTemplatesImportsTheTemplatesProvided{}
	if v := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvided {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvided{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvided {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvided{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesArray(ctx, key+".containing_templates", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_params_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_params_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_params_order")))) {
		request.CustomParamsOrder = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_policy")))) {
		request.FailurePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latest_version_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latest_version_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latest_version_time")))) {
		request.LatestVersionTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_template_id")))) {
		request.ParentTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_content")))) {
		request.RollbackTemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_type")))) {
		request.SoftwareType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_variant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_variant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_variant")))) {
		request.SoftwareVariant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors {
	request := dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsRollbackTemplateErrorsArray(ctx, key+".rollback_template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsTemplateErrorsArray(ctx, key+".template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_version")))) {
		request.TemplateVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsRollbackTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsRollbackTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors{}
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
		i := expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateImportTemplateImportsTheTemplatesProvidedItemValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenConfigurationTemplatesImportsTheTemplatesProvidedItem(item *dnacentersdkgo.ResponseConfigurationTemplatesImportsTheTemplatesProvidedResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
