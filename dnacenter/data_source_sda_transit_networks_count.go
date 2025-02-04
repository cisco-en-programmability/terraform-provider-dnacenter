package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaTransitNetworksCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of transit networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaTransitNetworksCountRead,
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Description: `type query parameter. Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
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
							Description: `Number of transit networks.
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

func dataSourceSdaTransitNetworksCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vType, okType := d.GetOk("type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransitNetworksCount")
		queryParams1 := dnacentersdkgo.GetTransitNetworksCountQueryParams{}

		if okType {
			queryParams1.Type = vType.(string)
		}

		response1, restyResp1, err := client.Sda.GetTransitNetworksCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTransitNetworksCount", err,
				"Failure at GetTransitNetworksCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetTransitNetworksCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransitNetworksCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetTransitNetworksCountItem(item *dnacentersdkgo.ResponseSdaGetTransitNetworksCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
