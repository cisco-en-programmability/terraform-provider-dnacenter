package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBusinessSdaHostonboardingSSIDIPpool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Get SSID to IP Pool Mapping
`,

		ReadContext: dataSourceBusinessSdaHostonboardingSSIDIPpoolRead,
		Schema: map[string]*schema.Schema{
			"site_name_hierarchy": &schema.Schema{
				Description: `siteNameHierarchy query parameter. Site Name Heirarchy
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_name": &schema.Schema{
				Description: `vlanName query parameter. VLAN Name
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"scalable_group_name": &schema.Schema{
										Description: `Scalable Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"vlan_name": &schema.Schema{
							Description: `VLAN Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVLANName := d.Get("vlan_name")
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSSIDToIPPoolMapping")
		queryParams1 := dnacentersdkgo.GetSSIDToIPPoolMappingQueryParams{}

		queryParams1.VLANName = vVLANName.(string)

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.FabricWireless.GetSSIDToIPPoolMapping(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSSIDToIPPoolMapping", err,
				"Failure at GetSSIDToIPPoolMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessGetSSIDToIPPoolMappingItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDToIPPoolMapping response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessGetSSIDToIPPoolMappingItem(item *dnacentersdkgo.ResponseFabricWirelessGetSSIDToIPPoolMapping) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["vlan_name"] = item.VLANName
	respItem["ssid_details"] = flattenFabricWirelessGetSSIDToIPPoolMappingItemSSIDDetails(item.SSIDDetails)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenFabricWirelessGetSSIDToIPPoolMappingItemSSIDDetails(items *[]dnacentersdkgo.ResponseFabricWirelessGetSSIDToIPPoolMappingSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["scalable_group_name"] = item.ScalableGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}
