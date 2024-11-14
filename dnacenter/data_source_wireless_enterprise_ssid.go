package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessEnterpriseSSID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Get Enterprise SSID
`,

		ReadContext: dataSourceWirelessEnterpriseSSIDRead,
		Schema: map[string]*schema.Schema{
			"ssid_name": &schema.Schema{
				Description: `ssidName query parameter. Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"group_uuid": &schema.Schema{
							Description: `Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_override": &schema.Schema{
										Description: `Aaa Override
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_server": &schema.Schema{
										Description: `Auth Server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"basic_service_set_client_idle_timeout": &schema.Schema{
										Description: `Basic Service Set ClientIdle Timeout
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"client_exclusion_timeout": &schema.Schema{
										Description: `Client Exclusion Timeout
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"client_rate_limit": &schema.Schema{
										Description: `Client Rate Limit. (in bits per second)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"coverage_hole_detection_enable": &schema.Schema{
										Description: `Coverage Hole Detection Enable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_basic_service_set_max_idle": &schema.Schema{
										Description: `Enable Basic Service Set Max Idle
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_broadcast_ssi_d": &schema.Schema{
										Description: `Enable Broadcast SSID
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_client_exclusion": &schema.Schema{
										Description: `Enable Client Exclusion
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_directed_multicast_service": &schema.Schema{
										Description: `Enable Directed MulticastService
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fast_lane": &schema.Schema{
										Description: `Enable Fast Lane
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_mac_filtering": &schema.Schema{
										Description: `Enable MAC Filtering
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_neighbor_list": &schema.Schema{
										Description: `Enable NeighborList
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_session_time_out": &schema.Schema{
										Description: `Enable Session Time Out
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fast_transition": &schema.Schema{
										Description: `Fast Transition
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_enabled": &schema.Schema{
										Description: `Is Enabled
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_fabric": &schema.Schema{
										Description: `Is Fabric
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mfp_client_protection": &schema.Schema{
										Description: `Mfp Client Protection
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"multi_psk_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"passphrase": &schema.Schema{
													Description: `Passphrase`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"passphrase_type": &schema.Schema{
													Description: `Passphrase Type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"priority": &schema.Schema{
													Description: `Priority
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nas_options": &schema.Schema{
										Description: `Nas Options`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"passphrase": &schema.Schema{
										Description: `Passphrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protected_management_frame": &schema.Schema{
										Description: `Protected Management Frame`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio_policy": &schema.Schema{
										Description: `Radio Policy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"security_level": &schema.Schema{
										Description: `Security Level
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"session_time_out": &schema.Schema{
										Description: `sessionTimeOut
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"traffic_type": &schema.Schema{
										Description: `Traffic Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_type": &schema.Schema{
										Description: `Wlan Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessEnterpriseSSIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSSIDName, okSSIDName := d.GetOk("ssid_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEnterpriseSSID")
		queryParams1 := dnacentersdkgo.GetEnterpriseSSIDQueryParams{}

		if okSSIDName {
			queryParams1.SSIDName = vSSIDName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetEnterpriseSSID(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEnterpriseSSID", err,
				"Failure at GetEnterpriseSSID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetEnterpriseSSIDItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEnterpriseSSID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetEnterpriseSSIDItems(items *dnacentersdkgo.ResponseWirelessGetEnterpriseSSID) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["version"] = item.Version
		respItem["ssid_details"] = flattenWirelessGetEnterpriseSSIDItemsSSIDDetails(item.SSIDDetails)
		respItem["group_uuid"] = item.GroupUUID
		respItem["inherited_group_uuid"] = item.InheritedGroupUUID
		respItem["inherited_group_name"] = item.InheritedGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetEnterpriseSSIDItemsSSIDDetails(items *[]dnacentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["wlan_type"] = item.WLANType
		respItem["enable_fast_lane"] = boolPtrToString(item.EnableFastLane)
		respItem["security_level"] = item.SecurityLevel
		respItem["auth_server"] = item.AuthServer
		respItem["passphrase"] = item.Passphrase
		respItem["traffic_type"] = item.TrafficType
		respItem["enable_mac_filtering"] = boolPtrToString(item.EnableMacFiltering)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["is_fabric"] = boolPtrToString(item.IsFabric)
		respItem["fast_transition"] = item.FastTransition
		respItem["radio_policy"] = item.RadioPolicy
		respItem["enable_broadcast_ssi_d"] = boolPtrToString(item.EnableBroadcastSSID)
		respItem["nas_options"] = item.NasOptions
		respItem["aaa_override"] = boolPtrToString(item.AAAOverride)
		respItem["coverage_hole_detection_enable"] = boolPtrToString(item.CoverageHoleDetectionEnable)
		respItem["protected_management_frame"] = item.ProtectedManagementFrame
		respItem["multi_psk_settings"] = flattenWirelessGetEnterpriseSSIDItemsSSIDDetailsMultipSKSettings(item.MultipSKSettings)
		respItem["client_rate_limit"] = item.ClientRateLimit
		respItem["enable_session_time_out"] = boolPtrToString(item.EnableSessionTimeOut)
		respItem["session_time_out"] = item.SessionTimeOut
		respItem["enable_client_exclusion"] = boolPtrToString(item.EnableClientExclusion)
		respItem["client_exclusion_timeout"] = item.ClientExclusionTimeout
		respItem["enable_basic_service_set_max_idle"] = boolPtrToString(item.EnableBasicServiceSetMaxIDle)
		respItem["basic_service_set_client_idle_timeout"] = item.BasicServiceSetClientIDleTimeout
		respItem["enable_directed_multicast_service"] = boolPtrToString(item.EnableDirectedMulticastService)
		respItem["enable_neighbor_list"] = boolPtrToString(item.EnableNeighborList)
		respItem["mfp_client_protection"] = item.MfpClientProtection
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetEnterpriseSSIDItemsSSIDDetailsMultipSKSettings(items *[]dnacentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["priority"] = item.Priority
		respItem["passphrase_type"] = item.PassphraseType
		respItem["passphrase"] = item.Passphrase
		respItems = append(respItems, respItem)
	}
	return respItems
}
