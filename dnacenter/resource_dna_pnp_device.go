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
							Type:     schema.TypeList,
							MaxItems: 1,
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
		if v, ok := credential["username"]; ok && v != nil {
			aaaCredentials.Username = v.(string)
		}
		if v, ok := credential["password"]; ok && v != nil {
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
		if v, ok := fsi["freespace"]; ok && v != nil {
			fileSystemItem.Freespace = v.(float64)
		}
		if v, ok := fsi["name"]; ok && v != nil {
			fileSystemItem.Name = v.(string)
		}
		if v, ok := fsi["readable"]; ok && v != nil {
			fileSystemItem.Readable = v.(bool)
		}
		if v, ok := fsi["size"]; ok && v != nil {
			fileSystemItem.Size = v.(float64)
		}
		if v, ok := fsi["type"]; ok && v != nil {
			fileSystemItem.Type = v.(string)
		}
		if v, ok := fsi["writeable"]; ok && v != nil {
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
		if v, ok := is["ipv4_address"]; ok && v != nil {
			interfaceItem.IPv4Address = v.(string)
		}
		if v, ok := is["ipv6_address_list"]; ok && v != nil {
			interfaceItem.IPv6AddressList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := is["mac_address"]; ok && v != nil {
			interfaceItem.MacAddress = v.(string)
		}
		if v, ok := is["name"]; ok && v != nil {
			interfaceItem.Name = v.(string)
		}
		if v, ok := is["status"]; ok && v != nil {
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
		if v, ok := location["address"]; ok && v != nil {
			result.Address = v.(string)
		}
		if v, ok := location["altitude"]; ok && v != nil {
			result.Altitude = v.(string)
		}
		if v, ok := location["latitude"]; ok && v != nil {
			result.Latitude = v.(string)
		}
		if v, ok := location["longitude"]; ok && v != nil {
			result.Longitude = v.(string)
		}
		if v, ok := location["siteId"]; ok && v != nil {
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
		if v, ok := ls["local_interface_name"]; ok && v != nil {
			linkItem.LocalInterfaceName = v.(string)
		}
		if v, ok := ls["local_mac_address"]; ok && v != nil {
			linkItem.LocalMacAddress = v.(string)
		}
		if v, ok := ls["local_short_interface_name"]; ok && v != nil {
			linkItem.LocalShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_device_name"]; ok && v != nil {
			linkItem.RemoteDeviceName = v.(string)
		}
		if v, ok := ls["remote_interface_name"]; ok && v != nil {
			linkItem.RemoteInterfaceName = v.(string)
		}
		if v, ok := ls["remote_mac_address"]; ok && v != nil {
			linkItem.RemoteMacAddress = v.(string)
		}
		if v, ok := ls["remote_platform"]; ok && v != nil {
			linkItem.RemotePlatform = v.(string)
		}
		if v, ok := ls["remote_short_interface_name"]; ok && v != nil {
			linkItem.RemoteShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_version"]; ok && v != nil {
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
		if v, ok := ps["created_by"]; ok && v != nil {
			profileItem.CreatedBy = v.(string)
		}
		if v, ok := ps["discovery_created"]; ok && v != nil {
			profileItem.DiscoveryCreated = v.(bool)
		}
		if v, ok := ps["primary_endpoint"]; ok && v != nil {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok && x != nil {
					profileItem.PrimaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok && x != nil {
					profileItem.PrimaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Protocol = x.(string)
				}
			}
		}
		if v, ok := ps["profile_name"]; ok && v != nil {
			profileItem.ProfileName = v.(string)
		}
		if v, ok := ps["secondary_endpoint"]; ok && v != nil {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok && x != nil {
					profileItem.SecondaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok && x != nil {
					profileItem.SecondaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok && x != nil {
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
		if v, ok := cos["cli"]; ok && v != nil {
			cliOutputItem.Cli = v.(string)
		}
		if v, ok := cos["cli_output"]; ok && v != nil {
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
		if v, ok := ms["hardware_version"]; ok && v != nil {
			memberItem.HardwareVersion = v.(string)
		}
		if v, ok := ms["license_level"]; ok && v != nil {
			memberItem.LicenseLevel = v.(string)
		}
		if v, ok := ms["license_type"]; ok && v != nil {
			memberItem.LicenseType = v.(string)
		}
		if v, ok := ms["mac_address"]; ok && v != nil {
			memberItem.MacAddress = v.(string)
		}
		if v, ok := ms["pid"]; ok && v != nil {
			memberItem.Pid = v.(string)
		}
		if v, ok := ms["priority"]; ok && v != nil {
			memberItem.Priority = v.(float64)
		}
		if v, ok := ms["role"]; ok && v != nil {
			memberItem.Role = v.(string)
		}
		if v, ok := ms["serial_number"]; ok && v != nil {
			memberItem.SerialNumber = v.(string)
		}
		if v, ok := ms["software_version"]; ok && v != nil {
			memberItem.SoftwareVersion = v.(string)
		}
		if v, ok := ms["stack_number"]; ok && v != nil {
			memberItem.StackNumber = v.(float64)
		}
		if v, ok := ms["state"]; ok && v != nil {
			memberItem.State = v.(string)
		}
		if v, ok := ms["sudi_serial_number"]; ok && v != nil {
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
		if v, ok := stackInfo["is_full_ring"]; ok && v != nil {
			result.IsFullRing = v.(bool)
		}
		if v, ok := stackInfo["stack_member_list"]; ok && v != nil {
			if w := constructAddDeviceInfoStackInfoStackMemberList(v.([]interface{})); w != nil {
				result.StackMemberList = *w
			}
		}
		if v, ok := stackInfo["stack_ring_protocol"]; ok && v != nil {
			result.StackRingProtocol = v.(string)
		}
		if v, ok := stackInfo["supports_stack_workflows"]; ok && v != nil {
			result.SupportsStackWorkflows = v.(bool)
		}
		if v, ok := stackInfo["total_member_count"]; ok && v != nil {
			result.TotalMemberCount = v.(float64)
		}
		if v, ok := stackInfo["valid_license_levels"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := taskInfo["addn_details"]; ok && v != nil {
			if w := constructAddRunSummaryListHistoryTaskInfoAddnDetails(v.([]interface{})); w != nil {
				result.AddnDetails = *w
			}
		}
		if v, ok := taskInfo["name"]; ok && v != nil {
			result.Name = v.(string)
		}
		if v, ok := taskInfo["time_taken"]; ok && v != nil {
			result.TimeTaken = v.(float64)
		}
		if v, ok := taskInfo["type"]; ok && v != nil {
			result.Type = v.(string)
		}
		if v, ok := taskInfo["work_item_list"]; ok && v != nil {
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
		if v, ok := deviceInfoItem["aaa_credentials"]; ok && v != nil {
			if w := constructAddDeviceInfoAAACredentials(v.([]interface{})); w != nil {
				deviceInfo.AAACredentials = *w
			}
		}
		if v, ok := deviceInfoItem["added_on"]; ok && v != nil {
			deviceInfo.AddedOn = v.(float64)
		}
		if v, ok := deviceInfoItem["addn_mac_addrs"]; ok && v != nil {
			deviceInfo.AddnMacAddrs = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["agent_type"]; ok && v != nil {
			deviceInfo.AgentType = v.(string)
		}
		if v, ok := deviceInfoItem["auth_status"]; ok && v != nil {
			deviceInfo.AuthStatus = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_mic_number"]; ok && v != nil {
			deviceInfo.AuthenticatedMicNumber = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_sudi_serial_no"]; ok && v != nil {
			deviceInfo.AuthenticatedSudiSerialNo = v.(string)
		}
		if v, ok := deviceInfoItem["capabilities_supported"]; ok && v != nil {
			deviceInfo.CapabilitiesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["cm_state"]; ok && v != nil {
			deviceInfo.CmState = v.(string)
		}
		if v, ok := deviceInfoItem["description"]; ok && v != nil {
			deviceInfo.Description = v.(string)
		}
		if v, ok := deviceInfoItem["device_sudi_serial_nos"]; ok && v != nil {
			deviceInfo.DeviceSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["device_type"]; ok && v != nil {
			deviceInfo.DeviceType = v.(string)
		}
		if v, ok := deviceInfoItem["features_supported"]; ok && v != nil {
			deviceInfo.FeaturesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["file_system_list"]; ok && v != nil {
			if w := constructAddDeviceInfoFileSystemList(v.([]interface{})); w != nil {
				deviceInfo.FileSystemList = *w
			}
		}
		if v, ok := deviceInfoItem["first_contact"]; ok && v != nil {
			deviceInfo.FirstContact = v.(float64)
		}
		if v, ok := deviceInfoItem["hostname"]; ok && v != nil {
			deviceInfo.Hostname = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok && v != nil {
			if w := constructAddDeviceInfoHTTPHeaders(v.([]interface{})); w != nil {
				deviceInfo.HTTPHeaders = *w
			}
		}
		if v, ok := deviceInfoItem["image_file"]; ok && v != nil {
			deviceInfo.ImageFile = v.(string)
		}
		if v, ok := deviceInfoItem["image_version"]; ok && v != nil {
			deviceInfo.ImageVersion = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok && v != nil {
			if w := constructAddDeviceInfoIPInterfaces(v.([]interface{})); w != nil {
				deviceInfo.IPInterfaces = *w
			}
		}
		if v, ok := deviceInfoItem["last_contact"]; ok && v != nil {
			deviceInfo.LastContact = v.(float64)
		}
		if v, ok := deviceInfoItem["last_sync_time"]; ok && v != nil {
			deviceInfo.LastSyncTime = v.(float64)
		}
		if v, ok := deviceInfoItem["last_update_on"]; ok && v != nil {
			deviceInfo.LastUpdateOn = v.(float64)
		}
		if v, ok := deviceInfoItem["location"]; ok && v != nil {
			if w := constructAddDeviceInfoLocation(v.([]interface{})); w != nil {
				deviceInfo.Location = *w
			}
		}
		if v, ok := deviceInfoItem["mac_address"]; ok && v != nil {
			deviceInfo.MacAddress = v.(string)
		}
		if v, ok := deviceInfoItem["mode"]; ok && v != nil {
			deviceInfo.Mode = v.(string)
		}
		if v, ok := deviceInfoItem["name"]; ok && v != nil {
			deviceInfo.Name = v.(string)
		}
		if v, ok := deviceInfoItem["neighbor_links"]; ok && v != nil {
			if w := constructAddDeviceInfoNeighborLinks(v.([]interface{})); w != nil {
				deviceInfo.NeighborLinks = *w
			}
		}
		if v, ok := deviceInfoItem["onb_state"]; ok && v != nil {
			deviceInfo.OnbState = v.(string)
		}
		if v, ok := deviceInfoItem["pid"]; ok && v != nil {
			deviceInfo.Pid = v.(string)
		}
		if v, ok := deviceInfoItem["pnp_profile_list"]; ok && v != nil {
			if w := constructAddDeviceInfoPnpProfileList(v.([]interface{})); w != nil {
				deviceInfo.PnpProfileList = *w
			}
		}
		if v, ok := deviceInfoItem["populate_inventory"]; ok && v != nil {
			deviceInfo.PopulateInventory = v.(bool)
		}
		if v, ok := deviceInfoItem["pre_workflow_cli_ouputs"]; ok && v != nil {
			if w := constructAddDeviceInfoPreWorkflowCliOuputs(v.([]interface{})); w != nil {
				deviceInfo.PreWorkflowCliOuputs = *w
			}
		}
		if v, ok := deviceInfoItem["project_id"]; ok && v != nil {
			deviceInfo.ProjectID = v.(string)
		}
		if v, ok := deviceInfoItem["project_name"]; ok && v != nil {
			deviceInfo.ProjectName = v.(string)
		}
		if v, ok := deviceInfoItem["reload_requested"]; ok && v != nil {
			deviceInfo.ReloadRequested = v.(bool)
		}
		if v, ok := deviceInfoItem["serial_number"]; ok && v != nil {
			deviceInfo.SerialNumber = v.(string)
		}
		if v, ok := deviceInfoItem["site_id"]; ok && v != nil {
			deviceInfo.SiteID = v.(string)
		}
		if v, ok := deviceInfoItem["site_name"]; ok && v != nil {
			deviceInfo.SiteName = v.(string)
		}
		if v, ok := deviceInfoItem["smart_account_id"]; ok && v != nil {
			deviceInfo.SmartAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["source"]; ok && v != nil {
			deviceInfo.Source = v.(string)
		}
		if v, ok := deviceInfoItem["stack"]; ok && v != nil {
			deviceInfo.Stack = v.(bool)
		}
		if v, ok := deviceInfoItem["stack_info"]; ok && v != nil {
			if w := constructAddDeviceInfoStackInfo(v.([]interface{})); w != nil {
				deviceInfo.StackInfo = *w
			}
		}
		if v, ok := deviceInfoItem["state"]; ok && v != nil {
			deviceInfo.State = v.(string)
		}
		if v, ok := deviceInfoItem["sudi_required"]; ok && v != nil {
			deviceInfo.SudiRequired = v.(bool)
		}
		if v, ok := deviceInfoItem["tags"]; ok && v != nil {
			deviceInfo.Tags = v.(string)
		}
		if v, ok := deviceInfoItem["user_mic_numbers"]; ok && v != nil {
			deviceInfo.UserMicNumbers = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["user_sudi_serial_nos"]; ok && v != nil {
			deviceInfo.UserSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["virtual_account_id"]; ok && v != nil {
			deviceInfo.VirtualAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_id"]; ok && v != nil {
			deviceInfo.WorkflowID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_name"]; ok && v != nil {
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
		if v, ok := rs["details"]; ok && v != nil {
			runSummaryItem.Details = v.(string)
		}
		if v, ok := rs["error_flag"]; ok && v != nil {
			runSummaryItem.ErrorFlag = v.(bool)
		}
		if v, ok := rs["history_task_info"]; ok && v != nil {
			if w := constructAddRunSummaryListHistoryTaskInfo(v.([]interface{})); w != nil {
				runSummaryItem.HistoryTaskInfo = *w
			}
		}
		if v, ok := rs["timestamp"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructAddSystemResetWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddSystemWorkflow(workflows []interface{}) *dnac.AddDeviceToPnpDatabaseRequestSystemWorkflow {
	var workflowItem dnac.AddDeviceToPnpDatabaseRequestSystemWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructAddSystemWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddWorkflow(workflows []interface{}) *dnac.AddDeviceToPnpDatabaseRequestWorkflow {
	var workflowItem dnac.AddDeviceToPnpDatabaseRequestWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructAddWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructAddWorkflowParams(params []interface{}) *dnac.AddDeviceToPnpDatabaseRequestWorkflowParameters {
	var workflowParameters dnac.AddDeviceToPnpDatabaseRequestWorkflowParameters
	for _, param := range params {
		ps := param.(map[string]interface{})

		if v, ok := ps["config_list"]; ok && v != nil {
			configListV := v.([]interface{})
			var configList []dnac.AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList
			for _, ci := range configListV {
				var configItem dnac.AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList
				cis := ci.(map[string]interface{})
				if v, ok := cis["config_id"]; ok && v != nil {
					configItem.ConfigID = v.(string)
				}
				if v, ok := cis["config_paramters"]; ok && v != nil {
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
		if v, ok := ps["license_level"]; ok && v != nil {
			workflowParameters.LicenseLevel = v.(string)
		}
		if v, ok := ps["license_type"]; ok && v != nil {
			workflowParameters.LicenseType = v.(string)
		}
		if v, ok := ps["top_of_stack_serial_number"]; ok && v != nil {
			workflowParameters.TopOfStackSerialNumber = v.(string)
		}
	}
	return &workflowParameters
}

func constructAddPnPDevice(pnpRequest map[string]interface{}) *dnac.AddDeviceToPnpDatabaseRequest {
	var request dnac.AddDeviceToPnpDatabaseRequest
	if v, ok := pnpRequest["id"]; ok && v != nil {
		request.TypeID = v.(string)
	}
	if v, ok := pnpRequest["day_zero_config"]; ok && v != nil {
		if w := constructAddZeroConfig(v.([]interface{})); w != nil {
			request.DayZeroConfig = *w
		}
	}
	if v, ok := pnpRequest["day_zero_config_preview"]; ok && v != nil {
		request.DayZeroConfigPreview = v.(string)
	}
	if v, ok := pnpRequest["device_info"]; ok && v != nil {
		if w := constructAddDeviceInfo(v.([]interface{})); w != nil {
			request.DeviceInfo = *w
		}
	}
	if v, ok := pnpRequest["run_summary_list"]; ok && v != nil {
		if w := constructAddRunSummaryList(v.([]interface{})); w != nil {
			request.RunSummaryList = *w
		}
	}
	if v, ok := pnpRequest["system_reset_workflow"]; ok && v != nil {
		if w := constructAddSystemResetWorkflow(v.([]interface{})); w != nil {
			request.SystemResetWorkflow = *w
		}
	}
	if v, ok := pnpRequest["system_workflow"]; ok && v != nil {
		if w := constructAddSystemWorkflow(v.([]interface{})); w != nil {
			request.SystemWorkflow = *w
		}
	}
	if v, ok := pnpRequest["tenant_id"]; ok && v != nil {
		request.TenantID = v.(string)
	}
	if v, ok := pnpRequest["version"]; ok && v != nil {
		request.Version = v.(float64)
	}
	if v, ok := pnpRequest["workflow"]; ok && v != nil {
		if w := constructAddWorkflow(v.([]interface{})); w != nil {
			request.Workflow = *w
		}
	}
	if v, ok := pnpRequest["workflow_parameters"]; ok && v != nil {
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
		if v, ok := credential["username"]; ok && v != nil {
			aaaCredentials.Username = v.(string)
		}
		if v, ok := credential["password"]; ok && v != nil {
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
		if v, ok := fsi["freespace"]; ok && v != nil {
			fileSystemItem.Freespace = v.(float64)
		}
		if v, ok := fsi["name"]; ok && v != nil {
			fileSystemItem.Name = v.(string)
		}
		if v, ok := fsi["readable"]; ok && v != nil {
			fileSystemItem.Readable = v.(bool)
		}
		if v, ok := fsi["size"]; ok && v != nil {
			fileSystemItem.Size = v.(float64)
		}
		if v, ok := fsi["type"]; ok && v != nil {
			fileSystemItem.Type = v.(string)
		}
		if v, ok := fsi["writeable"]; ok && v != nil {
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
		if v, ok := is["ipv4_address"]; ok && v != nil {
			interfaceItem.IPv4Address = v.(string)
		}
		if v, ok := is["ipv6_address_list"]; ok && v != nil {
			interfaceItem.IPv6AddressList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := is["mac_address"]; ok && v != nil {
			interfaceItem.MacAddress = v.(string)
		}
		if v, ok := is["name"]; ok && v != nil {
			interfaceItem.Name = v.(string)
		}
		if v, ok := is["status"]; ok && v != nil {
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
		if v, ok := location["address"]; ok && v != nil {
			result.Address = v.(string)
		}
		if v, ok := location["altitude"]; ok && v != nil {
			result.Altitude = v.(string)
		}
		if v, ok := location["latitude"]; ok && v != nil {
			result.Latitude = v.(string)
		}
		if v, ok := location["longitude"]; ok && v != nil {
			result.Longitude = v.(string)
		}
		if v, ok := location["siteId"]; ok && v != nil {
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
		if v, ok := ls["local_interface_name"]; ok && v != nil {
			linkItem.LocalInterfaceName = v.(string)
		}
		if v, ok := ls["local_mac_address"]; ok && v != nil {
			linkItem.LocalMacAddress = v.(string)
		}
		if v, ok := ls["local_short_interface_name"]; ok && v != nil {
			linkItem.LocalShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_device_name"]; ok && v != nil {
			linkItem.RemoteDeviceName = v.(string)
		}
		if v, ok := ls["remote_interface_name"]; ok && v != nil {
			linkItem.RemoteInterfaceName = v.(string)
		}
		if v, ok := ls["remote_mac_address"]; ok && v != nil {
			linkItem.RemoteMacAddress = v.(string)
		}
		if v, ok := ls["remote_platform"]; ok && v != nil {
			linkItem.RemotePlatform = v.(string)
		}
		if v, ok := ls["remote_short_interface_name"]; ok && v != nil {
			linkItem.RemoteShortInterfaceName = v.(string)
		}
		if v, ok := ls["remote_version"]; ok && v != nil {
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
		if v, ok := ps["created_by"]; ok && v != nil {
			profileItem.CreatedBy = v.(string)
		}
		if v, ok := ps["discovery_created"]; ok && v != nil {
			profileItem.DiscoveryCreated = v.(bool)
		}
		if v, ok := ps["primary_endpoint"]; ok && v != nil {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok && x != nil {
					profileItem.PrimaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok && x != nil {
					profileItem.PrimaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok && x != nil {
					profileItem.PrimaryEndpoint.Protocol = x.(string)
				}
			}
		}
		if v, ok := ps["profile_name"]; ok && v != nil {
			profileItem.ProfileName = v.(string)
		}
		if v, ok := ps["secondary_endpoint"]; ok && v != nil {
			if w := v.([]interface{}); len(w) > 0 {
				endpoint := w[0].(map[string]interface{})
				if x, ok := endpoint["certificate"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Certificate = x.(string)
				}
				if x, ok := endpoint["fqdn"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Fqdn = x.(string)
				}
				if x, ok := endpoint["ipv4_address"]; ok && x != nil {
					profileItem.SecondaryEndpoint.IPv4Address = x.(string)
				}
				if x, ok := endpoint["ipv6_address"]; ok && x != nil {
					profileItem.SecondaryEndpoint.IPv6Address = x.(string)
				}
				if x, ok := endpoint["port"]; ok && x != nil {
					profileItem.SecondaryEndpoint.Port = x.(float64)
				}
				if x, ok := endpoint["protocol"]; ok && x != nil {
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
		if v, ok := cos["cli"]; ok && v != nil {
			cliOutputItem.Cli = v.(string)
		}
		if v, ok := cos["cli_output"]; ok && v != nil {
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
		if v, ok := ms["hardware_version"]; ok && v != nil {
			memberItem.HardwareVersion = v.(string)
		}
		if v, ok := ms["license_level"]; ok && v != nil {
			memberItem.LicenseLevel = v.(string)
		}
		if v, ok := ms["license_type"]; ok && v != nil {
			memberItem.LicenseType = v.(string)
		}
		if v, ok := ms["mac_address"]; ok && v != nil {
			memberItem.MacAddress = v.(string)
		}
		if v, ok := ms["pid"]; ok && v != nil {
			memberItem.Pid = v.(string)
		}
		if v, ok := ms["priority"]; ok && v != nil {
			memberItem.Priority = v.(float64)
		}
		if v, ok := ms["role"]; ok && v != nil {
			memberItem.Role = v.(string)
		}
		if v, ok := ms["serial_number"]; ok && v != nil {
			memberItem.SerialNumber = v.(string)
		}
		if v, ok := ms["software_version"]; ok && v != nil {
			memberItem.SoftwareVersion = v.(string)
		}
		if v, ok := ms["stack_number"]; ok && v != nil {
			memberItem.StackNumber = v.(float64)
		}
		if v, ok := ms["state"]; ok && v != nil {
			memberItem.State = v.(string)
		}
		if v, ok := ms["sudi_serial_number"]; ok && v != nil {
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
		if v, ok := stackInfo["is_full_ring"]; ok && v != nil {
			result.IsFullRing = v.(bool)
		}
		if v, ok := stackInfo["stack_member_list"]; ok && v != nil {
			if w := constructUpdateDeviceInfoStackInfoStackMemberList(v.([]interface{})); w != nil {
				result.StackMemberList = *w
			}
		}
		if v, ok := stackInfo["stack_ring_protocol"]; ok && v != nil {
			result.StackRingProtocol = v.(string)
		}
		if v, ok := stackInfo["supports_stack_workflows"]; ok && v != nil {
			result.SupportsStackWorkflows = v.(bool)
		}
		if v, ok := stackInfo["total_member_count"]; ok && v != nil {
			result.TotalMemberCount = v.(float64)
		}
		if v, ok := stackInfo["valid_license_levels"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := taskInfo["addn_details"]; ok && v != nil {
			if w := constructUpdateRunSummaryListHistoryTaskInfoAddnDetails(v.([]interface{})); w != nil {
				result.AddnDetails = *w
			}
		}
		if v, ok := taskInfo["name"]; ok && v != nil {
			result.Name = v.(string)
		}
		if v, ok := taskInfo["time_taken"]; ok && v != nil {
			result.TimeTaken = v.(float64)
		}
		if v, ok := taskInfo["type"]; ok && v != nil {
			result.Type = v.(string)
		}
		if v, ok := taskInfo["work_item_list"]; ok && v != nil {
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
		if v, ok := deviceInfoItem["aaa_credentials"]; ok && v != nil {
			if w := constructUpdateDeviceInfoAAACredentials(v.([]interface{})); w != nil {
				deviceInfo.AAACredentials = *w
			}
		}
		if v, ok := deviceInfoItem["added_on"]; ok && v != nil {
			deviceInfo.AddedOn = v.(float64)
		}
		if v, ok := deviceInfoItem["addn_mac_addrs"]; ok && v != nil {
			deviceInfo.AddnMacAddrs = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["agent_type"]; ok && v != nil {
			deviceInfo.AgentType = v.(string)
		}
		if v, ok := deviceInfoItem["auth_status"]; ok && v != nil {
			deviceInfo.AuthStatus = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_mic_number"]; ok && v != nil {
			deviceInfo.AuthenticatedMicNumber = v.(string)
		}
		if v, ok := deviceInfoItem["authenticated_sudi_serial_no"]; ok && v != nil {
			deviceInfo.AuthenticatedSudiSerialNo = v.(string)
		}
		if v, ok := deviceInfoItem["capabilities_supported"]; ok && v != nil {
			deviceInfo.CapabilitiesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["cm_state"]; ok && v != nil {
			deviceInfo.CmState = v.(string)
		}
		if v, ok := deviceInfoItem["description"]; ok && v != nil {
			deviceInfo.Description = v.(string)
		}
		if v, ok := deviceInfoItem["device_sudi_serial_nos"]; ok && v != nil {
			deviceInfo.DeviceSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["device_type"]; ok && v != nil {
			deviceInfo.DeviceType = v.(string)
		}
		if v, ok := deviceInfoItem["features_supported"]; ok && v != nil {
			deviceInfo.FeaturesSupported = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["file_system_list"]; ok && v != nil {
			if w := constructUpdateDeviceInfoFileSystemList(v.([]interface{})); w != nil {
				deviceInfo.FileSystemList = *w
			}
		}
		if v, ok := deviceInfoItem["first_contact"]; ok && v != nil {
			deviceInfo.FirstContact = v.(float64)
		}
		if v, ok := deviceInfoItem["hostname"]; ok && v != nil {
			deviceInfo.Hostname = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok && v != nil {
			if w := constructUpdateDeviceInfoHTTPHeaders(v.([]interface{})); w != nil {
				deviceInfo.HTTPHeaders = *w
			}
		}
		if v, ok := deviceInfoItem["image_file"]; ok && v != nil {
			deviceInfo.ImageFile = v.(string)
		}
		if v, ok := deviceInfoItem["image_version"]; ok && v != nil {
			deviceInfo.ImageVersion = v.(string)
		}
		if v, ok := deviceInfoItem["http_headers"]; ok && v != nil {
			if w := constructUpdateDeviceInfoIPInterfaces(v.([]interface{})); w != nil {
				deviceInfo.IPInterfaces = *w
			}
		}
		if v, ok := deviceInfoItem["last_contact"]; ok && v != nil {
			deviceInfo.LastContact = v.(float64)
		}
		if v, ok := deviceInfoItem["last_sync_time"]; ok && v != nil {
			deviceInfo.LastSyncTime = v.(float64)
		}
		if v, ok := deviceInfoItem["last_update_on"]; ok && v != nil {
			deviceInfo.LastUpdateOn = v.(float64)
		}
		if v, ok := deviceInfoItem["location"]; ok && v != nil {
			if w := constructUpdateDeviceInfoLocation(v.([]interface{})); w != nil {
				deviceInfo.Location = *w
			}
		}
		if v, ok := deviceInfoItem["mac_address"]; ok && v != nil {
			deviceInfo.MacAddress = v.(string)
		}
		if v, ok := deviceInfoItem["mode"]; ok && v != nil {
			deviceInfo.Mode = v.(string)
		}
		if v, ok := deviceInfoItem["name"]; ok && v != nil {
			deviceInfo.Name = v.(string)
		}
		if v, ok := deviceInfoItem["neighbor_links"]; ok && v != nil {
			if w := constructUpdateDeviceInfoNeighborLinks(v.([]interface{})); w != nil {
				deviceInfo.NeighborLinks = *w
			}
		}
		if v, ok := deviceInfoItem["onb_state"]; ok && v != nil {
			deviceInfo.OnbState = v.(string)
		}
		if v, ok := deviceInfoItem["pid"]; ok && v != nil {
			deviceInfo.Pid = v.(string)
		}
		if v, ok := deviceInfoItem["pnp_profile_list"]; ok && v != nil {
			if w := constructUpdateDeviceInfoPnpProfileList(v.([]interface{})); w != nil {
				deviceInfo.PnpProfileList = *w
			}
		}
		if v, ok := deviceInfoItem["populate_inventory"]; ok && v != nil {
			deviceInfo.PopulateInventory = v.(bool)
		}
		if v, ok := deviceInfoItem["pre_workflow_cli_ouputs"]; ok && v != nil {
			if w := constructUpdateDeviceInfoPreWorkflowCliOuputs(v.([]interface{})); w != nil {
				deviceInfo.PreWorkflowCliOuputs = *w
			}
		}
		if v, ok := deviceInfoItem["project_id"]; ok && v != nil {
			deviceInfo.ProjectID = v.(string)
		}
		if v, ok := deviceInfoItem["project_name"]; ok && v != nil {
			deviceInfo.ProjectName = v.(string)
		}
		if v, ok := deviceInfoItem["reload_requested"]; ok && v != nil {
			deviceInfo.ReloadRequested = v.(bool)
		}
		if v, ok := deviceInfoItem["serial_number"]; ok && v != nil {
			deviceInfo.SerialNumber = v.(string)
		}
		if v, ok := deviceInfoItem["site_id"]; ok && v != nil {
			deviceInfo.SiteID = v.(string)
		}
		if v, ok := deviceInfoItem["site_name"]; ok && v != nil {
			deviceInfo.SiteName = v.(string)
		}
		if v, ok := deviceInfoItem["smart_account_id"]; ok && v != nil {
			deviceInfo.SmartAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["source"]; ok && v != nil {
			deviceInfo.Source = v.(string)
		}
		if v, ok := deviceInfoItem["stack"]; ok && v != nil {
			deviceInfo.Stack = v.(bool)
		}
		if v, ok := deviceInfoItem["stack_info"]; ok && v != nil {
			if w := constructUpdateDeviceInfoStackInfo(v.([]interface{})); w != nil {
				deviceInfo.StackInfo = *w
			}
		}
		if v, ok := deviceInfoItem["state"]; ok && v != nil {
			deviceInfo.State = v.(string)
		}
		if v, ok := deviceInfoItem["sudi_required"]; ok && v != nil {
			deviceInfo.SudiRequired = v.(bool)
		}
		if v, ok := deviceInfoItem["tags"]; ok && v != nil {
			deviceInfo.Tags = v.(string)
		}
		if v, ok := deviceInfoItem["user_mic_numbers"]; ok && v != nil {
			deviceInfo.UserMicNumbers = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["user_sudi_serial_nos"]; ok && v != nil {
			deviceInfo.UserSudiSerialNos = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := deviceInfoItem["virtual_account_id"]; ok && v != nil {
			deviceInfo.VirtualAccountID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_id"]; ok && v != nil {
			deviceInfo.WorkflowID = v.(string)
		}
		if v, ok := deviceInfoItem["workflow_name"]; ok && v != nil {
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
		if v, ok := rs["details"]; ok && v != nil {
			runSummaryItem.Details = v.(string)
		}
		if v, ok := rs["error_flag"]; ok && v != nil {
			runSummaryItem.ErrorFlag = v.(bool)
		}
		if v, ok := rs["history_task_info"]; ok && v != nil {
			if w := constructUpdateRunSummaryListHistoryTaskInfo(v.([]interface{})); w != nil {
				runSummaryItem.HistoryTaskInfo = *w
			}
		}
		if v, ok := rs["timestamp"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
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
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructUpdateSystemResetWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateSystemWorkflow(workflows []interface{}) *dnac.UpdateDeviceRequestSystemWorkflow {
	var workflowItem dnac.UpdateDeviceRequestSystemWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructUpdateSystemWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateWorkflow(workflows []interface{}) *dnac.UpdateDeviceRequestWorkflow {
	var workflowItem dnac.UpdateDeviceRequestWorkflow
	if len(workflows) > 0 {
		ws := workflows[0].(map[string]interface{})
		if v, ok := ws["id"]; ok && v != nil {
			workflowItem.TypeID = v.(string)
		}
		if v, ok := ws["add_to_inventory"]; ok && v != nil {
			workflowItem.AddToInventory = v.(bool)
		}
		if v, ok := ws["added_on"]; ok && v != nil {
			workflowItem.AddedOn = v.(float64)
		}
		if v, ok := ws["config_id"]; ok && v != nil {
			workflowItem.ConfigID = v.(string)
		}
		if v, ok := ws["curr_task_idx"]; ok && v != nil {
			workflowItem.CurrTaskIDx = v.(float64)
		}
		if v, ok := ws["description"]; ok && v != nil {
			workflowItem.Description = v.(string)
		}
		if v, ok := ws["end_time"]; ok && v != nil {
			workflowItem.EndTime = v.(int)
		}
		if v, ok := ws["exec_time"]; ok && v != nil {
			workflowItem.ExecTime = v.(float64)
		}
		if v, ok := ws["image_id"]; ok && v != nil {
			workflowItem.ImageID = v.(string)
		}
		if v, ok := ws["instance_type"]; ok && v != nil {
			workflowItem.InstanceType = v.(string)
		}
		if v, ok := ws["lastupdate_on"]; ok && v != nil {
			workflowItem.LastupdateOn = v.(float64)
		}
		if v, ok := ws["name"]; ok && v != nil {
			workflowItem.Name = v.(string)
		}
		if v, ok := ws["start_time"]; ok && v != nil {
			workflowItem.StartTime = v.(int)
		}
		if v, ok := ws["state"]; ok && v != nil {
			workflowItem.State = v.(string)
		}
		if v, ok := ws["tasks"]; ok && v != nil {
			if w := constructUpdateWorkflowTasks(v.([]interface{})); w != nil {
				workflowItem.Tasks = *w
			}
		}
		if v, ok := ws["tenant_id"]; ok && v != nil {
			workflowItem.TenantID = v.(string)
		}
		if v, ok := ws["type"]; ok && v != nil {
			workflowItem.Type = v.(string)
		}
		if v, ok := ws["use_state"]; ok && v != nil {
			workflowItem.UseState = v.(string)
		}
		if v, ok := ws["version"]; ok && v != nil {
			workflowItem.Version = v.(float64)
		}
	}
	return &workflowItem
}

func constructUpdateWorkflowParams(params []interface{}) *dnac.UpdateDeviceRequestWorkflowParameters {
	var workflowParameters dnac.UpdateDeviceRequestWorkflowParameters
	for _, param := range params {
		ps := param.(map[string]interface{})

		if v, ok := ps["config_list"]; ok && v != nil {
			configListV := v.([]interface{})
			var configList []dnac.UpdateDeviceRequestWorkflowParametersConfigList
			for _, ci := range configListV {
				var configItem dnac.UpdateDeviceRequestWorkflowParametersConfigList
				cis := ci.(map[string]interface{})
				if v, ok := cis["config_id"]; ok && v != nil {
					configItem.ConfigID = v.(string)
				}
				if v, ok := cis["config_paramters"]; ok && v != nil {
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
		if v, ok := ps["license_level"]; ok && v != nil {
			workflowParameters.LicenseLevel = v.(string)
		}
		if v, ok := ps["license_type"]; ok && v != nil {
			workflowParameters.LicenseType = v.(string)
		}
		if v, ok := ps["top_of_stack_serial_number"]; ok && v != nil {
			workflowParameters.TopOfStackSerialNumber = v.(string)
		}
	}
	return &workflowParameters
}

func constructUpdatePnPDevice(pnpRequest map[string]interface{}) *dnac.UpdateDeviceRequest {
	var request dnac.UpdateDeviceRequest
	if v, ok := pnpRequest["id"]; ok && v != nil {
		request.TypeID = v.(string)
	}
	if v, ok := pnpRequest["day_zero_config"]; ok && v != nil {
		if w := constructUpdateZeroConfig(v.([]interface{})); w != nil {
			request.DayZeroConfig = *w
		}
	}
	if v, ok := pnpRequest["day_zero_config_preview"]; ok && v != nil {
		request.DayZeroConfigPreview = v.(string)
	}
	if v, ok := pnpRequest["device_info"]; ok && v != nil {
		if w := constructUpdateDeviceInfo(v.([]interface{})); w != nil {
			request.DeviceInfo = *w
		}
	}
	if v, ok := pnpRequest["run_summary_list"]; ok && v != nil {
		if w := constructUpdateRunSummaryList(v.([]interface{})); w != nil {
			request.RunSummaryList = *w
		}
	}
	if v, ok := pnpRequest["system_reset_workflow"]; ok && v != nil {
		if w := constructUpdateSystemResetWorkflow(v.([]interface{})); w != nil {
			request.SystemResetWorkflow = *w
		}
	}
	if v, ok := pnpRequest["system_workflow"]; ok && v != nil {
		if w := constructUpdateSystemWorkflow(v.([]interface{})); w != nil {
			request.SystemWorkflow = *w
		}
	}
	if v, ok := pnpRequest["tenant_id"]; ok && v != nil {
		request.TenantID = v.(string)
	}
	if v, ok := pnpRequest["version"]; ok && v != nil {
		request.Version = v.(float64)
	}
	if v, ok := pnpRequest["workflow"]; ok && v != nil {
		if w := constructUpdateWorkflow(v.([]interface{})); w != nil {
			request.Workflow = *w
		}
	}
	if v, ok := pnpRequest["workflow_parameters"]; ok && v != nil {
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
	searchResponse, _, err := client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	deviceItem := flattenPnPDeviceReadItem(searchResponse)
	if err := d.Set("item", deviceItem); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePnPDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	deviceID := d.Id()
	searchResponse, _, err := client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
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
	searchResponse, _, err := client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil || searchResponse == nil {
		return diags
	}

	// Call function to delete application resource
	_, _, err = client.DeviceOnboardingPnP.DeleteDeviceByIDFromPnP(deviceID)
	if err != nil {
		return diag.FromErr(err)
	}

	searchResponse, _, err = client.DeviceOnboardingPnP.GetDeviceByID(deviceID)
	if err != nil || searchResponse == nil {
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to delete PnP device",
		Detail:   "",
	})

	return diags
}
