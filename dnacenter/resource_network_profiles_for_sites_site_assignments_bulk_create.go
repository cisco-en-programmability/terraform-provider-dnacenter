package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"errors"

	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.
`,

		CreateContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateCreate,
		ReadContext:   resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateRead,
		DeleteContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateDelete,
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

						"task_id": &schema.Schema{
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_id": &schema.Schema{
							Description: `profileId path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vProfileID := resourceItem["profile_id"]

	vvProfileID := vProfileID.(string)
	request1 := expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSites(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SiteDesign.AssignANetworkProfileForSitesToAListOfSites(vvProfileID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToAListOfSites", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToAListOfSites", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AssignANetworkProfileForSitesToAListOfSites", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSiteDesignAssignANetworkProfileForSitesToAListOfSitesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignANetworkProfileForSitesToAListOfSites response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSites(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSites {
	request := dnacentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSites{}
	request.Items = expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesItems(ctx, key, d)
	return &request
}

func expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesItems(ctx context.Context, key string, d *schema.ResourceData) []struct {
	ID string `json:"id,omitempty"`
} {
	var request []struct {
		ID string `json:"id,omitempty"`
	}
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
		i := expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return request
}

func expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesItem(ctx context.Context, key string, d *schema.ResourceData) *struct {
	ID string `json:"id,omitempty"`
} {
	var request struct {
		ID string `json:"id,omitempty"`
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func flattenSiteDesignAssignANetworkProfileForSitesToAListOfSitesItem(item *dnacentersdkgo.ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
