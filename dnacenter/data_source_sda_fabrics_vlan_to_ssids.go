package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricsVLANToSSIDs() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- It will return all vlan to SSID mapping across all the fabric site
`,

		ReadContext: dataSourceSdaFabricsVLANToSSIDsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Return only this many IP Pool to SSID Mapping
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Number of records to skip for pagination
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fabric_id": &schema.Schema{
							Description: `Fabric Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vlan_details": &schema.Schema{
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
													Description: `Name of the SSID.
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
				},
			},
		},
	}
}

func dataSourceSdaFabricsVLANToSSIDsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping")
		queryParams1 := dnacentersdkgo.ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.FabricWireless.ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping", err,
				"Failure at ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItems(items *[]dnacentersdkgo.ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["fabric_id"] = item.FabricID
		respItem["vlan_details"] = flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItemsVLANDetails(item.VLANDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItemsVLANDetails(items *[]dnacentersdkgo.ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["vlan_name"] = item.VLANName
		respItem["ssid_details"] = flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItemsVLANDetailsSSIDDetails(item.SSIDDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingItemsVLANDetailsSSIDDetails(items *[]dnacentersdkgo.ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetailsSSIDDetails) []map[string]interface{} {
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
