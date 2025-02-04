package dnacenter

import (
	"context"
	"strconv"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSite() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters are not given as an input.
`,

		ReadContext: dataSourceSiteRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of sites to be listed
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Site name hierarchy (E.g Global/USA/CA)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset/starting index for pagination. Indexed from 1.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Site type (Ex: area, building, floor)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_info": &schema.Schema{
							Description: `Additional Info`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
		},
	}
}

func dataSourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vSiteID, okSiteID := d.GetOk("site_id")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSite")
		queryParams1 := dnacentersdkgo.GetSiteQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		response1, restyResp1, err := client.Sites.GetSite(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSite", err,
				"Failure at GetSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetSiteItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteItems(items *[]dnacentersdkgo.ResponseSitesGetSiteResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["parent_id"] = item.ParentID
		respItem["name"] = item.Name
		respItem["additional_info"] = flattenSitesGetSiteItemsAdditionalInfo(item.AdditionalInfo, nil)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenSitesGetFloorItems(items *[]dnacentersdkgo.ResponseSitesGetFloorResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["parent_id"] = item.ParentID
		respItem["name"] = item.Name
		respItem["additional_info"] = flattenSitesGetFloorItemsAdditionalInfo(item.AdditionalInfo)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenSitesGetAreaItems(items *[]dnacentersdkgo.ResponseSitesGetAreaResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["parent_id"] = item.ParentID
		respItem["name"] = item.Name
		respItem["additional_info"] = flattenSitesGetAreaItemsAdditionalInfo(item.AdditionalInfo, nil)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenSitesGetSiteParams(items *[]dnacentersdkgo.ResponseSitesGetSiteResponse, parameters []interface{}) map[string]interface{} {
	respParams := make(map[string]interface{})
	buildings := make([]map[string]interface{}, 0)
	parentName := getParametersOfLastUpdatedBuilding(parameters, "parent_name", "building")

	for _, item := range *items {
		for _, additionalInfo := range item.AdditionalInfo {
			attributes := additionalInfo.Attributes
			latitudeStr := attributes.Latitude
			longitudeStr := attributes.Longitude
			latitude, err := strconv.ParseFloat(latitudeStr, 64)
			if err != nil {
				log.Printf("Error in parse float latitude")
			}

			longitude, err := strconv.ParseFloat(longitudeStr, 64)
			if err != nil {
				log.Printf("Error in parse float longitude")
			}
			building := map[string]interface{}{
				"address":     attributes.Address,
				"latitude":    latitude,
				"longitude":   longitude,
				"name":        item.Name,
				"parent_name": parentName,
				// "type":        attributes.Type,
			}
			buildings = append(buildings, building)
		}
	}

	respParams["site"] = []map[string]interface{}{
		{
			"building": buildings,
		},
	}
	respParams["type"] = "building"
	if len(parameters) > 0 {
		if parameters[0].(map[string]interface{})["site_id"] == nil || parameters[0].(map[string]interface{})["site_id"].(string) == "" {
			log.Print("[DEBUG] site_id Not setted")
		} else {
			respParams["site_id"] = parameters[0].(map[string]interface{})["site_id"].(string)
		}
	}

	return respParams

}

func flattenSitesGetFloorParams(items *[]dnacentersdkgo.ResponseSitesGetFloorResponse, parameters []interface{}) map[string]interface{} {
	respParams := make(map[string]interface{})
	parentName := getParametersOfLastUpdatedBuilding(parameters, "parent_name", "floor")
	rfModel := getParametersOfLastUpdatedBuilding(parameters, "rf_model", "floor")
	floors := make([]map[string]interface{}, 0)

	for _, item := range *items {
		var name string
		//var rfModel string
		var height float64
		var width float64
		var length float64
		var floorNumber float64
		//var respFloor []map[string]interface{}

		name = item.Name

		for _, additionalInfo := range item.AdditionalInfo {
			attributes := additionalInfo.Attributes
			if additionalInfo.Namespace == "mapGeometry" {
				height_, err := strconv.ParseFloat(attributes.Height, 64)
				if err != nil {
					log.Printf("Error in parse float height")
				}

				length_, err := strconv.ParseFloat(attributes.Length, 64)
				if err != nil {
					log.Printf("Error in parse float length")
				}

				width_, err := strconv.ParseFloat(attributes.Width, 64)
				if err != nil {
					log.Printf("Error in parse float width")
				}

				width = width_
				length = length_
				height = height_

			} else if additionalInfo.Namespace == "mapsSummary" {
				floorNumber_, err := strconv.ParseFloat(attributes.FloorIndex, 64)
				if err != nil {
					log.Printf("Error in parse float floorNumber")
				}
				floorNumber = floorNumber_
			}
		}
		floor := map[string]interface{}{
			"floor_number": floorNumber,
			"height":       height,
			"length":       length,
			"name":         name,
			"parent_name":  parentName,
			"rf_model":     rfModel,
			"width":        width,
		}
		floors = append(floors, floor)
	}
	respParams["site"] = []map[string]interface{}{
		{
			"floor": floors,
		},
	}
	respParams["type"] = "floor"
	if len(parameters) > 0 {
		if parameters[0].(map[string]interface{})["site_id"] == nil || parameters[0].(map[string]interface{})["site_id"].(string) == "" {
			log.Print("[DEBUG] site_id Not setted")
		} else {
			respParams["site_id"] = parameters[0].(map[string]interface{})["site_id"].(string)
		}
	}

	return respParams

}

func flattenSitesGetAreaParams(items *[]dnacentersdkgo.ResponseSitesGetAreaResponse, parameters []interface{}) map[string]interface{} {
	respParams := make(map[string]interface{})
	areas := make([]map[string]interface{}, 0)
	parentName := getParametersOfLastUpdatedBuilding(parameters, "parent_name", "area")
	for _, item := range *items {
		area := map[string]interface{}{
			"name":        item.Name,
			"parent_name": parentName,
			// "type":        attributes.Type,
		}
		areas = append(areas, area)
	}

	respParams["site"] = []map[string]interface{}{
		{
			"area": areas,
		},
	}
	respParams["type"] = "area"
	if len(parameters) > 0 {
		if parameters[0].(map[string]interface{})["site_id"] == nil || parameters[0].(map[string]interface{})["site_id"].(string) == "" {
			log.Print("[DEBUG] site_id Not setted")
		} else {
			respParams["site_id"] = parameters[0].(map[string]interface{})["site_id"].(string)
		}
	}
	return respParams

}

func flattenSitesGetSiteItemsAdditionalInfo(items []dnacentersdkgo.ResponseSitesGetSiteResponseAdditionalInfo, parameters []interface{}) []map[string]interface{} {
	var respItems []map[string]interface{}
	var parentName string
	if parameters != nil {
		parentName = getParametersOfLastUpdatedBuilding(parameters, "parent_name", "building")
	} else {
		parentName = ""
	}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name_space"] = item.Namespace
		respItem["attributes"] = []map[string]interface{}{
			{
				"country":              item.Attributes.Country,
				"address":              item.Attributes.Address,
				"latitude":             item.Attributes.Latitude,
				"addressinheritedfrom": item.Attributes.AddressInheritedFrom,
				"type":                 item.Attributes.Type,
				"longitude":            item.Attributes.Longitude,
				"offsetx":              item.Attributes.OffsetX,
				"offsety":              item.Attributes.OffsetY,
				"length":               item.Attributes.Length,
				"width":                item.Attributes.Width,
				"height":               item.Attributes.Height,
				"rfmodel":              item.Attributes.RfModel,
				"floorindex":           item.Attributes.FloorIndex,
				"parent_name":          parentName,
			},
		}
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetFloorItemsAdditionalInfo(items []dnacentersdkgo.ResponseSitesGetFloorResponseAdditionalInfo) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name_space"] = item.Namespace
		respItem["attributes"] = []map[string]interface{}{
			{
				"floor_number": item.Attributes.FloorIndex,
				"height":       item.Attributes.Height,
				"length":       item.Attributes.Length,
				"name":         item.Attributes.Name,
				"parent_name":  item.Attributes.ParentName,
				"rf_model":     item.Attributes.RfModel,
				"width":        item.Attributes.Width,
			},
		}
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetAreaItemsAdditionalInfo(items []dnacentersdkgo.ResponseSitesGetAreaResponseAdditionalInfo, parameters []interface{}) []map[string]interface{} {
	var respItems []map[string]interface{}
	var parentName string
	if parameters != nil {
		parentName = getParametersOfLastUpdatedBuilding(parameters, "parent_name", "area")
	} else {
		parentName = ""
	}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["name_space"] = item.Namespace
		respItem["attributes"] = []map[string]interface{}{
			{
				"addressinheritedfrom": item.Attributes.AddressInheritedFrom,
				"parent_name":          parentName,
				"type":                 item.Attributes.Type,
				"name":                 item.Attributes.Name,
			},
		}
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteItemsAdditionalInfoAtributes(item *dnacentersdkgo.ResponseSitesGetSiteResponseAdditionalInfoAttributes) map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["country"] = item.Country
	respItem["address"] = item.Address
	respItem["latitude"] = item.Latitude
	respItem["addressInheritedFrom"] = item.AddressInheritedFrom
	respItem["type"] = item.Longitude
	respItem["offsetX"] = item.OffsetX
	respItem["offsetY"] = item.OffsetY
	respItem["length"] = item.Length
	respItem["width"] = item.Width
	respItem["height"] = item.Height
	respItem["rfModel"] = item.RfModel
	respItem["rfModel"] = item.RfModel
	respItem["floorIndex"] = item.FloorIndex

	return respItem
}

func getParametersOfLastUpdatedBuilding(parameters []interface{}, searchValue string, typeSite string) string {
	for _, param := range parameters {
		paramData, ok := param.(map[string]interface{})
		if !ok {
			log.Printf("Error in paramData")
			continue
		}

		siteDataArray, ok := paramData["site"].([]interface{})
		if !ok || len(siteDataArray) == 0 {
			log.Printf("Error in enter site")
			continue
		}

		siteData, ok := siteDataArray[0].(map[string]interface{})
		if !ok {
			log.Printf("Error in enter data site ")
			continue
		}

		resourceDataArray, ok := siteData[typeSite].([]interface{})
		if !ok || len(resourceDataArray) == 0 {
			log.Printf("Error in enter resource ")
			continue
		}

		resourceData, ok := resourceDataArray[0].(map[string]interface{})
		if !ok {
			log.Printf("Error in enter data resource ")
			continue
		}

		value, ok := resourceData[searchValue].(string)
		if !ok {
			log.Printf("Error in get value  ")
			continue
		}

		return value
	}

	return ""

}
