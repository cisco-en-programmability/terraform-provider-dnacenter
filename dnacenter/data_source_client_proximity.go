package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientProximity() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- This intent API will provide client proximity information for a specific wireless user. Proximity is defined as
presence on the same floor at the same time as the specified wireless user. The Proximity workflow requires the
subscription to the following event (via the Event Notification workflow) prior to making this API call: NETWORK-
CLIENTS-3-506 Client Proximity Report.
`,

		ReadContext: dataSourceClientProximityRead,
		Schema: map[string]*schema.Schema{
			"number_days": &schema.Schema{
				Description: `number_days query parameter. Number of days to track proximity until current date. Defaults and maximum up to 14 days.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"time_resolution": &schema.Schema{
				Description: `time_resolution query parameter. Time interval (in minutes) to measure proximity. Defaults to 15 minutes with a minimum 5 minutes.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"username": &schema.Schema{
				Description: `username query parameter. Wireless client username for which proximity information is required
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClientProximityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vUsername := d.Get("username")
	vNumberDays, okNumberDays := d.GetOk("number_days")
	vTimeResolution, okTimeResolution := d.GetOk("time_resolution")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ClientProximity")
		queryParams1 := dnacentersdkgo.ClientProximityQueryParams{}

		queryParams1.Username = vUsername.(string)

		if okNumberDays {
			queryParams1.NumberDays = vNumberDays.(float64)
		}
		if okTimeResolution {
			queryParams1.TimeResolution = vTimeResolution.(float64)
		}

		response1, restyResp1, err := client.Clients.ClientProximity(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ClientProximity", err,
				"Failure at ClientProximity, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenClientsClientProximityItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ClientProximity response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsClientProximityItem(item *dnacentersdkgo.ResponseClientsClientProximity) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
