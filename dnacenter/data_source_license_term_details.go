package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseTermDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get license term details.
`,

		ReadContext: dataSourceLicenseTermDetailsRead,
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dna_level": &schema.Schema{
							Description: `Cisco DNA license level
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_license_expired": &schema.Schema{
							Description: `Is license expired
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"license_term_end_date": &schema.Schema{
							Description: `End date of license term
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"license_term_start_date": &schema.Schema{
							Description: `Start date of license term
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"model": &schema.Schema{
							Description: `Model of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"purchased_dna_license_count": &schema.Schema{
							Description: `Number of purchased Cisco DNA licenses
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_account_name": &schema.Schema{
							Description: `Name of virtual account
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

func dataSourceLicenseTermDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSmartAccountID := d.Get("smart_account_id")
	vVirtualAccountName := d.Get("virtual_account_name")
	vDeviceType := d.Get("device_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: LicenseTermDetails")
		vvSmartAccountID := vSmartAccountID.(string)
		vvVirtualAccountName := vVirtualAccountName.(string)
		queryParams1 := dnacentersdkgo.LicenseTermDetailsQueryParams{}

		queryParams1.DeviceType = vDeviceType.(string)

		response1, restyResp1, err := client.Licenses.LicenseTermDetails(vvSmartAccountID, vvVirtualAccountName, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LicenseTermDetails", err,
				"Failure at LicenseTermDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensesLicenseTermDetailsItems(response1.LicenseDetails)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LicenseTermDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesLicenseTermDetailsItems(items *[]dnacentersdkgo.ResponseLicensesLicenseTermDetailsLicenseDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["model"] = item.Model
		respItem["virtual_account_name"] = item.VirtualAccountName
		respItem["license_term_start_date"] = item.LicenseTermStartDate
		respItem["license_term_end_date"] = item.LicenseTermEndDate
		respItem["dna_level"] = item.DnaLevel
		respItem["purchased_dna_license_count"] = item.PurchasedDnaLicenseCount
		respItem["is_license_expired"] = item.IsLicenseExpired
		respItems = append(respItems, respItem)
	}
	return respItems
}
