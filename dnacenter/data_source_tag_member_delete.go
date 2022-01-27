package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTagMemberDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Tag.

- Removes Tag member from the tag specified by id
`,

		ReadContext: dataSourceTagMemberDeleteRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"member_id": &schema.Schema{
				Description: `memberId path parameter. TagMember id to be removed from tag
`,
				Type:     schema.TypeString,
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

func dataSourceTagMemberDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vMemberID := d.Get("member_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveTagMember")
		vvID := vID.(string)
		vvMemberID := vMemberID.(string)

		response1, restyResp1, err := client.Tag.RemoveTagMember(vvID, vvMemberID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveTagMember", err,
				"Failure at RemoveTagMember, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagRemoveTagMemberItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveTagMember response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagRemoveTagMemberItem(item *dnacentersdkgo.ResponseTagRemoveTagMemberResponse) []map[string]interface{} {
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
