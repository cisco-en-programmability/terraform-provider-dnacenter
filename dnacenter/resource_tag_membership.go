package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTagMembership() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s
`,

		CreateContext: resourceTagMembershipCreate,
		ReadContext:   resourceTagMembershipRead,
		UpdateContext: resourceTagMembershipUpdate,
		DeleteContext: resourceTagMembershipDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationSet`,
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id path parameter. Tag ID
			`,
							Type:     schema.TypeString,
							Required: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestAddMembersToTheTagRequest`,
							Type:        schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required: true,
						},
						"member_to_tags": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"member_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member_id": &schema.Schema{
							Description: `memberId path parameter. TagMember id to be removed from tag
			`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceTagMembershipCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ApplicationSets Create")
	client := m.(*dnacentersdkgo.Client)

	resourceItem := *getResourceItem(d.Get("parameters"))
	vID := resourceItem["id"]
	vvID := vID.(string)
	var diags diag.Diagnostics
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AddMembersToTheTag")

		request1 := expandRequestTagMemberCreateAddMembersToTheTag(ctx, "parameters", d)

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
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				diags = append(diags, diagError(
					"Failure when executing AddMembersToTheTag", err))
				return diags
			}
		}

	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceTagMembershipRead(ctx, d, m)
}

func resourceTagMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTagMembersByID")
		queryParams1 := dnacentersdkgo.GetTagMembersByIDQueryParams{}

		response1, restyResp1, err := client.Tag.GetTagMembersByID(vID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagMembersByID", err,
				"Failure at GetTagMembersByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTagGetTagMembersByIDItems(response1.Response)
		if err := d.Set("item", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagMembersByID response",
				err))
			return diags
		}
		return diags
	}
	return diags
}

func resourceTagMembershipUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				diags = append(diags, diagError(
					"Failure when executing UpdatesTagMembership", err))
				return diags
			}
		}
	}
	return resourceTagMembershipRead(ctx, d, m)
}

func resourceTagMembershipDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vMemberID := d.Get("member_id")
	vvMemberID := vMemberID.(string)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveTagMember")

		response1, restyResp1, err := client.Tag.RemoveTagMember(vID, vvMemberID)

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

	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

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
	for item_no := range objs {
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

func searchResourceTagMembership(m interface{}, queryParams dnacentersdkgo.GetApplicationSetsQueryParams) (*dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSets
	ite, _, err = client.ApplicationPolicy.GetApplicationSets(&queryParams)

	if ite == nil {
		return foundItem, err
	}

	if ite.Response == nil {
		return foundItem, err
	}

	items := ite.Response
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
