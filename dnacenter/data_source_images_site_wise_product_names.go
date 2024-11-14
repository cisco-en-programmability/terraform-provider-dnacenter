package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImagesSiteWiseProductNames() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns a list of network device product names and associated sites for a given image identifier. Refer
*/dna/intent/api/v1/images* API for obtaining *imageId*.
`,

		ReadContext: dataSourceImagesSiteWiseProductNamesRead,
		Schema: map[string]*schema.Schema{
			"assigned": &schema.Schema{
				Description: `assigned query parameter. Filter with the assigned/unassigned, *ASSIGNED* option will filter network device products that are associated with the given image. The *NOT_ASSIGNED* option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": &schema.Schema{
				Description: `imageId path parameter. Software image identifier. Refer */dna/intent/api/v1/images* API for obtaining *imageId*
`,
				Type:     schema.TypeString,
				Required: true,
			},
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
			"product_id": &schema.Schema{
				Description: `productId query parameter. Filter with product ID (PID)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Description: `productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommended": &schema.Schema{
				Description: `recommended query parameter. Filter with recommended source. If *CISCO* then the network device product assigned was recommended by Cisco and *USER* then the user has manually assigned. Available values: CISCO, USER
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Product name ordinal is unique value for each network device product
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"product_ids": &schema.Schema{
							Description: `Supported PIDs
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"product_name": &schema.Schema{
							Description: `Network device product name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"product_name_ordinal": &schema.Schema{
							Description: `Product name ordinal is unique value for each network device product
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"recommended": &schema.Schema{
							Description: `If 'CISCO' network device product recommandation came from Cisco.com and 'USER' manually assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_ids": &schema.Schema{
							Description: `Sites where all  this image is assigned
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceImagesSiteWiseProductNamesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vImageID := d.Get("image_id")
	vProductName, okProductName := d.GetOk("product_name")
	vProductID, okProductID := d.GetOk("product_id")
	vRecommended, okRecommended := d.GetOk("recommended")
	vAssigned, okAssigned := d.GetOk("assigned")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage")
		vvImageID := vImageID.(string)
		queryParams1 := dnacentersdkgo.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams{}

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
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(vvImageID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage", err,
				"Failure at RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["product_name"] = item.ProductName
		respItem["product_name_ordinal"] = item.ProductNameOrdinal
		respItem["product_ids"] = item.ProductIDs
		respItem["site_ids"] = item.SiteIDs
		respItem["recommended"] = item.Recommended
		respItems = append(respItems, respItem)
	}
	return respItems
}
