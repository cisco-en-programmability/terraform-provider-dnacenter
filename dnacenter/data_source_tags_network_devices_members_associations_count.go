package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagsNetworkDevicesMembersAssociationsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Fetches the count of network devices that are associated with at least one tag. A tag is a user-defined or system-
defined construct to group resources. When a device is tagged, it is called a member of the tag.
`,

		ReadContext: dataSourceTagsNetworkDevicesMembersAssociationsCountRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
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

func dataSourceTagsNetworkDevicesMembersAssociationsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag")

		response1, restyResp1, err := client.Tag.RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag", err,
				"Failure at RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagItem(item *dnacentersdkgo.ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
