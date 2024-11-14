package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProductNames() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Get the list of network device product names, their ordinal, and the support PIDs based on filter criteria.

- Get the network device product name, its ordinal, and supported PIDs.
`,

		ReadContext: dataSourceProductNamesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. The minimum value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
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
			"product_name_ordinal": &schema.Schema{
				Description: `productNameOrdinal path parameter. Product name ordinal is unique value for each network device product.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
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

						"product_ids": &schema.Schema{
							Description: `Supported PIDs
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
					},
				},
			},

			"items": &schema.Schema{
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

						"product_ids": &schema.Schema{
							Description: `Supported PIDs
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
					},
				},
			},
		},
	}
}

func dataSourceProductNamesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProductName, okProductName := d.GetOk("product_name")
	vProductID, okProductID := d.GetOk("product_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vProductNameOrdinal, okProductNameOrdinal := d.GetOk("product_name_ordinal")

	method1 := []bool{okProductName, okProductID, okOffset, okLimit}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okProductNameOrdinal}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfNetworkDeviceProductNames")
		queryParams1 := dnacentersdkgo.RetrievesTheListOfNetworkDeviceProductNamesQueryParams{}

		if okProductName {
			queryParams1.ProductName = vProductName.(string)
		}
		if okProductID {
			queryParams1.ProductID = vProductID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RetrievesTheListOfNetworkDeviceProductNames(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfNetworkDeviceProductNames", err,
				"Failure at RetrievesTheListOfNetworkDeviceProductNames, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfNetworkDeviceProductNames response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrieveNetworkDeviceProductName")
		vvProductNameOrdinal := vProductNameOrdinal.(float64)

		response2, restyResp2, err := client.SoftwareImageManagementSwim.RetrieveNetworkDeviceProductName(vvProductNameOrdinal)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveNetworkDeviceProductName", err,
				"Failure at RetrieveNetworkDeviceProductName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveNetworkDeviceProductName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["product_name"] = item.ProductName
		respItem["product_name_ordinal"] = item.ProductNameOrdinal
		respItem["product_ids"] = item.ProductIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["product_name"] = item.ProductName
	respItem["product_name_ordinal"] = item.ProductNameOrdinal
	respItem["product_ids"] = item.ProductIDs
	return []map[string]interface{}{
		respItem,
	}
}
