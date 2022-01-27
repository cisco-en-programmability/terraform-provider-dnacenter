package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplateProject() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- List the projects

- Get the details of the given project by its id.
`,

		ReadContext: dataSourceConfigurationTemplateProjectRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name query parameter. Name of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId path parameter. projectId(UUID) of project to get project details
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
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

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
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
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
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

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
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
					},
				},
			},
		},
	}
}

func dataSourceConfigurationTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vProjectID, okProjectID := d.GetOk("project_id")

	method1 := []bool{okName, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsAListOfProjects")
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsAListOfProjects", err,
				"Failure at GetsAListOfProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetsAListOfProjectsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAListOfProjects response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetsTheDetailsOfAGivenProject")
		vvProjectID := vProjectID.(string)

		response2, restyResp2, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheDetailsOfAGivenProject response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetsAListOfProjectsItems(items *dnacentersdkgo.ResponseConfigurationTemplatesGetsAListOfProjects) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(item.Tags)
		respItem["create_time"] = item.CreateTime
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["templates"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(item.Templates)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTags(item.Tags)
		respItem["author"] = item.Author
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["containing_templates"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplates(item.ContainingTemplates)
		respItem["create_time"] = item.CreateTime
		respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesDeviceTypes(item.DeviceTypes)
		respItem["document_database"] = boolPtrToString(item.DocumentDatabase)
		respItem["failure_policy"] = item.FailurePolicy
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["latest_version_time"] = item.LatestVersionTime
		respItem["name"] = item.Name
		respItem["parent_template_id"] = item.ParentTemplateID
		respItem["project_associated"] = boolPtrToString(item.ProjectAssociated)
		respItem["project_id"] = item.ProjectID
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_content"] = item.RollbackTemplateContent
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["software_type"] = item.SoftwareType
		respItem["software_variant"] = item.SoftwareVariant
		respItem["software_version"] = item.SoftwareVersion
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParams(item.TemplateParams)
		respItem["validation_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrors(item.ValidationErrors)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTags(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplates(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTags(item.Tags)
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTags(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrors(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsRollbackTemplateErrors(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsTemplateErrors(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(item.Tags)
	respItem["create_time"] = item.CreateTime
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["templates"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(item.Templates)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTags(item.Tags)
		respItem["author"] = item.Author
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["containing_templates"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplates(item.ContainingTemplates)
		respItem["create_time"] = item.CreateTime
		respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesDeviceTypes(item.DeviceTypes)
		respItem["document_database"] = boolPtrToString(item.DocumentDatabase)
		respItem["failure_policy"] = item.FailurePolicy
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["latest_version_time"] = item.LatestVersionTime
		respItem["name"] = item.Name
		respItem["parent_template_id"] = item.ParentTemplateID
		respItem["project_associated"] = boolPtrToString(item.ProjectAssociated)
		respItem["project_id"] = item.ProjectID
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_content"] = item.RollbackTemplateContent
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["software_type"] = item.SoftwareType
		respItem["software_variant"] = item.SoftwareVariant
		respItem["software_version"] = item.SoftwareVersion
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParams(item.TemplateParams)
		respItem["validation_errors"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrors(item.ValidationErrors)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplates(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTags(item.Tags)
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrors(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrorsRollbackTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplatesValidationErrorsTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
