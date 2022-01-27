package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceDiscoveryRangeDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Discovery.

- Stops discovery for the given range and removes them
`,

		ReadContext: dataSourceDiscoveryRangeDeleteRead,
		Schema: map[string]*schema.Schema{
			"records_to_delete": &schema.Schema{
				Description: `recordsToDelete path parameter. Number of records to delete
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_index": &schema.Schema{
				Description: `startIndex path parameter. Start index
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryRangeDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartIndex := d.Get("start_index")
	vRecordsToDelete := d.Get("records_to_delete")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteDiscoveryBySpecifiedRange")
		vvStartIndex := vStartIndex.(int)
		vvRecordsToDelete := vRecordsToDelete.(int)

		response1, restyResp1, err := client.Discovery.DeleteDiscoveryBySpecifiedRange(vvStartIndex, vvRecordsToDelete)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteDiscoveryBySpecifiedRange", err,
				"Failure at DeleteDiscoveryBySpecifiedRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryDeleteDiscoveryBySpecifiedRangeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteDiscoveryBySpecifiedRange response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryDeleteDiscoveryBySpecifiedRangeItem(item *dnacentersdkgo.ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
