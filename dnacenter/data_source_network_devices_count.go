package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Gets the total Network device counts. When there is no start and end time specified returns the latest interfaces
total count. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		ReadContext: dataSourceNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The List of Network Device model attributes. This is helps to specify the interested fields in the request.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"family": &schema.Schema{
				Description: `family query parameter. The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_score": &schema.Schema{
				Description: `healthScore query parameter. The list of entity health score categories Examples:healthScore=good,healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. The macAddress of the network device or client This field supports wildcard (***) character-based search.  Ex: **AB:AB:AB** or *AB:AB:AB** or **AB:AB:AB* Examples:
*macAddress=AB:AB:AB:CD:CD:CD* (single macAddress requested)
*macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE* (multiple macAddress requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_mode": &schema.Schema{
				Description: `maintenanceMode query parameter. The device maintenanceMode status true or false
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"management_ip_address": &schema.Schema{
				Description: `managementIpAddress query parameter. The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10")
This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Description: `role query parameter. The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. The list of network device serial numbers. This field supports wildcard (***) character-based search.  Ex: **MS1SV** or *MS1SV** or **MS1SV* Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San*
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (*) character search support. E.g. **uuid*, *uuid, uuid*
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid*
Examples:
*?siteId=id1* (single id requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter. The list of network device software version This field supports wildcard (***) character-based search. Ex: **17.8** or **17.8* or *17.8** Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *startTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. The list of network device type This field supports wildcard (***) character-based search. Ex: **9407R** or **9407R* or *9407R**Examples:type=SwitchesCisco Catalyst 9407R Switch (single network device types )type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The List of Network Device model views. Please refer to ***NetworkDeviceView*** for the supported list
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vID, okID := d.GetOk("id")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vManagementIPAddress, okManagementIPAddress := d.GetOk("management_ip_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vFamily, okFamily := d.GetOk("family")
	vType, okType := d.GetOk("type")
	vRole, okRole := d.GetOk("role")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vMaintenanceMode, okMaintenanceMode := d.GetOk("maintenance_mode")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vHealthScore, okHealthScore := d.GetOk("health_score")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters")
		queryParams1 := dnacentersdkgo.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okSiteHierarchy {
			queryParams1.SiteHierarchy = vSiteHierarchy.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okManagementIPAddress {
			queryParams1.ManagementIPAddress = vManagementIPAddress.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okMaintenanceMode {
			queryParams1.MaintenanceMode = vMaintenanceMode.(bool)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = vSoftwareVersion.(string)
		}
		if okHealthScore {
			queryParams1.HealthScore = vHealthScore.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}

		response1, restyResp1, err := client.Devices.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters", err,
				"Failure at GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersItem(item *dnacentersdkgo.ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
