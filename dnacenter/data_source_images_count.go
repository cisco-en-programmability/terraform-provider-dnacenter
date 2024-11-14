package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImagesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns the count of software images for given *siteId*. The default value of siteId is global
`,

		ReadContext: dataSourceImagesCountRead,
		Schema: map[string]*schema.Schema{
			"golden": &schema.Schema{
				Description: `golden query parameter. When set to *true*, it will retrieve the images marked tagged golden. When set to *false*, it will retrieve the images marked not tagged golden.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"has_addon_images": &schema.Schema{
				Description: `hasAddonImages query parameter. When set to *true*, it will retrieve the images which have add-on images. When set to *false*, it will retrieve the images which do not have add-on images.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"imported": &schema.Schema{
				Description: `imported query parameter. When the value is set to *true*, it will include physically imported images. Conversely, when the value is set to *false*, it will include image records from the cloud. The identifier for cloud images can be utilised to download images from Cisco.com to the disk.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"integrity": &schema.Schema{
				Description: `integrity query parameter. Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_addon_images": &schema.Schema{
				Description: `isAddonImages query parameter. When set to *true*, it will retrieve the images that an add-on image.  When set to *false*, it will retrieve the images that are not add-on images
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name_ordinal": &schema.Schema{
				Description: `productNameOrdinal query parameter. The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of the API */dna/intent/api/v1/siteWiseProductNames*.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"supervisor_product_name_ordinal": &schema.Schema{
				Description: `supervisorProductNameOrdinal query parameter. The supervisor engine module ordinal is a unique value for each supervisor module. The *supervisorProductNameOrdinal* can be obtained from the response of API */dna/intent/api/v1/siteWiseProductNames*
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"version": &schema.Schema{
				Description: `version query parameter. Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
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

func dataSourceImagesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vProductNameOrdinal, okProductNameOrdinal := d.GetOk("product_name_ordinal")
	vSupervisorProductNameOrdinal, okSupervisorProductNameOrdinal := d.GetOk("supervisor_product_name_ordinal")
	vImported, okImported := d.GetOk("imported")
	vName, okName := d.GetOk("name")
	vVersion, okVersion := d.GetOk("version")
	vGolden, okGolden := d.GetOk("golden")
	vIntegrity, okIntegrity := d.GetOk("integrity")
	vHasAddonImages, okHasAddonImages := d.GetOk("has_addon_images")
	vIsAddonImages, okIsAddonImages := d.GetOk("is_addon_images")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsCountOfSoftwareImages")
		queryParams1 := dnacentersdkgo.ReturnsCountOfSoftwareImagesQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okProductNameOrdinal {
			queryParams1.ProductNameOrdinal = vProductNameOrdinal.(float64)
		}
		if okSupervisorProductNameOrdinal {
			queryParams1.SupervisorProductNameOrdinal = vSupervisorProductNameOrdinal.(float64)
		}
		if okImported {
			queryParams1.Imported = vImported.(bool)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okVersion {
			queryParams1.Version = vVersion.(string)
		}
		if okGolden {
			queryParams1.Golden = vGolden.(string)
		}
		if okIntegrity {
			queryParams1.Integrity = vIntegrity.(string)
		}
		if okHasAddonImages {
			queryParams1.HasAddonImages = vHasAddonImages.(bool)
		}
		if okIsAddonImages {
			queryParams1.IsAddonImages = vIsAddonImages.(bool)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ReturnsCountOfSoftwareImages(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsCountOfSoftwareImages", err,
				"Failure at ReturnsCountOfSoftwareImages, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimReturnsCountOfSoftwareImagesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsCountOfSoftwareImages response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimReturnsCountOfSoftwareImagesItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
