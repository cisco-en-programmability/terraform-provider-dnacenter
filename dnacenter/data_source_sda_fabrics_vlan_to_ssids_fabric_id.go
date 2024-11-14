package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricsVLANToSSIDsFabricID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Retrieve the VLANs and SSIDs mapped to the VLAN, within a Fabric Site. The 'fabricId' represents the Fabric ID of a
particular Fabric Site.
`,

		ReadContext: dataSourceSdaFabricsVLANToSSIDsFabricIDRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
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
										Description: `Name of the SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"security_group_tag": &schema.Schema{
										Description: `Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"vlan_name": &schema.Schema{
							Description: `Vlan Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaFabricsVLANToSSIDsFabricIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite")
		vvFabricID := vFabricID.(string)
		queryParams1 := dnacentersdkgo.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.FabricWireless.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(vvFabricID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite", err,
				"Failure at RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteItems(items *[]dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["vlan_name"] = item.VLANName
		respItem["ssid_details"] = flattenFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteItemsSSIDDetails(item.SSIDDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteItemsSSIDDetails(items *[]dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponseSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["security_group_tag"] = item.SecurityGroupTag
		respItems = append(respItems, respItem)
	}
	return respItems
}
