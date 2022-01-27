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

func resourcePnpDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Onboarding (PnP).

- Adds a device to the PnP database.

- Updates device details specified by device id in PnP database

- Deletes specified device from PnP database
`,

		CreateContext: resourcePnpDeviceCreate,
		ReadContext:   resourcePnpDeviceRead,
		UpdateContext: resourcePnpDeviceUpdate,
		DeleteContext: resourcePnpDeviceDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_credentials": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
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
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ipv6_address_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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
										MaxItems: 1,
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
										Required: true,
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
													MaxItems: 1,
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
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
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
													MaxItems: 1,
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
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
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
										MaxItems: 1,
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
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
										MaxItems: 1,
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
							MaxItems: 1,
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
							MaxItems: 1,
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
							MaxItems: 1,
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

func resourcePnpDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPnpDeviceAddDevice(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.device_info"); ok {
			if _, ok := d.GetOk("parameters.0.device_info.0"); ok {
				if v, ok := d.GetOk("parameters.0.device_info.0.name"); ok {
					vName = interfaceToString(v)
				}
			}
		}
	}
	var vvName string
	vvName = vName
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpDeviceRead(ctx, d, m)
		}
	} else if vName != "" {
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}
		queryParams1.Name = append(queryParams1.Name, vvName)
		response1, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams1, vvName)
		if err == nil && response1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpDeviceRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddDevice(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDevice", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourcePnpDeviceRead(ctx, d, m)
}

func resourcePnpDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if vName != "" && okName {
		log.Printf("[DEBUG] Selected method 1: GetDeviceList2")
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}
		queryParams1.Name = append(queryParams1.Name, vName)
		response1, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams1, vName)
		if err != nil || response1 == nil {
			// log.Printf("[DEBUG] Error: %s%", err)
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDeviceList2 1", err,
			// 	"Failure at GetDeviceList2, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}
		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		vItems1 := flattenDeviceOnboardingPnpGetDeviceList2Item(response1)
		if err := d.Set("item", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceList2 response",
				err))
			return diags
		}
		return diags

	} else if vID != "" && okID {
		log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
		vvID := vID

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDeviceByID", err,
			// 	"Failure at GetDeviceByID, unexpected response", ""))
			// return diags
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

func resourcePnpDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vName, _ := resourceMap["name"]

	var vvID string

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }

	if vID != "" {
		log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
		vvID = vID

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
	} else if vName != "" {
		log.Printf("[DEBUG] Selected method 1: GetDeviceList2")
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}
		queryParams1.Name = append(queryParams1.Name, vName)
		respon1, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams1, vName)
		if err != nil || respon1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceList2", err,
				"Failure at GetDeviceList2, unexpected response", ""))
			return diags
		}
		vvID = respon1.ID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestPnpDeviceUpdateDevice(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DeviceOnboardingPnp.UpdateDevice(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDevice", err, restyResp1.String(),
					"Failure at UpdateDevice, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDevice", err,
				"Failure at UpdateDevice, unexpected response", ""))
			return diags
		}
	}

	return resourcePnpDeviceRead(ctx, d, m)
}

func resourcePnpDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vvID := vID
	vName := resourceMap["name"]
	if vName != "" {
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}
		queryParams1.Name = append(queryParams1.Name, vName)
		response1, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams1, vName)
		if err != nil || response1 == nil {
			// Assume that element it is already gone
			return diags
		}
		vvID = response1.ID
	}
	if vID != "" {
		response2, _, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)
		if err != nil || response2 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceOnboardingPnp.DeleteDeviceByIDFromPnp(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceByIDFromPnp", err, restyResp1.String(),
				"Failure at DeleteDeviceByIDFromPnp, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceByIDFromPnp", err,
			"Failure at DeleteDeviceByIDFromPnp, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPnpDeviceAddDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceAddDeviceDeviceInfo(ctx, key+".device_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_summary_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_summary_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_summary_list")))) {
		request.RunSummaryList = expandRequestPnpDeviceAddDeviceRunSummaryListArray(ctx, key+".run_summary_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_reset_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_reset_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_reset_workflow")))) {
		request.SystemResetWorkflow = expandRequestPnpDeviceAddDeviceSystemResetWorkflow(ctx, key+".system_reset_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_workflow")))) {
		request.SystemWorkflow = expandRequestPnpDeviceAddDeviceSystemWorkflow(ctx, key+".system_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow")))) {
		request.Workflow = expandRequestPnpDeviceAddDeviceWorkflow(ctx, key+".workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_parameters")))) {
		request.WorkflowParameters = expandRequestPnpDeviceAddDeviceWorkflowParameters(ctx, key+".workflow_parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpDeviceAddDeviceDeviceInfoAAACredentials(ctx, key+".aaa_credentials.0", d)
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
		request.FileSystemList = expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemListArray(ctx, key+".file_system_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_contact")))) {
		request.FirstContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_headers")))) {
		request.HTTPHeaders = expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeadersArray(ctx, key+".http_headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_file")))) {
		request.ImageFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_version")))) {
		request.ImageVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interfaces")))) {
		request.IPInterfaces = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesArray(ctx, key+".ip_interfaces", d)
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
		request.Location = expandRequestPnpDeviceAddDeviceDeviceInfoLocation(ctx, key+".location.0", d)
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
		request.NeighborLinks = expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinksArray(ctx, key+".neighbor_links", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".onb_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".onb_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".onb_state")))) {
		request.OnbState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pnp_profile_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pnp_profile_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pnp_profile_list")))) {
		request.PnpProfileList = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListArray(ctx, key+".pnp_profile_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_workflow_cli_ouputs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) {
		request.PreWorkflowCliOuputs = expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx, key+".pre_workflow_cli_ouputs", d)
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
		request.StackInfo = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestPnpDeviceAddDeviceDeviceInfoTags(ctx, key+".tags.0", d)
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

func expandRequestPnpDeviceAddDeviceDeviceInfoAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeaders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_list")))) {
		request.IPv6AddressList = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx, key+".ipv6_address_list", d)
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

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoLocation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".created_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".created_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".created_by")))) {
		request.CreatedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_created")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_created")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_created")))) {
		request.DiscoveryCreated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_endpoint")))) {
		request.PrimaryEndpoint = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx, key+".primary_endpoint.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_endpoint")))) {
		request.SecondaryEndpoint = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx, key+".secondary_endpoint.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
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

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
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

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
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

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList{}
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

func expandRequestPnpDeviceAddDeviceDeviceInfoTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".details")))) {
		request.Details = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_flag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_flag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_flag")))) {
		request.ErrorFlag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".history_task_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".history_task_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".history_task_info")))) {
		request.HistoryTaskInfo = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfo(ctx, key+".history_task_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp")))) {
		request.Timestamp = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_details")))) {
		request.AddnDetails = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx, key+".addn_details", d)
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceSystemResetWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceSystemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceSystemWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceSystemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceWorkflowParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListArray(ctx, key+".config_list", d)
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

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowParametersConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters{}
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

func expandRequestPnpDeviceUpdateDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx, key+".device_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_summary_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_summary_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_summary_list")))) {
		request.RunSummaryList = expandRequestPnpDeviceUpdateDeviceRunSummaryListArray(ctx, key+".run_summary_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_reset_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_reset_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_reset_workflow")))) {
		request.SystemResetWorkflow = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflow(ctx, key+".system_reset_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_workflow")))) {
		request.SystemWorkflow = expandRequestPnpDeviceUpdateDeviceSystemWorkflow(ctx, key+".system_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow")))) {
		request.Workflow = expandRequestPnpDeviceUpdateDeviceWorkflow(ctx, key+".workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_parameters")))) {
		request.WorkflowParameters = expandRequestPnpDeviceUpdateDeviceWorkflowParameters(ctx, key+".workflow_parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpDeviceUpdateDeviceDeviceInfoAAACredentials(ctx, key+".aaa_credentials.0", d)
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
		request.FileSystemList = expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemListArray(ctx, key+".file_system_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_contact")))) {
		request.FirstContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_headers")))) {
		request.HTTPHeaders = expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeadersArray(ctx, key+".http_headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_file")))) {
		request.ImageFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_version")))) {
		request.ImageVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interfaces")))) {
		request.IPInterfaces = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesArray(ctx, key+".ip_interfaces", d)
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
		request.Location = expandRequestPnpDeviceUpdateDeviceDeviceInfoLocation(ctx, key+".location.0", d)
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
		request.NeighborLinks = expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinksArray(ctx, key+".neighbor_links", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".onb_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".onb_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".onb_state")))) {
		request.OnbState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pnp_profile_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pnp_profile_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pnp_profile_list")))) {
		request.PnpProfileList = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListArray(ctx, key+".pnp_profile_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_workflow_cli_ouputs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) {
		request.PreWorkflowCliOuputs = expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx, key+".pre_workflow_cli_ouputs", d)
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
		request.StackInfo = expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestPnpDeviceUpdateDeviceDeviceInfoTags(ctx, key+".tags.0", d)
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeaders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_list")))) {
		request.IPv6AddressList = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx, key+".ipv6_address_list", d)
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoLocation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".created_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".created_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".created_by")))) {
		request.CreatedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_created")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_created")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_created")))) {
		request.DiscoveryCreated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_endpoint")))) {
		request.PrimaryEndpoint = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx, key+".primary_endpoint.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_endpoint")))) {
		request.SecondaryEndpoint = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx, key+".secondary_endpoint.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList{}
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

func expandRequestPnpDeviceUpdateDeviceDeviceInfoTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".details")))) {
		request.Details = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_flag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_flag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_flag")))) {
		request.ErrorFlag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".history_task_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".history_task_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".history_task_info")))) {
		request.HistoryTaskInfo = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfo(ctx, key+".history_task_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp")))) {
		request.Timestamp = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_details")))) {
		request.AddnDetails = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx, key+".addn_details", d)
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceSystemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceWorkflowParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListArray(ctx, key+".config_list", d)
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

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters{}
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

func searchDeviceOnboardingPnpGetDeviceList2(m interface{}, queryParams dnacentersdkgo.GetDeviceList2QueryParams, vName string) (*dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceList2, error) {
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
