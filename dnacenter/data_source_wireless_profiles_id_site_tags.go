package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfilesIDSiteTags() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This endpoint retrieves a list of all *Site Tags* associated with a specific *Wireless Profile*. This data source
requires the *id* of the *Wireless Profile* to be provided as a path parameter.
`,

		ReadContext: dataSourceWirelessProfilesIDSiteTagsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless profile id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"site_tag_name": &schema.Schema{
				Description: `siteTagName query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_profile_name": &schema.Schema{
							Description: `Ap Profile Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"flex_profile_name": &schema.Schema{
							Description: `Flex Profile Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_ids": &schema.Schema{
							Description: `Site Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"site_tag_id": &schema.Schema{
							Description: `Site Tag Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_tag_name": &schema.Schema{
							Description: `Site Tag Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessProfilesIDSiteTagsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSiteTagName, okSiteTagName := d.GetOk("site_tag_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveAllSiteTagsForAWirelessProfile")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.RetrieveAllSiteTagsForAWirelessProfileQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSiteTagName {
			queryParams1.SiteTagName = vSiteTagName.(string)
		}

		response1, restyResp1, err := client.Wireless.RetrieveAllSiteTagsForAWirelessProfile(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveAllSiteTagsForAWirelessProfile", err,
				"Failure at RetrieveAllSiteTagsForAWirelessProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessRetrieveAllSiteTagsForAWirelessProfileItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveAllSiteTagsForAWirelessProfile response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveAllSiteTagsForAWirelessProfileItems(items *[]dnacentersdkgo.ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_ids"] = item.SiteIDs
		respItem["site_tag_name"] = item.SiteTagName
		respItem["flex_profile_name"] = item.FlexProfileName
		respItem["ap_profile_name"] = item.ApProfileName
		respItem["site_tag_id"] = item.SiteTagID
		respItems = append(respItems, respItem)
	}
	return respItems
}
