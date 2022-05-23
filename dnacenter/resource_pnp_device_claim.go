package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePnpDeviceClaim() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Onboarding (PnP).

- Claim a device to the PnP database.

- Unclaim specified device from PnP database
`,

		CreateContext: resourcePnpDeviceClaimCreate,
		ReadContext:   resourcePnpDeviceClaimRead,
		UpdateContext: resourcePnpDeviceClaimUpdate,
		DeleteContext: resourcePnpDeviceClaimDelete,
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

										Type:     schema.TypeString,
										Computed: true,
									},
									"tags": &schema.Schema{
										Description: `Tags`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_file_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_claim_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
									"device_id": &schema.Schema{
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
									"top_of_stack_serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"file_service_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"populate_inventory": &schema.Schema{

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"workflow_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpDeviceClaimCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPnpDeviceClaimClaimDevice(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vDeviceID := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.device_claim_list"); ok {
			if _, ok := d.GetOk("parameters.0.device_claim_list.0"); ok {
				if v, ok := d.GetOk("parameters.0.device_claim_list.0.device_id"); ok {
					vDeviceID = interfaceToString(v)
				}
			}
		}
	}

	vvID := interfaceToString(vDeviceID)
	if vvID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)
		if err == nil && getResponse2 != nil && getResponse2.DeviceInfo != nil && getResponse2.DeviceInfo.State == "Claimed" {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpDeviceClaimRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.ClaimDevice(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ClaimDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ClaimDevice", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourcePnpDeviceClaimRead(ctx, d, m)
}

func resourcePnpDeviceClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if vID != "" {
		log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
		vvID := vID

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)

		if err != nil || response2 == nil || response2.DeviceInfo != nil && response2.DeviceInfo.State == "Unclaimed" {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenDeviceOnboardingPnpGetDeviceByIDItem(response2)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePnpDeviceClaimUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePnpDeviceClaimRead(ctx, d, m)
}

func resourcePnpDeviceClaimDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	if vID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetDeviceByID(vID)
		if err == nil && getResponse2 == nil && getResponse2.DeviceInfo != nil && getResponse2.DeviceInfo.State == "Unclaimed" {
			d.SetId("")
			return diags
		}
	} else {
		return diags
	}

	//var vvID string
	//var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	request1 := dnacentersdkgo.RequestDeviceOnboardingPnpUnClaimDevice{}
	request1.DeviceIDList = append(request1.DeviceIDList, vID)
	response1, restyResp1, err := client.DeviceOnboardingPnp.UnClaimDevice(&request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing UnClaimDevice", err, restyResp1.String(),
				"Failure at UnClaimDevice, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UnClaimDevice", err,
			"Failure at UnClaimDevice, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return diags
}
func expandRequestPnpDeviceClaimClaimDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpClaimDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_file_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_file_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_file_url")))) {
		request.ConfigFileURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_claim_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_claim_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_claim_list")))) {
		request.DeviceClaimList = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListArray(ctx, key+".device_claim_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_service_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_service_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_service_id")))) {
		request.FileServiceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_url")))) {
		request.ImageURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListArray(ctx, key+".config_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
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

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters{}
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
func searchDeviceOnboardingPnpGetDeviceClaimList2(m interface{}, queryParams dnacentersdkgo.GetDeviceList2QueryParams, vName string) (*dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2
	nResponse, _, err := client.DeviceOnboardingPnp.GetDeviceList2(nil)
	if nResponse == nil || err != nil {
		return foundItem, err
	}
	maxPageSize := len(*nResponse)
	for _, item := range *nResponse {
		if item.DeviceInfo != nil && vName == item.DeviceInfo.Name {
			foundItem = &item
			return foundItem, err
		}
	}
	queryParams.Limit = maxPageSize
	queryParams.Offset = maxPageSize
	nResponse, _, err = client.DeviceOnboardingPnp.GetDeviceList2(&queryParams)
	return foundItem, err
}

//ISSUE Pnp Device- ISSUE waiting for : [{}]  we have: {}
