package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessEnterpriseSSID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Gets either one or all the enterprise SSID
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

									"auth_server": &schema.Schema{
										Description: `Auth Server
`,
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

									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"passphrase": &schema.Schema{
										Description: `Passphrase
`,
										Type:     schema.TypeString,
										Computed: true,
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
		log.Printf("[DEBUG] Selected method 1: GetEnterpriseSSID")
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
				"Failure when executing GetEnterpriseSSID", err,
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
		respItems = append(respItems, respItem)
	}
	return respItems
}
