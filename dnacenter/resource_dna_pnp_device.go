package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpWorkItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"output_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_taken": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func pnpWorkflowTasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"curr_work_item_idx": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"end_time": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"start_time": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"state": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"task_seq_no": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"time_taken": &schema.Schema{
					Type:     schema.TypeFloat,
					Optional: true,
					Computed: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"work_item_list": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
					Elem:     pnpWorkItems(),
				},
			},
		},
	}
}

func pnpKeyValueMap() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func pnpSystemWorkflow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_to_inventory": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"added_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"curr_task_idx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exec_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lastupdate_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tasks": pnpWorkflowTasks(),
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePnPDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePnPDeviceCreate,
		ReadContext:   resourcePnPDeviceRead,
		UpdateContext: resourcePnPDeviceUpdate,
		DeleteContext: resourcePnPDeviceDelete,
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
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"day_zero_config": &schema.Schema{
							Type: schema.TypeList,
							// MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"day_zero_config_preview": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_info": &schema.Schema{
							Type: schema.TypeList,
							// MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aaa_credentials": &schema.Schema{
										Type: schema.TypeList,
										// MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"username": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"addn_mac_addrs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"agent_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"authenticated_mic_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"authenticated_sudi_serial_no": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"capabilities_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cm_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"features_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"file_system_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"freespace": &schema.Schema{
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"readable": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"size": &schema.Schema{
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"writeable": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"first_contact": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"http_headers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem:     pnpKeyValueMap(),
									},
									"image_file": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"image_version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ipv6_address_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"last_contact": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"last_sync_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"last_update_on": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"location": &schema.Schema{
										Type: schema.TypeList,
										// MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"altitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"latitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"longitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"site_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"neighbor_links": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"local_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"local_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_device_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_platform": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"remote_version": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"onb_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pnp_profile_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"created_by": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"discovery_created": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"primary_endpoint": &schema.Schema{
													Type: schema.TypeList,
													// MaxItems: 1,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"profile_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"secondary_endpoint": &schema.Schema{
													Type: schema.TypeList,
													// MaxItems: 1,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"populate_inventory": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"pre_workflow_cli_ouputs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cli": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"cli_output": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"project_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"project_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"reload_requested": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"site_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"site_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smart_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"source": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"stack": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"stack_info": &schema.Schema{
										Type: schema.TypeList,
										// MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_full_ring": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hardware_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"license_level": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"license_type": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"mac_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"pid": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"priority": &schema.Schema{
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"software_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"stack_number": &schema.Schema{
																Type:     schema.TypeFloat,
																Optional: true,
																Computed: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"sudi_serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"supports_stack_workflows": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"total_member_count": &schema.Schema{
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"valid_license_levels": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sudi_required": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_mic_numbers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"user_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"workflow_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"workflow_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"run_summary_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"details": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"error_flag": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"history_task_info": &schema.Schema{
										Type: schema.TypeList,
										// MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"addn_details": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem:     pnpKeyValueMap(),
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem:     pnpWorkItems(),
												},
											},
										},
									},
									"timestamp": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"system_reset_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							// MaxItems: 1,
							Elem: pnpSystemWorkflow(),
						},
						"system_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							// MaxItems: 1,
							Elem: pnpSystemWorkflow(),
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							// MaxItems: 1,
							Elem: pnpSystemWorkflow(),
						},
						"workflow_parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							// MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem:     pnpKeyValueMap(),
												},
											},
										},
									},
									"license_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"license_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"top_of_stack_serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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

func constructAddZeroConfig(zeroConfigs []interface{}) *dnac.AddDeviceToPnpDatabaseRequestDayZeroConfig {
	var zeroConfig dnac.AddDeviceToPnpDatabaseRequestDayZeroConfig
	if len(zeroConfigs) > 0 {
		zc := zeroConfigs[0].(map[string]interface{})
		zeroConfig.Config = zc["config"].(string)
	}
	return &zeroConfig
}

func constructAddDeviceInfoAAACredentials(credentials []interface{}) *dnac.AddDeviceToPnpDatabaseRequestDeviceInfoAAACredentials {
	var aaaCredentials dnac.AddDeviceToPnpDatabaseRequestDeviceInfoAAACredentials
	if len(credentials) > 0 {
		credential := credentials[0].(map[string]interface{})
		if v, ok := credential["username"]; ok {
			aaaCredentials.Username = v.(string)
		}
		if v, ok := credential["password"]; ok {
			aaaCredentials.Password = v.(string)
		}
	}
	return &aaaCredentials
}

func constructAddDeviceInfoFileSystemList(fileList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList {
	var fileSystemList []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList
	for _, fi := range fileList {
		fsi := fi.(map[string]interface{})
		fileSystemItem := dnac.AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList{}
		if v, ok := fsi["freespace"]; ok {
			fileSystemItem.Freespace = v.(float64)
		}
		if v, ok := fsi["name"]; ok {
			fileSystemItem.Name = v.(string)
		}
		if v, ok := fsi["readable"]; ok {
			fileSystemItem.Readable = v.(bool)
		}
		if v, ok := fsi["size"]; ok {
			fileSystemItem.Size = v.(float64)
		}
		if v, ok := fsi["type"]; ok {
			fileSystemItem.Type = v.(string)
		}
		if v, ok := fsi["writeable"]; ok {
			fileSystemItem.Writeable = v.(bool)
		}
		fileSystemList = append(fileSystemList, fileSystemItem)
	}
	return &fileSystemList
}

func constructAddDeviceInfoHTTPHeaders(headers []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders
	for _, header := range headers {
		hs := header.(map[string]interface{})
		result = append(result, dnac.AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders{
			Key:   hs["key"].(string),
			Value: hs["value"].(string),
		})
	}
	return &result
}

func constructAddDeviceInfoIPInterfaces(interfaces []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces
	for _, i := range interfaces {
		interfaceItem := dnac.AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces{}
		is := i.(map[string]interface{})
		if v, ok := is["ipv4_address"]; ok {
			interfaceItem.IPv4Address = v.(string)
		}
		if v, ok := is["ipv6_address_list"]; ok {
			interfaceItem.IPv6AddressList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := is["mac_address"]; ok {
			interfaceItem.MacAddress = v.(string)
		}
		if v, ok := is["name"]; ok {
			interfaceItem.Name = v.(string)
		}
		if v, ok := is["status"]; ok {
			interfaceItem.Status = v.(string)
		}
		result = append(result, interfaceItem)
	}
	return &result
}

func constructAddDeviceInfoLocation(locations []interface{}) *dnac.AddDeviceToPnpDatabaseRequestDeviceInfoLocation {
	var result dnac.AddDeviceToPnpDatabaseRequestDeviceInfoLocation
	if len(locations) > 0 {
		location := locations[0].(map[string]interface{})
		if v, ok := location["address"]; ok {
			result.Address = v.(string)
		}
		if v, ok := location["altitude"]; ok {
			result.Altitude = v.(string)
		}
		if v, ok := location["latitude"]; ok {
			result.Latitude = v.(string)
		}
		if v, ok := location["longitude"]; ok {
			result.Longitude = v.(string)
		}
		if v, ok := location["siteId"]; ok {
			result.SiteID = v.(string)
		}
	}
	return &result
}

func constructAddDeviceInfoNeighborLinks(links []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks
	for _, link := range links {
		var linkItem dnac.AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks
		ls := link.(map[string]interface{})
		if v, ok := ls["local_interface_name"]; ok {
			linkItem.LocalInterfaceName = v.(string)
		}
		if v, ok := ls["local_mac_address"]; ok {
			linkItem.LocalMacAddress = v.(string)
		}
		if v, ok := ls["local_short_interface_name"]; ok {
			linkItem.LocalShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_device_name"]; ok {
			linkItem.RemoteDeviceName = v.(string)
		}
		if v, ok := ls["remote_interface_name"]; ok {
			linkItem.RemoteInterfaceName = v.(string)
		}
		if v, ok := ls["remote_mac_address"]; ok {
			linkItem.RemoteMacAddress = v.(string)
		}
		if v, ok := ls["remote_platform"]; ok {
			linkItem.RemotePlatform = v.(string)
		}
		if v, ok := ls["remote_short_interface_name"]; ok {
			linkItem.RemoteShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_version"]; ok {
			linkItem.RemoteVersion = v.(string)
		}
		result = append(result, linkItem)
	}
	return &result
}

func constructAddDeviceInfoPnpProfileList(profiles []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList
	for _, profile := range profiles {
		var profileItem dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList
		ps := profile.(map[string]interface{})
		if v, ok := ps["created_by"]; ok {
			profileItem.CreatedBy = v.(string)
		}
		if v, ok := ps["discovery_created"]; ok {
			profileItem.DiscoveryCreated = v.(bool)
		}
		if v, ok := ps["primary_endpoint"]; ok {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok {
					profileItem.PrimaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok {
					profileItem.PrimaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok {
					profileItem.PrimaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok {
					profileItem.PrimaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok {
					profileItem.PrimaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok {
					profileItem.PrimaryEndpoint.Protocol = x.(string)
				}
			}
		}
		if v, ok := ps["profile_name"]; ok {
			profileItem.ProfileName = v.(string)
		}
		if v, ok := ps["secondary_endpoint"]; ok {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok {
					profileItem.SecondaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok {
					profileItem.SecondaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok {
					profileItem.SecondaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok {
					profileItem.SecondaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok {
					profileItem.SecondaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok {
					profileItem.SecondaryEndpoint.Protocol = x.(string)
				}
			}
		}
		result = append(result, profileItem)
	}
	return &result
}

func constructAddDeviceInfoPreWorkflowCliOuputs(cliOutputs []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs
	for _, cliOutput := range cliOutputs {
		var cliOutputItem dnac.AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs
		cos := cliOutput.(map[string]interface{})
		if v, ok := cos["cli"]; ok {
			cliOutputItem.Cli = v.(string)
		}
		if v, ok := cos["cli_output"]; ok {
			cliOutputItem.CliOutput = v.(string)
		}
		result = append(result, cliOutputItem)
	}
	return &result
}

func constructAddDeviceInfoStackInfoStackMemberList(memberList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList {
	var result []dnac.AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList
	for _, member := range memberList {
		var memberItem dnac.AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList
		ms := member.(map[string]interface{})
		if v, ok := ms["hardware_version"]; ok {
			memberItem.HardwareVersion = v.(string)
		}
		if v, ok := ms["license_level"]; ok {
			memberItem.LicenseLevel = v.(string)
		}
		if v, ok := ms["license_type"]; ok {
			memberItem.LicenseType = v.(string)
		}
		if v, ok := ms["mac_address"]; ok {
			memberItem.MacAddress = v.(string)
		}
		if v, ok := ms["pid"]; ok {
			memberItem.Pid = v.(string)
		}
		if v, ok := ms["priority"]; ok {
			memberItem.Priority = v.(float64)
		}
		if v, ok := ms["role"]; ok {
			memberItem.Role = v.(string)
		}
		if v, ok := ms["serial_number"]; ok {
			memberItem.SerialNumber = v.(string)
		}
		if v, ok := ms["software_version"]; ok {
			memberItem.SoftwareVersion = v.(string)
		}
		if v, ok := ms["stack_number"]; ok {
			memberItem.StackNumber = v.(float64)
		}
		if v, ok := ms["state"]; ok {
			memberItem.State = v.(string)
		}
		if v, ok := ms["sudi_serial_number"]; ok {
			memberItem.SudiSerialNumber = v.(string)
		}
		result = append(result, memberItem)
	}
	return &result
}

func constructAddDeviceInfoStackInfo(stackInfos []interface{}) *dnac.AddDeviceToPnpDatabaseRequestDeviceInfoStackInfo {
	var result dnac.AddDeviceToPnpDatabaseRequestDeviceInfoStackInfo
	if len(stackInfos) > 0 {
		stackInfo := stackInfos[0].(map[string]interface{})
		if v, ok := stackInfo["is_full_ring"]; ok {
			result.IsFullRing = v.(bool)
		}
		if v, ok := stackInfo["stack_member_list"]; ok {
			if w := constructAddDeviceInfoStackInfoStackMemberList(v.([]interface{})); w != nil {
				result.StackMemberList = *w
			}
		}
		if v, ok := stackInfo["stack_ring_protocol"]; ok {
			result.StackRingProtocol = v.(string)
		}
		if v, ok := stackInfo["supports_stack_workflows"]; ok {
			result.SupportsStackWorkflows = v.(bool)
		}
		if v, ok := stackInfo["total_member_count"]; ok {
			result.TotalMemberCount = v.(float64)
		}
		if v, ok := stackInfo["valid_license_levels"]; ok {
			result.ValidLicenseLevels = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
	}

	return &result
}

func constructAddRunSummaryListHistoryTaskInfoAddnDetails(details []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails {
	var result []dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails
	for _, detail := range details {
		hs := detail.(map[string]interface{})
		result = append(result, dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails{
			Key:   hs["key"].(string),
			Value: hs["value"].(string),
		})
	}
	return &result
}

func constructAddRunSummaryListHistoryTaskInfoWorkItemList(itemList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList {
	var result []dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructAddRunSummaryListHistoryTaskInfo(taskInfos []interface{}) *dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo {
	var result dnac.AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo
	if len(taskInfos) > 0 {
		taskInfo := taskInfos[0].(map[string]interface{})
		if v, ok := taskInfo["addn_details"]; ok {
			if w := constructAddRunSummaryListHistoryTaskInfoAddnDetails(v.([]interface{})); w != nil {
				result.AddnDetails = *w
			}
		}
		if v, ok := taskInfo["name"]; ok {
			result.Name = v.(string)
		}
		if v, ok := taskInfo["time_taken"]; ok {
			result.TimeTaken = v.(float64)
		}
		if v, ok := taskInfo["type"]; ok {
			result.Type = v.(string)
		}
		if v, ok := taskInfo["work_item_list"]; ok {
			if w := constructAddRunSummaryListHistoryTaskInfoWorkItemList(v.([]interface{})); w != nil {
				result.WorkItemList = *w
			}
		}
	}
	return &result
}

func constructAddDeviceInfo(deviceInfos []interface{}) *dnac.AddDeviceToPnpDatabaseRequestDeviceInfo {
	var deviceInfo dnac.AddDeviceToPnpDatabaseRequestDeviceInfo
	if len(deviceInfos) > 0 {
		deviceInfoItem := deviceInfos[0].(map[string]interface{})
		if v, ok := deviceInfoItem["aaa_credentials"]; ok {
			if w := constructAddDeviceInfoAAACredentials(v.([]interface{})); w != nil {
				deviceInfo.AAACredentials = *w
			}
		}
		if v, ok := deviceInfoItem["added_on"]; ok {
			deviceInfo.AddedOn = v.(float64)
		}
		if v, ok := deviceInfoItem["addn_mac_addrs"]; ok {
			deviceInfo.AddnMacAddrs = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["agent_type"]; ok {
			deviceInfo.AgentType = v.(string)
		}
		if v, ok := deviceInfoItem["auth_status"]; ok {
			deviceInfo.AuthStatus = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_mic_number"]; ok {
			deviceInfo.AuthenticatedMicNumber = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_sudi_serial_no"]; ok {
			deviceInfo.AuthenticatedSudiSerialNo = v.(string)
		}
		if v, ok := deviceInfoItem["capabilities_supported"]; ok {
			deviceInfo.CapabilitiesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["cm_state"]; ok {
			deviceInfo.CmState = v.(string)
		}
		if v, ok := deviceInfoItem["description"]; ok {
			deviceInfo.Description = v.(string)
		}
		if v, ok := deviceInfoItem["device_sudi_serial_nos"]; ok {
			deviceInfo.DeviceSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["device_type"]; ok {
			deviceInfo.DeviceType = v.(string)
		}
		if v, ok := deviceInfoItem["features_supported"]; ok {
			deviceInfo.FeaturesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["file_system_list"]; ok {
			if w := constructAddDeviceInfoFileSystemList(v.([]interface{})); w != nil {
				deviceInfo.FileSystemList = *w
			}
		}
		if v, ok := deviceInfoItem["first_contact"]; ok {
			deviceInfo.FirstContact = v.(float64)
		}
		if v, ok := deviceInfoItem["hostname"]; ok {
			deviceInfo.Hostname = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok {
			if w := constructAddDeviceInfoHTTPHeaders(v.([]interface{})); w != nil {
				deviceInfo.HTTPHeaders = *w
			}
		}
		if v, ok := deviceInfoItem["image_file"]; ok {
			deviceInfo.ImageFile = v.(string)
		}
		if v, ok := deviceInfoItem["image_version"]; ok {
			deviceInfo.ImageVersion = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok {
			if w := constructAddDeviceInfoIPInterfaces(v.([]interface{})); w != nil {
				deviceInfo.IPInterfaces = *w
			}
		}
		if v, ok := deviceInfoItem["last_contact"]; ok {
			deviceInfo.LastContact = v.(float64)
		}
		if v, ok := deviceInfoItem["last_sync_time"]; ok {
			deviceInfo.LastSyncTime = v.(float64)
		}
		if v, ok := deviceInfoItem["last_update_on"]; ok {
			deviceInfo.LastUpdateOn = v.(float64)
		}
		if v, ok := deviceInfoItem["location"]; ok {
			if w := constructAddDeviceInfoLocation(v.([]interface{})); w != nil {
				deviceInfo.Location = *w
			}
		}
		if v, ok := deviceInfoItem["mac_address"]; ok {
			deviceInfo.MacAddress = v.(string)
		}
		if v, ok := deviceInfoItem["mode"]; ok {
			deviceInfo.Mode = v.(string)
		}
		if v, ok := deviceInfoItem["name"]; ok {
			deviceInfo.Name = v.(string)
		}
		if v, ok := deviceInfoItem["neighbor_links"]; ok {
			if w := constructAddDeviceInfoNeighborLinks(v.([]interface{})); w != nil {
				deviceInfo.NeighborLinks = *w
			}
		}
		if v, ok := deviceInfoItem["onb_state"]; ok {
			deviceInfo.OnbState = v.(string)
		}
		if v, ok := deviceInfoItem["pid"]; ok {
			deviceInfo.Pid = v.(string)
		}
		if v, ok := deviceInfoItem["pnp_profile_list"]; ok {
			if w := constructAddDeviceInfoPnpProfileList(v.([]interface{})); w != nil {
				deviceInfo.PnpProfileList = *w
			}
		}
		if v, ok := deviceInfoItem["populate_inventory"]; ok {
			deviceInfo.PopulateInventory = v.(bool)
		}
		if v, ok := deviceInfoItem["pre_workflow_cli_ouputs"]; ok {
			if w := constructAddDeviceInfoPreWorkflowCliOuputs(v.([]interface{})); w != nil {
				deviceInfo.PreWorkflowCliOuputs = *w
			}
		}
		if v, ok := deviceInfoItem["project_id"]; ok {
			deviceInfo.ProjectID = v.(string)
		}
		if v, ok := deviceInfoItem["project_name"]; ok {
			deviceInfo.ProjectName = v.(string)
		}
		if v, ok := deviceInfoItem["reload_requested"]; ok {
			deviceInfo.ReloadRequested = v.(bool)
		}
		if v, ok := deviceInfoItem["serial_number"]; ok {
			deviceInfo.SerialNumber = v.(string)
		}
		if v, ok := deviceInfoItem["site_id"]; ok {
			deviceInfo.SiteID = v.(string)
		}
		if v, ok := deviceInfoItem["site_name"]; ok {
			deviceInfo.SiteName = v.(string)
		}
		if v, ok := deviceInfoItem["smart_account_id"]; ok {
			deviceInfo.SmartAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["source"]; ok {
			deviceInfo.Source = v.(string)
		}
		if v, ok := deviceInfoItem["stack"]; ok {
			deviceInfo.Stack = v.(bool)
		}
		if v, ok := deviceInfoItem["stack_info"]; ok {
			if w := constructAddDeviceInfoStackInfo(v.([]interface{})); w != nil {
				deviceInfo.StackInfo = *w
			}
		}
		if v, ok := deviceInfoItem["state"]; ok {
			deviceInfo.State = v.(string)
		}
		if v, ok := deviceInfoItem["sudi_required"]; ok {
			deviceInfo.SudiRequired = v.(bool)
		}
		if v, ok := deviceInfoItem["tags"]; ok {
			deviceInfo.Tags = v.(string)
		}
		if v, ok := deviceInfoItem["user_mic_numbers"]; ok {
			deviceInfo.UserMicNumbers = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["user_sudi_serial_nos"]; ok {
			deviceInfo.UserSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["virtual_account_id"]; ok {
			deviceInfo.VirtualAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_id"]; ok {
			deviceInfo.WorkflowID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_name"]; ok {
			deviceInfo.WorkflowName = v.(string)
		}
	}
	return &deviceInfo
}

func constructAddRunSummaryList(summaryList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestRunSummaryList {
	var runSummaryList []dnac.AddDeviceToPnpDatabaseRequestRunSummaryList
	for _, runSummary := range summaryList {
		var runSummaryItem dnac.AddDeviceToPnpDatabaseRequestRunSummaryList
		rs := runSummary.(map[string]interface{})
		if v, ok := rs["details"]; ok {
			runSummaryItem.Details = v.(string)
		}
		if v, ok := rs["error_flag"]; ok {
			runSummaryItem.ErrorFlag = v.(bool)
		}
		if v, ok := rs["history_task_info"]; ok {
			if w := constructAddRunSummaryListHistoryTaskInfo(v.([]interface{})); w != nil {
				runSummaryItem.HistoryTaskInfo = *w
			}
		}
		if v, ok := rs["timestamp"]; ok {
			runSummaryItem.Timestamp = v.(float64)
		}
		runSummaryList = append(runSummaryList, runSummaryItem)
	}
	return &runSummaryList
}

func constructAddSystemResetWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList {
	var result []dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructAddSystemWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList {
	var result []dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructAddWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList {
	var result []dnac.AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructAddSystemResetWorkflowTasks(wTasks []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks {
	var result []dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructAddSystemResetWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructAddSystemWorkflowTasks(wTasks []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasks {
	var result []dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.AddDeviceToPnpDatabaseRequestSystemWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructAddSystemWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructAddWorkflowTasks(wTasks []interface{}) *[]dnac.AddDeviceToPnpDatabaseRequestWorkflowTasks {
	var result []dnac.AddDeviceToPnpDatabaseRequestWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.AddDeviceToPnpDatabaseRequestWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructAddWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructAddSystemResetWorkflow(workflows []interface{}) *dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflow {
	var workflowItem dnac.AddDeviceToPnpDatabaseRequestSystemResetWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructAddSystemResetWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddSystemWorkflow(workflows []interface{}) *dnac.AddDeviceToPnpDatabaseRequestSystemWorkflow {
	var workflowItem dnac.AddDeviceToPnpDatabaseRequestSystemWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructAddSystemWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddWorkflow(workflows []interface{}) *dnac.AddDeviceToPnpDatabaseRequestWorkflow {
	var workflowItem dnac.AddDeviceToPnpDatabaseRequestWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructAddWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddWorkflowParams(params []interface{}) *dnac.AddDeviceToPnpDatabaseRequestWorkflowParameters {
	var workflowParameters dnac.AddDeviceToPnpDatabaseRequestWorkflowParameters
	for _, param := range params {
		ps := param.(map[string]interface{})

		if v, ok := ps["config_list"]; ok {
			configListV := v.([]interface{})
			var configList []dnac.AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList
			for _, ci := range configListV {
				var configItem dnac.AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList
				cis := ci.(map[string]interface{})
				if v, ok := cis["config_id"]; ok {
					configItem.ConfigID = v.(string)
				}
				if v, ok := cis["config_paramters"]; ok {
					configParameters := v.([]interface{})
					for _, configParam := range configParameters {
						cp := configParam.(map[string]interface{})
						configItem.ConfigParameters = append(configItem.ConfigParameters, dnac.AddDeviceToPnpDatabaseRequestWorkflowParametersConfigListConfigParameters{
							Key:   cp["key"].(string),
							Value: cp["value"].(string),
						})
					}
				}
				configList = append(configList, configItem)
			}
			workflowParameters.ConfigList = configList
		}
		if v, ok := ps["license_level"]; ok {
			workflowParameters.LicenseLevel = v.(string)
		}
		if v, ok := ps["license_type"]; ok {
			workflowParameters.LicenseType = v.(string)
		}
		if v, ok := ps["top_of_stack_serial_number"]; ok {
			workflowParameters.TopOfStackSerialNumber = v.(string)
		}
	}
	return &workflowParameters
}

func constructAddPnPDevice(pnpRequest map[string]interface{}) *dnac.AddDeviceToPnpDatabaseRequest {
	var request dnac.AddDeviceToPnpDatabaseRequest
	if v, ok := pnpRequest["id"]; ok {
		request.TypeID = v.(string)
	}
	if v, ok := pnpRequest["day_zero_config"]; ok {
		if w := constructAddZeroConfig(v.([]interface{})); w != nil {
			request.DayZeroConfig = *w
		}
	}
	if v, ok := pnpRequest["day_zero_config_preview"]; ok {
		request.DayZeroConfigPreview = v.(string)
	}
	if v, ok := pnpRequest["device_info"]; ok {
		if w := constructAddDeviceInfo(v.([]interface{})); w != nil {
			request.DeviceInfo = *w
		}
	}
	if v, ok := pnpRequest["run_summary_list"]; ok {
		if w := constructAddRunSummaryList(v.([]interface{})); w != nil {
			request.RunSummaryList = *w
		}
	}
	if v, ok := pnpRequest["system_reset_workflow"]; ok {
		if w := constructAddSystemResetWorkflow(v.([]interface{})); w != nil {
			request.SystemResetWorkflow = *w
		}
	}
	if v, ok := pnpRequest["system_workflow"]; ok {
		if w := constructAddSystemWorkflow(v.([]interface{})); w != nil {
			request.SystemWorkflow = *w
		}
	}
	if v, ok := pnpRequest["tenant_id"]; ok {
		request.TenantID = v.(string)
	}
	if v, ok := pnpRequest["version"]; ok {
		request.Version = v.(float64)
	}
	if v, ok := pnpRequest["workflow"]; ok {
		if w := constructAddWorkflow(v.([]interface{})); w != nil {
			request.Workflow = *w
		}
	}
	if v, ok := pnpRequest["workflow_parameters"]; ok {
		if w := constructAddWorkflowParams(v.([]interface{})); w != nil {
			request.WorkflowParameters = *w
		}
	}
	return &request
}

////// start update
func constructUpdateZeroConfig(zeroConfigs []interface{}) *dnac.UpdateDeviceRequestDayZeroConfig {
	var zeroConfig dnac.UpdateDeviceRequestDayZeroConfig
	if len(zeroConfigs) > 0 {
		zc := zeroConfigs[0].(map[string]interface{})
		zeroConfig.Config = zc["config"].(string)
	}
	return &zeroConfig
}

func constructUpdateDeviceInfoAAACredentials(credentials []interface{}) *dnac.UpdateDeviceRequestDeviceInfoAAACredentials {
	var aaaCredentials dnac.UpdateDeviceRequestDeviceInfoAAACredentials
	if len(credentials) > 0 {
		credential := credentials[0].(map[string]interface{})
		if v, ok := credential["username"]; ok {
			aaaCredentials.Username = v.(string)
		}
		if v, ok := credential["password"]; ok {
			aaaCredentials.Password = v.(string)
		}
	}
	return &aaaCredentials
}

func constructUpdateDeviceInfoFileSystemList(fileList []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoFileSystemList {
	var fileSystemList []dnac.UpdateDeviceRequestDeviceInfoFileSystemList
	for _, fi := range fileList {
		fsi := fi.(map[string]interface{})
		fileSystemItem := dnac.UpdateDeviceRequestDeviceInfoFileSystemList{}
		if v, ok := fsi["freespace"]; ok {
			fileSystemItem.Freespace = v.(float64)
		}
		if v, ok := fsi["name"]; ok {
			fileSystemItem.Name = v.(string)
		}
		if v, ok := fsi["readable"]; ok {
			fileSystemItem.Readable = v.(bool)
		}
		if v, ok := fsi["size"]; ok {
			fileSystemItem.Size = v.(float64)
		}
		if v, ok := fsi["type"]; ok {
			fileSystemItem.Type = v.(string)
		}
		if v, ok := fsi["writeable"]; ok {
			fileSystemItem.Writeable = v.(bool)
		}
		fileSystemList = append(fileSystemList, fileSystemItem)
	}
	return &fileSystemList
}

func constructUpdateDeviceInfoHTTPHeaders(headers []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoHTTPHeaders {
	var result []dnac.UpdateDeviceRequestDeviceInfoHTTPHeaders
	for _, header := range headers {
		hs := header.(map[string]interface{})
		result = append(result, dnac.UpdateDeviceRequestDeviceInfoHTTPHeaders{
			Key:   hs["key"].(string),
			Value: hs["value"].(string),
		})
	}
	return &result
}

func constructUpdateDeviceInfoIPInterfaces(interfaces []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoIPInterfaces {
	var result []dnac.UpdateDeviceRequestDeviceInfoIPInterfaces
	for _, i := range interfaces {
		interfaceItem := dnac.UpdateDeviceRequestDeviceInfoIPInterfaces{}
		is := i.(map[string]interface{})
		if v, ok := is["ipv4_address"]; ok {
			interfaceItem.IPv4Address = v.(string)
		}
		if v, ok := is["ipv6_address_list"]; ok {
			interfaceItem.IPv6AddressList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := is["mac_address"]; ok {
			interfaceItem.MacAddress = v.(string)
		}
		if v, ok := is["name"]; ok {
			interfaceItem.Name = v.(string)
		}
		if v, ok := is["status"]; ok {
			interfaceItem.Status = v.(string)
		}
		result = append(result, interfaceItem)
	}
	return &result
}

func constructUpdateDeviceInfoLocation(locations []interface{}) *dnac.UpdateDeviceRequestDeviceInfoLocation {
	var result dnac.UpdateDeviceRequestDeviceInfoLocation
	if len(locations) > 0 {
		location := locations[0].(map[string]interface{})
		if v, ok := location["address"]; ok {
			result.Address = v.(string)
		}
		if v, ok := location["altitude"]; ok {
			result.Altitude = v.(string)
		}
		if v, ok := location["latitude"]; ok {
			result.Latitude = v.(string)
		}
		if v, ok := location["longitude"]; ok {
			result.Longitude = v.(string)
		}
		if v, ok := location["siteId"]; ok {
			result.SiteID = v.(string)
		}
	}
	return &result
}

func constructUpdateDeviceInfoNeighborLinks(links []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoNeighborLinks {
	var result []dnac.UpdateDeviceRequestDeviceInfoNeighborLinks
	for _, link := range links {
		var linkItem dnac.UpdateDeviceRequestDeviceInfoNeighborLinks
		ls := link.(map[string]interface{})
		if v, ok := ls["local_interface_name"]; ok {
			linkItem.LocalInterfaceName = v.(string)
		}
		if v, ok := ls["local_mac_address"]; ok {
			linkItem.LocalMacAddress = v.(string)
		}
		if v, ok := ls["local_short_interface_name"]; ok {
			linkItem.LocalShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_device_name"]; ok {
			linkItem.RemoteDeviceName = v.(string)
		}
		if v, ok := ls["remote_interface_name"]; ok {
			linkItem.RemoteInterfaceName = v.(string)
		}
		if v, ok := ls["remote_mac_address"]; ok {
			linkItem.RemoteMacAddress = v.(string)
		}
		if v, ok := ls["remote_platform"]; ok {
			linkItem.RemotePlatform = v.(string)
		}
		if v, ok := ls["remote_short_interface_name"]; ok {
			linkItem.RemoteShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_version"]; ok {
			linkItem.RemoteVersion = v.(string)
		}
		result = append(result, linkItem)
	}
	return &result
}

func constructUpdateDeviceInfoPnpProfileList(profiles []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoPnpProfileList {
	var result []dnac.UpdateDeviceRequestDeviceInfoPnpProfileList
	for _, profile := range profiles {
		var profileItem dnac.UpdateDeviceRequestDeviceInfoPnpProfileList
		ps := profile.(map[string]interface{})
		if v, ok := ps["created_by"]; ok {
			profileItem.CreatedBy = v.(string)
		}
		if v, ok := ps["discovery_created"]; ok {
			profileItem.DiscoveryCreated = v.(bool)
		}
		if v, ok := ps["primary_endpoint"]; ok {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok {
					profileItem.PrimaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok {
					profileItem.PrimaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok {
					profileItem.PrimaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok {
					profileItem.PrimaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok {
					profileItem.PrimaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok {
					profileItem.PrimaryEndpoint.Protocol = x.(string)
				}
			}
		}
		if v, ok := ps["profile_name"]; ok {
			profileItem.ProfileName = v.(string)
		}
		if v, ok := ps["secondary_endpoint"]; ok {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok {
					profileItem.SecondaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok {
					profileItem.SecondaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok {
					profileItem.SecondaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok {
					profileItem.SecondaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok {
					profileItem.SecondaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok {
					profileItem.SecondaryEndpoint.Protocol = x.(string)
				}
			}
		}
		result = append(result, profileItem)
	}
	return &result
}

func constructUpdateDeviceInfoPreWorkflowCliOuputs(cliOutputs []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs {
	var result []dnac.UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs
	for _, cliOutput := range cliOutputs {
		var cliOutputItem dnac.UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs
		cos := cliOutput.(map[string]interface{})
		if v, ok := cos["cli"]; ok {
			cliOutputItem.Cli = v.(string)
		}
		if v, ok := cos["cli_output"]; ok {
			cliOutputItem.CliOutput = v.(string)
		}
		result = append(result, cliOutputItem)
	}
	return &result
}

func constructUpdateDeviceInfoStackInfoStackMemberList(memberList []interface{}) *[]dnac.UpdateDeviceRequestDeviceInfoStackInfoStackMemberList {
	var result []dnac.UpdateDeviceRequestDeviceInfoStackInfoStackMemberList
	for _, member := range memberList {
		var memberItem dnac.UpdateDeviceRequestDeviceInfoStackInfoStackMemberList
		ms := member.(map[string]interface{})
		if v, ok := ms["hardware_version"]; ok {
			memberItem.HardwareVersion = v.(string)
		}
		if v, ok := ms["license_level"]; ok {
			memberItem.LicenseLevel = v.(string)
		}
		if v, ok := ms["license_type"]; ok {
			memberItem.LicenseType = v.(string)
		}
		if v, ok := ms["mac_address"]; ok {
			memberItem.MacAddress = v.(string)
		}
		if v, ok := ms["pid"]; ok {
			memberItem.Pid = v.(string)
		}
		if v, ok := ms["priority"]; ok {
			memberItem.Priority = v.(float64)
		}
		if v, ok := ms["role"]; ok {
			memberItem.Role = v.(string)
		}
		if v, ok := ms["serial_number"]; ok {
			memberItem.SerialNumber = v.(string)
		}
		if v, ok := ms["software_version"]; ok {
			memberItem.SoftwareVersion = v.(string)
		}
		if v, ok := ms["stack_number"]; ok {
			memberItem.StackNumber = v.(float64)
		}
		if v, ok := ms["state"]; ok {
			memberItem.State = v.(string)
		}
		if v, ok := ms["sudi_serial_number"]; ok {
			memberItem.SudiSerialNumber = v.(string)
		}
		result = append(result, memberItem)
	}
	return &result
}

func constructUpdateDeviceInfoStackInfo(stackInfos []interface{}) *dnac.UpdateDeviceRequestDeviceInfoStackInfo {
	var result dnac.UpdateDeviceRequestDeviceInfoStackInfo
	if len(stackInfos) > 0 {
		stackInfo := stackInfos[0].(map[string]interface{})
		if v, ok := stackInfo["is_full_ring"]; ok {
			result.IsFullRing = v.(bool)
		}
		if v, ok := stackInfo["stack_member_list"]; ok {
			if w := constructUpdateDeviceInfoStackInfoStackMemberList(v.([]interface{})); w != nil {
				result.StackMemberList = *w
			}
		}
		if v, ok := stackInfo["stack_ring_protocol"]; ok {
			result.StackRingProtocol = v.(string)
		}
		if v, ok := stackInfo["supports_stack_workflows"]; ok {
			result.SupportsStackWorkflows = v.(bool)
		}
		if v, ok := stackInfo["total_member_count"]; ok {
			result.TotalMemberCount = v.(float64)
		}
		if v, ok := stackInfo["valid_license_levels"]; ok {
			result.ValidLicenseLevels = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
	}

	return &result
}

func constructUpdateRunSummaryListHistoryTaskInfoAddnDetails(details []interface{}) *[]dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails {
	var result []dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails
	for _, detail := range details {
		hs := detail.(map[string]interface{})
		result = append(result, dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails{
			Key:   hs["key"].(string),
			Value: hs["value"].(string),
		})
	}
	return &result
}

func constructUpdateRunSummaryListHistoryTaskInfoWorkItemList(itemList []interface{}) *[]dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList {
	var result []dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructUpdateRunSummaryListHistoryTaskInfo(taskInfos []interface{}) *dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfo {
	var result dnac.UpdateDeviceRequestRunSummaryListHistoryTaskInfo
	if len(taskInfos) > 0 {
		taskInfo := taskInfos[0].(map[string]interface{})
		if v, ok := taskInfo["addn_details"]; ok {
			if w := constructUpdateRunSummaryListHistoryTaskInfoAddnDetails(v.([]interface{})); w != nil {
				result.AddnDetails = *w
			}
		}
		if v, ok := taskInfo["name"]; ok {
			result.Name = v.(string)
		}
		if v, ok := taskInfo["time_taken"]; ok {
			result.TimeTaken = v.(float64)
		}
		if v, ok := taskInfo["type"]; ok {
			result.Type = v.(string)
		}
		if v, ok := taskInfo["work_item_list"]; ok {
			if w := constructUpdateRunSummaryListHistoryTaskInfoWorkItemList(v.([]interface{})); w != nil {
				result.WorkItemList = *w
			}
		}
	}
	return &result
}

func constructUpdateDeviceInfo(deviceInfos []interface{}) *dnac.UpdateDeviceRequestDeviceInfo {
	var deviceInfo dnac.UpdateDeviceRequestDeviceInfo
	if len(deviceInfos) > 0 {
		deviceInfoItem := deviceInfos[0].(map[string]interface{})
		if v, ok := deviceInfoItem["aaa_credentials"]; ok {
			if w := constructUpdateDeviceInfoAAACredentials(v.([]interface{})); w != nil {
				deviceInfo.AAACredentials = *w
			}
		}
		if v, ok := deviceInfoItem["added_on"]; ok {
			deviceInfo.AddedOn = v.(float64)
		}
		if v, ok := deviceInfoItem["addn_mac_addrs"]; ok {
			deviceInfo.AddnMacAddrs = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["agent_type"]; ok {
			deviceInfo.AgentType = v.(string)
		}
		if v, ok := deviceInfoItem["auth_status"]; ok {
			deviceInfo.AuthStatus = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_mic_number"]; ok {
			deviceInfo.AuthenticatedMicNumber = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_sudi_serial_no"]; ok {
			deviceInfo.AuthenticatedSudiSerialNo = v.(string)
		}
		if v, ok := deviceInfoItem["capabilities_supported"]; ok {
			deviceInfo.CapabilitiesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["cm_state"]; ok {
			deviceInfo.CmState = v.(string)
		}
		if v, ok := deviceInfoItem["description"]; ok {
			deviceInfo.Description = v.(string)
		}
		if v, ok := deviceInfoItem["device_sudi_serial_nos"]; ok {
			deviceInfo.DeviceSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["device_type"]; ok {
			deviceInfo.DeviceType = v.(string)
		}
		if v, ok := deviceInfoItem["features_supported"]; ok {
			deviceInfo.FeaturesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["file_system_list"]; ok {
			if w := constructUpdateDeviceInfoFileSystemList(v.([]interface{})); w != nil {
				deviceInfo.FileSystemList = *w
			}
		}
		if v, ok := deviceInfoItem["first_contact"]; ok {
			deviceInfo.FirstContact = v.(float64)
		}
		if v, ok := deviceInfoItem["hostname"]; ok {
			deviceInfo.Hostname = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok {
			if w := constructUpdateDeviceInfoHTTPHeaders(v.([]interface{})); w != nil {
				deviceInfo.HTTPHeaders = *w
			}
		}
		if v, ok := deviceInfoItem["image_file"]; ok {
			deviceInfo.ImageFile = v.(string)
		}
		if v, ok := deviceInfoItem["image_version"]; ok {
			deviceInfo.ImageVersion = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok {
			if w := constructUpdateDeviceInfoIPInterfaces(v.([]interface{})); w != nil {
				deviceInfo.IPInterfaces = *w
			}
		}
		if v, ok := deviceInfoItem["last_contact"]; ok {
			deviceInfo.LastContact = v.(float64)
		}
		if v, ok := deviceInfoItem["last_sync_time"]; ok {
			deviceInfo.LastSyncTime = v.(float64)
		}
		if v, ok := deviceInfoItem["last_update_on"]; ok {
			deviceInfo.LastUpdateOn = v.(float64)
		}
		if v, ok := deviceInfoItem["location"]; ok {
			if w := constructUpdateDeviceInfoLocation(v.([]interface{})); w != nil {
				deviceInfo.Location = *w
			}
		}
		if v, ok := deviceInfoItem["mac_address"]; ok {
			deviceInfo.MacAddress = v.(string)
		}
		if v, ok := deviceInfoItem["mode"]; ok {
			deviceInfo.Mode = v.(string)
		}
		if v, ok := deviceInfoItem["name"]; ok {
			deviceInfo.Name = v.(string)
		}
		if v, ok := deviceInfoItem["neighbor_links"]; ok {
			if w := constructUpdateDeviceInfoNeighborLinks(v.([]interface{})); w != nil {
				deviceInfo.NeighborLinks = *w
			}
		}
		if v, ok := deviceInfoItem["onb_state"]; ok {
			deviceInfo.OnbState = v.(string)
		}
		if v, ok := deviceInfoItem["pid"]; ok {
			deviceInfo.Pid = v.(string)
		}
		if v, ok := deviceInfoItem["pnp_profile_list"]; ok {
			if w := constructUpdateDeviceInfoPnpProfileList(v.([]interface{})); w != nil {
				deviceInfo.PnpProfileList = *w
			}
		}
		if v, ok := deviceInfoItem["populate_inventory"]; ok {
			deviceInfo.PopulateInventory = v.(bool)
		}
		if v, ok := deviceInfoItem["pre_workflow_cli_ouputs"]; ok {
			if w := constructUpdateDeviceInfoPreWorkflowCliOuputs(v.([]interface{})); w != nil {
				deviceInfo.PreWorkflowCliOuputs = *w
			}
		}
		if v, ok := deviceInfoItem["project_id"]; ok {
			deviceInfo.ProjectID = v.(string)
		}
		if v, ok := deviceInfoItem["project_name"]; ok {
			deviceInfo.ProjectName = v.(string)
		}
		if v, ok := deviceInfoItem["reload_requested"]; ok {
			deviceInfo.ReloadRequested = v.(bool)
		}
		if v, ok := deviceInfoItem["serial_number"]; ok {
			deviceInfo.SerialNumber = v.(string)
		}
		if v, ok := deviceInfoItem["site_id"]; ok {
			deviceInfo.SiteID = v.(string)
		}
		if v, ok := deviceInfoItem["site_name"]; ok {
			deviceInfo.SiteName = v.(string)
		}
		if v, ok := deviceInfoItem["smart_account_id"]; ok {
			deviceInfo.SmartAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["source"]; ok {
			deviceInfo.Source = v.(string)
		}
		if v, ok := deviceInfoItem["stack"]; ok {
			deviceInfo.Stack = v.(bool)
		}
		if v, ok := deviceInfoItem["stack_info"]; ok {
			if w := constructUpdateDeviceInfoStackInfo(v.([]interface{})); w != nil {
				deviceInfo.StackInfo = *w
			}
		}
		if v, ok := deviceInfoItem["state"]; ok {
			deviceInfo.State = v.(string)
		}
		if v, ok := deviceInfoItem["sudi_required"]; ok {
			deviceInfo.SudiRequired = v.(bool)
		}
		if v, ok := deviceInfoItem["tags"]; ok {
			deviceInfo.Tags = v.(string)
		}
		if v, ok := deviceInfoItem["user_mic_numbers"]; ok {
			deviceInfo.UserMicNumbers = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["user_sudi_serial_nos"]; ok {
			deviceInfo.UserSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["virtual_account_id"]; ok {
			deviceInfo.VirtualAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_id"]; ok {
			deviceInfo.WorkflowID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_name"]; ok {
			deviceInfo.WorkflowName = v.(string)
		}
	}
	return &deviceInfo
}

func constructUpdateRunSummaryList(summaryList []interface{}) *[]dnac.UpdateDeviceRequestRunSummaryList {
	var runSummaryList []dnac.UpdateDeviceRequestRunSummaryList
	for _, runSummary := range summaryList {
		var runSummaryItem dnac.UpdateDeviceRequestRunSummaryList
		rs := runSummary.(map[string]interface{})
		if v, ok := rs["details"]; ok {
			runSummaryItem.Details = v.(string)
		}
		if v, ok := rs["error_flag"]; ok {
			runSummaryItem.ErrorFlag = v.(bool)
		}
		if v, ok := rs["history_task_info"]; ok {
			if w := constructUpdateRunSummaryListHistoryTaskInfo(v.([]interface{})); w != nil {
				runSummaryItem.HistoryTaskInfo = *w
			}
		}
		if v, ok := rs["timestamp"]; ok {
			runSummaryItem.Timestamp = v.(float64)
		}
		runSummaryList = append(runSummaryList, runSummaryItem)
	}
	return &runSummaryList
}

func constructUpdateSystemResetWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList {
	var result []dnac.UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructUpdateSystemWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.UpdateDeviceRequestSystemWorkflowTasksWorkItemList {
	var result []dnac.UpdateDeviceRequestSystemWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.UpdateDeviceRequestSystemWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructUpdateWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.UpdateDeviceRequestWorkflowTasksWorkItemList {
	var result []dnac.UpdateDeviceRequestWorkflowTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.UpdateDeviceRequestWorkflowTasksWorkItemList
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructUpdateSystemResetWorkflowTasks(wTasks []interface{}) *[]dnac.UpdateDeviceRequestSystemResetWorkflowTasks {
	var result []dnac.UpdateDeviceRequestSystemResetWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.UpdateDeviceRequestSystemResetWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructUpdateSystemResetWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructUpdateSystemWorkflowTasks(wTasks []interface{}) *[]dnac.UpdateDeviceRequestSystemWorkflowTasks {
	var result []dnac.UpdateDeviceRequestSystemWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.UpdateDeviceRequestSystemWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructUpdateSystemWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructUpdateWorkflowTasks(wTasks []interface{}) *[]dnac.UpdateDeviceRequestWorkflowTasks {
	var result []dnac.UpdateDeviceRequestWorkflowTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.UpdateDeviceRequestWorkflowTasks
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructUpdateWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructUpdateSystemResetWorkflow(workflows []interface{}) *dnac.UpdateDeviceRequestSystemResetWorkflow {
	var workflowItem dnac.UpdateDeviceRequestSystemResetWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructUpdateSystemResetWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateSystemWorkflow(workflows []interface{}) *dnac.UpdateDeviceRequestSystemWorkflow {
	var workflowItem dnac.UpdateDeviceRequestSystemWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructUpdateSystemWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateWorkflow(workflows []interface{}) *dnac.UpdateDeviceRequestWorkflow {
	var workflowItem dnac.UpdateDeviceRequestWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok {
			if w := constructUpdateWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateWorkflowParams(params []interface{}) *dnac.UpdateDeviceRequestWorkflowParameters {
	var workflowParameters dnac.UpdateDeviceRequestWorkflowParameters
	for _, param := range params {
		ps := param.(map[string]interface{})

		if v, ok := ps["config_list"]; ok {
			configListV := v.([]interface{})
			var configList []dnac.UpdateDeviceRequestWorkflowParametersConfigList
			for _, ci := range configListV {
				var configItem dnac.UpdateDeviceRequestWorkflowParametersConfigList
				cis := ci.(map[string]interface{})
				if v, ok := cis["config_id"]; ok {
					configItem.ConfigID = v.(string)
				}
				if v, ok := cis["config_paramters"]; ok {
					configParameters := v.([]interface{})
					for _, configParam := range configParameters {
						cp := configParam.(map[string]interface{})
						configItem.ConfigParameters = append(configItem.ConfigParameters, dnac.UpdateDeviceRequestWorkflowParametersConfigListConfigParameters{
							Key:   cp["key"].(string),
							Value: cp["value"].(string),
						})
					}
				}
				configList = append(configList, configItem)
			}
			workflowParameters.ConfigList = configList
		}
		if v, ok := ps["license_level"]; ok {
			workflowParameters.LicenseLevel = v.(string)
		}
		if v, ok := ps["license_type"]; ok {
			workflowParameters.LicenseType = v.(string)
		}
		if v, ok := ps["top_of_stack_serial_number"]; ok {
			workflowParameters.TopOfStackSerialNumber = v.(string)
		}
	}
	return &workflowParameters
}

func constructUpdatePnPDevice(pnpRequest map[string]interface{}) *dnac.UpdateDeviceRequest {
	var request dnac.UpdateDeviceRequest
	if v, ok := pnpRequest["id"]; ok {
		request.TypeID = v.(string)
	}
	if v, ok := pnpRequest["day_zero_config"]; ok {
		if w := constructUpdateZeroConfig(v.([]interface{})); w != nil {
			request.DayZeroConfig = *w
		}
	}
	if v, ok := pnpRequest["day_zero_config_preview"]; ok {
		request.DayZeroConfigPreview = v.(string)
	}
	if v, ok := pnpRequest["device_info"]; ok {
		if w := constructUpdateDeviceInfo(v.([]interface{})); w != nil {
			request.DeviceInfo = *w
		}
	}
	if v, ok := pnpRequest["run_summary_list"]; ok {
		if w := constructUpdateRunSummaryList(v.([]interface{})); w != nil {
			request.RunSummaryList = *w
		}
	}
	if v, ok := pnpRequest["system_reset_workflow"]; ok {
		if w := constructUpdateSystemResetWorkflow(v.([]interface{})); w != nil {
			request.SystemResetWorkflow = *w
		}
	}
	if v, ok := pnpRequest["system_workflow"]; ok {
		if w := constructUpdateSystemWorkflow(v.([]interface{})); w != nil {
			request.SystemWorkflow = *w
		}
	}
	if v, ok := pnpRequest["tenant_id"]; ok {
		request.TenantID = v.(string)
	}
	if v, ok := pnpRequest["version"]; ok {
		request.Version = v.(float64)
	}
	if v, ok := pnpRequest["workflow"]; ok {
		if w := constructUpdateWorkflow(v.([]interface{})); w != nil {
			request.Workflow = *w
		}
	}
	if v, ok := pnpRequest["workflow_parameters"]; ok {
		if w := constructUpdateWorkflowParams(v.([]interface{})); w != nil {
			request.WorkflowParameters = *w
		}
	}
	return &request
}

////// end update

func resourcePnPDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	pnpRequest := item.(map[string]interface{})

	request := constructAddPnPDevice(pnpRequest)
	response, _, err := client.DeviceOnboardingPnP.AddDeviceToPnpDatabase(request)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update resource id
	d.SetId(response.TypeID)
	// Update resource on Terraform
	resourcePnPDeviceRead(ctx, d, m)
	return diags
}

func resourcePnPDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	deviceID := d.Id()
	response, _, err := client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if response == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	deviceItem := flattenPnPDeviceReadItem(response)
	if err := d.Set("item", deviceItem); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePnPDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		deviceID := d.Id()
		item := d.Get("item").([]interface{})[0]
		pnpRequest := item.(map[string]interface{})
		request := constructUpdatePnPDevice(pnpRequest)
		_, _, err := client.DeviceOnboardingPnP.UpdateDevice(deviceID, request)
		if err != nil {
			return diag.FromErr(err)
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourcePnPDeviceRead(ctx, d, m)
}

func resourcePnPDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deviceID := d.Id()
	// Call function to delete application resource
	_, _, err := client.DeviceOnboardingPnP.DeleteDeviceByIDFromPnP(deviceID)
	if err != nil {
		return diag.FromErr(err)
	}

	response, _, err := client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil || response == nil {
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to delete PnP device",
		Detail:   "",
	})

	return diags
}
