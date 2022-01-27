package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSwimImageDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns software image list based on a filter criteria. For example: "filterbyName = cat3k%"
`,

		ReadContext: dataSourceSwimImageDetailsRead,
		Schema: map[string]*schema.Schema{
			"application_type": &schema.Schema{
				Description: `applicationType query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"created_time": &schema.Schema{
				Description: `createdTime query parameter. time in milliseconds (epoch format)
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"family": &schema.Schema{
				Description: `family query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image_integrity_status": &schema.Schema{
				Description: `imageIntegrityStatus query parameter. imageIntegrityStatus FAILURE, UNKNOWN, VERIFIED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_name": &schema.Schema{
				Description: `imageName query parameter. image Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_series": &schema.Schema{
				Description: `imageSeries query parameter. image Series
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_size_greater_than": &schema.Schema{
				Description: `imageSizeGreaterThan query parameter. size in bytes
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"image_size_lesser_than": &schema.Schema{
				Description: `imageSizeLesserThan query parameter. size in bytes
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"image_uuid": &schema.Schema{
				Description: `imageUuid query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_cco_latest": &schema.Schema{
				Description: `isCCOLatest query parameter. is latest from cisco.com
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_cco_recommended": &schema.Schema{
				Description: `isCCORecommended query parameter. is recommended from cisco.com
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_tagged_golden": &schema.Schema{
				Description: `isTaggedGolden query parameter. is Tagged Golden
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. sort results by this field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. sort order 'asc' or 'des'. Default is asc
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Description: `version query parameter. software Image Version
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"applicable_devices_for_image": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdf_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"product_id": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"product_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"application_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"created_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"extended_attributes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"feature": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"file_service_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"file_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_integrity_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_series": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"image_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"import_source_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_tagged_golden": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"md5_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"extended_attributes": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"memory": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"product_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"profile_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"shares": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"v_cpu": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"sha_check_sum": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSwimImageDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vImageUUID, okImageUUID := d.GetOk("image_uuid")
	vName, okName := d.GetOk("name")
	vFamily, okFamily := d.GetOk("family")
	vApplicationType, okApplicationType := d.GetOk("application_type")
	vImageIntegrityStatus, okImageIntegrityStatus := d.GetOk("image_integrity_status")
	vVersion, okVersion := d.GetOk("version")
	vImageSeries, okImageSeries := d.GetOk("image_series")
	vImageName, okImageName := d.GetOk("image_name")
	vIsTaggedGolden, okIsTaggedGolden := d.GetOk("is_tagged_golden")
	vIsCCORecommended, okIsCCORecommended := d.GetOk("is_cco_recommended")
	vIsCCOLatest, okIsCCOLatest := d.GetOk("is_cco_latest")
	vCreatedTime, okCreatedTime := d.GetOk("created_time")
	vImageSizeGreaterThan, okImageSizeGreaterThan := d.GetOk("image_size_greater_than")
	vImageSizeLesserThan, okImageSizeLesserThan := d.GetOk("image_size_lesser_than")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSoftwareImageDetails")
		queryParams1 := dnacentersdkgo.GetSoftwareImageDetailsQueryParams{}

		if okImageUUID {
			queryParams1.ImageUUID = vImageUUID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okApplicationType {
			queryParams1.ApplicationType = vApplicationType.(string)
		}
		if okImageIntegrityStatus {
			queryParams1.ImageIntegrityStatus = vImageIntegrityStatus.(string)
		}
		if okVersion {
			queryParams1.Version = vVersion.(string)
		}
		if okImageSeries {
			queryParams1.ImageSeries = vImageSeries.(string)
		}
		if okImageName {
			queryParams1.ImageName = vImageName.(string)
		}
		if okIsTaggedGolden {
			queryParams1.IsTaggedGolden = vIsTaggedGolden.(bool)
		}
		if okIsCCORecommended {
			queryParams1.IsCCORecommended = vIsCCORecommended.(bool)
		}
		if okIsCCOLatest {
			queryParams1.IsCCOLatest = vIsCCOLatest.(bool)
		}
		if okCreatedTime {
			queryParams1.CreatedTime = vCreatedTime.(int)
		}
		if okImageSizeGreaterThan {
			queryParams1.ImageSizeGreaterThan = vImageSizeGreaterThan.(int)
		}
		if okImageSizeLesserThan {
			queryParams1.ImageSizeLesserThan = vImageSizeLesserThan.(int)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSoftwareImageDetails", err,
				"Failure at GetSoftwareImageDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSoftwareImageDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["applicable_devices_for_image"] = flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsApplicableDevicesForImage(item.ApplicableDevicesForImage)
		respItem["application_type"] = item.ApplicationType
		respItem["created_time"] = item.CreatedTime
		respItem["extended_attributes"] = flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsExtendedAttributes(item.ExtendedAttributes)
		respItem["family"] = item.Family
		respItem["feature"] = item.Feature
		respItem["file_service_id"] = item.FileServiceID
		respItem["file_size"] = item.FileSize
		respItem["image_integrity_status"] = item.ImageIntegrityStatus
		respItem["image_name"] = item.ImageName
		respItem["image_series"] = item.ImageSeries
		respItem["image_source"] = item.ImageSource
		respItem["image_type"] = item.ImageType
		respItem["image_uuid"] = item.ImageUUID
		respItem["import_source_type"] = item.ImportSourceType
		respItem["is_tagged_golden"] = boolPtrToString(item.IsTaggedGolden)
		respItem["md5_checksum"] = item.Md5Checksum
		respItem["name"] = item.Name
		respItem["profile_info"] = flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsProfileInfo(item.ProfileInfo)
		respItem["sha_check_sum"] = item.ShaCheckSum
		respItem["vendor"] = item.Vendor
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsApplicableDevicesForImage(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseApplicableDevicesForImage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mdf_id"] = item.MdfID
		respItem["product_id"] = item.ProductID
		respItem["product_name"] = item.ProductName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsExtendedAttributes(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseExtendedAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsProfileInfo(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["description"] = item.Description
		respItem["extended_attributes"] = flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsProfileInfoExtendedAttributes(item.ExtendedAttributes)
		respItem["memory"] = item.Memory
		respItem["product_type"] = item.ProductType
		respItem["profile_name"] = item.ProfileName
		respItem["shares"] = item.Shares
		respItem["v_cpu"] = item.VCPU
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItemsProfileInfoExtendedAttributes(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfoExtendedAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
