package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get template(s) details
`,

		ReadContext: dataSourceTemplatesDetailsRead,
		Schema: map[string]*schema.Schema{
			"all_template_attributes": &schema.Schema{
				Description: `allTemplateAttributes query parameter. Return all template attributes
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"filter_conflicting_templates": &schema.Schema{
				Description: `filterConflictingTemplates query parameter. Filter template(s) based on confliting templates
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of template to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_version_details": &schema.Schema{
				Description: `includeVersionDetails query parameter. Include template version details
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Limits number of results
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of template to be searched
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
			"product_family": &schema.Schema{
				Description: `productFamily query parameter. Filter template(s) based on device family
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_series": &schema.Schema{
				Description: `productSeries query parameter. Filter template(s) based on device series
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Description: `productType query parameter. Filter template(s) based on device type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId query parameter. Filter template(s) based on project id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_name": &schema.Schema{
				Description: `projectName query parameter. Filter template(s) based on project name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_type": &schema.Schema{
				Description: `softwareType query parameter. Filter template(s) based software type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter. Filter template(s) based softwareVersion
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (dsc)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Description: `tags query parameter. Filter template(s) based on tags
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"un_committed": &schema.Schema{
				Description: `unCommitted query parameter. Return uncommitted template
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"items": &schema.Schema{
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
							// Type:     schema.TypeBool,
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

						"versions_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_comment": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_time": &schema.Schema{
										Type:     schema.TypeInt,
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

func dataSourceTemplatesDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vProjectID, okProjectID := d.GetOk("project_id")
	vProjectName, okProjectName := d.GetOk("project_name")
	vSoftwareType, okSoftwareType := d.GetOk("software_type")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vProductFamily, okProductFamily := d.GetOk("product_family")
	vProductSeries, okProductSeries := d.GetOk("product_series")
	vProductType, okProductType := d.GetOk("product_type")
	vFilterConflictingTemplates, okFilterConflictingTemplates := d.GetOk("filter_conflicting_templates")
	vTags, okTags := d.GetOk("tags")
	vUnCommitted, okUnCommitted := d.GetOk("un_committed")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vAllTemplateAttributes, okAllTemplateAttributes := d.GetOk("all_template_attributes")
	vIncludeVersionDetails, okIncludeVersionDetails := d.GetOk("include_version_details")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTemplatesDetails")
		queryParams1 := dnacentersdkgo.GetTemplatesDetailsQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okProjectID {
			queryParams1.ProjectID = vProjectID.(string)
		}
		if okProjectName {
			queryParams1.ProjectName = vProjectName.(string)
		}
		if okSoftwareType {
			queryParams1.SoftwareType = vSoftwareType.(string)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = vSoftwareVersion.(string)
		}
		if okProductFamily {
			queryParams1.ProductFamily = vProductFamily.(string)
		}
		if okProductSeries {
			queryParams1.ProductSeries = vProductSeries.(string)
		}
		if okProductType {
			queryParams1.ProductType = vProductType.(string)
		}
		if okFilterConflictingTemplates {
			queryParams1.FilterConflictingTemplates = vFilterConflictingTemplates.(bool)
		}
		if okTags {
			queryParams1.Tags = interfaceToSliceString(vTags)
		}
		if okUnCommitted {
			queryParams1.UnCommitted = vUnCommitted.(bool)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okAllTemplateAttributes {
			queryParams1.AllTemplateAttributes = vAllTemplateAttributes.(bool)
		}
		if okIncludeVersionDetails {
			queryParams1.IncludeVersionDetails = vIncludeVersionDetails.(bool)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplatesDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTemplatesDetails", err,
				"Failure at GetTemplatesDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetTemplatesDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplatesDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplatesDetailsItems(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["author"] = item.Author
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["containing_templates"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplates(item.ContainingTemplates)
		respItem["create_time"] = item.CreateTime
		respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsDeviceTypes(item.DeviceTypes)
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
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["software_type"] = item.SoftwareType
		respItem["software_variant"] = item.SoftwareVariant
		respItem["software_version"] = item.SoftwareVersion
		respItem["tags"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsTags(item.Tags)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParams(item.TemplateParams)
		respItem["validation_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrors(item.ValidationErrors)
		respItem["version"] = item.Version
		respItem["versions_info"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsVersionsInfo(item.VersionsInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplates(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["tags"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTags(item.Tags)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsContainingTemplatesTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsDeviceTypes(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsRollbackTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParams(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsRange(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsSelection(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsTemplateParamsSelectionSelectionValues(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrors(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrorsRollbackTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsRollbackTemplateErrors) []interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsValidationErrorsTemplateErrors(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsTemplateErrors) []interface{} {
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

func flattenConfigurationTemplatesGetTemplatesDetailsItemsVersionsInfo(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsResponseVersionsInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["author"] = item.Author
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["version"] = item.Version
		respItem["version_comment"] = item.VersionComment
		respItem["version_time"] = item.VersionTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
