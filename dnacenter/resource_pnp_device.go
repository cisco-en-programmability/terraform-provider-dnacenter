package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

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
						"type_id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
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
							Type:        schema.TypeString, //TEST,
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
													Type:        schema.TypeString, //TEST,
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
																Type:        schema.TypeString, //TEST,
																Computed:    true,
															},
															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString, //TEST,
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
																Type:        schema.TypeString, //TEST,
																Computed:    true,
															},
															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeString, //TEST,
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
										Type:        schema.TypeString, //TEST,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Description: `Device Sudi Serial Nos`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hostname": &schema.Schema{
										Description: `Hostname`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"mac_address": &schema.Schema{
										Description: `Mac Address`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"pid": &schema.Schema{
										Description: `Pid`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"serial_number": &schema.Schema{
										Description: `Serial Number`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"stack": &schema.Schema{
										Description: `Stack`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"stack_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"is_full_ring": &schema.Schema{
													Description: `Is Full Ring`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hardware_version": &schema.Schema{
																Description: `Hardware Version`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"license_level": &schema.Schema{
																Description: `License Level`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"license_type": &schema.Schema{
																Description: `License Type`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pid": &schema.Schema{
																Description: `Pid`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"priority": &schema.Schema{
																Description: `Priority`,
																Type:        schema.TypeFloat,
																Optional:    true,
																Computed:    true,
															},
															"role": &schema.Schema{
																Description: `Role`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"serial_number": &schema.Schema{
																Description: `Serial Number`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"software_version": &schema.Schema{
																Description: `Software Version`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"stack_number": &schema.Schema{
																Description: `Stack Number`,
																Type:        schema.TypeFloat,
																Optional:    true,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"sudi_serial_number": &schema.Schema{
																Description: `Sudi Serial Number`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Description: `Stack Ring Protocol`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"supports_stack_workflows": &schema.Schema{
													Description: `Supports Stack Workflows`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"total_member_count": &schema.Schema{
													Description: `Total Member Count`,
													Type:        schema.TypeFloat,
													Optional:    true,
													Computed:    true,
												},
												"valid_license_levels": &schema.Schema{
													Description: `Valid License Levels`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"sudi_required": &schema.Schema{
										Description: `Is Sudi Required`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"user_mic_numbers": &schema.Schema{
										Description: `User Mic Numbers`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"user_sudi_serial_nos": &schema.Schema{
										Description: `List of Secure Unique Device Identifier (SUDI) serial numbers to perform SUDI authorization, Required if sudiRequired is true.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"workflow_id": &schema.Schema{
										Description: `Workflow Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"workflow_name": &schema.Schema{
										Description: `Workflow Name`,
										Type:        schema.TypeString,
										Optional:    true,
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
				if v, ok := d.GetOk("parameters.0.device_info.0.hostname"); ok {
					vName = interfaceToString(v)
				} else if v, ok := d.GetOk("parameters.0.device_info.0.serial_number"); ok {
					vName = interfaceToString(v)
				}
			}
		}
	}
	log.Printf("[DEBUG] vName: %s", vName)
	log.Printf("[DEBUG] vID: %s", vID)
	vvName := vName
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
		queryParams1 := dnacentersdkgo.GetDeviceListSiteManagementQueryParams{}
		response1, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams1, vvName)
		if err == nil && response1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response1.ID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourcePnpDeviceRead(ctx, d, m)
		}
	}
	log.Printf("[DEBUG] ANTES: %s", responseInterfaceToString(*request1))
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddDevice(request1)
	log.Printf("[DEBUG] DESPUES: %s", vID)
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

	queryParams3 := dnacentersdkgo.GetDeviceListSiteManagementQueryParams{}
	response2, err := searchDeviceOnboardingPnpGetDeviceList2(m, queryParams3, vvName)
	if err != nil {
		diags = append(diags, diagError(
			"Failure when executing searchDeviceOnboardingPnpGetDeviceList2", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = response2.ID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourcePnpDeviceRead(ctx, d, m)
}

func resourcePnpDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
	vvID := vID

	response2, restyResp2, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)

	if err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetDeviceByID response",
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

	vItemID2 := flattenDeviceOnboardingPnpGetDeviceByIDItem(response2)
	if err := d.Set("item", vItemID2); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetDeviceByID response",
			err))
		return diags
	}
	return diags

}

func resourcePnpDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	// NOTE: Consider adding getAllItems and search function to get missing params

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
	request.DeviceInfo = expandRequestPnpDeviceAddDeviceDeviceInfo(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) {
		request.DeviceSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_mic_numbers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_mic_numbers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_mic_numbers")))) {
		request.UserMicNumbers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_name")))) {
		request.WorkflowName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_info")))) {
		request.StackInfo = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supports_stack_workflows")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supports_stack_workflows")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supports_stack_workflows")))) {
		request.SupportsStackWorkflows = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_ring_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_ring_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_ring_protocol")))) {
		request.StackRingProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_license_levels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_license_levels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_license_levels")))) {
		request.ValidLicenseLevels = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_member_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_member_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_member_count")))) {
		request.TotalMemberCount = interfaceToFloat64Ptr(v)
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
	for item_no := range objs {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_serial_number")))) {
		request.SudiSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_version")))) {
		request.HardwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_number")))) {
		request.StackNumber = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPnpDeviceUpdateDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx, key+".device_info.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDeviceOnboardingPnpGetDeviceList2(m interface{}, queryParams dnacentersdkgo.GetDeviceListSiteManagementQueryParams, vName string) (*dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagement, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagement
	queryParams.Offset = 0
	for {
		log.Println("[DEBUG] INSIDE THE LOOP")
		nResponse, _, err := client.DeviceOnboardingPnp.GetDeviceListSiteManagement(&queryParams)
		if err != nil {
			return foundItem, err
		}
		if nResponse == nil || len(*nResponse) == 0 {
			break
		}
		for _, item := range *nResponse {
			if item.DeviceInfo != nil && (vName == item.DeviceInfo.Hostname || vName == item.DeviceInfo.SerialNumber) {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Offset += len(*nResponse)
		queryParams.Limit = len(*nResponse)
	}
	return foundItem, err
}
