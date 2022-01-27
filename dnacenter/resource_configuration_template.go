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

func resourceConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Configuration Templates.

- API to update a template.

- Deletes the template by its id
`,

		CreateContext: resourceConfigurationTemplateCreate,
		ReadContext:   resourceConfigurationTemplateRead,
		UpdateContext: resourceConfigurationTemplateUpdate,
		DeleteContext: resourceConfigurationTemplateDelete,
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

						"author": &schema.Schema{
							Description: `Author of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
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
											},
										},
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
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
											},
										},
									},
									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time of template
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"custom_params_order": &schema.Schema{
							Description: `Custom Params Order
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"product_family": &schema.Schema{
										Description: `Device family
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_series": &schema.Schema{
										Description: `Device series
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_type": &schema.Schema{
										Description: `Device type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"document_database": &schema.Schema{
							Description: `Document Database
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"failure_policy": &schema.Schema{
							Description: `Define failure policy if template provisioning fails
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"language": &schema.Schema{
							Description: `Template language (JINJA or VELOCITY)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of template
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"latest_version_time": &schema.Schema{
							Description: `Latest versioned template time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_template_id": &schema.Schema{
							Description: `Parent templateID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_associated": &schema.Schema{
							Description: `Project Associated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Description: `Project UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Description: `Project name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rollback_template_content": &schema.Schema{
							Description: `Rollback template content
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
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
								},
							},
						},
						"software_type": &schema.Schema{
							Description: `Applicable device software type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_variant": &schema.Schema{
							Description: `Applicable device software variant
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Description: `Applicable device software version
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"template_content": &schema.Schema{
							Description: `Template content
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
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
								},
							},
						},
						"validation_errors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"template_id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"template_version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Current version of template
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
													MaxItems: 1,
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
													MaxItems: 1,
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
							Required: true,
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
										MaxItems: 1,
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
						"template_id": &schema.Schema{
							Description: `templateId path parameter. templateId(UUID) of template to be deleted
`,
							Type:     schema.TypeString,
							Required: true,
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
										MaxItems: 1,
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
						"validation_errors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func resourceConfigurationTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceConfigurationTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vTemplateID := resourceMap["id"]
	vTemplateName := resourceMap["name"]

	if vTemplateName != "" {
		log.Printf("[DEBUG] Selected method 1: GetsTheTemplatesAvailable")
		queryParams1 := dnacentersdkgo.GetsTheTemplatesAvailableQueryParams{}

		response1, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParams1, vTemplateName)

		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetsTheTemplatesAvailable", err,
			// 	"Failure at GetsTheTemplatesAvailable, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		queryParams2 := dnacentersdkgo.GetsDetailsOfAGivenTemplateQueryParams{}

		response2, restyResp2, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(response1.TemplateID, &queryParams2)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsDetailsOfAGivenTemplate", err,
				"Failure at GetsDetailsOfAGivenTemplate, unexpected response", ""))
			return diags
		}
		vItems1 := flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateItem(response2)
		if err := d.Set("item", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTemplatesAvailable response",
				err))
			return diags
		}
		return diags
	} else if vTemplateID != "" {
		log.Printf("[DEBUG] Selected method 2: GetsDetailsOfAGivenTemplate")
		vvTemplateID := vTemplateID
		queryParams2 := dnacentersdkgo.GetsDetailsOfAGivenTemplateQueryParams{}

		response2, restyResp2, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetsDetailsOfAGivenTemplate", err,
			// 	"Failure at GetsDetailsOfAGivenTemplate, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsDetailsOfAGivenTemplate response",
				err))
			return diags
		}
		return diags
	}
	return diags
}

func resourceConfigurationTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vTemplateID := resourceMap["id"]
	vTemplateName := resourceMap["name"]
	var vvTemplateID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if vTemplateID != "" {
		vvTemplateID = vTemplateID
		queryParams2 := dnacentersdkgo.GetsDetailsOfAGivenTemplateQueryParams{}

		getResp, _, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, &queryParams2)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsDetailsOfAGivenTemplate", err,
				"Failure at GetsDetailsOfAGivenTemplate, unexpected response", ""))
			return diags
		}
	} else if vTemplateName != "" {
		queryParams1 := dnacentersdkgo.GetsTheTemplatesAvailableQueryParams{}

		response1, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParams1, vTemplateName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheTemplatesAvailable", err,
				"Failure at GetsTheTemplatesAvailable, unexpected response", ""))
			return diags
		}
		vvTemplateID = response1.TemplateID
	}

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vTemplateName)
		request1 := expandRequestConfigurationTemplateUpdateTemplate(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil && request1.ID == "" {
			request1.ID = vvTemplateID
		}
		response1, restyResp1, err := client.ConfigurationTemplates.UpdateTemplate(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTemplate", err, restyResp1.String(),
					"Failure at UpdateTemplate, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTemplate", err,
				"Failure at UpdateTemplate, unexpected response", ""))
			return diags
		}
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateTemplate", err))
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
					"Failure when executing CreateApplicationSet", err))
				return diags
			}
		}
	}

	return resourceConfigurationTemplateRead(ctx, d, m)
}

func resourceConfigurationTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vTemplateID := resourceMap["id"]
	vTemplateName := resourceMap["name"]

	var vvTemplateID string
	// REVIEW: Add getAllItems and search function to get missing params
	if vTemplateID != "" {
		vvTemplateID = vTemplateID
		queryParams2 := dnacentersdkgo.GetsDetailsOfAGivenTemplateQueryParams{}

		getResp, _, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, &queryParams2)
		if err != nil || getResp == nil {
			return diags
		}
	}
	if vTemplateName != "" {
		queryParams1 := dnacentersdkgo.GetsTheTemplatesAvailableQueryParams{}

		response1, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParams1, vTemplateName)

		if err != nil || response1 == nil {
			return diags
		}
		vvTemplateID = response1.TemplateID
	}
	response1, restyResp1, err := client.ConfigurationTemplates.DeletesTheTemplate(vvTemplateID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesTheTemplate", err, restyResp1.String(),
				"Failure at DeletesTheTemplate, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesTheTemplate", err,
			"Failure at DeletesTheTemplate, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestConfigurationTemplateUpdateTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateUpdateTemplateTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesArray(ctx, key+".containing_templates", d)
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
		request.DeviceTypes = expandRequestConfigurationTemplateUpdateTemplateDeviceTypesArray(ctx, key+".device_types", d)
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
		request.RollbackTemplateParams = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
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
		request.TemplateParams = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags{}
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

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
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
		request.RollbackTemplateParams = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags{}
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

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes{}
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

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange{}
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

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes{}
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

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange{}
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

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange{}
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

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrors {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrorsArray(ctx, key+".rollback_template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrorsArray(ctx, key+".template_errors", d)
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

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchConfigurationTemplatesGetsTheTemplatesAvailable(m interface{}, queryParams dnacentersdkgo.GetsTheTemplatesAvailableQueryParams, vName string) (*dnacentersdkgo.ResponseItemConfigurationTemplatesGetsTheTemplatesAvailable, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsTheTemplatesAvailable
	//var allItems []*dnacenterskgo.ResponseItemApplicationPolicyGetApplications
	nResponse, _, err := client.ConfigurationTemplates.GetsTheTemplatesAvailable(nil)

	if err != nil {
		return foundItem, err
	}
	//maxPageSize := 10

	for _, item := range *nResponse {
		if vName == item.Name {
			foundItem = &item
			return foundItem, err
		}

	}
	return foundItem, err
}
