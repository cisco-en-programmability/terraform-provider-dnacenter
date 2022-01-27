package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTopologyNetworkHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Topology.

- Returns Overall Network Health information by Device category (Access, Distribution, Core, Router, Wireless) for any
given point of time
`,

		ReadContext: dataSourceTopologyNetworkHealthRead,
		Schema: map[string]*schema.Schema{
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the Network health data is required
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bad_count": &schema.Schema{
							Description: `Bad Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"entity": &schema.Schema{
							Description: `Entity`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fair_count": &schema.Schema{
							Description: `Fair Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"good_count": &schema.Schema{
							Description: `Good Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"health_score": &schema.Schema{
							Description: `Health Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"time": &schema.Schema{
							Description: `Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timein_millis": &schema.Schema{
							Description: `Timein Millis`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_count": &schema.Schema{
							Description: `Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"unmon_count": &schema.Schema{
							Description: `Unmon Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTopologyNetworkHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTimestamp, okTimestamp := d.GetOk("timestamp")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetOverallNetworkHealth")
		queryParams1 := dnacentersdkgo.GetOverallNetworkHealthQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(string)
		}

		response1, restyResp1, err := client.Topology.GetOverallNetworkHealth(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetOverallNetworkHealth", err,
				"Failure at GetOverallNetworkHealth, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTopologyGetOverallNetworkHealthItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetOverallNetworkHealth response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTopologyGetOverallNetworkHealthItems(items *[]dnacentersdkgo.ResponseTopologyGetOverallNetworkHealthResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["time"] = item.Time
		respItem["health_score"] = item.HealthScore
		respItem["total_count"] = item.TotalCount
		respItem["good_count"] = item.GoodCount
		respItem["unmon_count"] = item.UnmonCount
		respItem["fair_count"] = item.FairCount
		respItem["bad_count"] = item.BadCount
		respItem["entity"] = flattenTopologyGetOverallNetworkHealthItemsEntity(item.Entity)
		respItem["timein_millis"] = item.TimeinMillis
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTopologyGetOverallNetworkHealthItemsEntity(item *dnacentersdkgo.ResponseTopologyGetOverallNetworkHealthResponseEntity) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
