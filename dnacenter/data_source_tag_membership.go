package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTagMembership() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Tag.

- Updates tag membership. As part of the request payload through this API, only the specified members are added /
retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using
the /tag/member/type API
`,

		ReadContext: dataSourceTagMembershipRead,
		Schema: map[string]*schema.Schema{
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
			"member_to_tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"key": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"member_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceTagMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdatesTagMembership")
		request1 := expandRequestTagMembershipUpdatesTagMembership(ctx, "", d)

		response1, restyResp1, err := client.Tag.UpdatesTagMembership(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesTagMembership", err,
				"Failure at UpdatesTagMembership, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagUpdatesTagMembershipItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdatesTagMembership response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTagMembershipUpdatesTagMembership(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdatesTagMembership {
	request := dnacentersdkgo.RequestTagUpdatesTagMembership{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_to_tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_to_tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_to_tags")))) {
		request.MemberToTags = expandRequestTagMembershipUpdatesTagMembershipMemberToTagsArray(ctx, key+".member_to_tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagMembershipUpdatesTagMembershipMemberToTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagUpdatesTagMembershipMemberToTags {
	request := []dnacentersdkgo.RequestTagUpdatesTagMembershipMemberToTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestTagMembershipUpdatesTagMembershipMemberToTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagMembershipUpdatesTagMembershipMemberToTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdatesTagMembershipMemberToTags {
	request := dnacentersdkgo.RequestTagUpdatesTagMembershipMemberToTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenTagUpdatesTagMembershipItem(item *dnacentersdkgo.ResponseTagUpdatesTagMembershipResponse) []map[string]interface{} {
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
