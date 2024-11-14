package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnalyticsEndpointsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AI Endpoint Analytics.

- Fetch the total count of endpoints that match the given filter criteria.
`,

		ReadContext: dataSourceAnalyticsEndpointsCountRead,
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
			"operating_system": &schema.Schema{
				Description: `operatingSystem query parameter. Operating system to search for. Partial string is allowed.
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

						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAnalyticsEndpointsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: FetchTheCountOfEndpoints")
		queryParams1 := dnacentersdkgo.FetchTheCountOfEndpointsQueryParams{}

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

		response1, restyResp1, err := client.AIEndpointAnalytics.FetchTheCountOfEndpoints(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 FetchTheCountOfEndpoints", err,
				"Failure at FetchTheCountOfEndpoints, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAIEndpointAnalyticsFetchTheCountOfEndpointsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting FetchTheCountOfEndpoints response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAIEndpointAnalyticsFetchTheCountOfEndpointsItem(item *dnacentersdkgo.ResponseAIEndpointAnalyticsFetchTheCountOfEndpoints) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
