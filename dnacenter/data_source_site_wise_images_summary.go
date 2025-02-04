package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteWiseImagesSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns aggregate counts of network device product names, golden and non-golden tagged products, imported images,
golden images tagged, and advisor for a specific site provide, the default value of *siteId* is set to global.
`,

		ReadContext: dataSourceSiteWiseImagesSummaryRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site identifier to get the aggreagte counts products under the site. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for *siteId*
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"golden_image_count": &schema.Schema{
							Description: `Count of images marked as golden
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"imported_image_count": &schema.Schema{
							Description: `Count of images imported
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"installed_image_advisor_count": &schema.Schema{
							Description: `Advisor count of installed images
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"installed_image_count": &schema.Schema{
							Description: `Count of installed images
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"non_golden_image_count": &schema.Schema{
							Description: `Count of images not marked as golden
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"product_count": &schema.Schema{
							Description: `Count of Network device product name
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"products_with_golden_count": &schema.Schema{
							Description: `Count of Network device product name marked as golden
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"products_without_golden_count": &schema.Schema{
							Description: `Count of Network device product name not marked as golden
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteWiseImagesSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsTheImageSummaryForTheGivenSite")
		queryParams1 := dnacentersdkgo.ReturnsTheImageSummaryForTheGivenSiteQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ReturnsTheImageSummaryForTheGivenSite(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsTheImageSummaryForTheGivenSite", err,
				"Failure at ReturnsTheImageSummaryForTheGivenSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsTheImageSummaryForTheGivenSite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSiteItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["imported_image_count"] = item.ImportedImageCount
	respItem["installed_image_count"] = item.InstalledImageCount
	respItem["golden_image_count"] = item.GoldenImageCount
	respItem["non_golden_image_count"] = item.NonGoldenImageCount
	respItem["installed_image_advisor_count"] = item.InstalledImageAdvisorCount
	respItem["product_count"] = item.ProductCount
	respItem["products_with_golden_count"] = item.ProductsWithGoldenCount
	respItem["products_without_golden_count"] = item.ProductsWithoutGoldenCount
	return []map[string]interface{}{
		respItem,
	}
}
