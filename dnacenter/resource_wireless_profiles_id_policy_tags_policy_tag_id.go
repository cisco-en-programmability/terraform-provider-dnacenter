package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProfilesIDPolicyTagsPolicyTagID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Wireless.

- This endpoint allows for the deletion of a specific *Policy Tag* associated with a given *Wireless Profile*. This
resource requires the *id* of the *Wireless Profile* and the *policyTagId* of the *Policy Tag* to be provided as path
parameters.

- This endpoint allows updating the details of a specific *Policy Tag* associated with a given *Wireless Profile*. The
*id* of the *Wireless Profile* and the *policyTagId* of the Policy Tag must be provided as path parameters, and the
request body should contain the updated details of the *Policy Tag*. The *policyTagName* cannot be modified through this
endpoint. Note: When updating a Policy Tag, if the same set of AP Zones (apZones) is used for the same site or its
parent site, the existing Policy Tag will be overridden by the new one.
`,

		CreateContext: resourceWirelessProfilesIDPolicyTagsPolicyTagIDCreate,
		ReadContext:   resourceWirelessProfilesIDPolicyTagsPolicyTagIDRead,
		UpdateContext: resourceWirelessProfilesIDPolicyTagsPolicyTagIDUpdate,
		DeleteContext: resourceWirelessProfilesIDPolicyTagsPolicyTagIDDelete,
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

						"ap_zones": &schema.Schema{
							Description: `Ap Zones`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"policy_tag_id": &schema.Schema{
							Description: `Policy Tag Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"policy_tag_name": &schema.Schema{
							Description: `Policy Tag Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_ids": &schema.Schema{
							Description: `Site Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_zones": &schema.Schema{
							Description: `Ap Zones`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Wireless Profile Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"policy_tag_id": &schema.Schema{
							Description: `policyTagId path parameter. Policy Tag Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"policy_tag_name": &schema.Schema{
							Description: `Policy Tag Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"site_ids": &schema.Schema{
							Description: `Site Ids`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceWirelessProfilesIDPolicyTagsPolicyTagIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	resourceMap["policy_tag_id"] = interfaceToString(resourceItem["policy_tag_id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceWirelessProfilesIDPolicyTagsPolicyTagIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	vPolicyTagID := resourceMap["policy_tag_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveASpecificPolicyTagForAWirelessProfile")
		vvID := vID
		vvPolicyTagID := vPolicyTagID

		response1, restyResp1, err := client.Wireless.RetrieveASpecificPolicyTagForAWirelessProfile(vvID, vvPolicyTagID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveASpecificPolicyTagForAWirelessProfileItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveASpecificPolicyTagForAWirelessProfile response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceWirelessProfilesIDPolicyTagsPolicyTagIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vPolicyTagID := resourceMap["policy_tag_id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vID)
		request1 := expandRequestWirelessProfilesIDPolicyTagsPolicyTagIDUpdateASpecificPolicyTagForAWirelessProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateASpecificPolicyTagForAWirelessProfile(vID, vPolicyTagID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateASpecificPolicyTagForAWirelessProfile", err, restyResp1.String(),
					"Failure at UpdateASpecificPolicyTagForAWirelessProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateASpecificPolicyTagForAWirelessProfile", err,
				"Failure at UpdateASpecificPolicyTagForAWirelessProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateASpecificPolicyTagForAWirelessProfile", err))
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateASpecificPolicyTagForAWirelessProfile", err1))
				return diags
			}
		}

	}

	return resourceWirelessProfilesIDPolicyTagsPolicyTagIDRead(ctx, d, m)
}

func resourceWirelessProfilesIDPolicyTagsPolicyTagIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vPolicyTagID := resourceMap["policy_tag_id"]

	response1, restyResp1, err := client.Wireless.DeleteASpecificPolicyTagFromAWirelessProfile(vID, vPolicyTagID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteASpecificPolicyTagFromAWirelessProfile", err, restyResp1.String(),
				"Failure at DeleteASpecificPolicyTagFromAWirelessProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteASpecificPolicyTagFromAWirelessProfile", err,
			"Failure at DeleteASpecificPolicyTagFromAWirelessProfile, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteASpecificPolicyTagFromAWirelessProfile", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteASpecificPolicyTagFromAWirelessProfile", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessProfilesIDPolicyTagsPolicyTagIDUpdateASpecificPolicyTagForAWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateASpecificPolicyTagForAWirelessProfile {
	request := dnacentersdkgo.RequestWirelessUpdateASpecificPolicyTagForAWirelessProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_tag_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_tag_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_tag_name")))) {
		request.PolicyTagName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_zones")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_zones")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_zones")))) {
		request.ApZones = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
