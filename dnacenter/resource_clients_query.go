package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceClientsQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Clients.

- Retrieves the list of clients by applying complex filters while also supporting aggregate attributes. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml
`,

		CreateContext: resourceClientsQueryCreate,
		ReadContext:   resourceClientsQueryRead,
		DeleteContext: resourceClientsQueryDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"connected_network_device": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connected_network_device_id": &schema.Schema{
													Description: `Connected Network Device Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_network_device_mac": &schema.Schema{
													Description: `Connected Network Device Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_network_device_management_ip": &schema.Schema{
													Description: `Connected Network Device Management Ip`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_network_device_name": &schema.Schema{
													Description: `Connected Network Device Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_network_device_type": &schema.Schema{
													Description: `Connected Network Device Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"duplex_mode": &schema.Schema{
													Description: `Duplex Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"interface_name": &schema.Schema{
													Description: `Interface Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"interface_speed": &schema.Schema{
													Description: `Interface Speed`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"connection": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ap_ethernet_mac": &schema.Schema{
													Description: `Ap Ethernet Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_mac": &schema.Schema{
													Description: `Ap Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_mode": &schema.Schema{
													Description: `Ap Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"auth_type": &schema.Schema{
													Description: `Auth Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"band": &schema.Schema{
													Description: `Band`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"bridge_vmmode": &schema.Schema{
													Description: `Bridge V M Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"channel": &schema.Schema{
													Description: `Channel`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"channel_width": &schema.Schema{
													Description: `Channel Width`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"data_rate": &schema.Schema{
													Description: `Data Rate`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"is_ios_analytics_capable": &schema.Schema{
													Description: `Is Ios Analytics Capable`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"l2_vn": &schema.Schema{
													Description: `L2 Vn`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"l3_vn": &schema.Schema{
													Description: `L3 Vn`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"link_speed": &schema.Schema{
													Description: `Link Speed`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"protocol_capability": &schema.Schema{
													Description: `Protocol Capability`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"radio_id": &schema.Schema{
													Description: `Radio Id`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rssi": &schema.Schema{
													Description: `Rssi`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"security_group_tag": &schema.Schema{
													Description: `Security Group Tag`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"session_duration": &schema.Schema{
													Description: `Session Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"snr": &schema.Schema{
													Description: `Snr`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"ssid": &schema.Schema{
													Description: `Ssid`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"upn_duid": &schema.Schema{
													Description: `Upn Duid`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"upn_id": &schema.Schema{
													Description: `Upn Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"upn_name": &schema.Schema{
													Description: `Upn Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"upn_owner": &schema.Schema{
													Description: `Upn Owner`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"vlan_id": &schema.Schema{
													Description: `Vlan Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"vn_id": &schema.Schema{
													Description: `Vn Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"wlc_id": &schema.Schema{
													Description: `Wlc Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"wlc_name": &schema.Schema{
													Description: `Wlc Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"connection_status": &schema.Schema{
										Description: `Connection Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"form_factor": &schema.Schema{
										Description: `Form Factor`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"health": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connected_score": &schema.Schema{
													Description: `Connected Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"is_link_error_included": &schema.Schema{
													Description: `Is Link Error Included`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"is_rssi_included": &schema.Schema{
													Description: `Is Rssi Included`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"is_snr_included": &schema.Schema{
													Description: `Is Snr Included`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"link_error_percentage_threshold": &schema.Schema{
													Description: `Link Error Percentage Threshold`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"onboarding_score": &schema.Schema{
													Description: `Onboarding Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"overall_score": &schema.Schema{
													Description: `Overall Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rssi_threshold": &schema.Schema{
													Description: `Rssi Threshold`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"snr_threshold": &schema.Schema{
													Description: `Snr Threshold`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv4_address": &schema.Schema{
										Description: `Ipv4 Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv6_addresses": &schema.Schema{
										Description: `Ipv6 Addresses`,
										Type:        schema.TypeList,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_private_mac_address": &schema.Schema{
										Description: `Is Private Mac Address`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"last_updated_time": &schema.Schema{
										Description: `Last Updated Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"latency": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"background": &schema.Schema{
													Description: `Background`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"best_effort": &schema.Schema{
													Description: `Best Effort`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"video": &schema.Schema{
													Description: `Video`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"voice": &schema.Schema{
													Description: `Voice`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Description: `Mac Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"onboarding": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa_failure_reason": &schema.Schema{
													Description: `Aaa Failure Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"aaa_server_ip": &schema.Schema{
													Description: `Aaa Server Ip`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"assoc_done_time": &schema.Schema{
													Description: `Assoc Done Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"assoc_failure_reason": &schema.Schema{
													Description: `Assoc Failure Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"auth_done_time": &schema.Schema{
													Description: `Auth Done Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"avg_assoc_duration": &schema.Schema{
													Description: `Avg Assoc Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"avg_auth_duration": &schema.Schema{
													Description: `Avg Auth Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"avg_dhcp_duration": &schema.Schema{
													Description: `Avg Dhcp Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"avg_run_duration": &schema.Schema{
													Description: `Avg Run Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"dhcp_done_time": &schema.Schema{
													Description: `Dhcp Done Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"dhcp_failure_reason": &schema.Schema{
													Description: `Dhcp Failure Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"dhcp_server_ip": &schema.Schema{
													Description: `Dhcp Server Ip`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"failed_roaming_count": &schema.Schema{
													Description: `Failed Roaming Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"latest_failure_reason": &schema.Schema{
													Description: `Latest Failure Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"max_assoc_duration": &schema.Schema{
													Description: `Max Assoc Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"max_auth_duration": &schema.Schema{
													Description: `Max Auth Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"max_dhcp_duration": &schema.Schema{
													Description: `Max Dhcp Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"max_roaming_duration": &schema.Schema{
													Description: `Max Roaming Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"max_run_duration": &schema.Schema{
													Description: `Max Run Duration`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"onboarding_time": &schema.Schema{
													Description: `Onboarding Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"other_failure_reason": &schema.Schema{
													Description: `Other Failure Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"roaming_time": &schema.Schema{
													Description: `Roaming Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"successful_roaming_count": &schema.Schema{
													Description: `Successful Roaming Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"total_roaming_attempts": &schema.Schema{
													Description: `Total Roaming Attempts`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"os_type": &schema.Schema{
										Description: `Os Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"os_version": &schema.Schema{
										Description: `Os Version`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy_id": &schema.Schema{
										Description: `Site Hierarchy Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"tracked": &schema.Schema{
										Description: `Tracked`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"traffic": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dns_request_count": &schema.Schema{
													Description: `Dns Request Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"dns_response_count": &schema.Schema{
													Description: `Dns Response Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_bytes": &schema.Schema{
													Description: `Rx Bytes`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_link_error_percentage": &schema.Schema{
													Description: `Rx Link Error Percentage`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_packets": &schema.Schema{
													Description: `Rx Packets`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_rate": &schema.Schema{
													Description: `Rx Rate`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_retries": &schema.Schema{
													Description: `Rx Retries`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"rx_retry_percentage": &schema.Schema{
													Description: `Rx Retry Percentage`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_bytes": &schema.Schema{
													Description: `Tx Bytes`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_drop_percentage": &schema.Schema{
													Description: `Tx Drop Percentage`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_drops": &schema.Schema{
													Description: `Tx Drops`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_link_error_percentage": &schema.Schema{
													Description: `Tx Link Error Percentage`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_packets": &schema.Schema{
													Description: `Tx Packets`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"tx_rate": &schema.Schema{
													Description: `Tx Rate`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"usage": &schema.Schema{
													Description: `Usage`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"vendor": &schema.Schema{
										Description: `Vendor`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"offset": &schema.Schema{
										Description: `Offset`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"order": &schema.Schema{
													Description: `Order`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"views": &schema.Schema{
							Description: `Views`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceClientsQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Clients.RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceClientsQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceClientsQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes {
	request := dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters {
	request := []dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters{}
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
		i := expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters {
	request := dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes {
	request := []dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes{}
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
		i := expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes {
	request := dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage {
	request := dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy {
	request := []dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy{}
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
		i := expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsQueryRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy {
	request := dnacentersdkgo.RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItems(items *[]dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["mac_address"] = item.MacAddress
		respItem["type"] = item.Type
		respItem["name"] = item.Name
		respItem["user_id"] = item.UserID
		respItem["username"] = item.Username
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_addresses"] = item.IPv6Addresses
		respItem["vendor"] = item.Vendor
		respItem["os_type"] = item.OsType
		respItem["os_version"] = item.OsVersion
		respItem["form_factor"] = item.FormFactor
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["site_id"] = item.SiteID
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItem["connection_status"] = item.ConnectionStatus
		respItem["tracked"] = item.Tracked
		respItem["is_private_mac_address"] = boolPtrToString(item.IsPrivateMacAddress)
		respItem["health"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsHealth(item.Health)
		respItem["traffic"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsTraffic(item.Traffic)
		respItem["connected_network_device"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsConnectedNetworkDevice(item.ConnectedNetworkDevice)
		respItem["connection"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsConnection(item.Connection)
		respItem["onboarding"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsOnboarding(item.Onboarding)
		respItem["latency"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsLatency(item.Latency)
		respItem["aggregate_attributes"] = flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsHealth(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["overall_score"] = item.OverallScore
	respItem["onboarding_score"] = item.OnboardingScore
	respItem["connected_score"] = item.ConnectedScore
	respItem["link_error_percentage_threshold"] = item.LinkErrorPercentageThreshold
	respItem["is_link_error_included"] = boolPtrToString(item.IsLinkErrorIncluded)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["snr_threshold"] = item.SnrThreshold
	respItem["is_rssi_included"] = boolPtrToString(item.IsRssiIncluded)
	respItem["is_snr_included"] = boolPtrToString(item.IsSnrIncluded)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsTraffic(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseTraffic) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tx_bytes"] = item.TxBytes
	respItem["rx_bytes"] = item.RxBytes
	respItem["usage"] = item.Usage
	respItem["rx_packets"] = item.RxPackets
	respItem["tx_packets"] = item.TxPackets
	respItem["rx_rate"] = item.RxRate
	respItem["tx_rate"] = item.TxRate
	respItem["rx_link_error_percentage"] = item.RxLinkErrorPercentage
	respItem["tx_link_error_percentage"] = item.TxLinkErrorPercentage
	respItem["rx_retries"] = item.RxRetries
	respItem["rx_retry_percentage"] = item.RxRetryPercentage
	respItem["tx_drops"] = item.TxDrops
	respItem["tx_drop_percentage"] = item.TxDropPercentage
	respItem["dns_request_count"] = item.DNSRequestCount
	respItem["dns_response_count"] = item.DNSResponseCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsConnectedNetworkDevice(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnectedNetworkDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connected_network_device_id"] = item.ConnectedNetworkDeviceID
	respItem["connected_network_device_name"] = item.ConnectedNetworkDeviceName
	respItem["connected_network_device_management_ip"] = item.ConnectedNetworkDeviceManagementIP
	respItem["connected_network_device_mac"] = item.ConnectedNetworkDeviceMac
	respItem["connected_network_device_type"] = item.ConnectedNetworkDeviceType
	respItem["interface_name"] = item.InterfaceName
	respItem["interface_speed"] = item.InterfaceSpeed
	respItem["duplex_mode"] = item.DuplexMode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsConnection(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["vlan_id"] = item.VLANID
	respItem["session_duration"] = item.SessionDuration
	respItem["vn_id"] = item.VnID
	respItem["l2_vn"] = item.L2Vn
	respItem["l3_vn"] = item.L3Vn
	respItem["security_group_tag"] = item.SecurityGroupTag
	respItem["link_speed"] = item.LinkSpeed
	respItem["bridge_vmmode"] = item.BridgeVMMode
	respItem["band"] = item.Band
	respItem["ssid"] = item.SSID
	respItem["auth_type"] = item.AuthType
	respItem["wlc_name"] = item.WlcName
	respItem["wlc_id"] = item.WlcID
	respItem["ap_mac"] = item.ApMac
	respItem["ap_ethernet_mac"] = item.ApEthernetMac
	respItem["ap_mode"] = item.ApMode
	respItem["radio_id"] = item.RadioID
	respItem["channel"] = item.Channel
	respItem["channel_width"] = item.ChannelWidth
	respItem["protocol"] = item.Protocol
	respItem["protocol_capability"] = item.ProtocolCapability
	respItem["upn_id"] = item.UpnID
	respItem["upn_name"] = item.UpnName
	respItem["upn_owner"] = item.UpnOwner
	respItem["upn_duid"] = item.UpnDuid
	respItem["rssi"] = item.Rssi
	respItem["snr"] = item.Snr
	respItem["data_rate"] = item.DataRate
	respItem["is_ios_analytics_capable"] = boolPtrToString(item.IsIosAnalyticsCapable)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsOnboarding(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["avg_run_duration"] = item.AvgRunDuration
	respItem["max_run_duration"] = item.MaxRunDuration
	respItem["avg_assoc_duration"] = item.AvgAssocDuration
	respItem["max_assoc_duration"] = item.MaxAssocDuration
	respItem["avg_auth_duration"] = item.AvgAuthDuration
	respItem["max_auth_duration"] = item.MaxAuthDuration
	respItem["avg_dhcp_duration"] = item.AvgDhcpDuration
	respItem["max_dhcp_duration"] = item.MaxDhcpDuration
	respItem["max_roaming_duration"] = item.MaxRoamingDuration
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["dhcp_server_ip"] = item.DhcpServerIP
	respItem["onboarding_time"] = item.OnboardingTime
	respItem["auth_done_time"] = item.AuthDoneTime
	respItem["assoc_done_time"] = item.AssocDoneTime
	respItem["dhcp_done_time"] = item.DhcpDoneTime
	respItem["roaming_time"] = item.RoamingTime
	respItem["failed_roaming_count"] = item.FailedRoamingCount
	respItem["successful_roaming_count"] = item.SuccessfulRoamingCount
	respItem["total_roaming_attempts"] = item.TotalRoamingAttempts
	respItem["assoc_failure_reason"] = item.AssocFailureReason
	respItem["aaa_failure_reason"] = item.AAAFailureReason
	respItem["dhcp_failure_reason"] = item.DhcpFailureReason
	respItem["other_failure_reason"] = item.OtherFailureReason
	respItem["latest_failure_reason"] = item.LatestFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsLatency(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseLatency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["video"] = item.Video
	respItem["voice"] = item.Voice
	respItem["best_effort"] = item.BestEffort
	respItem["background"] = item.Background

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["function"] = item.Function
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
