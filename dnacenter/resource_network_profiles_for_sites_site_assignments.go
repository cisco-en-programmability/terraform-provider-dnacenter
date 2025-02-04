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

func resourceNetworkProfilesForSitesSiteAssignments() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Site Design.

- Assigns a given network profile for sites to a given site. Also assigns the profile to child sites.

- Unassigns a given network profile for sites from a site. The profile must be removed from parent sites first,
otherwise this operation will not ulimately  unassign the profile.
`,

		CreateContext: resourceNetworkProfilesForSitesSiteAssignmentsCreate,
		ReadContext:   resourceNetworkProfilesForSitesSiteAssignmentsRead,
		UpdateContext: resourceNetworkProfilesForSitesSiteAssignmentsUpdate,
		DeleteContext: resourceNetworkProfilesForSitesSiteAssignmentsDelete,
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

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"profile_id": &schema.Schema{
							Description: `profileId path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkProfilesForSitesSiteAssignmentsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkProfilesForSitesSiteAssignmentsAssignANetworkProfileForSitesToTheGivenSite(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vProfileID := resourceItem["profile_id"]
	vvProfileID := interfaceToString(vProfileID)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParamImport := dnacentersdkgo.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams{}
	item2, err := searchSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(m, queryParamImport, vvProfileID, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["profile_id"] = vvProfileID
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceNetworkProfilesForSitesSiteAssignmentsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.SiteDesign.AssignANetworkProfileForSitesToTheGivenSite(vvProfileID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AssignANetworkProfileForSitesToTheGivenSite", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToTheGivenSite", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToTheGivenSite", err))
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
				"Failure when executing AssignANetworkProfileForSitesToTheGivenSite", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams{}
	item3, err := searchSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(m, queryParamValidate, vvProfileID, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AssignANetworkProfileForSitesToTheGivenSite", err,
			"Failure at AssignANetworkProfileForSitesToTheGivenSite, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["profile_id"] = vvProfileID
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkProfilesForSitesSiteAssignmentsRead(ctx, d, m)
}

func resourceNetworkProfilesForSitesSiteAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vProfileID := resourceMap["profile_id"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo")
		vvProfileID := vProfileID
		queryParams1 := dnacentersdkgo.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams{}

		item1, err := searchSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(m, queryParams1, vvProfileID, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}

		items := []dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse{
			*item1,
		}

		// Review flatten function used
		vItem1 := flattenSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceNetworkProfilesForSitesSiteAssignmentsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkProfilesForSitesSiteAssignmentsRead(ctx, d, m)
}

func resourceNetworkProfilesForSitesSiteAssignmentsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvProfileID := resourceMap["profile_id"]
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.SiteDesign.UnassignsANetworkProfileForSitesFromASite(vvProfileID, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing UnassignsANetworkProfileForSitesFromASite", err, restyResp1.String(),
				"Failure at UnassignsANetworkProfileForSitesFromASite, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UnassignsANetworkProfileForSitesFromASite", err,
			"Failure at UnassignsANetworkProfileForSitesFromASite, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UnassignsANetworkProfileForSitesFromASite", err))
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
				"Failure when executing UnassignsANetworkProfileForSitesFromASite", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkProfilesForSitesSiteAssignmentsAssignANetworkProfileForSitesToTheGivenSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSite {
	request := dnacentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(m interface{}, queryParams dnacentersdkgo.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams, vProfileID string, vID string) (*dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse
	// var ite *dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.SiteDesign.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(vProfileID, nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.SiteDesign.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(vProfileID, &queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	}
	return foundItem, err
}
