package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagMemberCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Returns the number of members in a given tag
`,

		ReadContext: dataSourceTagMemberCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"level": &schema.Schema{
				Description: `level query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"member_association_type": &schema.Schema{
				Description: `memberAssociationType query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"member_type": &schema.Schema{
				Description: `memberType query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagMemberCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vMemberType := d.Get("member_type")
	vMemberAssociationType, okMemberAssociationType := d.GetOk("member_association_type")
	vLevel, okLevel := d.GetOk("level")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTagMemberCount")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetTagMemberCountQueryParams{}

		queryParams1.MemberType = vMemberType.(string)

		if okMemberAssociationType {
			queryParams1.MemberAssociationType = vMemberAssociationType.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}

		response1, restyResp1, err := client.Tag.GetTagMemberCount(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagMemberCount", err,
				"Failure at GetTagMemberCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagGetTagMemberCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagMemberCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagGetTagMemberCountItem(item *dnacentersdkgo.ResponseTagGetTagMemberCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
