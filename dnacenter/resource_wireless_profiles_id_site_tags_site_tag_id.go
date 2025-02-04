package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProfilesIDSiteTagsSiteTagID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Wireless.

- This endpoint allows updating the details of a specific *Site Tag* associated with a given *Wireless Profile*. The
*id* of the *Wireless Profile* and the *siteTagId* of the Site Tag must be provided as path parameters, and the request
body should contain the updated *Site Tag* details.  The *siteTagName* cannot be modified through this endpoint. Note:
When updating a Site Tag (siteTag), if the siteId already has an associated siteTag and the same siteId is included in
the update request, the existing siteTag for that siteId will be overridden by the new one. For Flex-enabled Wireless
Profiles (i.e., a Wireless Profile with one or more Flex SSIDs), a non-default Flex Profile Name (flexProfileName) will
be used. If no custom flexProfileName is provided, the System will automatically generate one and configure it in the
controller.

- This endpoint enables the deletion of a specific *Site Tag* associated with a given *Wireless Profile*. This resource
requires the *id* of the *Wireless Profile* and the *siteTagId* of the *Site Tag* to be provided as path parameters.
`,

		CreateContext: resourceWirelessProfilesIDSiteTagsSiteTagIDCreate,
		ReadContext:   resourceWirelessProfilesIDSiteTagsSiteTagIDRead,
		UpdateContext: resourceWirelessProfilesIDSiteTagsSiteTagIDUpdate,
		DeleteContext: resourceWirelessProfilesIDSiteTagsSiteTagIDDelete,
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

						"ap_profile_name": &schema.Schema{
							Description: `Ap Profile Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"flex_profile_name": &schema.Schema{
							Description: `Flex Profile Name`,
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
						"site_tag_id": &schema.Schema{
							Description: `Site Tag Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_tag_name": &schema.Schema{
							Description: `Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space
`,
							Type:     schema.TypeString,
							Computed: true,
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

						"ap_profile_name": &schema.Schema{
							Description: `Ap Profile Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"flex_profile_name": &schema.Schema{
							Description: `Flex Profile Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Wireless Profile Id
`,
							Type:     schema.TypeString,
							Required: true,
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
						"site_tag_id": &schema.Schema{
							Description: `siteTagId path parameter. Site Tag Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"site_tag_name": &schema.Schema{
							Description: `Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessProfilesIDSiteTagsSiteTagIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["site_tag_id"] = interfaceToString(resourceItem["site_tag_id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceWirelessProfilesIDSiteTagsSiteTagIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vSiteTagID := resourceMap["site_tag_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveASpecificSiteTagForAWirelessProfile")
		vvID := vID
		vvSiteTagID := vSiteTagID

		response1, restyResp1, err := client.Wireless.RetrieveASpecificSiteTagForAWirelessProfile(vvID, vvSiteTagID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveASpecificSiteTagForAWirelessProfileItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveASpecificSiteTagForAWirelessProfile response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceWirelessProfilesIDSiteTagsSiteTagIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	vSiteTagID := resourceMap["site_tag_id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vID)
		request1 := expandRequestWirelessProfilesIDSiteTagsSiteTagIDUpdateASpecificSiteTagForAWirelessProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateASpecificSiteTagForAWirelessProfile(vID, vSiteTagID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateASpecificSiteTagForAWirelessProfile", err, restyResp1.String(),
					"Failure at UpdateASpecificSiteTagForAWirelessProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateASpecificSiteTagForAWirelessProfile", err,
				"Failure at UpdateASpecificSiteTagForAWirelessProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateASpecificSiteTagForAWirelessProfile", err))
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
					"Failure when executing UpdateASpecificSiteTagForAWirelessProfile", err1))
				return diags
			}
		}

	}

	return resourceWirelessProfilesIDSiteTagsSiteTagIDRead(ctx, d, m)
}

func resourceWirelessProfilesIDSiteTagsSiteTagIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vSiteTagID := resourceMap["site_tag_id"]
	response1, restyResp1, err := client.Wireless.DeleteASpecificSiteTagFromAWirelessProfile(vID, vSiteTagID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteASpecificSiteTagFromAWirelessProfile", err, restyResp1.String(),
				"Failure at DeleteASpecificSiteTagFromAWirelessProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteASpecificSiteTagFromAWirelessProfile", err,
			"Failure at DeleteASpecificSiteTagFromAWirelessProfile, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteASpecificSiteTagFromAWirelessProfile", err))
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
				"Failure when executing DeleteASpecificSiteTagFromAWirelessProfile", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessProfilesIDSiteTagsSiteTagIDUpdateASpecificSiteTagForAWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateASpecificSiteTagForAWirelessProfile {
	request := dnacentersdkgo.RequestWirelessUpdateASpecificSiteTagForAWirelessProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_tag_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_tag_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_tag_name")))) {
		request.SiteTagName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_profile_name")))) {
		request.FlexProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_profile_name")))) {
		request.ApProfileName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
