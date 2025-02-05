package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnalyticsEndpoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AI Endpoint Analytics.

- Query the endpoints, optionally using various filter and pagination criteria. 'GET /endpoints/count' API can be used
to find out the total number of endpoints matching the filter criteria.

- Fetches details of the endpoint for the given unique identifier 'epId'.
`,

		ReadContext: dataSourceAnalyticsEndpointsRead,
		Schema: map[string]*schema.Schema{
			"ai_spoofing_trust_level": &schema.Schema{
				Description: `aiSpoofingTrustLevel query parameter. Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"anc_policy": &schema.Schema{
				Description: `ancPolicy query parameter. ANC policy. Only exact match will be returned.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_method": &schema.Schema{
				Description: `authMethod query parameter. Authentication method. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"changed_profile_trust_level": &schema.Schema{
				Description: `changedProfileTrustLevel query parameter. Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"concurrent_mac_trust_level": &schema.Schema{
				Description: `concurrentMacTrustLevel query parameter. Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. Type of device to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_id": &schema.Schema{
				Description: `epId path parameter. Unique identifier for the endpoint.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardware_manufacturer": &schema.Schema{
				Description: `hardwareManufacturer query parameter. Hardware manufacturer to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardware_model": &schema.Schema{
				Description: `hardwareModel query parameter. Hardware model to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"include": &schema.Schema{
				Description: `include query parameter. The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' Include all attributes. 'CDP', 'DHCP', etc. Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' Include ANC policy related details. 'TRUST' Include trust score details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Description: `ip query parameter. IP address to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_blocklist_detected": &schema.Schema{
				Description: `ipBlocklistDetected query parameter. Flag to fetch endpoints hitting IP blocklist or not.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to be fetched. If not provided, 50 records will be fetched by default. Maximum 1000 records can be fetched at a time. Use pagination if more records need to be fetched.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. MAC address to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addresses": &schema.Schema{
				Description: `macAddresses query parameter. List of MAC addresses to filter on. Only exact matches will be returned.
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nat_trust_level": &schema.Schema{
				Description: `natTrustLevel query parameter. Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Record offset to start data fetch at. Offset starts at zero.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"operating_system": &schema.Schema{
				Description: `operatingSystem query parameter. Operating system to search for. Partial string is allowed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order to be used for sorting. Possible values are 'asc', 'desc'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"posture_status": &schema.Schema{
				Description: `postureStatus query parameter. Posture status.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profiling_status": &schema.Schema{
				Description: `profilingStatus query parameter. Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"random_mac": &schema.Schema{
				Description: `randomMac query parameter. Flag to fetch endpoints having randomized MAC or not.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"registered": &schema.Schema{
				Description: `registered query parameter. Flag to fetch manually registered or non-registered endpoints.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested. Possible values are 'macAddress', 'ip'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_score": &schema.Schema{
				Description: `trustScore query parameter. Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"unauth_port_detected": &schema.Schema{
				Description: `unauthPortDetected query parameter. Flag to fetch endpoints exposing unauthorized ports or not.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"weak_cred_detected": &schema.Schema{
				Description: `weakCredDetected query parameter. Flag to fetch endpoints having weak credentials or not.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"anc_policy": &schema.Schema{
							Description: `ANC policy currently applied to the endpoint in ISE.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"attributes": &schema.Schema{
							Description: `Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"device_type": &schema.Schema{
							Description: `Type of the device represented by this endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"duid": &schema.Schema{
							Description: `Unique DUID.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"granular_anc_policy": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name of the granular ANC policy.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nas_ip_address": &schema.Schema{
										Description: `IP address of the network device where endpoint is attached.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"hardware_manufacturer": &schema.Schema{
							Description: `Hardware manufacturer for the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"hardware_model": &schema.Schema{
							Description: `Hardware model of the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Unique identifier for the endpoint.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_probe_collection_timestamp": &schema.Schema{
							Description: `Last probe collection timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of the endpoint.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"operating_system": &schema.Schema{
							Description: `Operating system of the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"random_mac": &schema.Schema{
							Description: `Flag to indicate whether MAC address is a randomized one or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"registered": &schema.Schema{
							Description: `Flag to indicate whether this is a manually registered endpoint or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"trust_data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ai_spoofing_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to AI spoofing.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_method": &schema.Schema{
										Description: `Authentication method.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"changed_profile_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to changing profile labels.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"concurrent_mac_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to concurrent MAC address.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_blocklist_detected": &schema.Schema{
										Description: `Flag to fetch endpoints hitting IP blocklist or not.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nat_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to NAT access.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"posture_status": &schema.Schema{
										Description: `Posture status.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"trust_score": &schema.Schema{
										Description: `Overall trust score of the endpoint.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"unauth_port_detected": &schema.Schema{
										Description: `Flag to fetch endpoints exposing unauthorized ports or not.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"weak_cred_detected": &schema.Schema{
										Description: `Flag to fetch endpoints having weak credentials or not.
`,
										// Type:        schema.TypeBool,
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

func dataSourceAnalyticsEndpointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProfilingStatus, okProfilingStatus := d.GetOk("profiling_status")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vMacAddresses, okMacAddresses := d.GetOk("mac_addresses")
	vIP, okIP := d.GetOk("ip")
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vHardwareManufacturer, okHardwareManufacturer := d.GetOk("hardware_manufacturer")
	vHardwareModel, okHardwareModel := d.GetOk("hardware_model")
	vOperatingSystem, okOperatingSystem := d.GetOk("operating_system")
	vRegistered, okRegistered := d.GetOk("registered")
	vRandomMac, okRandomMac := d.GetOk("random_mac")
	vTrustScore, okTrustScore := d.GetOk("trust_score")
	vAuthMethod, okAuthMethod := d.GetOk("auth_method")
	vPostureStatus, okPostureStatus := d.GetOk("posture_status")
	vAiSpoofingTrustLevel, okAiSpoofingTrustLevel := d.GetOk("ai_spoofing_trust_level")
	vChangedProfileTrustLevel, okChangedProfileTrustLevel := d.GetOk("changed_profile_trust_level")
	vNatTrustLevel, okNatTrustLevel := d.GetOk("nat_trust_level")
	vConcurrentMacTrustLevel, okConcurrentMacTrustLevel := d.GetOk("concurrent_mac_trust_level")
	vIPBlocklistDetected, okIPBlocklistDetected := d.GetOk("ip_blocklist_detected")
	vUnauthPortDetected, okUnauthPortDetected := d.GetOk("unauth_port_detected")
	vWeakCredDetected, okWeakCredDetected := d.GetOk("weak_cred_detected")
	vAncPolicy, okAncPolicy := d.GetOk("anc_policy")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vInclude, okInclude := d.GetOk("include")
	vEpID, okEpID := d.GetOk("ep_id")

	method1 := []bool{okProfilingStatus, okMacAddress, okMacAddresses, okIP, okDeviceType, okHardwareManufacturer, okHardwareModel, okOperatingSystem, okRegistered, okRandomMac, okTrustScore, okAuthMethod, okPostureStatus, okAiSpoofingTrustLevel, okChangedProfileTrustLevel, okNatTrustLevel, okConcurrentMacTrustLevel, okIPBlocklistDetected, okUnauthPortDetected, okWeakCredDetected, okAncPolicy, okLimit, okOffset, okSortBy, okOrder, okInclude}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okEpID, okInclude}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: QueryTheEndpoints")
		queryParams1 := dnacentersdkgo.QueryTheEndpointsQueryParams{}

		if okProfilingStatus {
			queryParams1.ProfilingStatus = vProfilingStatus.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okMacAddresses {
			queryParams1.MacAddresses = interfaceToSliceString(vMacAddresses)
		}
		if okIP {
			queryParams1.IP = vIP.(string)
		}
		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okHardwareManufacturer {
			queryParams1.HardwareManufacturer = vHardwareManufacturer.(string)
		}
		if okHardwareModel {
			queryParams1.HardwareModel = vHardwareModel.(string)
		}
		if okOperatingSystem {
			queryParams1.OperatingSystem = vOperatingSystem.(string)
		}
		if okRegistered {
			queryParams1.Registered = vRegistered.(bool)
		}
		if okRandomMac {
			queryParams1.RandomMac = vRandomMac.(bool)
		}
		if okTrustScore {
			queryParams1.TrustScore = vTrustScore.(string)
		}
		if okAuthMethod {
			queryParams1.AuthMethod = vAuthMethod.(string)
		}
		if okPostureStatus {
			queryParams1.PostureStatus = vPostureStatus.(string)
		}
		if okAiSpoofingTrustLevel {
			queryParams1.AiSpoofingTrustLevel = vAiSpoofingTrustLevel.(string)
		}
		if okChangedProfileTrustLevel {
			queryParams1.ChangedProfileTrustLevel = vChangedProfileTrustLevel.(string)
		}
		if okNatTrustLevel {
			queryParams1.NatTrustLevel = vNatTrustLevel.(string)
		}
		if okConcurrentMacTrustLevel {
			queryParams1.ConcurrentMacTrustLevel = vConcurrentMacTrustLevel.(string)
		}
		if okIPBlocklistDetected {
			queryParams1.IPBlocklistDetected = vIPBlocklistDetected.(bool)
		}
		if okUnauthPortDetected {
			queryParams1.UnauthPortDetected = vUnauthPortDetected.(bool)
		}
		if okWeakCredDetected {
			queryParams1.WeakCredDetected = vWeakCredDetected.(bool)
		}
		if okAncPolicy {
			queryParams1.AncPolicy = vAncPolicy.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okInclude {
			queryParams1.Include = vInclude.(string)
		}

		response1, restyResp1, err := client.AIEndpointAnalytics.QueryTheEndpoints(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 QueryTheEndpoints", err,
				"Failure at QueryTheEndpoints, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetEndpointDetails")
		vvEpID := vEpID.(string)
		queryParams2 := dnacentersdkgo.GetEndpointDetailsQueryParams{}

		if okInclude {
			queryParams2.Include = vInclude.(string)
		}

		response2, restyResp2, err := client.AIEndpointAnalytics.GetEndpointDetails(vvEpID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEndpointDetails", err,
				"Failure at GetEndpointDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenAIEndpointAnalyticsGetEndpointDetailsItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAIEndpointAnalyticsGetEndpointDetailsItem(item *dnacentersdkgo.ResponseAIEndpointAnalyticsGetEndpointDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["duid"] = item.Duid
	respItem["mac_address"] = item.MacAddress
	respItem["device_type"] = item.DeviceType
	respItem["hardware_manufacturer"] = item.HardwareManufacturer
	respItem["hardware_model"] = item.HardwareModel
	respItem["operating_system"] = item.OperatingSystem
	respItem["last_probe_collection_timestamp"] = item.LastProbeCollectionTimestamp
	respItem["random_mac"] = boolPtrToString(item.RandomMac)
	respItem["registered"] = boolPtrToString(item.Registered)
	respItem["attributes"] = flattenAIEndpointAnalyticsGetEndpointDetailsItemAttributes(item.Attributes)
	respItem["trust_data"] = flattenAIEndpointAnalyticsGetEndpointDetailsItemTrustData(item.TrustData)
	respItem["anc_policy"] = item.AncPolicy
	respItem["granular_anc_policy"] = flattenAIEndpointAnalyticsGetEndpointDetailsItemGranularAncPolicy(item.GranularAncPolicy)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAIEndpointAnalyticsGetEndpointDetailsItemAttributes(item *dnacentersdkgo.ResponseAIEndpointAnalyticsGetEndpointDetailsAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenAIEndpointAnalyticsGetEndpointDetailsItemTrustData(item *dnacentersdkgo.ResponseAIEndpointAnalyticsGetEndpointDetailsTrustData) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["trust_score"] = item.TrustScore
	respItem["auth_method"] = item.AuthMethod
	respItem["posture_status"] = item.PostureStatus
	respItem["ai_spoofing_trust_level"] = item.AiSpoofingTrustLevel
	respItem["changed_profile_trust_level"] = item.ChangedProfileTrustLevel
	respItem["nat_trust_level"] = item.NatTrustLevel
	respItem["concurrent_mac_trust_level"] = item.ConcurrentMacTrustLevel
	respItem["ip_blocklist_detected"] = boolPtrToString(item.IPBlocklistDetected)
	respItem["unauth_port_detected"] = boolPtrToString(item.UnauthPortDetected)
	respItem["weak_cred_detected"] = boolPtrToString(item.WeakCredDetected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAIEndpointAnalyticsGetEndpointDetailsItemGranularAncPolicy(items *[]dnacentersdkgo.ResponseAIEndpointAnalyticsGetEndpointDetailsGranularAncPolicy) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["nas_ip_address"] = item.NasIPAddress
		respItems = append(respItems, respItem)
	}
	return respItems
}
