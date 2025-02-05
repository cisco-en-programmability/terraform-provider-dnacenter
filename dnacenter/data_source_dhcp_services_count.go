package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDhcpServicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieves the total number of DHCP Services for given parameters. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceDhcpServicesCountRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. The device UUID.

 Examples:
 *deviceId=6bef213c-19ca-4170-8375-b694e251101c* (single deviceId is requested)
 *deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9* (multiple networkDeviceIds with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_name": &schema.Schema{
				Description: `deviceName query parameter. Name of the device. This parameter supports wildcard (***) character -based search. Example: *wnbu-sjc* or *wnbu-sjc* or *wnbu-sjc* Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_site_hierarchy": &schema.Schema{
				Description: `deviceSiteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (***) character search support. E.g. */San*, */San*, /San**
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_site_hierarchy_id": &schema.Schema{
				Description: `deviceSiteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. *uuid*, *uuid*, *uuid*
Examples:
*?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_site_id": &schema.Schema{
				Description: `deviceSiteId query parameter. The UUID of the site. (Ex. *flooruuid*)
Examples:
*?deviceSiteIds=id1* (single id requested)
*?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3* (multiple ids requested)
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
			"server_ip": &schema.Schema{
				Description: `serverIp query parameter. IP Address of the DHCP Server. This parameter supports wildcard (***) character -based search. Example: *10.76.81.* or *56.78* or *50.2* Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceDhcpServicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vServerIP, okServerIP := d.GetOk("server_ip")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vDeviceName, okDeviceName := d.GetOk("device_name")
	vDeviceSiteHierarchy, okDeviceSiteHierarchy := d.GetOk("device_site_hierarchy")
	vDeviceSiteHierarchyID, okDeviceSiteHierarchyID := d.GetOk("device_site_hierarchy_id")
	vDeviceSiteID, okDeviceSiteID := d.GetOk("device_site_id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheTotalNumberOfDHCPServicesForGivenParameters")

		headerParams1 := dnacentersdkgo.RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okServerIP {
			queryParams1.ServerIP = vServerIP.(string)
		}
		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okDeviceName {
			queryParams1.DeviceName = vDeviceName.(string)
		}
		if okDeviceSiteHierarchy {
			queryParams1.DeviceSiteHierarchy = vDeviceSiteHierarchy.(string)
		}
		if okDeviceSiteHierarchyID {
			queryParams1.DeviceSiteHierarchyID = vDeviceSiteHierarchyID.(string)
		}
		if okDeviceSiteID {
			queryParams1.DeviceSiteID = vDeviceSiteID.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.RetrievesTheTotalNumberOfDHCPServicesForGivenParameters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheTotalNumberOfDHCPServicesForGivenParameters", err,
				"Failure at RetrievesTheTotalNumberOfDHCPServicesForGivenParameters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheTotalNumberOfDHCPServicesForGivenParameters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersItem(item *dnacentersdkgo.ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
