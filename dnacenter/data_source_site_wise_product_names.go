package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteWiseProductNames() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Provides network device product names for a site. The default value of *siteId* is global. The response will include
the network device count and image summary.
`,

		ReadContext: dataSourceSiteWiseProductNamesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. The minimum value is 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Description: `productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The unique identifier for the record is the *id*. If there is no supervisor engine involved, the *id* will be the same as the *productNameOrdinal*. However, if the supervisor engine is applicable, the *id* will be in the form of *<productNameOrdinal>-<supervisorProductNameOrdinal>*.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"golden_image_count": &schema.Schema{
										Description: `Count of golden tagged images
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"installed_image_advisor_count": &schema.Schema{
										Description: `Count of advisor on installed images
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"installed_image_count": &schema.Schema{
										Description: `Count of installed images
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"network_device_count": &schema.Schema{
							Description: `Count of network devices
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"product_name": &schema.Schema{
							Description: `Name of product
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"product_name_ordinal": &schema.Schema{
							Description: `Product name ordinal
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"supervisor_product_name": &schema.Schema{
							Description: `Name of the Supervisor Engine Module, supported by the *productName*.                  Example: The *Cisco Catalyst 9404R Switch* chassis is capable of supporting  different supervisor engine modules: the *Cisco Catalyst 9400 Supervisor Engine-1*, the *Cisco Catalyst 9400 Supervisor Engine-1XL*, the *Cisco Catalyst 9400 Supervisor Engine-1XL-Y*, etc.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"supervisor_product_name_ordinal": &schema.Schema{
							Description: `Supervisor Engine Module Ordinal, supported by the *productNameOrdinal*. Example: The *286315691* chassis ordinal is capable of supporting  different supervisor engine module ordinals: *286316172*, *286316710*, *286320394* etc.
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

func dataSourceSiteWiseProductNamesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vProductName, okProductName := d.GetOk("product_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsNetworkDeviceProductNamesForASite")
		queryParams1 := dnacentersdkgo.ReturnsNetworkDeviceProductNamesForASiteQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okProductName {
			queryParams1.ProductName = vProductName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ReturnsNetworkDeviceProductNamesForASite(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsNetworkDeviceProductNamesForASite", err,
				"Failure at ReturnsNetworkDeviceProductNamesForASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsNetworkDeviceProductNamesForASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["product_name_ordinal"] = item.ProductNameOrdinal
	respItem["product_name"] = item.ProductName
	respItem["supervisor_product_name"] = item.SupervisorProductName
	respItem["supervisor_product_name_ordinal"] = item.SupervisorProductNameOrdinal
	respItem["network_device_count"] = item.NetworkDeviceCount
	respItem["image_summary"] = flattenSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteItemImageSummary(item.ImageSummary)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteItemImageSummary(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponseImageSummary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["installed_image_count"] = item.InstalledImageCount
	respItem["golden_image_count"] = item.GoldenImageCount
	respItem["installed_image_advisor_count"] = item.InstalledImageAdvisorCount

	return []map[string]interface{}{
		respItem,
	}

}
