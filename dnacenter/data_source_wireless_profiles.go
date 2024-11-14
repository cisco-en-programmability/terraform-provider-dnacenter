package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all Wireless Network Profiles

- This data source allows the user to get a Wireless Network Profile by ID
`,

		ReadContext: dataSourceWirelessProfilesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless Profile Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_flex_connect": &schema.Schema{
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"interface_name": &schema.Schema{
										Description: `Interface Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"policy_profile_name": &schema.Schema{
										Description: `Policy Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Profile Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_flex_connect": &schema.Schema{
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"interface_name": &schema.Schema{
										Description: `Interface Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"policy_profile_name": &schema.Schema{
										Description: `Policy Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Profile Name
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

func dataSourceWirelessProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfiles")
		queryParams1 := dnacentersdkgo.GetWirelessProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetWirelessProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessProfiles", err,
				"Failure at GetWirelessProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetWirelessProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfileByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Wireless.GetWirelessProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessProfileByID", err,
				"Failure at GetWirelessProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetWirelessProfileByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfileByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetWirelessProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGetWirelessProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["wireless_profile_name"] = item.WirelessProfileName
		respItem["ssid_details"] = flattenWirelessGetWirelessProfilesItemsSSIDDetails(item.SSIDDetails)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfilesItemsSSIDDetails(items *[]dnacentersdkgo.ResponseWirelessGetWirelessProfilesResponseSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid_name"] = item.SSIDName
		respItem["flex_connect"] = flattenWirelessGetWirelessProfilesItemsSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["wlan_profile_name"] = item.WLANProfileName
		respItem["interface_name"] = item.InterfaceName
		respItem["policy_profile_name"] = item.PolicyProfileName
		respItem["dot11be_profile_id"] = item.Dot11BeProfileID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfilesItemsSSIDDetailsFlexConnect(item *dnacentersdkgo.ResponseWirelessGetWirelessProfilesResponseSSIDDetailsFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_flex_connect"] = boolPtrToString(item.EnableFlexConnect)
	respItem["local_to_vlan"] = item.LocalToVLAN

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetWirelessProfileByIDItem(item *dnacentersdkgo.ResponseWirelessGetWirelessProfileByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["wireless_profile_name"] = item.WirelessProfileName
	respItem["ssid_details"] = flattenWirelessGetWirelessProfileByIDItemSSIDDetails(item.SSIDDetails)
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetWirelessProfileByIDItemSSIDDetails(items *[]dnacentersdkgo.ResponseWirelessGetWirelessProfileByIDResponseSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid_name"] = item.SSIDName
		respItem["flex_connect"] = flattenWirelessGetWirelessProfileByIDItemSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["wlan_profile_name"] = item.WLANProfileName
		respItem["interface_name"] = item.InterfaceName
		respItem["policy_profile_name"] = item.PolicyProfileName
		respItem["dot11be_profile_id"] = item.Dot11BeProfileID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileByIDItemSSIDDetailsFlexConnect(item *dnacentersdkgo.ResponseWirelessGetWirelessProfileByIDResponseSSIDDetailsFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_flex_connect"] = boolPtrToString(item.EnableFlexConnect)
	respItem["local_to_vlan"] = item.LocalToVLAN

	return []map[string]interface{}{
		respItem,
	}

}
