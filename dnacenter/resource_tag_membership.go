package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTagMembership() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Tag member.

	•	Adds members to the tag specified by id
	•	Removes Tag member from the tag specified by id
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
						"tag_id": &schema.Schema{
							Description: `id path parameter. Tag ID
			`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"member_type": &schema.Schema{
							Description: ``,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"member_id": &schema.Schema{
							Description: `memberId path parameter. TagMember id to be removed from tag
			`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
	vID := resourceItem["tag_id"]
	vvID := vID.(string)
	vMemberId := resourceItem["member_id"]
	vvMemberId := vMemberId.(string)
	vMemberType := resourceItem["member_type"]
	vvMemberType := vMemberType.(string)
	var diags diag.Diagnostics
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AddMembersToTheTag")

		request1 := expandRequestTagMemberCreateAddMembersToTheTag(ctx, "parameters.0", d)

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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing AddMembersToTheTag", err1))
				return diags
			}
		}

	}
	resourceMap := make(map[string]string)
	resourceMap["tag_id"] = vvID
	resourceMap["member_type"] = vvMemberType
	resourceMap["member_id"] = vvMemberId
	d.SetId(joinResourceID(resourceMap))
	return resourceTagMembershipRead(ctx, d, m)
}

func resourceTagMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["tag_id"]
	vMemberType := resourceMap["member_type"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTagMembersByID")
		queryParams1 := dnacentersdkgo.GetTagMembersByIDQueryParams{}
		queryParams1.MemberType = vMemberType
		response1, restyResp1, err := client.Tag.GetTagMembersByID(vID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
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
		request1 := expandRequestTagMembershipUpdatesTagMembership(ctx, "parameters.0", d)

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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdatesTagMembership", err1))
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
	vID := resourceMap["tag_id"]
	vMemberID := resourceMap["member_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveTagMember")

		response1, restyResp1, err := client.Tag.RemoveTagMember(vID, vMemberID)

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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing DeleteTagMembership", err1))
				return diags
			}
		}
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestTagMemberCreateAddMembersToTheTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagAddMembersToTheTag {
	request := dnacentersdkgo.RequestTagAddMembersToTheTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		if v2, ok := d.GetOkExists(fixKeyAccess(key + ".member_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_id")))) {
			x := make(map[string][]string)
			x[v.(string)] = append(x[v.(string)], v2.(string))
			request = x
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagMembershipUpdatesTagMembership(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdatesTagMembership {
	request := dnacentersdkgo.RequestTagUpdatesTagMembership{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag_id")))) {
		if v2, ok := d.GetOkExists(fixKeyAccess(key + ".member_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_id")))) {
			x := make(map[string][]string)
			x[v2.(string)] = append(x[v2.(string)], v.(string))
			request.MemberToTags = x
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
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
