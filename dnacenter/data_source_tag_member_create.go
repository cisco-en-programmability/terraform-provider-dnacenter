package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTagMemberCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Tag.

- Adds members to the tag specified by id
`,

		ReadContext: dataSourceTagMemberCreateRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
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
			"payload": &schema.Schema{
				Description: `Array of RequestAddMembersToTheTagRequest`,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func dataSourceTagMemberCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AddMembersToTheTag")
		vvID := vID.(string)
		request1 := expandRequestTagMemberCreateAddMembersToTheTag(ctx, "", d)

		response1, restyResp1, err := client.Tag.AddMembersToTheTag(vvID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddMembersToTheTag", err,
				"Failure at AddMembersToTheTag, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagAddMembersToTheTagItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AddMembersToTheTag response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTagMemberCreateAddMembersToTheTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagAddMembersToTheTag {
	request := dnacentersdkgo.RequestTagAddMembersToTheTag{}
	if v := expandRequestItemTagMemberAddMembersToTheTag(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestItemTagMemberAddMembersToTheTag(ctx context.Context, key string, d *schema.ResourceData) *map[string][]string {
	var request map[string][]string
	o := d.Get(fixKeyAccess(key))
	if o == nil {
		return nil
	}
	request = o.(map[string][]string)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenTagAddMembersToTheTagItem(item *dnacentersdkgo.ResponseTagAddMembersToTheTagResponse) []map[string]interface{} {
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
