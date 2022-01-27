package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get total number of managed device(s).
`,

		ReadContext: dataSourceLicenseDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `device_type query parameter. Type of device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"dna_level": &schema.Schema{
				Description: `dna_level query parameter. Device Cisco DNA license level
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"registration_status": &schema.Schema{
				Description: `registration_status query parameter. Smart license registration status of device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"smart_account_id": &schema.Schema{
				Description: `smart_account_id query parameter. Id of smart account
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_account_name": &schema.Schema{
				Description: `virtual_account_name query parameter. Name of virtual account
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Total number of managed device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version
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

func dataSourceLicenseDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vRegistrationStatus, okRegistrationStatus := d.GetOk("registration_status")
	vDnaLevel, okDnaLevel := d.GetOk("dna_level")
	vVirtualAccountName, okVirtualAccountName := d.GetOk("virtual_account_name")
	vSmartAccountID, okSmartAccountID := d.GetOk("smart_account_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeviceCountDetails")
		queryParams1 := dnacentersdkgo.DeviceCountDetailsQueryParams{}

		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okRegistrationStatus {
			queryParams1.RegistrationStatus = vRegistrationStatus.(string)
		}
		if okDnaLevel {
			queryParams1.DnaLevel = vDnaLevel.(string)
		}
		if okVirtualAccountName {
			queryParams1.VirtualAccountName = vVirtualAccountName.(string)
		}
		if okSmartAccountID {
			queryParams1.SmartAccountID = vSmartAccountID.(string)
		}

		response1, restyResp1, err := client.Licenses.DeviceCountDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceCountDetails", err,
				"Failure at DeviceCountDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesDeviceCountDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeviceCountDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesDeviceCountDetailsItem(item *dnacentersdkgo.ResponseLicensesDeviceCountDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
