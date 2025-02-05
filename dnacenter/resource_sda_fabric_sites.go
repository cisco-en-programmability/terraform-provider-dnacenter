package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricSites() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds a fabric site based on user input.

- Updates a fabric site based on user input.

- Deletes a fabric site based on id.
`,

		CreateContext: resourceSdaFabricSitesCreate,
		ReadContext:   resourceSdaFabricSitesRead,
		UpdateContext: resourceSdaFabricSitesUpdate,
		DeleteContext: resourceSdaFabricSitesDelete,
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

						"authentication_profile_name": &schema.Schema{
							Description: `Authentication profile used for this fabric.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the fabric site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pub_sub_enabled": &schema.Schema{
							Description: `Specifies whether this fabric site will use pub/sub for control nodes.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `ID of the network hierarchy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddFabricSite`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_profile_name": &schema.Schema{
										Description: `Authentication profile used for this fabric.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the fabric site (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_pub_sub_enabled": &schema.Schema{
										Description: `Specifies whether this fabric site will use pub/sub for control nodes.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"site_id": &schema.Schema{
										Description: `ID of the network hierarchy.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

func resourceSdaFabricSitesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaFabricSitesAddFabricSite(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	queryParamImport := dnacentersdkgo.GetFabricSitesQueryParams{}
	queryParamImport.SiteID = vvSiteID
	item2, err := searchSdaGetFabricSites(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_id"] = item2.SiteID
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricSitesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddFabricSite(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabricSite", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabricSite", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddFabricSite", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
				"Failure when executing AddFabricSite", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetFabricSitesQueryParams{}
	queryParamValidate.SiteID = vvSiteID
	item3, err := searchSdaGetFabricSites(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddFabricSite", err,
			"Failure at AddFabricSite, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["site_id"] = item3.SiteID
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricSitesRead(ctx, d, m)
}

func resourceSdaFabricSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	vvSiteID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricSites")
		queryParams1 := dnacentersdkgo.GetFabricSitesQueryParams{}
		queryParams1.SiteID = vvSiteID

		item1, err := searchSdaGetFabricSites(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseSdaGetFabricSitesResponse{
			*item1,
		}

		vItem1 := flattenSdaGetFabricSitesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricSites search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricSitesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricSitesUpdateFabricSite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateFabricSite(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFabricSite", err, restyResp1.String(),
					"Failure at UpdateFabricSite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFabricSite", err,
				"Failure at UpdateFabricSite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateFabricSite", err))
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
					"Failure when executing UpdateFabricSite", err1))
				return diags
			}
		}

	}

	return resourceSdaFabricSitesRead(ctx, d, m)
}

func resourceSdaFabricSitesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteFabricSiteByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFabricSiteByID", err, restyResp1.String(),
				"Failure at DeleteFabricSiteByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFabricSiteByID", err,
			"Failure at DeleteFabricSiteByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricSiteByID", err))
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
				"Failure when executing DeleteFabricSiteByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricSitesAddFabricSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddFabricSite {
	request := dnacentersdkgo.RequestSdaAddFabricSite{}
	if v := expandRequestSdaFabricSitesAddFabricSiteItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricSitesAddFabricSiteItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddFabricSite {
	request := []dnacentersdkgo.RequestItemSdaAddFabricSite{}
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
		i := expandRequestSdaFabricSitesAddFabricSiteItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricSitesAddFabricSiteItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricSite {
	request := dnacentersdkgo.RequestItemSdaAddFabricSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_profile_name")))) {
		request.AuthenticationProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_pub_sub_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_pub_sub_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_pub_sub_enabled")))) {
		request.IsPubSubEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricSitesUpdateFabricSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateFabricSite {
	request := dnacentersdkgo.RequestSdaUpdateFabricSite{}
	if v := expandRequestSdaFabricSitesUpdateFabricSiteItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricSitesUpdateFabricSiteItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateFabricSite {
	request := []dnacentersdkgo.RequestItemSdaUpdateFabricSite{}
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
		i := expandRequestSdaFabricSitesUpdateFabricSiteItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricSitesUpdateFabricSiteItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricSite {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_profile_name")))) {
		request.AuthenticationProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_pub_sub_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_pub_sub_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_pub_sub_enabled")))) {
		request.IsPubSubEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetFabricSites(m interface{}, queryParams dnacentersdkgo.GetFabricSitesQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetFabricSitesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetFabricSitesResponse
	var ite *dnacentersdkgo.ResponseSdaGetFabricSites

	ite, _, err = client.Sda.GetFabricSites(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err
	}
	itemsCopy := *ite.Response
	if itemsCopy == nil {
		return foundItem, err
	}
	for _, item := range itemsCopy {
		if item.SiteID == queryParams.SiteID {
			foundItem = &item
			return foundItem, err
		}
	}
	return foundItem, err

}
