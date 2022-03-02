package dnacenter

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Sites.

- Creates site with area/building/floor with specified hierarchy.

- Update site area/building/floor with specified hierarchy and new values

- Delete site with area/building/floor by siteId.
`,

		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,
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
							Description: `Additional Info`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"address_inherited_from": &schema.Schema{
													Description: `addressInheritedFrom`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"country": &schema.Schema{
													Description: `country`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"floor_index": &schema.Schema{
													Description: `floorIndex`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"height": &schema.Schema{
													Description: `height`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"latitude": &schema.Schema{
													Description: `latitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"length": &schema.Schema{
													Description: `length`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"longitude": &schema.Schema{
													Description: `longitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"offset_x": &schema.Schema{
													Description: `offsetX`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"offset_y": &schema.Schema{
													Description: `offsetY`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"rf_model": &schema.Schema{
													Description: `rfModel`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"type": &schema.Schema{
													Description: `type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"width": &schema.Schema{
													Description: `width`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"namespace": &schema.Schema{
										Description: `namespace`,
										Type:        schema.TypeString,
										Computed:    true,
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
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"area": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name of the area (eg: Area1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of the area to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"building": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `Address of the building to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"latitude": &schema.Schema{
													Description: `Latitude coordinate of the building (eg:37.338)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"longitude": &schema.Schema{
													Description: `Longitude coordinate of the building (eg:-121.832)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the building (eg: building1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of building to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"floor": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"height": &schema.Schema{
													Description: `Height of the floor (eg: 15)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"length": &schema.Schema{
													Description: `Length of the floor (eg: 100)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the floor (eg:floor-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of the floor to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"rf_model": &schema.Schema{
													Description: `Type of floor. Allowed values are 'Cubes And Walled Offices', 'Drywall Office Only', 'Indoor High Ceiling', 'Outdoor Open Space'.
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"width": &schema.Schema{
													Description: `Width of the floor (eg:100)
`,
													Type:     schema.TypeFloat,
													Optional: true,
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
						},
						"type": &schema.Schema{
							Description: `Type of site to create (eg: area, building, floor)
`,
							Type:     schema.TypeString,
							Required: true,
						},
						/*"runsync": &schema.Schema{
													Description: `HeaderParam
						`,
													Type:     schema.TypeString,
													Required: true,
												},
												"timeout": &schema.Schema{
													Description: `HeaderParam
						`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"persistbapioutput": &schema.Schema{
													Description: `HeaderParam
						`,
													Type:     schema.TypeString,
													Required: true,
												},*/
					},
				},
			},
		},
	}
}

func resourceSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteCreateSite(ctx, "parameters.0", d)

	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vType := resourceItem["type"]
	vvType := interfaceToString(vType)
	/*vTimeout, okTimeout := resourceItem["timeout"]
	vvTimeout := interfaceToString(vTimeout)
	vRunsync := resourceItem["runsync"]
	vvRunsync := interfaceToString(vRunsync)
	vPersistbapioutput := resourceItem["persistbapioutput"]
	vvPersistbapioutput := interfaceToString(vPersistbapioutput)*/

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
	log.Printf("[DEBUG] newName => %s", newName)
	item, err := searchSitesGetSite(m, queryParams1)
	if err == nil || item != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_id"] = item.ID
		//resourceMap["type"] = item.AdditionalInfo
		resourceMap["name"] = item.SiteNameHierarchy
		/*resourceMap["runsync"] = vvRunsync
		resourceMap["persistbapioutput"] = vvPersistbapioutput
		resourceMap["timeout"] = vvTimeout*/
		d.SetId(joinResourceID(resourceMap))
		return resourceSiteRead(ctx, d, m)
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSite", err,
				"Failure at CreateSite execution", bapiError))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	//resourceMap["type"] = vvType
	resourceMap["name"] = newName
	//resourceMap["parent_name"] = vvParentName
	/*resourceMap["runsync"] = vvRunsync
	resourceMap["persistbapioutput"] = vvPersistbapioutput
	resourceMap["timeout"] = vvTimeout*/
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteRead(ctx, d, m)
}

func resourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]
	log.Printf("[DEBUG] Read SiteNameHierarchy => %s", vName)
	//vSiteID := resourceMap["site_id"]
	//vParentName := resourceMap["parent_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		//pathName := []string{vParentName, vName}
		//newName := strings.Join(pathName, "/")
		//log.Printf("[DEBUG] Selected method 1: GetSite")
		queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
		queryParams1.Name = vName
		//queryParams1.SiteID = vSiteID
		log.Printf("[DEBUG] Read name => %s", queryParams1.Name)
		log.Printf("[DEBUG] Read site => %s", queryParams1.SiteID)
		response1, restyResp1, err := client.Sites.GetSite(&queryParams1)

		if err != nil || response1 == nil {
			log.Printf("[DEBUG] Error => %s", err.Error())
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response3 %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetSiteItems(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSite search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]
	vSiteID := resourceMap["site_id"]
	//vParentName := resourceMap["parent_name"]
	/*vRunsync := resourceMap["runsync"]
	vPersistbapioutput := resourceMap["persistbapioutput"]
	vTimeout := resourceMap["timeout"]*/
	//pathName := []string{vParentName, vName}
	//newName := strings.Join(pathName, "/")
	queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
	queryParams1.Name = vName
	item, err := searchSitesGetSite(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetSite", err,
			"Failure at GetSite, unexpected response", ""))
		return diags
	}
	if vSiteID != item.ID {
		vSiteID = item.ID
	}
	var vvID string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
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

	return resourceSiteRead(ctx, d, m)
}

func resourceSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]
	vSiteID := resourceMap["site_id"]

	queryParams1 := dnacentersdkgo.GetSiteQueryParams{}
	queryParams1.Name = vName
	item, err := searchSitesGetSite(m, queryParams1)
	if err != nil || item == nil {
		d.SetId("")
		return diags
	}

	if vSiteID != item.ID {
		vSiteID = item.ID
	}

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

	if response1.Status == "FAILURE" {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSite", err,
			"Failure at DeleteSite, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSiteCreateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSite {
	request := dnacentersdkgo.RequestSitesCreateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteCreateSiteSite(ctx, key+".site.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteCreateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSiteSite {
	request := dnacentersdkgo.RequestSitesCreateSiteSite{}
	var typeStr string
	if typeS, ok := d.GetOkExists(fixKeyAccess("parameters.0.type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess("parameters.0.type")))) && (ok || !reflect.DeepEqual(typeS, d.Get(fixKeyAccess("parameters.0.type")))) {
		typeStr = interfaceToString(typeS)
	} else {
		return nil
	}

	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		if typeStr == "area" {
			request.Area = expandRequestSiteCreateSiteSiteArea(ctx, key+".area.0", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
		if typeStr == "building" {
			request.Building = expandRequestSiteCreateSiteSiteBuilding(ctx, key+".building.0", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
		if typeStr == "floor" {
			request.Floor = expandRequestSiteCreateSiteSiteFloor(ctx, key+".floor.0", d)
		}
	}
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

func expandRequestSiteCreateSiteSiteBuilding(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSiteSiteBuilding {
	request := dnacentersdkgo.RequestSitesCreateSiteSiteBuilding{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteCreateSiteSiteFloor(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesCreateSiteSiteFloor {
	request := dnacentersdkgo.RequestSitesCreateSiteSiteFloor{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_model")))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".width")))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".length")))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".height")))) {
		request.Height = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteUpdateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSite {
	request := dnacentersdkgo.RequestSitesUpdateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteUpdateSiteSite(ctx, key+".site.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteUpdateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSite {
	request := dnacentersdkgo.RequestSitesUpdateSiteSite{}
	var typeStr string
	if typeS, ok := d.GetOkExists(fixKeyAccess("parameters.0.type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess("parameters.0.type")))) && (ok || !reflect.DeepEqual(typeS, d.Get(fixKeyAccess("parameters.0.type")))) {
		typeStr = interfaceToString(typeS)
	} else {
		return nil
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		if typeStr == "area" {
			request.Area = expandRequestSiteUpdateSiteSiteArea(ctx, key+".area.0", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
		if typeStr == "building" {
			request.Building = expandRequestSiteUpdateSiteSiteBuilding(ctx, key+".building.0", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
		if typeStr == "floor" {
			request.Floor = expandRequestSiteUpdateSiteSiteFloor(ctx, key+".floor.0", d)
		}
	}
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

func expandRequestSiteUpdateSiteSiteBuilding(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteBuilding {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteBuilding{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteUpdateSiteSiteFloor(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteFloor {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteFloor{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_model")))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".width")))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".length")))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".height")))) {
		request.Height = interfaceToFloat64Ptr(v)
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
	var ite *dnacentersdkgo.ResponseSitesGetSite
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
	return foundItem, err
}
