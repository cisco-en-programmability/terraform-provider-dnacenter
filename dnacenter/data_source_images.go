package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImages() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- A list of available images for the specified site is provided. The default value of the site is set to global. The
list includes images that have been imported onto the disk, as well as the latest and suggested images from Cisco.com.
`,

		ReadContext: dataSourceImagesRead,
		Schema: map[string]*schema.Schema{
			"golden": &schema.Schema{
				Description: `golden query parameter. When set to *true*, it will retrieve the images marked as tagged golden. When set to *false*, it will retrieve the images marked as not tagged golden.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"has_addon_images": &schema.Schema{
				Description: `hasAddonImages query parameter. When set to *true*, it will retrieve the images which have add-on images. When set to *false*, it will retrieve the images which do not have add-on images.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"imported": &schema.Schema{
				Description: `imported query parameter. When the value is set to *true*, it will include physically imported images. Conversely, when the value is set to *false*, it will include image records from the cloud. The identifier for cloud images can be utilized to download images from Cisco.com to the disk.
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
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. The minimum value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"product_name_ordinal": &schema.Schema{
				Description: `productNameOrdinal query parameter. The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of API */dna/intent/api/v1/siteWiseProductNames*
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for *siteId*
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
				Description: `version query parameter. Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cisco_latest": &schema.Schema{
							Description: `*true* if the image is latest/suggested from Cisco.com
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"golden_tagging_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_roles": &schema.Schema{
										Description: `Golden tagging based on the device roles
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"device_tags": &schema.Schema{
										Description: `Golden tagging based on the device tags
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inherited_site_id": &schema.Schema{
										Description: `The Site Id of the site that this setting is inherited from.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `The name of the site that this setting is inherited from
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"has_addon_images": &schema.Schema{
							Description: `Software images that have an applicable list of add-on images. The value of *true* will return software images with add-on images, while the value of *false* will return software images without add-on images
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Software image identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_type": &schema.Schema{
							Description: `Software image type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"imported": &schema.Schema{
							Description: `Flag for image info whether it is imported image or cloud image
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"integrity_status": &schema.Schema{
							Description: `Image Integrity verification status with Known Good Verification
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_addon_image": &schema.Schema{
							Description: `The value of *true* will indicate the image as an add-on image, while the value of *false* will indicate software image
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_golden_tagged": &schema.Schema{
							Description: `The value of *true* will indicate the image marked as golden, while the value of *false* will indicate the image not marked as golden
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the software image
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"product_names": &schema.Schema{
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

									"supervisor_product_name": &schema.Schema{
										Description: `Name of the Supervisor Engine Module, supported by the *productName*.                  Example: The *Cisco Catalyst 9404R Switch* chassis is capable of supporting  different supervisor engine modules: the *Cisco Catalyst 9400 Supervisor Engine-1*, the *Cisco Catalyst 9400 Supervisor Engine-1XL*, the *Cisco Catalyst 9400 Supervisor Engine-1XL-Y*, etc.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"supervisor_product_name_ordinal": &schema.Schema{
										Description: `Supervisor Engine Module Ordinal, supported by the *productNameOrdinal*. Example: The *286315691* chassis ordinal is capable of supporting different supervisor engine module ordinals: *286316172*, *286316710*, *286320394* etc.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"recommended": &schema.Schema{
							Description: `CISCO if the image is recommended from Cisco.com
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Software image  version
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

func dataSourceImagesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsListOfSoftwareImages")
		queryParams1 := dnacentersdkgo.ReturnsListOfSoftwareImagesQueryParams{}

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
			queryParams1.Golden = vGolden.(bool)
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
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ReturnsListOfSoftwareImages(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsListOfSoftwareImages", err,
				"Failure at ReturnsListOfSoftwareImages, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsListOfSoftwareImages response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["imported"] = boolPtrToString(item.Imported)
		respItem["name"] = item.Name
		respItem["version"] = item.Version
		respItem["image_type"] = item.ImageType
		respItem["recommended"] = item.Recommended
		respItem["cisco_latest"] = boolPtrToString(item.CiscoLatest)
		respItem["integrity_status"] = item.IntegrityStatus
		respItem["is_addon_image"] = boolPtrToString(item.IsAddonImage)
		respItem["has_addon_images"] = boolPtrToString(item.HasAddonImages)
		respItem["golden_tagging_details"] = flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItemsGoldenTaggingDetails(item.GoldenTaggingDetails)
		respItem["product_names"] = flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItemsProductNames(item.ProductNames)
		respItem["is_golden_tagged"] = boolPtrToString(item.IsGoldenTagged)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItemsGoldenTaggingDetails(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseGoldenTaggingDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_roles"] = item.DeviceRoles
		respItem["device_tags"] = item.DeviceTags
		respItem["inherited_site_id"] = item.InheritedSiteID
		respItem["inherited_site_name"] = item.InheritedSiteName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimReturnsListOfSoftwareImagesItemsProductNames(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseProductNames) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["product_name"] = item.ProductName
		respItem["product_name_ordinal"] = item.ProductNameOrdinal
		respItem["supervisor_product_name"] = item.SupervisorProductName
		respItem["supervisor_product_name_ordinal"] = item.SupervisorProductNameOrdinal
		respItems = append(respItems, respItem)
	}
	return respItems
}
