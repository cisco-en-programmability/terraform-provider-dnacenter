package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImagesAddonImages() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Retrieves the list of applicable add-on images if available for the given software image. *id* can be obtained from
the response of API [ /dna/intent/api/v1/images?hasAddonImages=true ].
`,

		ReadContext: dataSourceImagesAddonImagesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Software image identifier. Check */dna/intent/api/v1/images?hasAddonImages=true* API to get the same.
`,
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceImagesAddonImagesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveApplicableAddOnImagesForTheGivenSoftwareImage")
		vvID := vID.(string)

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RetrieveApplicableAddOnImagesForTheGivenSoftwareImage(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveApplicableAddOnImagesForTheGivenSoftwareImage", err,
				"Failure at RetrieveApplicableAddOnImagesForTheGivenSoftwareImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveApplicableAddOnImagesForTheGivenSoftwareImage response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponse) []map[string]interface{} {
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
		respItem["golden_tagging_details"] = flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItemsGoldenTaggingDetails(item.GoldenTaggingDetails)
		respItem["product_names"] = flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItemsProductNames(item.ProductNames)
		respItem["is_golden_tagged"] = boolPtrToString(item.IsGoldenTagged)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItemsGoldenTaggingDetails(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseGoldenTaggingDetails) []map[string]interface{} {
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

func flattenSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageItemsProductNames(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseProductNames) []map[string]interface{} {
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
