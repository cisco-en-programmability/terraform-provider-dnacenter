package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Configuration Templates.

- API to create a template by project id.

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
																Type:     schema.TypeString, //TEST,
																Computed: true,
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
																Type:     schema.TypeString, //TEST,
																Computed: true,
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
													Type:     schema.TypeString, //TEST,
													Computed: true,
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
													Type:     schema.TypeString, //TEST,
													Computed: true,
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
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"author": &schema.Schema{
							Description: `Author of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
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
													Computed:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Optional: true,
																Computed: true,
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
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
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
													Computed:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Optional: true,
																Computed: true,
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
										Computed: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time of template
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"custom_params_order": &schema.Schema{
							Description: `Custom Params Order
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Description: `Description of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"product_family": &schema.Schema{
										Description: `Device family
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_series": &schema.Schema{
										Description: `Device series
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_type": &schema.Schema{
										Description: `Device type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"failure_policy": &schema.Schema{
							Description: `Define failure policy if template provisioning fails
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"language": &schema.Schema{
							Description: `Template language (JINJA or VELOCITY)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of template
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"latest_version_time": &schema.Schema{
							Description: `Latest versioned template time
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_template_id": &schema.Schema{
							Description: `Parent templateID
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Description: `Project UUID
`,
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"project_name": &schema.Schema{
							Description: `Project name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rollback_template_content": &schema.Schema{
							Description: `Rollback template content
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
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
										Computed:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
													Computed: true,
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
							Computed: true,
						},
						"software_variant": &schema.Schema{
							Description: `Applicable device software variant
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Description: `Applicable device software version
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"template_content": &schema.Schema{
							Description: `Template content
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `templateId path parameter. templateId(UUID) of template to be deleted
`,
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
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
										Computed:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
													Computed: true,
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
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										Computed: true,
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										Computed: true,
									},
									"template_id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"template_version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Current version of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceConfigurationTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestConfigurationTemplateCreateTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTemplateID := resourceItem["template_id"]
	vvTemplateID := interfaceToString(vTemplateID)
	vProjectID := resourceItem["project_id"]
	vvProjectID := interfaceToString(vProjectID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if vvTemplateID != "" {
		getResponse2, _, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, nil)
		// log.Printf("[DEBUG] getResponse2 sent => %v", responseInterfaceToString(*getResponse2))
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["template_id"] = vvTemplateID
			resourceMap["project_id"] = vvProjectID
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.GetsTheTemplatesAvailableQueryParams{}
		queryParamImport.ProjectID = vvProjectID

		response2, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParamImport, vvName)
		// log.Printf("[DEBUG] response2 sent => %v", responseInterfaceToString(*response2))
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["template_id"] = response2.TemplateID
			resourceMap["project_id"] = response2.ProjectID
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.CreateTemplate(vvProjectID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTemplate", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateTemplate", err))
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
				"Failure when executing CreateTemplate", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetsTheTemplatesAvailableQueryParams{}
	queryParamValidate.ProjectID = vvProjectID
	item3, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateTemplate", err,
			"Failure at CreateTemplate, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["template_id"] = item3.TemplateID
	resourceMap["project_id"] = item3.ProjectID
	d.SetId(joinResourceID(resourceMap))
	return resourceConfigurationTemplateRead(ctx, d, m)
}

func resourceConfigurationTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vTemplateID := resourceMap["template_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsDetailsOfAGivenTemplate")
		vvTemplateID := vTemplateID
		queryParams1 := dnacentersdkgo.GetsDetailsOfAGivenTemplateQueryParams{}

		response1, restyResp1, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		// Review flatten function used
		vItem1 := flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTemplatesAvailable search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceConfigurationTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["template_id"]
	if d.HasChange("parameters") {
		request1 := expandRequestConfigurationTemplateUpdateTemplate(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && request1.ID == "" {
			request1.ID = vID
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateTemplate", err1))
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
	vvTemplateID := resourceMap["template_id"]

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

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesTheTemplate", err))
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
				"Failure when executing DeletesTheTemplate", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestConfigurationTemplateCreateTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateCreateTemplateTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesArray(ctx, key+".containing_templates", d)
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
		request.DeviceTypes = expandRequestConfigurationTemplateCreateTemplateDeviceTypesArray(ctx, key+".device_types", d)
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
		request.RollbackTemplateParams = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
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
		request.TemplateParams = expandRequestConfigurationTemplateCreateTemplateTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTags{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTags{}
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

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
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
		request.RollbackTemplateParams = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags{}
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

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes{}
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

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange{}
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

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes{}
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
		i := expandRequestConfigurationTemplateCreateTemplateDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes{}
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

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange{}
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

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams{}
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
		request.Range = expandRequestConfigurationTemplateCreateTemplateTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange{}
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

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection{}
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
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrors {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrorsRollbackTemplateErrors(ctx, key+".rollback_template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrorsTemplateErrors(ctx, key+".template_errors.0", d)
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

func expandRequestConfigurationTemplateCreateTemplateValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
	for item_no := range objs {
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
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx, key+".rollback_template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrors(ctx, key+".template_errors.0", d)
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

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
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
	nResponse, _, err := client.ConfigurationTemplates.GetsTheTemplatesAvailable(nil)

	if err != nil {
		log.Printf("[DEBUG] Error when search => %s", err.Error())
		return foundItem, err
	}
	//maxPageSize := 10
	log.Printf("[DEBUG] Start to find a concidence")
	for _, item := range *nResponse {
		if vName == item.Name {
			foundItem = &item
			return foundItem, err
		}

	}
	return foundItem, err
}
