package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProductNamesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Count of product names based on filter criteria
`,

		ReadContext: dataSourceProductNamesCountRead,
		Schema: map[string]*schema.Schema{
			"product_id": &schema.Schema{
				Description: `productId query parameter. Filter with product ID (PID)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": &schema.Schema{
				Description: `productName query parameter. Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
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

func dataSourceProductNamesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProductName, okProductName := d.GetOk("product_name")
	vProductID, okProductID := d.GetOk("product_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountOfNetworkProductNames")
		queryParams1 := dnacentersdkgo.CountOfNetworkProductNamesQueryParams{}

		if okProductName {
			queryParams1.ProductName = vProductName.(string)
		}
		if okProductID {
			queryParams1.ProductID = vProductID.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.CountOfNetworkProductNames(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountOfNetworkProductNames", err,
				"Failure at CountOfNetworkProductNames, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimCountOfNetworkProductNamesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountOfNetworkProductNames response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimCountOfNetworkProductNamesItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
