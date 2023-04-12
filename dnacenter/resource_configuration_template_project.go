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

func resourceConfigurationTemplateProject() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Configuration Templates.

- This resource is used to create a new project.

- This resource is used to update an existing project.

- Deletes the project by its id
`,

		CreateContext: resourceConfigurationTemplateProjectCreate,
		ReadContext:   resourceConfigurationTemplateProjectRead,
		UpdateContext: resourceConfigurationTemplateProjectUpdate,
		DeleteContext: resourceConfigurationTemplateProjectDelete,
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
							Description: `Is deletable`,

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
																			Type:     schema.TypeString,
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
																			Type:     schema.TypeString,
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
																Type:     schema.TypeString,
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
																Type:     schema.TypeString,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"project_id": &schema.Schema{
							Description: `projectId path parameter. projectId(UUID) of project to be deleted
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
						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeList,
							Optional: true,
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
										Default:  "",
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
				},
			},
		},
	}
}

func resourceConfigurationTemplateProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestConfigurationTemplateProjectCreateProject(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vProjectID, okProjectID := resourceItem["project_id"]
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vvProjectID := interfaceToString(vProjectID)
	if okProjectID && vvProjectID != "" {
		getResponse2, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["project_id"] = vvProjectID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateProjectRead(ctx, d, m)
		}
	} else {
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}
		queryParams1.Name = vvName
		item2, err := searchConfigurationTemplatesGetsAListOfProjects(m, queryParams1)
		if err == nil && item2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["project_id"] = vvProjectID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateProjectRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.CreateProject(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateProject", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateProject", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateProject", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateConfigurationTemplateProject", err1))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["project_id"] = vvProjectID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceConfigurationTemplateProjectRead(ctx, d, m)
}

func resourceConfigurationTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vProjectID, okProjectID := resourceMap["project_id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsAListOfProjects")
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName
		}
		response1, restyResp1, _ := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams1)

		/*		if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetsTheDetailsOfAGivenProject", err,
					"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
				return diags
			}*/
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item2, err := searchConfigurationTemplatesGetsAListOfProjects(m, queryParams1)
		if err != nil && item2 == nil {
			d.SetId("")
			return diags
		}
		response2, restyResp2, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(item2.ID)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		vItem1 := flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(response2)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheDetailsOfAGivenProject response",
				err))
			return diags
		}
		return diags
	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetsTheDetailsOfAGivenProject")
		vvProjectID := vProjectID

		response2, restyResp2, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheDetailsOfAGivenProject response",
				err))
			return diags
		}
		if response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourceConfigurationTemplateProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vProjectID := resourceMap["project_id"]
	vName := resourceMap["name"]

	// NOTE: Consider adding getAllItems and search function to get missing params
	if vProjectID != "" {
		getResp, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vProjectID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}
	} else if vName != "" {
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}
		queryParams1.Name = vName
		item2, err := searchConfigurationTemplatesGetsAListOfProjects(m, queryParams1)
		if err != nil || item2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}
		vProjectID = item2.ID
	}

	if d.HasChange("parameters") {
		//log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestConfigurationTemplateProjectUpdateProject(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil && request1.ID == "" {
			request1.ID = vProjectID
		}
		response1, restyResp1, err := client.ConfigurationTemplates.UpdateProject(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateProject", err, restyResp1.String(),
					"Failure at UpdateProject, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateProject", err,
				"Failure at UpdateProject, unexpected response", ""))
			return diags
		}
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateProject", err))
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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UdpateConfigurationTemplateProject", err1))
				return diags
			}
		}
	}

	return resourceConfigurationTemplateProjectRead(ctx, d, m)
}

func resourceConfigurationTemplateProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vProjectID, okProjectID := resourceMap["project_id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	var vvID string
	if selectedMethod == 1 {
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}
		queryParams1.Name = vName
		item1, err := searchConfigurationTemplatesGetsAListOfProjects(m, queryParams1)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vProjectID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vProjectID
		}
	}
	if selectedMethod == 2 {
		vvID = vProjectID
		getResp, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.ConfigurationTemplates.DeletesTheProject(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesTheProject", err, restyResp1.String(),
				"Failure at DeletesTheProject, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesTheProject", err,
			"Failure at DeletesTheProject, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestConfigurationTemplateProjectCreateProject(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProject {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProject{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectCreateProjectTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".templates")))) {
		request.Templates = expandRequestConfigurationTemplateProjectCreateProjectTemplates(ctx, key+".templates.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags{}
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
		i := expandRequestConfigurationTemplateProjectCreateProjectTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags{}
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

func expandRequestConfigurationTemplateProjectCreateProjectTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProject(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProject {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProject{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectUpdateProjectTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".templates")))) {
		request.Templates = expandRequestConfigurationTemplateProjectUpdateProjectTemplates(ctx, key+".templates.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags{}
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
		i := expandRequestConfigurationTemplateProjectUpdateProjectTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags{}
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

func expandRequestConfigurationTemplateProjectUpdateProjectTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchConfigurationTemplatesGetsAListOfProjects(m interface{}, queryParams dnacentersdkgo.GetsAListOfProjectsQueryParams) (*dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjects, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error

	var foundItem *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjects
	var ite *dnacentersdkgo.ResponseConfigurationTemplatesGetsAListOfProjects
	ite, _, err = client.ConfigurationTemplates.GetsAListOfProjects(&queryParams)
	if err != nil {
		return foundItem, err
	}
	if ite == nil {
		return foundItem, err
	}

	items := ite

	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjects
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
