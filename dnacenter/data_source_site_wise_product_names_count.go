package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteWiseProductNamesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns the count of network device product names for given filters. The default value of *siteId* is global.
`,

		ReadContext: dataSourceSiteWiseProductNamesCountRead,
		Schema: map[string]*schema.Schema{
			"product_name": &schema.Schema{
				Description: `productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site identifier to get the list of all available products under the site. The default value is global site id. See https://developer.cisco.com/docs/dna-center/get-site/ for siteId
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Reports a count, for example, a total count of records in a given resource.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteWiseProductNamesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vProductName, okProductName := d.GetOk("product_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsTheCountOfNetworkDeviceProductNamesForASite")
		queryParams1 := dnacentersdkgo.ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okProductName {
			queryParams1.ProductName = vProductName.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ReturnsTheCountOfNetworkDeviceProductNamesForASite(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsTheCountOfNetworkDeviceProductNamesForASite", err,
				"Failure at ReturnsTheCountOfNetworkDeviceProductNamesForASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsTheCountOfNetworkDeviceProductNamesForASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
