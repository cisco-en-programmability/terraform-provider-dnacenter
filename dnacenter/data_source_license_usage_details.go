package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

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
				Description: `virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.
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
		log.Printf("[DEBUG] Selected method: LicenseUsageDetails2")
		vvSmartAccountID := vSmartAccountID.(string)
		vvVirtualAccountName := vVirtualAccountName.(string)
		queryParams1 := dnacentersdkgo.LicenseUsageDetails2QueryParams{}

		queryParams1.DeviceType = vDeviceType.(string)

		response1, restyResp1, err := client.Licenses.LicenseUsageDetails2(vvSmartAccountID, vvVirtualAccountName, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LicenseUsageDetails2", err,
				"Failure at LicenseUsageDetails2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesLicenseUsageDetails2Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LicenseUsageDetails2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesLicenseUsageDetails2Item(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["purchased_dna_license"] = flattenLicensesLicenseUsageDetails2ItemPurchasedDnaLicense(item.PurchasedDnaLicense)
	respItem["purchased_network_license"] = flattenLicensesLicenseUsageDetails2ItemPurchasedNetworkLicense(item.PurchasedNetworkLicense)
	respItem["used_dna_license"] = flattenLicensesLicenseUsageDetails2ItemUsedDnaLicense(item.UsedDnaLicense)
	respItem["used_network_license"] = flattenLicensesLicenseUsageDetails2ItemUsedNetworkLicense(item.UsedNetworkLicense)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesLicenseUsageDetails2ItemPurchasedDnaLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails2PurchasedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetails2ItemPurchasedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetails2ItemPurchasedDnaLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetails2PurchasedDnaLicenseLicenseCountByType) []map[string]interface{} {
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

func flattenLicensesLicenseUsageDetails2ItemPurchasedNetworkLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetails2ItemPurchasedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetails2ItemPurchasedNetworkLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicenseLicenseCountByType) []map[string]interface{} {
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

func flattenLicensesLicenseUsageDetails2ItemUsedDnaLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails2UsedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetails2ItemUsedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetails2ItemUsedDnaLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetails2UsedDnaLicenseLicenseCountByType) []map[string]interface{} {
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

func flattenLicensesLicenseUsageDetails2ItemUsedNetworkLicense(item *dnacentersdkgo.ResponseLicensesLicenseUsageDetails2UsedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetails2ItemUsedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetails2ItemUsedNetworkLicenseLicenseCountByType(items *[]dnacentersdkgo.ResponseLicensesLicenseUsageDetails2UsedNetworkLicenseLicenseCountByType) []map[string]interface{} {
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
