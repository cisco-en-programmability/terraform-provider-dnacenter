package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPendingFabricEvents() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of pending fabric events that match the provided query parameters.
`,

		ReadContext: dataSourceSdaPendingFabricEventsRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return. The maximum number of objects supported in a single request is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detail": &schema.Schema{
							Description: `Detail of the pending fabric event.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the pending fabric event.
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

func dataSourceSdaPendingFabricEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPendingFabricEvents")
		queryParams1 := dnacentersdkgo.GetPendingFabricEventsQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetPendingFabricEvents(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPendingFabricEvents", err,
				"Failure at GetPendingFabricEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetPendingFabricEventsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPendingFabricEvents response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPendingFabricEventsItems(items *[]dnacentersdkgo.ResponseSdaGetPendingFabricEventsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["detail"] = item.Detail
		respItems = append(respItems, respItem)
	}
	return respItems
}
