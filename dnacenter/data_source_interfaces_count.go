package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterfacesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Gets the total Network device interface counts. For detailed information about the usage of the API, please refer to
the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml
`,

		ReadContext: dataSourceInterfacesCountRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"interface_id": &schema.Schema{
				Description: `interfaceId query parameter. The list of Interface Uuids. (Ex. *6bef213c-19ca-4170-8375-b694e251101c*)
Examples:
*interfaceId=6bef213c-19ca-4170-8375-b694e251101c* (single interface uuid )
*interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0* (multiple Interface uuid with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_name": &schema.Schema{
				Description: `interfaceName query parameter. The list of Interface name (Ex. *GigabitEthernet1/0/1*) This field supports wildcard (***) character-based search.  Ex: **1/0/1** or *1/0/1** or **1/0/1*
Examples:
*interfaceNames=GigabitEthernet1/0/1* (single interface name)
*interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1* (multiple interface names with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. The list of Network Device Uuids. (Ex. *6bef213c-19ca-4170-8375-b694e251101c*)
Examples:
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c* (single networkDeviceId requested)
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0* (multiple networkDeviceIds with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_ip_address": &schema.Schema{
				Description: `networkDeviceIpAddress query parameter. The list of Network Device management IP Address. (Ex. *121.1.1.10*)
This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples:
*networkDeviceIpAddress=121.1.1.10*
*networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10* (multiple networkDevice IP Address with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_mac_address": &schema.Schema{
				Description: `networkDeviceMacAddress query parameter. The list of Network Device MAC Address. (Ex. *64:f6:9d:07:9a:00*)
This field supports wildcard (***) character-based search.  Ex: **AB:AB:AB** or *AB:AB:AB** or **AB:AB:AB*
Examples:
*networkDeviceMacAddress=64:f6:9d:07:9a:00*
*networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77* (multiple networkDevice MAC addresses with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (***) character search support. E.g. **/San*, */San, /San**
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. **uuid*, *uuid, uuid**
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
Examples:
*?siteId=id1* (single id requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple ids requested)
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

func dataSourceInterfacesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vNetworkDeviceIPAddress, okNetworkDeviceIPAddress := d.GetOk("network_device_ip_address")
	vNetworkDeviceMacAddress, okNetworkDeviceMacAddress := d.GetOk("network_device_mac_address")
	vInterfaceID, okInterfaceID := d.GetOk("interface_id")
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount")
		queryParams1 := dnacentersdkgo.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
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
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okNetworkDeviceIPAddress {
			queryParams1.NetworkDeviceIPAddress = vNetworkDeviceIPAddress.(string)
		}
		if okNetworkDeviceMacAddress {
			queryParams1.NetworkDeviceMacAddress = vNetworkDeviceMacAddress.(string)
		}
		if okInterfaceID {
			queryParams1.InterfaceID = vInterfaceID.(string)
		}
		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName.(string)
		}

		response1, restyResp1, err := client.Devices.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount", err,
				"Failure at GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountItem(item *dnacentersdkgo.ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
