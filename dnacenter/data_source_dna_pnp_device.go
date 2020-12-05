package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpDevicesWorkItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_taken": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func pnpDevicesWorkflowTasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"curr_work_item_idx": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"end_time": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"start_time": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"state": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"task_seq_no": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"time_taken": &schema.Schema{
					Type:     schema.TypeFloat,
					Computed: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"work_item_list": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem:     pnpDevicesWorkItems(),
				},
			},
		},
	}
}

func pnpDevicesKeyValueMap() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func pnpDevicesSystemWorkflow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"add_to_inventory": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"added_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"curr_task_idx": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"exec_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastupdate_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tasks": pnpDevicesWorkflowTasks(),
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceRead,
		Schema: map[string]*schema.Schema{
			"sort": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"asc", "des"}),
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"onb_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cm_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"smart_account_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_account_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_contact": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"day_zero_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"day_zero_config_preview": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aaa_credentials": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"username": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"addn_mac_addrs": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"agent_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"authenticated_mic_number": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"authenticated_sudi_serial_no": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"capabilities_supported": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cm_state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"features_supported": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"file_system_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"freespace": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"readable": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"size": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"writeable": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"first_contact": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"http_headers": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     pnpDevicesKeyValueMap(),
									},
									"image_file": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv6_address_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"last_contact": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"last_sync_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"last_update_on": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"altitude": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"latitude": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"longitude": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"site_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"neighbor_links": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"local_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"local_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_device_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_platform": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"remote_version": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"onb_state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"pid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"pnp_profile_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"created_by": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"discovery_created": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"primary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"profile_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"secondary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
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
										Computed: true,
									},
									"pre_workflow_cli_ouputs": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cli": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"cli_output": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"project_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"reload_requested": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"smart_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"stack": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"stack_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_full_ring": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hardware_version": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"license_level": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"license_type": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"mac_address": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"pid": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"priority": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"software_version": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"stack_number": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"sudi_serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"supports_stack_workflows": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"total_member_count": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"valid_license_levels": &schema.Schema{
													Type:     schema.TypeList,
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
										Computed: true,
									},
									"sudi_required": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_mic_numbers": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"user_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"workflow_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"workflow_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"run_summary_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"details": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_flag": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"history_task_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"addn_details": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem:     pnpDevicesKeyValueMap(),
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem:     pnpDevicesWorkItems(),
												},
											},
										},
									},
									"timestamp": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"system_reset_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     pnpDevicesSystemWorkflow(),
						},
						"system_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     pnpDevicesSystemWorkflow(),
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"workflow": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     pnpDevicesSystemWorkflow(),
						},
						"workflow_parameters": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"config_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem:     pnpDevicesKeyValueMap(),
												},
											},
										},
									},
									"license_level": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"license_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"top_of_stack_serial_number": &schema.Schema{
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

func dataSourcePnPDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParam := dnac.GetPnpDeviceListQueryParams{}

	if v, ok := d.GetOk("limit"); ok {
		queryParam.Limit = v.(int)
	}
	if v, ok := d.GetOk("offset"); ok {
		queryParam.Offset = v.(int)
	}
	if v, ok := d.GetOk("sort"); ok {
		queryParam.Sort = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("sort_order"); ok {
		queryParam.SortOrder = v.(string)
	}
	if v, ok := d.GetOk("serial_number"); ok {
		queryParam.SerialNumber = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("state"); ok {
		queryParam.State = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("onb_state"); ok {
		queryParam.OnbState = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("cm_state"); ok {
		queryParam.CmState = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("name"); ok {
		queryParam.Name = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("pid"); ok {
		queryParam.Pid = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("source"); ok {
		queryParam.Source = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("project_id"); ok {
		queryParam.ProjectID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflow_id"); ok {
		queryParam.WorkflowID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("project_name"); ok {
		queryParam.ProjectName = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflow_name"); ok {
		queryParam.WorkflowName = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("smart_account_id"); ok {
		queryParam.SmartAccountID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("virtual_account_i"); ok {
		queryParam.VirtualAccountID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("last_contact"); ok {
		queryParam.LastContact = v.(bool)
	}
	if v, ok := d.GetOk("mac_address"); ok {
		queryParam.MacAddress = v.(string)
	}
	if v, ok := d.GetOk("hostname"); ok {
		queryParam.Hostname = v.(string)
	}
	if v, ok := d.GetOk("site_name"); ok {
		queryParam.SiteName = v.(string)
	}

	// Prepare Request
	response, _, err := client.DeviceOnboardingPnP.GetPnpDeviceList(&queryParam)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source

	items := flattenPnPDevicesReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
