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

func resourcePathTrace() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Path Trace.

- Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to
get results and follow progress.

- Deletes a flow analysis request by its id
`,

		CreateContext: resourcePathTraceCreate,
		ReadContext:   resourcePathTraceRead,
		UpdateContext: resourcePathTraceUpdate,
		DeleteContext: resourcePathTraceDelete,
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

						"detailed_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"acl_trace_calculation": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"acl_trace_calculation_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"last_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_elements": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"egress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"egress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"wireless_lan_controller_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ingress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ingress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"network_elements_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"egress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"wireless_lan_controller_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ingress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_monitor_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"request": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_path": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"create_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"dest_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dest_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"inclusions": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"last_update_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"periodic_refresh": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_port": &schema.Schema{
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
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"control_path": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"dest_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dest_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"flow_analysis_id": &schema.Schema{
							Description: `flowAnalysisId path parameter. Flow analysis request id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"inclusions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"periodic_refresh": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePathTraceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPathTraceInitiateANewPathtrace(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vFlowAnalysisID, okFlowAnalysisID := resourceItem["flow_analysis_id"]
	vvFlowAnalysisID := interfaceToString(vFlowAnalysisID)
	if okFlowAnalysisID && vvFlowAnalysisID != "" {
		getResponse2, _, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["flow_analysis_id"] = vvFlowAnalysisID
			d.SetId(joinResourceID(resourceMap))
			return resourcePathTraceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.PathTrace.RetrivesAllPreviousPathtracesSummary(nil)
		if response2 != nil && err == nil {
			item2, err := searchPathTraceRetrivesAllPreviousPathtracesSummary(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["flow_analysis_id"] = vvFlowAnalysisID
				d.SetId(joinResourceID(resourceMap))
				return resourcePathTraceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.PathTrace.InitiateANewPathtrace(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing InitiateANewPathtrace", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing InitiateANewPathtrace", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["flow_analysis_id"] = vvFlowAnalysisID
	d.SetId(joinResourceID(resourceMap))
	return resourcePathTraceRead(ctx, d, m)
}

func resourcePathTraceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPeriodicRefresh, okPeriodicRefresh := resourceMap["periodic_refresh"]
	vSourceIP, okSourceIP := resourceMap["source_ip"]
	vDestIP, okDestIP := resourceMap["dest_ip"]
	vSourcePort, okSourcePort := resourceMap["source_port"]
	vDestPort, okDestPort := resourceMap["dest_port"]
	vGtCreateTime, okGtCreateTime := resourceMap["gt_create_time"]
	vLtCreateTime, okLtCreateTime := resourceMap["lt_create_time"]
	vProtocol, okProtocol := resourceMap["protocol"]
	vStatus, okStatus := resourceMap["status"]
	vTaskID, okTaskID := resourceMap["task_id"]
	vLastUpdateTime, okLastUpdateTime := resourceMap["last_update_time"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vOrder, okOrder := resourceMap["order"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vFlowAnalysisID, okFlowAnalysisID := resourceMap["flow_analysis_id"]

	method1 := []bool{okPeriodicRefresh, okSourceIP, okDestIP, okSourcePort, okDestPort, okGtCreateTime, okLtCreateTime, okProtocol, okStatus, okTaskID, okLastUpdateTime, okLimit, okOffset, okOrder, okSortBy}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okFlowAnalysisID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RetrivesAllPreviousPathtracesSummary")
		queryParams1 := dnacentersdkgo.RetrivesAllPreviousPathtracesSummaryQueryParams{}

		if okPeriodicRefresh {
			queryParams1.PeriodicRefresh = *stringToBooleanPtr(vPeriodicRefresh)
		}
		if okSourceIP {
			queryParams1.SourceIP = vSourceIP
		}
		if okDestIP {
			queryParams1.DestIP = vDestIP
		}
		if okSourcePort {
			queryParams1.SourcePort = vSourcePort
		}
		if okDestPort {
			queryParams1.DestPort = vDestPort
		}
		if okGtCreateTime {
			queryParams1.GtCreateTime = vGtCreateTime
		}
		if okLtCreateTime {
			queryParams1.LtCreateTime = vLtCreateTime
		}
		if okProtocol {
			queryParams1.Protocol = vProtocol
		}
		if okStatus {
			queryParams1.Status = vStatus
		}
		if okTaskID {
			queryParams1.TaskID = vTaskID
		}
		if okLastUpdateTime {
			queryParams1.LastUpdateTime = vLastUpdateTime
		}
		if okLimit {
			queryParams1.Limit = vLimit
		}
		if okOffset {
			queryParams1.Offset = vOffset
		}
		if okOrder {
			queryParams1.Order = vOrder
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy
		}

		response1, restyResp1, err := client.PathTrace.RetrivesAllPreviousPathtracesSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetrivesAllPreviousPathtracesSummary", err,
				"Failure at RetrivesAllPreviousPathtracesSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenPathTraceRetrivesAllPreviousPathtracesSummaryItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrivesAllPreviousPathtracesSummary search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: RetrievesPreviousPathtrace")
		vvFlowAnalysisID := vFlowAnalysisID

		response2, restyResp2, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetrievesPreviousPathtrace", err,
				"Failure at RetrievesPreviousPathtrace, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPathTraceRetrievesPreviousPathtraceItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesPreviousPathtrace response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourcePathTraceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePathTraceRead(ctx, d, m)
}

func resourcePathTraceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPeriodicRefresh, okPeriodicRefresh := resourceMap["periodic_refresh"]
	vSourceIP, okSourceIP := resourceMap["source_ip"]
	vDestIP, okDestIP := resourceMap["dest_ip"]
	vSourcePort, okSourcePort := resourceMap["source_port"]
	vDestPort, okDestPort := resourceMap["dest_port"]
	vGtCreateTime, okGtCreateTime := resourceMap["gt_create_time"]
	vLtCreateTime, okLtCreateTime := resourceMap["lt_create_time"]
	vProtocol, okProtocol := resourceMap["protocol"]
	vStatus, okStatus := resourceMap["status"]
	vTaskID, okTaskID := resourceMap["task_id"]
	vLastUpdateTime, okLastUpdateTime := resourceMap["last_update_time"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vOrder, okOrder := resourceMap["order"]
	vSortBy, okSortBy := resourceMap["sort_by"]

	queryParams1 := dnacentersdkgo.RetrivesAllPreviousPathtracesSummaryQueryParams
	queryParams1.PeriodicRefresh = *stringToBooleanPtr(vPeriodicRefresh)
	queryParams1.SourceIP = vSourceIP
	queryParams1.DestIP = vDestIP
	queryParams1.SourcePort = vSourcePort
	queryParams1.DestPort = vDestPort
	queryParams1.GtCreateTime = vGtCreateTime
	queryParams1.LtCreateTime = vLtCreateTime
	queryParams1.Protocol = vProtocol
	queryParams1.Status = vStatus
	queryParams1.TaskID = vTaskID
	queryParams1.LastUpdateTime = vLastUpdateTime
	queryParams1.Limit = vLimit
	queryParams1.Offset = vOffset
	queryParams1.Order = vOrder
	queryParams1.SortBy = vSortBy
	item, err := searchPathTraceRetrivesAllPreviousPathtracesSummary(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RetrivesAllPreviousPathtracesSummary", err,
			"Failure at RetrivesAllPreviousPathtracesSummary, unexpected response", ""))
		return diags
	}

	vFlowAnalysisID, okFlowAnalysisID := resourceMap["flow_analysis_id"]

	method1 := []bool{okPeriodicRefresh, okSourceIP, okDestIP, okSourcePort, okDestPort, okGtCreateTime, okLtCreateTime, okProtocol, okStatus, okTaskID, okLastUpdateTime, okLimit, okOffset, okOrder, okSortBy}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okFlowAnalysisID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.PathTrace.RetrivesAllPreviousPathtracesSummary(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsPathTraceRetrivesAllPreviousPathtracesSummary(m, getResp1, nil)
		item1, err := searchPathTraceRetrivesAllPreviousPathtracesSummary(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.PathTrace.DeletesPathtraceByID(vvFlowAnalysisID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesPathtraceByID", err, restyResp1.String(),
				"Failure at DeletesPathtraceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesPathtraceByID", err,
			"Failure at DeletesPathtraceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestPathTraceInitiateANewPathtrace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPathTraceInitiateANewPathtrace {
	request := dnacentersdkgo.RequestPathTraceInitiateANewPathtrace{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".control_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".control_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".control_path")))) {
		request.ControlPath = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_ip")))) {
		request.DestIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_port")))) {
		request.DestPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".inclusions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".inclusions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".inclusions")))) {
		request.Inclusions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".periodic_refresh")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".periodic_refresh")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".periodic_refresh")))) {
		request.PeriodicRefresh = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_ip")))) {
		request.SourceIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_port")))) {
		request.SourcePort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchPathTraceRetrivesAllPreviousPathtracesSummary(m interface{}, queryParams dnacentersdkgo.RetrivesAllPreviousPathtracesSummaryQueryParams) (*dnacentersdkgo.ResponseItemPathTraceRetrivesAllPreviousPathtracesSummary, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemPathTraceRetrivesAllPreviousPathtracesSummary
	var ite *dnacentersdkgo.ResponsePathTraceRetrivesAllPreviousPathtracesSummary
	ite, _, err = client.PathTrace.RetrivesAllPreviousPathtracesSummary(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemPathTraceRetrivesAllPreviousPathtracesSummary
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
