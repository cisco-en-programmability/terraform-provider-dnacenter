package dnacenter

import (
	"context"
	"errors"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevicesNetworkProfilesForSites() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and delete operations on Site Design.

- Deletes a network profile for sites.
`,

		CreateContext: resourceNetworkDevicesNetworkProfilesForSitesCreate,
		ReadContext:   resourceNetworkDevicesNetworkProfilesForSitesRead,
		UpdateContext: resourceNetworkDevicesNetworkProfilesForSitesUpdate,
		DeleteContext: resourceNetworkDevicesNetworkProfilesForSitesDelete,
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
							Description: `The ID of this network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The name of the network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
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

func resourceNetworkDevicesNetworkProfilesForSitesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDevicesNetworkProfilesForSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveANetworkProfileForSitesByID")
		vvID := vID

		response1, restyResp1, err := client.SiteDesign.RetrieveANetworkProfileForSitesByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenSiteDesignRetrieveANetworkProfileForSitesByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfNetworkProfilesForSites search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceNetworkDevicesNetworkProfilesForSitesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDevicesNetworkProfilesForSitesRead(ctx, d, m)
}

func resourceNetworkDevicesNetworkProfilesForSitesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.SiteDesign.DeletesANetworkProfileForSites(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesANetworkProfileForSites", err, restyResp1.String(),
				"Failure at DeletesANetworkProfileForSites, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesANetworkProfileForSites", err,
			"Failure at DeletesANetworkProfileForSites, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesANetworkProfileForSites", err))
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
				"Failure when executing DeletesANetworkProfileForSites", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

// func searchSiteDesignRetrievesTheListOfNetworkProfilesForSites(m interface{}, queryParams dnacentersdkgo.RetrievesTheListOfNetworkProfilesForSitesQueryParams, vID string) (*dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesResponse, error) {
// 	client := m.(*dnacentersdkgo.Client)
// 	var err error
// 	var foundItem *dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesResponse
// 	var ite *dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSites
// 	if vID != "" {
// 		queryParams.Offset = 1
// 		nResponse, _, err := client.SiteDesign.RetrievesTheListOfNetworkProfilesForSites(nil)
// 		maxPageSize := len(*nResponse.Response)
// 		for len(*nResponse.Response) > 0 {
// 			time.Sleep(15 * time.Second)
// 			for _, item := range *nResponse.Response {
// 				if vID == item.ID {
// 					foundItem = &item
// 					return foundItem, err
// 				}
// 			}
// 			queryParams.Limit = float64(maxPageSize)
// 			queryParams.Offset += float64(maxPageSize)
// 			nResponse, _, err = client.SiteDesign.RetrievesTheListOfNetworkProfilesForSites(&queryParams)
// if nResponse == nil || nResponse.Response == nil {
//                 break
//             }
// 		}
// 		return nil, err
// 	} else if queryParams.Name != "" {
// 		ite, _, err = client.SiteDesign.RetrievesTheListOfNetworkProfilesForSites(&queryParams)
// 		if err != nil || ite == nil {
// 			return foundItem, err
// 		}
// 		itemsCopy := *ite.Response
// 		if itemsCopy == nil {
// 			return foundItem, err
// 		}
// 		for _, item := range itemsCopy {
// 			if item.Name == queryParams.Name {
// 				foundItem = &item
// 				return foundItem, err
// 			}
// 		}
// 		return foundItem, err
// 	}
// 	return foundItem, err
// }
