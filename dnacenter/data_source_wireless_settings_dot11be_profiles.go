package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsDot11BeProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all 802.11be Profile(s) configured under Wireless Settings

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
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAll80211BeProfiles")
		queryParams1 := dnacentersdkgo.GetAll80211BeProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetAll80211BeProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAll80211BeProfiles", err,
				"Failure at GetAll80211BeProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetAll80211BeProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAll80211BeProfiles response",
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

func flattenWirelessGetAll80211BeProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGetAll80211BeProfilesResponse) []map[string]interface{} {
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
