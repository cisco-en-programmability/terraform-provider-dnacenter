package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsDot11BeProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get 802.11be Profile(s) configured under Wireless Settings

- This data source allows the user to get 802.11be Profile by ID
`,

		ReadContext: dataSourceWirelessSettingsDot11BeProfilesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. 802.11be Profile ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_mu_mimo_down_link": &schema.Schema{
				Description: `isMuMimoDownLink query parameter. MU-MIMO Downlink
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_mu_mimo_up_link": &schema.Schema{
				Description: `isMuMimoUpLink query parameter. MU-MIMO Uplink
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_of_dma_down_link": &schema.Schema{
				Description: `isOfDmaDownLink query parameter. OFDMA Downlink
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_of_dma_multi_ru": &schema.Schema{
				Description: `isOfDmaMultiRu query parameter. OFDMA Multi-RU
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_of_dma_up_link": &schema.Schema{
				Description: `isOfDmaUpLink query parameter. OFDMA Uplink
`,
				Type:     schema.TypeBool,
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

						"default": &schema.Schema{
							Description: `Is 802.11be Profile marked as default in System . (Read only field)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `802.11be Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mu_mimo_down_link": &schema.Schema{
							Description: `MU-MIMO Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mu_mimo_up_link": &schema.Schema{
							Description: `MU-MIMO Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_down_link": &schema.Schema{
							Description: `OFDMA Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_multi_ru": &schema.Schema{
							Description: `OFDMA Multi-RU
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_up_link": &schema.Schema{
							Description: `OFDMA Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `802.11be Profile Name
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

						"default": &schema.Schema{
							Description: `802.11be Profile is marked default or custom (Read only field)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `802.11be Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mu_mimo_down_link": &schema.Schema{
							Description: `MU-MIMO Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mu_mimo_up_link": &schema.Schema{
							Description: `MU-MIMO Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_down_link": &schema.Schema{
							Description: `OFDMA Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_multi_ru": &schema.Schema{
							Description: `OFDMA Multi-RU
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ofdma_up_link": &schema.Schema{
							Description: `OFDMA Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `802.11be Profile Name
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

func dataSourceWirelessSettingsDot11BeProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vProfileName, okProfileName := d.GetOk("profile_name")
	vIsOfDmaDownLink, okIsOfDmaDownLink := d.GetOk("is_of_dma_down_link")
	vIsOfDmaUpLink, okIsOfDmaUpLink := d.GetOk("is_of_dma_up_link")
	vIsMuMimoUpLink, okIsMuMimoUpLink := d.GetOk("is_mu_mimo_up_link")
	vIsMuMimoDownLink, okIsMuMimoDownLink := d.GetOk("is_mu_mimo_down_link")
	vIsOfDmaMultiRu, okIsOfDmaMultiRu := d.GetOk("is_of_dma_multi_ru")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okProfileName, okIsOfDmaDownLink, okIsOfDmaUpLink, okIsMuMimoUpLink, okIsMuMimoDownLink, okIsOfDmaMultiRu}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: Get80211BeProfiles")
		queryParams1 := dnacentersdkgo.Get80211BeProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okProfileName {
			queryParams1.ProfileName = vProfileName.(string)
		}
		if okIsOfDmaDownLink {
			queryParams1.IsOfDmaDownLink = vIsOfDmaDownLink.(bool)
		}
		if okIsOfDmaUpLink {
			queryParams1.IsOfDmaUpLink = vIsOfDmaUpLink.(bool)
		}
		if okIsMuMimoUpLink {
			queryParams1.IsMuMimoUpLink = vIsMuMimoUpLink.(bool)
		}
		if okIsMuMimoDownLink {
			queryParams1.IsMuMimoDownLink = vIsMuMimoDownLink.(bool)
		}
		if okIsOfDmaMultiRu {
			queryParams1.IsOfDmaMultiRu = vIsOfDmaMultiRu.(bool)
		}

		response1, restyResp1, err := client.Wireless.Get80211BeProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 Get80211BeProfiles", err,
				"Failure at Get80211BeProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGet80211BeProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Get80211BeProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: Get80211BeProfileByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Wireless.Get80211BeProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 Get80211BeProfileByID", err,
				"Failure at Get80211BeProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGet80211BeProfileByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Get80211BeProfileByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGet80211BeProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGet80211BeProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["profile_name"] = item.ProfileName
		respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
		respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
		respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
		respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
		respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)
		respItem["default"] = boolPtrToString(item.Default)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGet80211BeProfileByIDItem(item *dnacentersdkgo.ResponseWirelessGet80211BeProfileByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["profile_name"] = item.ProfileName
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)
	respItem["default"] = boolPtrToString(item.Default)
	return []map[string]interface{}{
		respItem,
	}
}
