package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaAnycastGateways() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of anycast gateways that match the provided query parameters.
`,

		ReadContext: dataSourceSdaAnycastGatewaysRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the anycast gateway is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. ID of the anycast gateway.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_pool_name": &schema.Schema{
				Description: `ipPoolName query parameter. Name of the IP pool associated with the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"virtual_network_name": &schema.Schema{
				Description: `virtualNetworkName query parameter. Name of the virtual network associated with the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Description: `vlanId query parameter. VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Description: `vlanName query parameter. VLAN name of the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this anycast gateway is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the anycast gateway.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Name of the IP pool associated with the anycast gateway.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_critical_pool": &schema.Schema{
							Description: `Enable/disable critical VLAN (not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_group_based_policy_enforcement_enabled": &schema.Schema{
							Description: `Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_intra_subnet_routing_enabled": &schema.Schema{
							Description: `Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_ip_directed_broadcast": &schema.Schema{
							Description: `Enable/disable IP-directed broadcast (not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_layer2_flooding_enabled": &schema.Schema{
							Description: `Enable/disable layer 2 flooding (not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_multiple_ip_to_mac_addresses": &schema.Schema{
							Description: `Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_supplicant_based_extended_node_onboarding": &schema.Schema{
							Description: `Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_wireless_pool": &schema.Schema{
							Description: `Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pool_type": &schema.Schema{
							Description: `The pool type of the anycast gateway (applicable only to INFRA_VN).
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"security_group_name": &schema.Schema{
							Description: `Name of the associated Security Group (not applicable to INFRA_VN).
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tcp_mss_adjustment": &schema.Schema{
							Description: `TCP maximum segment size adjustment.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"traffic_type": &schema.Schema{
							Description: `The type of traffic the anycast gateway serves.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Name of the layer 3 virtual network associated with the anycast gateway.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `ID of the VLAN of the anycast gateway.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"vlan_name": &schema.Schema{
							Description: `Name of the VLAN of the anycast gateway.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaAnycastGatewaysRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vVirtualNetworkName, okVirtualNetworkName := d.GetOk("virtual_network_name")
	vIPPoolName, okIPPoolName := d.GetOk("ip_pool_name")
	vVLANName, okVLANName := d.GetOk("vlan_name")
	vVLANID, okVLANID := d.GetOk("vlan_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnycastGateways")
		queryParams1 := dnacentersdkgo.GetAnycastGatewaysQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okVirtualNetworkName {
			queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)
		}
		if okIPPoolName {
			queryParams1.IPPoolName = vIPPoolName.(string)
		}
		if okVLANName {
			queryParams1.VLANName = vVLANName.(string)
		}
		if okVLANID {
			queryParams1.VLANID = vVLANID.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetAnycastGateways(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAnycastGateways", err,
				"Failure at GetAnycastGateways, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetAnycastGatewaysItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnycastGateways response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetAnycastGatewaysItems(items *[]dnacentersdkgo.ResponseSdaGetAnycastGatewaysResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["tcp_mss_adjustment"] = item.TCPMssAdjustment
		respItem["vlan_name"] = item.VLANName
		respItem["vlan_id"] = item.VLANID
		respItem["traffic_type"] = item.TrafficType
		respItem["pool_type"] = item.PoolType
		respItem["security_group_name"] = item.SecurityGroupName
		respItem["is_critical_pool"] = boolPtrToString(item.IsCriticalPool)
		respItem["is_layer2_flooding_enabled"] = boolPtrToString(item.IsLayer2FloodingEnabled)
		respItem["is_wireless_pool"] = boolPtrToString(item.IsWirelessPool)
		respItem["is_ip_directed_broadcast"] = boolPtrToString(item.IsIPDirectedBroadcast)
		respItem["is_intra_subnet_routing_enabled"] = boolPtrToString(item.IsIntraSubnetRoutingEnabled)
		respItem["is_multiple_ip_to_mac_addresses"] = boolPtrToString(item.IsMultipleIPToMacAddresses)
		respItem["is_supplicant_based_extended_node_onboarding"] = boolPtrToString(item.IsSupplicantBasedExtendedNodeOnboarding)
		respItem["is_group_based_policy_enforcement_enabled"] = boolPtrToString(item.IsGroupBasedPolicyEnforcementEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}
