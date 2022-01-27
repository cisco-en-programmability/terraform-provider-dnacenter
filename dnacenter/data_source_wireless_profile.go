package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Gets either one or all the wireless network profiles if no name is provided for network-profile.
`,

		ReadContext: dataSourceWirelessProfileRead,
		Schema: map[string]*schema.Schema{
			"profile_name": &schema.Schema{
				Description: `profileName query parameter. Wireless Network Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"profile_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sites": &schema.Schema{
										Description: `array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_fabric": &schema.Schema{
													Description: `true if fabric is enabled else false
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
																Description: `true if flex connect is enabled else false
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"local_to_vlan": &schema.Schema{
																Description: `Local To VLAN ID
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

												"name": &schema.Schema{
													Description: `SSID Name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `SSID Type(enum: Enterprise/Guest)
`,
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
				},
			},
		},
	}
}

func dataSourceWirelessProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileName, okProfileName := d.GetOk("profile_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetWirelessProfile")
		queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}

		if okProfileName {
			queryParams1.ProfileName = vProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetWirelessProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWirelessProfile", err,
				"Failure at GetWirelessProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetWirelessProfileItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetWirelessProfileItems(items *dnacentersdkgo.ResponseWirelessGetWirelessProfile) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_details"] = flattenWirelessGetWirelessProfileItemsProfileDetails(item.ProfileDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileItemsProfileDetails(item *dnacentersdkgo.ResponseItemWirelessGetWirelessProfileProfileDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["sites"] = item.Sites
	respItem["ssid_details"] = flattenWirelessGetWirelessProfileItemsProfileDetailsSSIDDetails(item.SSIDDetails)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetWirelessProfileItemsProfileDetailsSSIDDetails(items *[]dnacentersdkgo.ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["flex_connect"] = flattenWirelessGetWirelessProfileItemsProfileDetailsSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["interface_name"] = item.InterfaceName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileItemsProfileDetailsSSIDDetailsFlexConnect(item *dnacentersdkgo.ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetailsFlexConnect) []map[string]interface{} {
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
