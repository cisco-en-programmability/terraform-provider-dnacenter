package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseUsageDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get count of purchased and in use Cisco DNA and Network licenses.
`,

		ReadContext: dataSourceLicenseUsageDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `device_type query parameter. Type of device like router, switch, wireless or ise
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"smart_account_id": &schema.Schema{
				Description: `smart_account_id path parameter. Id of smart account
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_account_name": &schema.Schema{
				Description: `virtual_account_name path parameter. Name of virtual account. Putting "All" will give license usage detail for all virtual accounts.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"purchased_dna_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"purchased_network_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"used_dna_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"used_network_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceLicenseUsageDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSmartAccountID := d.Get("smart_account_id")
	vVirtualAccountName := d.Get("virtual_account_name")
	vDeviceType := d.Get("device_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: LicenseUsageDetails")
		vvSmartAccountID := vSmartAccountID.(string)
		vvVirtualAccountName := vVirtualAccountName.(string)
		queryParams1 := dnacentersdkgo.LicenseUsageDetailsQueryParams{}

		queryParams1.DeviceType = vDeviceType.(string)

		response1, restyResp1, err := client.Licenses.LicenseUsageDetails(vvSmartAccountID, vvVirtualAccountName, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LicenseUsageDetails", err,
				"Failure at LicenseUsageDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesLicenseUsageDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LicenseUsageDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesLicenseUsageDetailsItem(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["purchased_dna_license"] = flattenLicensesLicenseUsageDetailsItemPurchasedDnaLicense(item.PurchasedDnaLicense)
	respItem["purchased_network_license"] = flattenLicensesLicenseUsageDetailsItemPurchasedNetworkLicense(item.PurchasedNetworkLicense)
	respItem["used_dna_license"] = flattenLicensesLicenseUsageDetailsItemUsedDnaLicense(item.UsedDnaLicense)
	respItem["used_network_license"] = flattenLicensesLicenseUsageDetailsItemUsedNetworkLicense(item.UsedNetworkLicense)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesLicenseUsageDetailsItemPurchasedDnaLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetailsPurchasedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsItemPurchasedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsItemPurchasedDnaLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetailsPurchasedDnaLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsItemPurchasedNetworkLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsItemPurchasedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsItemPurchasedNetworkLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsItemUsedDnaLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetailsUsedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsItemUsedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsItemUsedDnaLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetailsUsedDnaLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsItemUsedNetworkLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetailsUsedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsItemUsedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsItemUsedNetworkLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetailsUsedNetworkLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}
