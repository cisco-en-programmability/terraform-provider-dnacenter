package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseVirtualAccountDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get virtual account details of a smart account.
`,

		ReadContext: dataSourceLicenseVirtualAccountDetailsRead,
		Schema: map[string]*schema.Schema{
			"smart_account_id": &schema.Schema{
				Description: `smart_account_id path parameter. Id of smart account
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"smart_account_id": &schema.Schema{
							Description: `Id of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"smart_account_name": &schema.Schema{
							Description: `Name of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_account_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"virtual_account_id": &schema.Schema{
										Description: `Id of virtual account
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
				},
			},
		},
	}
}

func dataSourceLicenseVirtualAccountDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSmartAccountID := d.Get("smart_account_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: VirtualAccountDetails2")
		vvSmartAccountID := vSmartAccountID.(string)

		response1, restyResp1, err := client.Licenses.VirtualAccountDetails2(vvSmartAccountID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing VirtualAccountDetails2", err,
				"Failure at VirtualAccountDetails2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesVirtualAccountDetails2Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting VirtualAccountDetails2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesVirtualAccountDetails2Item(item *dnacentersdkgo.ResponseLicensesVirtualAccountDetails2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["smart_account_name"] = item.SmartAccountName
	respItem["virtual_account_details"] = flattenLicensesVirtualAccountDetails2ItemVirtualAccountDetails(item.VirtualAccountDetails)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesVirtualAccountDetails2ItemVirtualAccountDetails(items *[]dnacentersdkgo.ResponseLicensesVirtualAccountDetails2VirtualAccountDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["virtual_account_id"] = item.VirtualAccountID
		respItem["virtual_account_name"] = item.VirtualAccountName
		respItems = append(respItems, respItem)
	}
	return respItems
}
