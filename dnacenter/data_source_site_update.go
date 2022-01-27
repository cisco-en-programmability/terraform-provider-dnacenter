package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSiteUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sites.

- Update site area/building/floor with specified hierarchy and new values
`,

		ReadContext: dataSourceSiteUpdateRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site id to which site details to be updated.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"runsync": &schema.Schema{
				Description: `__runsync header parameter. Enable this parameter to execute the API and return a response synchronously
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"timeout": &schema.Schema{
				Description: `__timeout header parameter. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
			`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Description: `Data`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"is_error": &schema.Schema{
							Description: `Is Error`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"operation_id_list": &schema.Schema{
							Description: `Operation Id List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"progress": &schema.Schema{
							Description: `Progress`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"root_id": &schema.Schema{
							Description: `Root Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"service_type": &schema.Schema{
							Description: `Service Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"site": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"area": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"parent_name": &schema.Schema{
										Description: `Parent Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"building": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": &schema.Schema{
										Description: `Address`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"latitude": &schema.Schema{
										Description: `Latitude`,
										Type:        schema.TypeFloat,
										Optional:    true,
									},
									"longitude": &schema.Schema{
										Description: `Longitude`,
										Type:        schema.TypeFloat,
										Optional:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"parent_name": &schema.Schema{
										Description: `Parent Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"floor": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"height": &schema.Schema{
										Description: `Height`,
										Type:        schema.TypeFloat,
										Optional:    true,
									},
									"length": &schema.Schema{
										Description: `Length`,
										Type:        schema.TypeFloat,
										Optional:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"rf_model": &schema.Schema{
										Description: `Rf Model. Allowed values are 'Cubes And Walled Offices', 'Drywall Office Only', 'Indoor High Ceiling', 'Outdoor Open Space'
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"width": &schema.Schema{
										Description: `Width`,
										Type:        schema.TypeFloat,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"type": &schema.Schema{
				Description: `Type`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSiteUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vRunsync, okRunsync := d.GetOk("runsync")
	vTimeout, okTimeout := d.GetOk("timeout")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateSite")
		vvSiteID := vSiteID.(string)
		request1 := expandRequestSiteUpdateUpdateSite(ctx, "", d)

		headerParams1 := dnacentersdkgo.UpdateSiteHeaderParams{}
		if okRunsync {
			headerParams1.Runsync = vRunsync.(string)
		}
		if okTimeout {
			headerParams1.Timeout = vTimeout.(string)
		}
		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.Sites.UpdateSite(vvSiteID, request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSite", err,
				"Failure at UpdateSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesUpdateSiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateSite response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSiteUpdateUpdateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSite {
	request := dnacentersdkgo.RequestSitesUpdateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteUpdateUpdateSiteSite(ctx, key+".site.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteUpdateUpdateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSite {
	request := dnacentersdkgo.RequestSitesUpdateSiteSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		request.Area = expandRequestSiteUpdateUpdateSiteSiteArea(ctx, key+".area.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
		request.Building = expandRequestSiteUpdateUpdateSiteSiteBuilding(ctx, key+".building.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
		request.Floor = expandRequestSiteUpdateUpdateSiteSiteFloor(ctx, key+".floor.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteUpdateUpdateSiteSiteArea(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteArea {
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

func expandRequestSiteUpdateUpdateSiteSiteBuilding(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteBuilding {
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

func expandRequestSiteUpdateUpdateSiteSiteFloor(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteFloor {
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

func flattenSitesUpdateSiteItem(item *dnacentersdkgo.ResponseSitesUpdateSiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["version"] = item.Version
	respItem["start_time"] = item.StartTime
	respItem["progress"] = item.Progress
	respItem["data"] = item.Data
	respItem["service_type"] = item.ServiceType
	respItem["operation_id_list"] = item.OperationIDList
	respItem["is_error"] = item.IsError
	respItem["root_id"] = item.RootID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
