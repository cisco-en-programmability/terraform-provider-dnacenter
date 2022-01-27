package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceByIDDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return compliance detailed report for a device.
`,

		ReadContext: dataSourceComplianceDeviceByIDDetailRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Description: `category query parameter. complianceCategory can have any value among 'INTENT', 'RUNNING_CONFIG'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. complianceType can have any value among 'NETWORK_DESIGN', 'NETWORK_PROFILE', 'FABRIC', 'POLICY', 'RUNNING_CONFIG'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"diff_list": &schema.Schema{
				Description: `diffList query parameter. diff list [ pass true to fetch the diff list ]
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"key": &schema.Schema{
				Description: `key query parameter. extended attribute key
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Description: `value query parameter. extended attribute value
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_data_url": &schema.Schema{
							Description: `Additional Data U R L`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"compliance_type": &schema.Schema{
							Description: `Compliance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_uuid": &schema.Schema{
							Description: `Device Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_sync_time": &schema.Schema{
							Description: `Last Sync Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"source_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_name": &schema.Schema{
										Description: `App Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"business_key": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"business_key_attributes": &schema.Schema{
													Description: `Business Key Attributes`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"other_attributes": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cfs_attributes": &schema.Schema{
																Description: `Cfs Attributes`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"resource_name": &schema.Schema{
													Description: `Resource Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"count": &schema.Schema{
										Description: `Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"diff_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"business_key": &schema.Schema{
													Description: `Business Key`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"configured_value": &schema.Schema{
													Description: `Configured Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"display_name": &schema.Schema{
													Description: `Display Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"extended_attributes": &schema.Schema{
													Description: `Extended Attributes`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"intended_value": &schema.Schema{
													Description: `Intended Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"move_from_path": &schema.Schema{
													Description: `Move From Path`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"op": &schema.Schema{
													Description: `Op`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"path": &schema.Schema{
													Description: `Path`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"license_app_name": &schema.Schema{
										Description: `License App Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name_with_business_key": &schema.Schema{
										Description: `Name With Business Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"network_profile_name": &schema.Schema{
										Description: `Network Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"provisioning_area": &schema.Schema{
										Description: `Provisioning Area`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"source_enum": &schema.Schema{
										Description: `Source Enum`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"state": &schema.Schema{
							Description: `State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceComplianceDeviceByIDDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")
	vCategory, okCategory := d.GetOk("category")
	vComplianceType, okComplianceType := d.GetOk("compliance_type")
	vDiffList, okDiffList := d.GetOk("diff_list")
	vKey, okKey := d.GetOk("key")
	vValue, okValue := d.GetOk("value")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ComplianceDetailsOfDevice")
		vvDeviceUUID := vDeviceUUID.(string)
		queryParams1 := dnacentersdkgo.ComplianceDetailsOfDeviceQueryParams{}

		if okCategory {
			queryParams1.Category = vCategory.(string)
		}
		if okComplianceType {
			queryParams1.ComplianceType = vComplianceType.(string)
		}
		if okDiffList {
			queryParams1.DiffList = vDiffList.(bool)
		}
		if okKey {
			queryParams1.Key = vKey.(string)
		}
		if okValue {
			queryParams1.Value = vValue.(string)
		}

		response1, restyResp1, err := client.Compliance.ComplianceDetailsOfDevice(vvDeviceUUID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ComplianceDetailsOfDevice", err,
				"Failure at ComplianceDetailsOfDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceComplianceDetailsOfDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ComplianceDetailsOfDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceComplianceDetailsOfDeviceItems(items *[]dnacentersdkgo.ResponseComplianceComplianceDetailsOfDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["compliance_type"] = item.ComplianceType
		respItem["last_sync_time"] = item.LastSyncTime
		respItem["additional_data_url"] = item.AdditionalDataURL
		respItem["source_info_list"] = flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoList(item.SourceInfoList)
		respItem["device_uuid"] = item.DeviceUUID
		respItem["message"] = item.Message
		respItem["state"] = item.State
		respItem["status"] = item.Status
		respItem["category"] = item.Category
		respItem["last_update_time"] = item.LastUpdateTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoList(items *[]dnacentersdkgo.ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["count"] = item.Count
		respItem["display_name"] = item.DisplayName
		respItem["diff_list"] = flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListDiffList(item.DiffList)
		respItem["source_enum"] = item.SourceEnum
		respItem["license_app_name"] = item.LicenseAppName
		respItem["provisioning_area"] = item.ProvisioningArea
		respItem["network_profile_name"] = item.NetworkProfileName
		respItem["name_with_business_key"] = item.NameWithBusinessKey
		respItem["app_name"] = item.AppName
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["business_key"] = flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListBusinessKey(item.BusinessKey)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListDiffList(items *[]dnacentersdkgo.ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["move_from_path"] = item.MoveFromPath
		respItem["op"] = item.Op
		respItem["configured_value"] = item.ConfiguredValue
		respItem["intended_value"] = item.IntendedValue
		respItem["path"] = item.Path
		respItem["business_key"] = item.BusinessKey
		respItem["extended_attributes"] = item.ExtendedAttributes
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListBusinessKey(item *dnacentersdkgo.ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["other_attributes"] = flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListBusinessKeyOtherAttributes(item.OtherAttributes)
	respItem["resource_name"] = item.ResourceName
	respItem["business_key_attributes"] = item.BusinessKeyAttributes

	return []map[string]interface{}{
		respItem,
	}

}

func flattenComplianceComplianceDetailsOfDeviceItemsSourceInfoListBusinessKeyOtherAttributes(item *dnacentersdkgo.ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cfs_attributes"] = item.CfsAttributes
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}
