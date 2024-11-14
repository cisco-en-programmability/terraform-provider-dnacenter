package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImagesSiteWiseProductNamesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns count of assigned network device product for a given image identifier. Refer */dna/intent/api/v1/images* API
for obtaining *imageId*
`,

		ReadContext: dataSourceImagesSiteWiseProductNamesCountRead,
		Schema: map[string]*schema.Schema{
			"assigned": &schema.Schema{
				Description: `assigned query parameter. Filter with the assigned/unassigned, *ASSIGNED* option will filter network device products that are associated with the given image. The *NOT_ASSIGNED* option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": &schema.Schema{
				Description: `imageId path parameter. Software image identifier. Refer */dna/intent/api/v/images* API for obtaining *imageId*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"product_id": &schema.Schema{
				Description: `productId query parameter. Filter with product ID (PID)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Description: `productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommended": &schema.Schema{
				Description: `recommended query parameter. Filter with recommended source. If *CISCO* then the network device product assigned was recommended by Cisco and *USER* then the user has manually assigned. Available values : CISCO, USER
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

func dataSourceImagesSiteWiseProductNamesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vImageID := d.Get("image_id")
	vProductName, okProductName := d.GetOk("product_name")
	vProductID, okProductID := d.GetOk("product_id")
	vRecommended, okRecommended := d.GetOk("recommended")
	vAssigned, okAssigned := d.GetOk("assigned")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfAssignedNetworkDeviceProducts")
		vvImageID := vImageID.(string)
		queryParams1 := dnacentersdkgo.RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams{}

		if okProductName {
			queryParams1.ProductName = vProductName.(string)
		}
		if okProductID {
			queryParams1.ProductID = vProductID.(string)
		}
		if okRecommended {
			queryParams1.Recommended = vRecommended.(string)
		}
		if okAssigned {
			queryParams1.Assigned = vAssigned.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RetrievesTheCountOfAssignedNetworkDeviceProducts(vvImageID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfAssignedNetworkDeviceProducts", err,
				"Failure at RetrievesTheCountOfAssignedNetworkDeviceProducts, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfAssignedNetworkDeviceProducts response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
