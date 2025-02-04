package dnacenter

import (
	"context"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceArea() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Sites.

- Creates site with area with specified hierarchy.

- Update site area with specified hierarchy and new values

- Delete site with area by siteId.
`,

		CreateContext: resourceAreaCreate,
		ReadContext:   resourceAreaRead,
		UpdateContext: resourceAreaUpdate,
		DeleteContext: resourceAreaDelete,
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

						"additional_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"addressinheritedfrom": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"parent_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"name_space": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"parent_id": &schema.Schema{
							Description: `Parent Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
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

						"site": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"area": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name of the area (eg: Area1)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of the area to be created
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
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id to which site details to be updated.
`,
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"type": &schema.Schema{
							Description: `Type of site to create (eg: area, building, floor)
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

func resourceAreaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteCreateSite(ctx, "parameters.0", d)

	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vType := resourceItem["type"]
	vvType := interfaceToString(vType)

	vvName := ""
	vvParentName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.site"); ok {
			if _, ok := d.GetOk("parameters.0.site.0"); ok {
				if _, ok := d.GetOk("parameters.0.site.0." + vvType); ok {
					if v, ok := d.GetOk("parameters.0.site.0." + vvType + ".0.name"); ok {
						vvName = interfaceToString(v)
					}
					if v2, ok := d.GetOk("parameters.0.site.0." + vvType + ".0.parent_name"); ok {
						vvParentName = interfaceToString(v2)
					}
				}
			}
		}
	}

	pathName := []string{vvParentName, vvName}
	newName := strings.Join(pathName, "/")
	if !strings.Contains(newName, "Global/") {
		newPathName := []string{"Global", newName}
		newName = strings.Join(newPathName, "/")
	}
	queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
	queryParams1.Name = newName
	queryParams1.SiteID = vvSiteID
	log.Printf("[DEBUG] newName => %s", newName)
	item, err := searchSitesGetSite(m, queryParams1)
	if err == nil || item != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_id"] = item.ID
		resourceMap["name"] = item.SiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceAreaRead(ctx, d, m)
	}
	headers := dnacentersdkgo.CreateSiteHeaderParams{}
	headers.Persistbapioutput = "false"
	headers.Runsync = "false"
	/*if okTimeout {
		headers.Timeout = vvTimeout
	}*/
	resp1, restyResp1, err := client.Sites.CreateSite(request1, &headers)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSite", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSite", err))
		return diags
	}

	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response1 %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response2 %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			if strings.Contains(response2.BapiError, "Rate Limit") {
				return resourceAreaCreate(ctx, d, m)
			} else {
				bapiError := response2.BapiError
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing CreateArea", err,
					"Failure at CreateArea execution", bapiError))
			}
			return diags
		}
	}
	queryParams2 := dnacentersdkgo.GetSiteQueryParams{}
	queryParams2.Name = newName
	// queryParams2.SiteID = vvSiteID
	item2, err := searchSitesGetSite(m, queryParams2)
	if err != nil || item2 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateArea", err,
			"Failure at CreateArea execution", ""))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = item2.ID
	resourceMap["name"] = item2.SiteNameHierarchy
	d.SetId(joinResourceID(resourceMap))
	return resourceAreaRead(ctx, d, m)
}

func resourceAreaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]
	log.Printf("[DEBUG] Read SiteNameHierarchy => %s", vName)
	// vSiteID := resourceMap["site_id"]
	//vParentName := resourceMap["parent_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
		queryParams1.Name = vName
		// queryParams1.SiteID = vSiteID
		log.Printf("[DEBUG] Read name => %s", queryParams1.Name)
		log.Printf("[DEBUG] Read site => %s", queryParams1.SiteID)
		response1, restyResp1, err := client.Sites.GetArea(&queryParams1)
		if err != nil || response1 == nil {
			log.Printf("[DEBUG] Error => %s", err.Error())
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response3 %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}
		parameters := d.Get("parameters").([]interface{})
		vItem1 := flattenSitesGetAreaParams(response1.Response, parameters)
		log.Printf("[DEBUG] response flatten sent => %v", responseInterfaceToString(vItem1))
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSite search response",
				err))
			return diags
		}

		vItem2 := flattenSitesGetAreaParams(response1.Response, parameters)
		if err := d.Set("parameters", []map[string]interface{}{vItem2}); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSite search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAreaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	// vName := resourceMap["name"]
	vSiteID := resourceMap["site_id"]
	log.Printf("[DEBUG] ID used for update operation %s", vSiteID)
	if d.HasChange("parameters") {
		request1 := expandRequestSiteUpdateSite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		headers := dnacentersdkgo.UpdateSiteHeaderParams{}
		headers.Persistbapioutput = "false"
		headers.Runsync = "false"
		/*if vTimeout != "" {
			headers.Timeout = vTimeout
		}*/

		response1, restyResp1, err := client.Sites.UpdateSite(vSiteID, request1, &headers)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSite", err, restyResp1.String(),
					"Failure at UpdateSite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSite", err,
				"Failure at UpdateSite, unexpected response", ""))
			return diags
		}
		if response1.Response != nil {
			errorResult, _ := strconv.ParseBool(response1.Response.IsError)
			if errorResult {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing UpdateSite", err,
					"Failure at UpdateSite, unexpected response", ""))
				return diags
			}
		}
	}

	return resourceAreaRead(ctx, d, m)
}

func resourceAreaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	// time.Sleep(1 * time.Minute)
	// queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
	// queryParams1.Name = vName
	// queryParams1.SiteID = vSiteID
	// item, err := searchSitesGetSite(m, queryParams1)
	// if err != nil || item == nil {
	// 	d.SetId("")
	// 	return diags
	// }

	// if vSiteID != item.ID {
	// 	vSiteID = item.ID
	// }

	response1, restyResp1, err := client.Sites.DeleteSite(vSiteID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSite", err, restyResp1.String(),
				"Failure at DeleteSite, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSite", err,
			"Failure at DeleteSite, unexpected response", ""))
		return diags
	}

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response1 %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response2 %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}

		if statusIsFailure(response2.Status) {
			if strings.Contains(response2.BapiError, "Rate Limit") {
				return resourceAreaDelete(ctx, d, m)
			} else {
				bapiError := response2.BapiError
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing DeleteArea", err,
					"Failure at DeleteArea execution", bapiError))
			}
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

// fixKeyAccess(key + ".type") now is fixKeyAccess("area.type")
func expandRequestSiteCreateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSite {
	request := dnacentersdkgo.RequestSitesCreateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess("area.type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess("area.type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess("area.type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteCreateSiteSite(ctx, key+".site.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = "area"
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteCreateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSiteSite {
	request := dnacentersdkgo.RequestSitesCreateSiteSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		request.Area = expandRequestSiteCreateSiteSiteArea(ctx, key+".area.0", d)
	} /*
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
			request.Building = expandRequestSiteCreateSiteSiteBuilding(ctx, key+".building.0", d)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
			request.Floor = expandRequestSiteCreateSiteSiteFloor(ctx, key+".floor.0", d)
		}*/
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteCreateSiteSiteArea(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSiteSiteArea {
	request := dnacentersdkgo.RequestSitesCreateSiteSiteArea{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

// GetOkExists(fixKeyAccess(key + ".type")) now is GetOkExists(fixKeyAccess("area.type"))
func expandRequestSiteUpdateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSite {
	request := dnacentersdkgo.RequestSitesUpdateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess("area.type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess("area.type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess("area.type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteUpdateSiteSite(ctx, key+".site.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = "area"
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteUpdateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSite {
	request := dnacentersdkgo.RequestSitesUpdateSiteSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		request.Area = expandRequestSiteUpdateSiteSiteArea(ctx, key+".area.0", d)
	} /*
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
			request.Building = expandRequestSiteUpdateSiteSiteBuilding(ctx, key+".building.0", d)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
			request.Floor = expandRequestSiteUpdateSiteSiteFloor(ctx, key+".floor.0", d)
		}*/
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteUpdateSiteSiteArea(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteArea {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteArea{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSitesGetSite(m interface{}, queryParams dnacentersdkgo.GetSiteQueryParams) (*dnacentersdkgo.ResponseSitesGetSiteResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSitesGetSiteResponse
	// var ite *dnacentersdkgo.ResponseSitesGetSite
	if queryParams.SiteID == "" {
		ite, restyResp1, err := client.Sites.GetSite(&queryParams)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			}
			log.Printf("[DEBUG] Error =>%s", err.Error())
			return foundItem, err
		}
		items := ite.Response
		if items == nil {
			return foundItem, err
		}
		itemsCopy := *items
		for _, item := range itemsCopy {
			// Call get by _ method and set value to foundItem and return
			if item.SiteNameHierarchy == queryParams.Name {
				var getItem *dnacentersdkgo.ResponseSitesGetSiteResponse
				getItem = &item
				foundItem = getItem
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
