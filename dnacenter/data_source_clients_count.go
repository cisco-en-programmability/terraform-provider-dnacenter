package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Retrieves the number of clients by applying basic filtering. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceClientsCountRead,
		Schema: map[string]*schema.Schema{
			"band": &schema.Schema{
				Description: `band query parameter. WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz GHz Examples:
*band=5GHZ* (single band requested)
*band=2.4GHZ&band=6GHZ* (multiple band requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"connected_network_device_name": &schema.Schema{
				Description: `connectedNetworkDeviceName query parameter. Name of the neighbor network device that client is connected to. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search. Ex: **ap-25** or *ap-25** or **ap-25*
Examples:
*connectedNetworkDeviceName=ap-25* (single connectedNetworkDeviceName requested)
*connectedNetworkDeviceName=ap-25&ap-34* (multiple connectedNetworkDeviceName requested)
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
			"ipv4_address": &schema.Schema{
				Description: `ipv4Address query parameter. IPv4 Address of the network entity either network device or client This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples:
*ipv4Address=1.1.1.1* (single ipv4Address requested)
*ipv4Address=1.1.1.1&ipv4Address=2.2.2.2* (multiple ipv4Address requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": &schema.Schema{
				Description: `ipv6Address query parameter. IPv6 Address of the network entity either network device or client This field supports wildcard (***) character-based search. Ex: **2001:db8** or *2001:db8** or **2001:db8*
Examples:
*ipv6Address=2001:db8:0:0:0:0:2:1* (single ipv6Address requested)
*ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348* (multiple ipv6Address requested)
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
			"os_type": &schema.Schema{
				Description: `osType query parameter. Client device operating system type. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **iOS** or *iOS** or **iOS* Examples:
*osType=iOS* (single osType requested)
*osType=iOS&osType=Android* (multiple osType requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_version": &schema.Schema{
				Description: `osVersion query parameter. Client device operating system version This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **14.3** or *14.3** or **14.3* Examples:
*osVersion=14.3* (single osVersion requested)
*osVersion=14.3&osVersion=10.1* (multiple osVersion requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **BuildingName** or *BuildingName** or **BuildingName* Examples: *siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested) *siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2* (multiple siteHierarchy requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (***) character-based search.  Ex: **buildingUuid** or *buildingUuid** or **buildingUuid* Examples: *siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid* (single siteHierarchyId requested) *siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2* (multiple siteHierarchyId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: *siteId=floorUuid* (single siteId requested) *siteId=floorUuid1&siteId=floorUuid2* (multiple siteId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID Wireless Local Area Network Identifier. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **Alpha** or *Alpha** or **Alpha*
Examples:
*ssid=Alpha* (single ssid requested)
*ssid=Alpha&ssid=Guest* (multiple ssid requested)
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
				Description: `type query parameter. The client device type whether client is connected to network through Wired or Wireless medium.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlc_name": &schema.Schema{
				Description: `wlcName query parameter. Wireless Controller name that reports the wireless client. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search. Ex: **wlc-25** or *wlc-25** or **wlc-25*
Examples:
*wlcName=wlc-25* (single wlcName requested)
*wlcName=wlc-25&wlc-34* (multiple wlcName requested)
`,
				Type:     schema.TypeString,
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

func dataSourceClientsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vType, okType := d.GetOk("type")
	vOsType, okOsType := d.GetOk("os_type")
	vOsVersion, okOsVersion := d.GetOk("os_version")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vIPv4Address, okIPv4Address := d.GetOk("ipv4_address")
	vIPv6Address, okIPv6Address := d.GetOk("ipv6_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vWlcName, okWlcName := d.GetOk("wlc_name")
	vConnectedNetworkDeviceName, okConnectedNetworkDeviceName := d.GetOk("connected_network_device_name")
	vSSID, okSSID := d.GetOk("ssid")
	vBand, okBand := d.GetOk("band")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheTotalCountOfClientsByApplyingBasicFiltering")

		headerParams1 := dnacentersdkgo.RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOsType {
			queryParams1.OsType = vOsType.(string)
		}
		if okOsVersion {
			queryParams1.OsVersion = vOsVersion.(string)
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
		if okIPv4Address {
			queryParams1.IPv4Address = vIPv4Address.(string)
		}
		if okIPv6Address {
			queryParams1.IPv6Address = vIPv6Address.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okWlcName {
			queryParams1.WlcName = vWlcName.(string)
		}
		if okConnectedNetworkDeviceName {
			queryParams1.ConnectedNetworkDeviceName = vConnectedNetworkDeviceName.(string)
		}
		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okBand {
			queryParams1.Band = vBand.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Clients.RetrievesTheTotalCountOfClientsByApplyingBasicFiltering(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheTotalCountOfClientsByApplyingBasicFiltering", err,
				"Failure at RetrievesTheTotalCountOfClientsByApplyingBasicFiltering, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheTotalCountOfClientsByApplyingBasicFiltering response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringItem(item *dnacentersdkgo.ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
