package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns list of devices based on filter crieteria. If a limit is not specified, it will default to return 50 devices.
Pagination and sorting are also supported by this endpoint

- Returns device details specified by device id
`,

		ReadContext: dataSourcePnpDeviceRead,
		Schema: map[string]*schema.Schema{
			"cm_state": &schema.Schema{
				Description: `cmState query parameter. Device Connection Manager State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hostname": &schema.Schema{
				Description: `hostname query parameter. Device Hostname
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_contact": &schema.Schema{
				Description: `lastContact query parameter. Device Has Contacted lastContact > 0
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
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Device Mac Address
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Device Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Index of first result
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"onb_state": &schema.Schema{
				Description: `onbState query parameter. Device Onboarding State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pid": &schema.Schema{
				Description: `pid query parameter. Device ProductId
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_id": &schema.Schema{
				Description: `projectId query parameter. Device Project Id
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_name": &schema.Schema{
				Description: `projectName query parameter. Device Project Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Device Serial Number
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"site_name": &schema.Schema{
				Description: `siteName query parameter. Device Site Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"smart_account_id": &schema.Schema{
				Description: `smartAccountId query parameter. Device Smart Account
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort": &schema.Schema{
				Description: `sort query parameter. Comma seperated list of fields to sort on
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Description: `source query parameter. Device Source
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Description: `state query parameter. Device State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_account_id": &schema.Schema{
				Description: `virtualAccountId query parameter. Device Virtual Account
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_id": &schema.Schema{
				Description: `workflowId query parameter. Device Workflow Id
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_name": &schema.Schema{
				Description: `workflowName query parameter. Device Workflow Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
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
							Type:        schema.TypeString,
							Computed:    true,
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
													Type:        schema.TypeString,
													Computed:    true,
												},

												"ipv6_address_list": &schema.Schema{
													Description: `Ipv6 Address List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString,
																Computed:    true,
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
										Type:        schema.TypeString,
										Computed:    true,
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

									"type_id": &schema.Schema{
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

									"type_id": &schema.Schema{
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

									"type_id": &schema.Schema{
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

			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
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
							Type:        schema.TypeString,
							Computed:    true,
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
													Type:        schema.TypeString,
													Computed:    true,
												},

												"ipv6_address_list": &schema.Schema{
													Description: `Ipv6 Address List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString,
																Computed:    true,
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
										Type:        schema.TypeString,
										Computed:    true,
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

									"type_id": &schema.Schema{
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

									"type_id": &schema.Schema{
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

									"type_id": &schema.Schema{
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
	}
}

func dataSourcePnpDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSort, okSort := d.GetOk("sort")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vState, okState := d.GetOk("state")
	vOnbState, okOnbState := d.GetOk("onb_state")
	vCmState, okCmState := d.GetOk("cm_state")
	vName, okName := d.GetOk("name")
	vPid, okPid := d.GetOk("pid")
	vSource, okSource := d.GetOk("source")
	vProjectID, okProjectID := d.GetOk("project_id")
	vWorkflowID, okWorkflowID := d.GetOk("workflow_id")
	vProjectName, okProjectName := d.GetOk("project_name")
	vWorkflowName, okWorkflowName := d.GetOk("workflow_name")
	vSmartAccountID, okSmartAccountID := d.GetOk("smart_account_id")
	vVirtualAccountID, okVirtualAccountID := d.GetOk("virtual_account_id")
	vLastContact, okLastContact := d.GetOk("last_contact")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vHostname, okHostname := d.GetOk("hostname")
	vSiteName, okSiteName := d.GetOk("site_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okSerialNumber, okState, okOnbState, okCmState, okName, okPid, okSource, okProjectID, okWorkflowID, okProjectName, okWorkflowName, okSmartAccountID, okVirtualAccountID, okLastContact, okMacAddress, okHostname, okSiteName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceList2")
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okSort {
			queryParams1.Sort = interfaceToSliceString(vSort)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
		}
		if okState {
			queryParams1.State = interfaceToSliceString(vState)
		}
		if okOnbState {
			queryParams1.OnbState = interfaceToSliceString(vOnbState)
		}
		if okCmState {
			queryParams1.CmState = interfaceToSliceString(vCmState)
		}
		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}
		if okPid {
			queryParams1.Pid = interfaceToSliceString(vPid)
		}
		if okSource {
			queryParams1.Source = interfaceToSliceString(vSource)
		}
		if okProjectID {
			queryParams1.ProjectID = interfaceToSliceString(vProjectID)
		}
		if okWorkflowID {
			queryParams1.WorkflowID = interfaceToSliceString(vWorkflowID)
		}
		if okProjectName {
			queryParams1.ProjectName = interfaceToSliceString(vProjectName)
		}
		if okWorkflowName {
			queryParams1.WorkflowName = interfaceToSliceString(vWorkflowName)
		}
		if okSmartAccountID {
			queryParams1.SmartAccountID = interfaceToSliceString(vSmartAccountID)
		}
		if okVirtualAccountID {
			queryParams1.VirtualAccountID = interfaceToSliceString(vVirtualAccountID)
		}
		if okLastContact {
			queryParams1.LastContact = vLastContact.(bool)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okHostname {
			queryParams1.Hostname = vHostname.(string)
		}
		if okSiteName {
			queryParams1.SiteName = vSiteName.(string)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetDeviceList2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceList2", err,
				"Failure at GetDeviceList2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetDeviceList2Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceList2 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceByID", err,
				"Failure at GetDeviceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceOnboardingPnpGetDeviceByIDItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetDeviceList2Items(items *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceList2) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_info"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfo(item.DeviceInfo)
		respItem["system_reset_workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflow(item.SystemResetWorkflow)
		respItem["system_workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflow(item.SystemWorkflow)
		respItem["workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflow(item.Workflow)
		respItem["run_summary_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryList(item.RunSummaryList)
		respItem["workflow_parameters"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParameters(item.WorkflowParameters)
		respItem["day_zero_config"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfig(item.DayZeroConfig)
		respItem["day_zero_config_preview"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfigPreview(item.DayZeroConfigPreview)
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2Item(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_info"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfo(item.DeviceInfo)
	respItem["system_reset_workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflow(item.SystemResetWorkflow)
	respItem["system_workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflow(item.SystemWorkflow)
	respItem["workflow"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflow(item.Workflow)
	respItem["run_summary_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryList(item.RunSummaryList)
	respItem["workflow_parameters"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParameters(item.WorkflowParameters)
	respItem["day_zero_config"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfig(item.DayZeroConfig)
	respItem["day_zero_config_preview"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfigPreview(item.DayZeroConfigPreview)
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfo(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["source"] = item.Source
	respItem["serial_number"] = item.SerialNumber
	respItem["stack"] = boolPtrToString(item.Stack)
	respItem["mode"] = item.Mode
	respItem["state"] = item.State
	respItem["location"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoLocation(item.Location)
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
	respItem["file_system_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoFileSystemList(item.FileSystemList)
	respItem["pnp_profile_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileList(item.PnpProfileList)
	respItem["image_file"] = item.ImageFile
	respItem["http_headers"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoHTTPHeaders(item.HTTPHeaders)
	respItem["neighbor_links"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoNeighborLinks(item.NeighborLinks)
	respItem["last_sync_time"] = item.LastSyncTime
	respItem["ip_interfaces"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfaces(item.IPInterfaces)
	respItem["hostname"] = item.Hostname
	respItem["auth_status"] = item.AuthStatus
	respItem["stack_info"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoStackInfo(item.StackInfo)
	respItem["reload_requested"] = boolPtrToString(item.ReloadRequested)
	respItem["added_on"] = item.AddedOn
	respItem["site_id"] = item.SiteID
	respItem["aaa_credentials"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoAAACredentials(item.AAACredentials)
	respItem["user_mic_numbers"] = item.UserMicNumbers
	respItem["user_sudi_serial_nos"] = item.UserSudiSerialNos
	respItem["addn_mac_addrs"] = item.AddnMacAddrs
	respItem["pre_workflow_cli_ouputs"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPreWorkflowCliOuputs(item.PreWorkflowCliOuputs)
	respItem["tags"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoTags(item.Tags)
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoLocation(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoLocation) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoFileSystemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoFileSystemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_name"] = item.ProfileName
		respItem["discovery_created"] = boolPtrToString(item.DiscoveryCreated)
		respItem["created_by"] = item.CreatedBy
		respItem["primary_endpoint"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpoint(item.PrimaryEndpoint)
		respItem["secondary_endpoint"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpoint(item.SecondaryEndpoint)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpoint(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpoint(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoHTTPHeaders(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoHTTPHeaders) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoNeighborLinks(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoNeighborLinks) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfaces(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["status"] = item.Status
		respItem["mac_address"] = item.MacAddress
		respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfacesIPv4Address(item.IPv4Address)
		respItem["ipv6_address_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfacesIPv6AddressList(item.IPv6AddressList)
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfacesIPv4Address(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoIPInterfacesIPv6AddressList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv6AddressList) []interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoStackInfo(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["supports_stack_workflows"] = boolPtrToString(item.SupportsStackWorkflows)
	respItem["is_full_ring"] = boolPtrToString(item.IsFullRing)
	respItem["stack_member_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoStackInfoStackMemberList(item.StackMemberList)
	respItem["stack_ring_protocol"] = item.StackRingProtocol
	respItem["valid_license_levels"] = item.ValidLicenseLevels
	respItem["total_member_count"] = item.TotalMemberCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoStackInfoStackMemberList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfoStackMemberList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoAAACredentials(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoAAACredentials) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoPreWorkflowCliOuputs(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPreWorkflowCliOuputs) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDeviceInfoTags(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoTags) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflow(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflowTasks(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemResetWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflow(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflowTasks(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsSystemWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflow(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2Workflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowTasks(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["details"] = item.Details
		respItem["history_task_info"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfo(item.HistoryTaskInfo)
		respItem["error_flag"] = boolPtrToString(item.ErrorFlag)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfo(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfoWorkItemList(item.WorkItemList)
	respItem["time_taken"] = item.TimeTaken
	respItem["addn_details"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfoAddnDetails(item.AddnDetails)
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfoWorkItemList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsRunSummaryListHistoryTaskInfoAddnDetails(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoAddnDetails) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParameters(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["top_of_stack_serial_number"] = item.TopOfStackSerialNumber
	respItem["license_level"] = item.LicenseLevel
	respItem["license_type"] = item.LicenseType
	respItem["config_list"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParametersConfigList(item.ConfigList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParametersConfigList(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["config_parameters"] = flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParametersConfigListConfigParameters(item.ConfigParameters)
		respItem["config_id"] = item.ConfigID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsWorkflowParametersConfigListConfigParameters(items *[]dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigListConfigParameters) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfig(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["config"] = item.Config

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceList2ItemsDayZeroConfigPreview(item *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfigPreview) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["device_info"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfo(item.DeviceInfo)
	respItem["system_reset_workflow"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflow(item.SystemResetWorkflow)
	respItem["system_workflow"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflow(item.SystemWorkflow)
	respItem["workflow"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflow(item.Workflow)
	respItem["run_summary_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryList(item.RunSummaryList)
	respItem["workflow_parameters"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParameters(item.WorkflowParameters)
	respItem["day_zero_config"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDayZeroConfig(item.DayZeroConfig)
	respItem["day_zero_config_preview"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDayZeroConfigPreview(item.DayZeroConfigPreview)
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["source"] = item.Source
	respItem["serial_number"] = item.SerialNumber
	respItem["stack"] = boolPtrToString(item.Stack)
	respItem["mode"] = item.Mode
	respItem["state"] = item.State
	respItem["location"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoLocation(item.Location)
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
	respItem["file_system_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoFileSystemList(item.FileSystemList)
	respItem["pnp_profile_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileList(item.PnpProfileList)
	respItem["image_file"] = item.ImageFile
	respItem["http_headers"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoHTTPHeaders(item.HTTPHeaders)
	respItem["neighbor_links"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoNeighborLinks(item.NeighborLinks)
	respItem["last_sync_time"] = item.LastSyncTime
	respItem["ip_interfaces"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfaces(item.IPInterfaces)
	respItem["hostname"] = item.Hostname
	respItem["auth_status"] = item.AuthStatus
	respItem["stack_info"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoStackInfo(item.StackInfo)
	respItem["reload_requested"] = boolPtrToString(item.ReloadRequested)
	respItem["added_on"] = item.AddedOn
	respItem["site_id"] = item.SiteID
	respItem["aaa_credentials"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoAAACredentials(item.AAACredentials)
	respItem["user_mic_numbers"] = item.UserMicNumbers
	respItem["user_sudi_serial_nos"] = item.UserSudiSerialNos
	respItem["addn_mac_addrs"] = item.AddnMacAddrs
	respItem["pre_workflow_cli_ouputs"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPreWorkflowCliOuputs(item.PreWorkflowCliOuputs)
	respItem["tags"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoTags(item.Tags)
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoLocation(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoLocation) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoFileSystemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoFileSystemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_name"] = item.ProfileName
		respItem["discovery_created"] = boolPtrToString(item.DiscoveryCreated)
		respItem["created_by"] = item.CreatedBy
		respItem["primary_endpoint"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpoint(item.PrimaryEndpoint)
		respItem["secondary_endpoint"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpoint(item.SecondaryEndpoint)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpoint(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpoint(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoHTTPHeaders(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoHTTPHeaders) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoNeighborLinks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoNeighborLinks) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfaces(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["status"] = item.Status
		respItem["mac_address"] = item.MacAddress
		respItem["ipv4_address"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfacesIPv4Address(item.IPv4Address)
		respItem["ipv6_address_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfacesIPv6AddressList(item.IPv6AddressList)
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfacesIPv4Address(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoIPInterfacesIPv6AddressList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv6AddressList) []interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoStackInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["supports_stack_workflows"] = boolPtrToString(item.SupportsStackWorkflows)
	respItem["is_full_ring"] = boolPtrToString(item.IsFullRing)
	respItem["stack_member_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoStackInfoStackMemberList(item.StackMemberList)
	respItem["stack_ring_protocol"] = item.StackRingProtocol
	respItem["valid_license_levels"] = item.ValidLicenseLevels
	respItem["total_member_count"] = item.TotalMemberCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoStackInfoStackMemberList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfoStackMemberList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoAAACredentials(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoAAACredentials) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoPreWorkflowCliOuputs(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPreWorkflowCliOuputs) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDeviceInfoTags(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoTags) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemResetWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemSystemWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflow(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflow) []map[string]interface{} {
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
	respItem["tasks"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowTasks(item.Tasks)
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowTasks(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasks) []map[string]interface{} {
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
		respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowTasksWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasksWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["details"] = item.Details
		respItem["history_task_info"] = flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfo(item.HistoryTaskInfo)
		respItem["error_flag"] = boolPtrToString(item.ErrorFlag)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfo(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfoWorkItemList(item.WorkItemList)
	respItem["time_taken"] = item.TimeTaken
	respItem["addn_details"] = flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfoAddnDetails(item.AddnDetails)
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfoWorkItemList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoWorkItemList) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemRunSummaryListHistoryTaskInfoAddnDetails(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoAddnDetails) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParameters(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["top_of_stack_serial_number"] = item.TopOfStackSerialNumber
	respItem["license_level"] = item.LicenseLevel
	respItem["license_type"] = item.LicenseType
	respItem["config_list"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParametersConfigList(item.ConfigList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParametersConfigList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["config_parameters"] = flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParametersConfigListConfigParameters(item.ConfigParameters)
		respItem["config_id"] = item.ConfigID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceByIDItemWorkflowParametersConfigListConfigParameters(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigListConfigParameters) []map[string]interface{} {
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

func flattenDeviceOnboardingPnpGetDeviceByIDItemDayZeroConfig(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["config"] = item.Config

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceByIDItemDayZeroConfigPreview(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfigPreview) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
