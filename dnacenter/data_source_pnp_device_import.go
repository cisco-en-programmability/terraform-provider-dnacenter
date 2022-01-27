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
func dataSourcePnpDeviceImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Add devices to PnP in bulk
`,

		ReadContext: dataSourcePnpDeviceImportRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failure_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"index": &schema.Schema{
										Description: `Index`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"msg": &schema.Schema{
										Description: `Msg`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"serial_num": &schema.Schema{
										Description: `Serial Num`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"success_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"day_zero_config": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Description: `Config`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"day_zero_config_preview": &schema.Schema{
										Description: `Day Zero Config Preview`,
										Type:        schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
																Description: `Password`,
																Type:        schema.TypeString,
																Sensitive:   true,
																Computed:    true,
															},
															"username": &schema.Schema{
																Description: `Username`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"addn_mac_addrs": &schema.Schema{
													Description: `Addn Mac Addrs`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"agent_type": &schema.Schema{
													Description: `Agent Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"auth_status": &schema.Schema{
													Description: `Auth Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"authenticated_mic_number": &schema.Schema{
													Description: `Authenticated Mic Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"authenticated_sudi_serial_no": &schema.Schema{
													Description: `Authenticated Sudi Serial No`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"capabilities_supported": &schema.Schema{
													Description: `Capabilities Supported`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"cm_state": &schema.Schema{
													Description: `Cm State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"device_sudi_serial_nos": &schema.Schema{
													Description: `Device Sudi Serial Nos`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"device_type": &schema.Schema{
													Description: `Device Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"features_supported": &schema.Schema{
													Description: `Features Supported`,
													Type:        schema.TypeList,
													Computed:    true,
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
																Description: `Freespace`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"readable": &schema.Schema{
																Description: `Readable`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"size": &schema.Schema{
																Description: `Size`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"writeable": &schema.Schema{
																Description: `Writeable`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"first_contact": &schema.Schema{
													Description: `First Contact`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"hostname": &schema.Schema{
													Description: `Hostname`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"http_headers": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Description: `Key`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"image_file": &schema.Schema{
													Description: `Image File`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"image_version": &schema.Schema{
													Description: `Image Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"ip_interfaces": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4_address": &schema.Schema{
																Description: `Ipv4 Address`,
																Type:        schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Computed: true,
															},
															"ipv6_address_list": &schema.Schema{
																Description: `Ipv6 Address List`,
																Type:        schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"status": &schema.Schema{
																Description: `Status`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"last_contact": &schema.Schema{
													Description: `Last Contact`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"last_sync_time": &schema.Schema{
													Description: `Last Sync Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"last_update_on": &schema.Schema{
													Description: `Last Update On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"location": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"address": &schema.Schema{
																Description: `Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"altitude": &schema.Schema{
																Description: `Altitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"latitude": &schema.Schema{
																Description: `Latitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"longitude": &schema.Schema{
																Description: `Longitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"site_id": &schema.Schema{
																Description: `Site Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"mac_address": &schema.Schema{
													Description: `Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"mode": &schema.Schema{
													Description: `Mode`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"neighbor_links": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"local_interface_name": &schema.Schema{
																Description: `Local Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"local_mac_address": &schema.Schema{
																Description: `Local Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"local_short_interface_name": &schema.Schema{
																Description: `Local Short Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_device_name": &schema.Schema{
																Description: `Remote Device Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_interface_name": &schema.Schema{
																Description: `Remote Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_mac_address": &schema.Schema{
																Description: `Remote Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_platform": &schema.Schema{
																Description: `Remote Platform`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_short_interface_name": &schema.Schema{
																Description: `Remote Short Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_version": &schema.Schema{
																Description: `Remote Version`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"onb_state": &schema.Schema{
													Description: `Onb State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"pid": &schema.Schema{
													Description: `Pid`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"pnp_profile_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"created_by": &schema.Schema{
																Description: `Created By`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"discovery_created": &schema.Schema{
																Description: `Discovery Created`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"primary_endpoint": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"certificate": &schema.Schema{
																			Description: `Certificate`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"fqdn": &schema.Schema{
																			Description: `Fqdn`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"ipv4_address": &schema.Schema{
																			Description: `Ipv4 Address`,
																			Type:        schema.TypeList,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																			Computed: true,
																		},
																		"ipv6_address": &schema.Schema{
																			Description: `Ipv6 Address`,
																			Type:        schema.TypeList,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																			Computed: true,
																		},
																		"port": &schema.Schema{
																			Description: `Port`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"protocol": &schema.Schema{
																			Description: `Protocol`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
															"profile_name": &schema.Schema{
																Description: `Profile Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"secondary_endpoint": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"certificate": &schema.Schema{
																			Description: `Certificate`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"fqdn": &schema.Schema{
																			Description: `Fqdn`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"ipv4_address": &schema.Schema{
																			Description: `Ipv4 Address`,
																			Type:        schema.TypeList,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																			Computed: true,
																		},
																		"ipv6_address": &schema.Schema{
																			Description: `Ipv6 Address`,
																			Type:        schema.TypeList,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																			Computed: true,
																		},
																		"port": &schema.Schema{
																			Description: `Port`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"protocol": &schema.Schema{
																			Description: `Protocol`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"populate_inventory": &schema.Schema{
													Description: `Populate Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"pre_workflow_cli_ouputs": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cli": &schema.Schema{
																Description: `Cli`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"cli_output": &schema.Schema{
																Description: `Cli Output`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"project_id": &schema.Schema{
													Description: `Project Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"project_name": &schema.Schema{
													Description: `Project Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"reload_requested": &schema.Schema{
													Description: `Reload Requested`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"serial_number": &schema.Schema{
													Description: `Serial Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"site_id": &schema.Schema{
													Description: `Site Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"site_name": &schema.Schema{
													Description: `Site Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"smart_account_id": &schema.Schema{
													Description: `Smart Account Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"source": &schema.Schema{
													Description: `Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"stack": &schema.Schema{
													Description: `Stack`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"stack_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"is_full_ring": &schema.Schema{
																Description: `Is Full Ring`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"stack_member_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"hardware_version": &schema.Schema{
																			Description: `Hardware Version`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"license_level": &schema.Schema{
																			Description: `License Level`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"license_type": &schema.Schema{
																			Description: `License Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"mac_address": &schema.Schema{
																			Description: `Mac Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"pid": &schema.Schema{
																			Description: `Pid`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"priority": &schema.Schema{
																			Description: `Priority`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"role": &schema.Schema{
																			Description: `Role`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"serial_number": &schema.Schema{
																			Description: `Serial Number`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"software_version": &schema.Schema{
																			Description: `Software Version`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"stack_number": &schema.Schema{
																			Description: `Stack Number`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"sudi_serial_number": &schema.Schema{
																			Description: `Sudi Serial Number`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
															"stack_ring_protocol": &schema.Schema{
																Description: `Stack Ring Protocol`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"supports_stack_workflows": &schema.Schema{
																Description: `Supports Stack Workflows`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"total_member_count": &schema.Schema{
																Description: `Total Member Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"valid_license_levels": &schema.Schema{
																Description: `Valid License Levels`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sudi_required": &schema.Schema{
													Description: `Sudi Required`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"tags": &schema.Schema{
													Description: `Tags`,
													Type:        schema.TypeList,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Computed: true,
												},
												"user_mic_numbers": &schema.Schema{
													Description: `User Mic Numbers`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"user_sudi_serial_nos": &schema.Schema{
													Description: `User Sudi Serial Nos`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"virtual_account_id": &schema.Schema{
													Description: `Virtual Account Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"workflow_id": &schema.Schema{
													Description: `Workflow Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"workflow_name": &schema.Schema{
													Description: `Workflow Name`,
													Type:        schema.TypeString,
													Computed:    true,
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
													Description: `Details`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"error_flag": &schema.Schema{
													Description: `Error Flag`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
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
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"key": &schema.Schema{
																			Description: `Key`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"value": &schema.Schema{
																			Description: `Value`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"timestamp": &schema.Schema{
													Description: `Timestamp`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"system_reset_workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"system_workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
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
																Description: `Config Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"config_parameters": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"key": &schema.Schema{
																			Description: `Key`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"value": &schema.Schema{
																			Description: `Value`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"license_level": &schema.Schema{
													Description: `License Level`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"license_type": &schema.Schema{
													Description: `License Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"top_of_stack_serial_number": &schema.Schema{
													Description: `Top Of Stack Serial Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"payload": &schema.Schema{
				Description: `Array of RequestDeviceOnboardingPnpImportDevicesInBulk`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_credentials": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"password": &schema.Schema{
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"username": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"addn_mac_addrs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"agent_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"authenticated_sudi_serial_no": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"capabilities_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cm_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"features_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"file_system_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"freespace": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"readable": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"size": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"writeable": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"first_contact": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"http_headers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"image_file": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"image_version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4_address": &schema.Schema{
													Type: schema.TypeList,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Optional: true,
												},
												"ipv6_address_list": &schema.Schema{
													Type: schema.TypeList,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Optional: true,
												},
												"mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"last_contact": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"last_sync_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"last_update_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"altitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"latitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"longitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"site_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"neighbor_links": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"local_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"local_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"local_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_device_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_platform": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_version": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"onb_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pnp_profile_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"created_by": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"discovery_created": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"primary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"ipv4_address": &schema.Schema{
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
															"ipv6_address": &schema.Schema{
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"profile_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"secondary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"ipv4_address": &schema.Schema{
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
															"ipv6_address": &schema.Schema{
																Type: schema.TypeList,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Optional: true,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"populate_inventory": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"pre_workflow_cli_ouputs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cli": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"cli_output": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"project_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"project_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"reload_requested": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"smart_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"source": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"stack": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"stack_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"is_full_ring": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hardware_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"license_level": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"license_type": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"mac_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"pid": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"priority": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"software_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"stack_number": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"sudi_serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"supports_stack_workflows": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"total_member_count": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"valid_license_levels": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
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
									},
									"sudi_required": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"tags": &schema.Schema{
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional: true,
									},
									"user_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"workflow_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"workflow_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"run_summary_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"details": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"error_flag": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"history_task_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addn_details": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"command": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"end_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"output_str": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"time_taken": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"timestamp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"system_reset_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"add_to_inventory": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"config_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"curr_task_idx": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"exec_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"image_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"lastupdate_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tasks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"curr_work_item_idx": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"end_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"state": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"task_seq_no": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"command": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"end_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"output_str": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"time_taken": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"use_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"system_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"add_to_inventory": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"config_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"curr_task_idx": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"exec_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"image_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"lastupdate_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tasks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"curr_work_item_idx": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"end_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"state": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"task_seq_no": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"command": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"end_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"output_str": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"time_taken": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"use_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"add_to_inventory": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"config_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"curr_task_idx": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"exec_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"image_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"lastupdate_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tasks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"curr_work_item_idx": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"end_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"state": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"task_seq_no": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"work_item_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"command": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"end_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"output_str": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"time_taken": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"use_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"workflow_parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"license_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"license_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"top_of_stack_serial_number": &schema.Schema{
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

func dataSourcePnpDeviceImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportDevicesInBulk")
		request1 := expandRequestPnpDeviceImportImportDevicesInBulk(ctx, "", d)

		response1, restyResp1, err := client.DeviceOnboardingPnp.ImportDevicesInBulk(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportDevicesInBulk", err,
				"Failure at ImportDevicesInBulk, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpImportDevicesInBulkItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportDevicesInBulk response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPnpDeviceImportImportDevicesInBulk(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpImportDevicesInBulk {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpImportDevicesInBulk{}
	if v := expandRequestPnpDeviceImportImportDevicesInBulkItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfo(ctx, key+".device_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_summary_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_summary_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_summary_list")))) {
		request.RunSummaryList = expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListArray(ctx, key+".run_summary_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_reset_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_reset_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_reset_workflow")))) {
		request.SystemResetWorkflow = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflow(ctx, key+".system_reset_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_workflow")))) {
		request.SystemWorkflow = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflow(ctx, key+".system_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow")))) {
		request.Workflow = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflow(ctx, key+".workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_parameters")))) {
		request.WorkflowParameters = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParameters(ctx, key+".workflow_parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoAAACredentials(ctx, key+".aaa_credentials.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_mac_addrs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_mac_addrs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_mac_addrs")))) {
		request.AddnMacAddrs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".agent_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".agent_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".agent_type")))) {
		request.AgentType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_status")))) {
		request.AuthStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticated_sudi_serial_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) {
		request.AuthenticatedSudiSerialNo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".capabilities_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".capabilities_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".capabilities_supported")))) {
		request.CapabilitiesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cm_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cm_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cm_state")))) {
		request.CmState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) {
		request.DeviceSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".features_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".features_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".features_supported")))) {
		request.FeaturesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_system_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_system_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_system_list")))) {
		request.FileSystemList = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoFileSystemListArray(ctx, key+".file_system_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_contact")))) {
		request.FirstContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_headers")))) {
		request.HTTPHeaders = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoHTTPHeadersArray(ctx, key+".http_headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_file")))) {
		request.ImageFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_version")))) {
		request.ImageVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interfaces")))) {
		request.IPInterfaces = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesArray(ctx, key+".ip_interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_contact")))) {
		request.LastContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync_time")))) {
		request.LastSyncTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_on")))) {
		request.LastUpdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoLocation(ctx, key+".location.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode")))) {
		request.Mode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbor_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbor_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbor_links")))) {
		request.NeighborLinks = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoNeighborLinksArray(ctx, key+".neighbor_links", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".onb_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".onb_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".onb_state")))) {
		request.OnbState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pnp_profile_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pnp_profile_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pnp_profile_list")))) {
		request.PnpProfileList = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListArray(ctx, key+".pnp_profile_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_workflow_cli_ouputs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) {
		request.PreWorkflowCliOuputs = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPreWorkflowCliOuputsArray(ctx, key+".pre_workflow_cli_ouputs", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reload_requested")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reload_requested")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reload_requested")))) {
		request.ReloadRequested = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source")))) {
		request.Source = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_info")))) {
		request.StackInfo = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoTags(ctx, key+".tags.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_name")))) {
		request.WorkflowName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoAAACredentials {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoAAACredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoFileSystemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoFileSystemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoFileSystemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".freespace")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".freespace")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".freespace")))) {
		request.Freespace = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".readable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".readable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".readable")))) {
		request.Readable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".size")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".size")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".size")))) {
		request.Size = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".writeable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".writeable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".writeable")))) {
		request.Writeable = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoHTTPHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoHTTPHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoHTTPHeaders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_list")))) {
		request.IPv6AddressList = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv6AddressListArray(ctx, key+".ipv6_address_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv4Address {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv6AddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv6AddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoIPInterfacesIPv6AddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoLocation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoLocation {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoLocation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".altitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".altitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".altitude")))) {
		request.Altitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoNeighborLinksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoNeighborLinks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoNeighborLinks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_interface_name")))) {
		request.LocalInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_mac_address")))) {
		request.LocalMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_short_interface_name")))) {
		request.LocalShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_device_name")))) {
		request.RemoteDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_interface_name")))) {
		request.RemoteInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_mac_address")))) {
		request.RemoteMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_platform")))) {
		request.RemotePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_short_interface_name")))) {
		request.RemoteShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_version")))) {
		request.RemoteVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".created_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".created_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".created_by")))) {
		request.CreatedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_created")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_created")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_created")))) {
		request.DiscoveryCreated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_endpoint")))) {
		request.PrimaryEndpoint = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpoint(ctx, key+".primary_endpoint.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_endpoint")))) {
		request.SecondaryEndpoint = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpoint(ctx, key+".secondary_endpoint.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpoint {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpoint {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPreWorkflowCliOuputsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPreWorkflowCliOuputs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoPreWorkflowCliOuputs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli")))) {
		request.Cli = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_output")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_output")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_output")))) {
		request.CliOutput = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_ring_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_ring_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_ring_protocol")))) {
		request.StackRingProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supports_stack_workflows")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supports_stack_workflows")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supports_stack_workflows")))) {
		request.SupportsStackWorkflows = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_member_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_member_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_member_count")))) {
		request.TotalMemberCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_license_levels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_license_levels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_license_levels")))) {
		request.ValidLicenseLevels = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_version")))) {
		request.HardwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_number")))) {
		request.StackNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_serial_number")))) {
		request.SudiSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoTags {
	var request dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoTags
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".details")))) {
		request.Details = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_flag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_flag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_flag")))) {
		request.ErrorFlag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".history_task_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".history_task_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".history_task_info")))) {
		request.HistoryTaskInfo = expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfo(ctx, key+".history_task_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp")))) {
		request.Timestamp = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfo {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_details")))) {
		request.AddnDetails = expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx, key+".addn_details", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoAddnDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoAddnDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemRunSummaryListHistoryTaskInfoWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflow {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemResetWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflow {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemSystemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflow {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParameters {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListArray(ctx, key+".config_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_of_stack_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) {
		request.TopOfStackSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters {
	request := []dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemWorkflowParametersConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters {
	request := dnacentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulk) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessList(item.SuccessList)
	respItem["failure_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemFailureList(item.FailureList)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type_id"] = item.TypeID
		respItem["device_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfo(item.DeviceInfo)
		respItem["system_reset_workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflow(item.SystemResetWorkflow)
		respItem["system_workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflow(item.SystemWorkflow)
		respItem["workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflow(item.Workflow)
		respItem["run_summary_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryList(item.RunSummaryList)
		respItem["workflow_parameters"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParameters(item.WorkflowParameters)
		respItem["day_zero_config"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfig(item.DayZeroConfig)
		respItem["day_zero_config_preview"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfigPreview(item.DayZeroConfigPreview)
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["source"] = item.Source
	respItem["serial_number"] = item.SerialNumber
	respItem["stack"] = boolPtrToString(item.Stack)
	respItem["mode"] = item.Mode
	respItem["state"] = item.State
	respItem["location"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoLocation(item.Location)
	respItem["description"] = item.Description
	respItem["onb_state"] = item.OnbState
	respItem["authenticated_mic_number"] = item.AuthenticatedMicNumber
	respItem["authenticated_sudi_serial_no"] = item.AuthenticatedSudiSerialNo
	respItem["capabilities_supported"] = item.CapabilitiesSupported
	respItem["features_supported"] = item.FeaturesSupported
	respItem["cm_state"] = item.CmState
	respItem["first_contact"] = item.FirstContact
	respItem["last_contact"] = item.LastContact
	respItem["mac_address"] = item.MacAddress
	respItem["pid"] = item.Pid
	respItem["device_sudi_serial_nos"] = item.DeviceSudiSerialNos
	respItem["last_update_on"] = item.LastUpdateOn
	respItem["workflow_id"] = item.WorkflowID
	respItem["workflow_name"] = item.WorkflowName
	respItem["project_id"] = item.ProjectID
	respItem["project_name"] = item.ProjectName
	respItem["device_type"] = item.DeviceType
	respItem["agent_type"] = item.AgentType
	respItem["image_version"] = item.ImageVersion
	respItem["file_system_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoFileSystemList(item.FileSystemList)
	respItem["pnp_profile_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileList(item.PnpProfileList)
	respItem["image_file"] = item.ImageFile
	respItem["http_headers"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoHTTPHeaders(item.HTTPHeaders)
	respItem["neighbor_links"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoNeighborLinks(item.NeighborLinks)
	respItem["last_sync_time"] = item.LastSyncTime
	respItem["ip_interfaces"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfaces(item.IPInterfaces)
	respItem["hostname"] = item.Hostname
	respItem["auth_status"] = item.AuthStatus
	respItem["stack_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfo(item.StackInfo)
	respItem["reload_requested"] = boolPtrToString(item.ReloadRequested)
	respItem["added_on"] = item.AddedOn
	respItem["site_id"] = item.SiteID
	respItem["aaa_credentials"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoAAACredentials(item.AAACredentials)
	respItem["user_mic_numbers"] = item.UserMicNumbers
	respItem["user_sudi_serial_nos"] = item.UserSudiSerialNos
	respItem["addn_mac_addrs"] = item.AddnMacAddrs
	respItem["pre_workflow_cli_ouputs"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPreWorkflowCliOuputs(item.PreWorkflowCliOuputs)
	respItem["tags"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoTags(item.Tags)
	respItem["sudi_required"] = boolPtrToString(item.SudiRequired)
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["populate_inventory"] = boolPtrToString(item.PopulateInventory)
	respItem["site_name"] = item.SiteName
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoLocation(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_id"] = item.SiteID
	respItem["address"] = item.Address
	respItem["latitude"] = item.Latitude
	respItem["longitude"] = item.Longitude
	respItem["altitude"] = item.Altitude

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoFileSystemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoFileSystemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["writeable"] = boolPtrToString(item.Writeable)
		respItem["freespace"] = item.Freespace
		respItem["name"] = item.Name
		respItem["readable"] = boolPtrToString(item.Readable)
		respItem["size"] = item.Size
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_name"] = item.ProfileName
		respItem["discovery_created"] = boolPtrToString(item.DiscoveryCreated)
		respItem["created_by"] = item.CreatedBy
		respItem["primary_endpoint"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpoint(item.PrimaryEndpoint)
		respItem["secondary_endpoint"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpoint(item.SecondaryEndpoint)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpoint(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpoint(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoHTTPHeaders(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoHTTPHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoNeighborLinks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoNeighborLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["local_interface_name"] = item.LocalInterfaceName
		respItem["local_short_interface_name"] = item.LocalShortInterfaceName
		respItem["local_mac_address"] = item.LocalMacAddress
		respItem["remote_interface_name"] = item.RemoteInterfaceName
		respItem["remote_short_interface_name"] = item.RemoteShortInterfaceName
		respItem["remote_mac_address"] = item.RemoteMacAddress
		respItem["remote_device_name"] = item.RemoteDeviceName
		respItem["remote_platform"] = item.RemotePlatform
		respItem["remote_version"] = item.RemoteVersion
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfaces(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["status"] = item.Status
		respItem["mac_address"] = item.MacAddress
		respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv4Address(item.IPv4Address)
		respItem["ipv6_address_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv6AddressList(item.IPv6AddressList)
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv6AddressList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv6AddressList) []interface{} {
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

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["supports_stack_workflows"] = boolPtrToString(item.SupportsStackWorkflows)
	respItem["is_full_ring"] = boolPtrToString(item.IsFullRing)
	respItem["stack_member_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfoStackMemberList(item.StackMemberList)
	respItem["stack_ring_protocol"] = item.StackRingProtocol
	respItem["valid_license_levels"] = item.ValidLicenseLevels
	respItem["total_member_count"] = item.TotalMemberCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfoStackMemberList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfoStackMemberList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["role"] = item.Role
		respItem["mac_address"] = item.MacAddress
		respItem["pid"] = item.Pid
		respItem["license_level"] = item.LicenseLevel
		respItem["license_type"] = item.LicenseType
		respItem["sudi_serial_number"] = item.SudiSerialNumber
		respItem["hardware_version"] = item.HardwareVersion
		respItem["stack_number"] = item.StackNumber
		respItem["software_version"] = item.SoftwareVersion
		respItem["priority"] = item.Priority
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoAAACredentials(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoAAACredentials) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["password"] = item.Password
	respItem["username"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPreWorkflowCliOuputs(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPreWorkflowCliOuputs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["cli"] = item.Cli
		respItem["cli_output"] = item.CliOutput
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoTags(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoTags) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["details"] = item.Details
		respItem["history_task_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfo(item.HistoryTaskInfo)
		respItem["error_flag"] = boolPtrToString(item.ErrorFlag)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoWorkItemList(item.WorkItemList)
	respItem["time_taken"] = item.TimeTaken
	respItem["addn_details"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoAddnDetails(item.AddnDetails)
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoAddnDetails(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoAddnDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParameters(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["top_of_stack_serial_number"] = item.TopOfStackSerialNumber
	respItem["license_level"] = item.LicenseLevel
	respItem["license_type"] = item.LicenseType
	respItem["config_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigList(item.ConfigList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["config_parameters"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigListConfigParameters(item.ConfigParameters)
		respItem["config_id"] = item.ConfigID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigListConfigParameters(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigListConfigParameters) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfig(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["config"] = item.Config

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfigPreview(item *dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfigPreview) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemFailureList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkFailureList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["index"] = item.Index
		respItem["serial_num"] = item.SerialNum
		respItem["id"] = item.ID
		respItem["msg"] = item.Msg
		respItems = append(respItems, respItem)
	}
	return respItems
}
