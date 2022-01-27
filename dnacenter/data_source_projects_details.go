package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProjectsDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get project(s) details
`,

		ReadContext: dataSourceProjectsDetailsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. Id of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Limits number of results
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Index of first result
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (dsc)
`,
				Type:     schema.TypeString,
				Optional: true,
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

						"is_deletable": &schema.Schema{
							Description: `Flag to check if project is deletable or not(for internal use only)
`,
							// Type:        schema.TypeBool,
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

func dataSourceProjectsDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortOrder, okSortOrder := d.GetOk("sort_order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetProjectsDetails")
		queryParams1 := dnacentersdkgo.GetProjectsDetailsQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetProjectsDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetProjectsDetails", err,
				"Failure at GetProjectsDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetProjectsDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProjectsDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetProjectsDetailsItems(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["create_time"] = item.CreateTime
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["is_deletable"] = boolPtrToString(item.IsDeletable)
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["tags"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTags(item.Tags)
		respItem["templates"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplates(item.Templates)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplates(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTags(item.Tags)
		respItem["author"] = item.Author
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["containing_templates"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplates(item.ContainingTemplates)
		respItem["create_time"] = item.CreateTime
		respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesDeviceTypes(item.DeviceTypes)
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
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["software_type"] = item.SoftwareType
		respItem["software_variant"] = item.SoftwareVariant
		respItem["software_version"] = item.SoftwareVersion
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParams(item.TemplateParams)
		respItem["validation_errors"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrors(item.ValidationErrors)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplates(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTags(item.Tags)
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrors(item *dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrorsRollbackTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsRollbackTemplateErrors) []interface{} {
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

func flattenConfigurationTemplatesGetProjectsDetailsItemsTemplatesValidationErrorsTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsTemplateErrors) []interface{} {
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
