package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieves the list of the interfaces from all network devices based on the provided query parameters. The latest
interfaces data in the specified start and end time range will be returned. When there is no start and end time
specified returns the latest available data.
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default
value: name.

 The supported sorting options are: name, adminStatus, description, duplexConfig,
duplexOper,interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus,portChannelId, portMode, portType,speed,
vlanId. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
interfaces-1.0.2-resolved.yaml

- Returns the interface data for the given interface instance Uuid along with the statistics data. The latest interface
data in the specified start and end time range will be returned. When there is no start and end time specified returns
the latest available data for the given interface Id. For detailed information about the usage of the API, please refer
to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml
`,

		ReadContext: dataSourceInterfacesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The following list of attributes can be provided in the attribute field
[id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]
If length of attribute list is too long, please use 'views' param instead.
Examples:
attributes=name (single attribute requested)
attributes=name,description,duplexOper (multiple attributes with comma separator)
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
			"id": &schema.Schema{
				Description: `id path parameter. The interface Uuid
`,
				Type:     schema.TypeString,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
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
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
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
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
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
			"view": &schema.Schema{
				Description: `view query parameter. Interface data model views
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Description: `Admin Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duplex_config": &schema.Schema{
							Description: `Duplex Config`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duplex_oper": &schema.Schema{
							Description: `Duplex Oper`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interface_if_index": &schema.Schema{
							Description: `Interface If Index`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"interface_type": &schema.Schema{
							Description: `Interface Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv4_address": &schema.Schema{
							Description: `Ipv4 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6_address_list": &schema.Schema{
							Description: `Ipv6 Address List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"is_l3_interface": &schema.Schema{
							Description: `Is L3 Interface`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_wan": &schema.Schema{
							Description: `Is Wan`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_addr": &schema.Schema{
							Description: `Mac Addr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"media_type": &schema.Schema{
							Description: `Media Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_ip_address": &schema.Schema{
							Description: `Network Device Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_mac_address": &schema.Schema{
							Description: `Network Device Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"oper_status": &schema.Schema{
							Description: `Oper Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"peer_stack_member": &schema.Schema{
							Description: `Peer Stack Member`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"peer_stack_port": &schema.Schema{
							Description: `Peer Stack Port`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_channel_id": &schema.Schema{
							Description: `Port Channel Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_mode": &schema.Schema{
							Description: `Port Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_type": &schema.Schema{
							Description: `Port Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rx_discards": &schema.Schema{
							Description: `Rx Discards`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"rx_error": &schema.Schema{
							Description: `Rx Error`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"rx_rate": &schema.Schema{
							Description: `Rx Rate`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"rx_utilization": &schema.Schema{
							Description: `Rx Utilization`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"speed": &schema.Schema{
							Description: `Speed`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"stack_port_type": &schema.Schema{
							Description: `Stack Port Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"tx_discards": &schema.Schema{
							Description: `Tx Discards`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"tx_error": &schema.Schema{
							Description: `Tx Error`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"tx_rate": &schema.Schema{
							Description: `Tx Rate`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"tx_utilization": &schema.Schema{
							Description: `Tx Utilization`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Description: `Admin Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duplex_config": &schema.Schema{
							Description: `Duplex Config`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duplex_oper": &schema.Schema{
							Description: `Duplex Oper`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interface_if_index": &schema.Schema{
							Description: `Interface If Index`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"interface_type": &schema.Schema{
							Description: `Interface Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv4_address": &schema.Schema{
							Description: `Ipv4 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6_address_list": &schema.Schema{
							Description: `Ipv6 Address List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"is_l3_interface": &schema.Schema{
							Description: `Is L3 Interface`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_wan": &schema.Schema{
							Description: `Is Wan`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_addr": &schema.Schema{
							Description: `Mac Addr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"media_type": &schema.Schema{
							Description: `Media Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_ip_address": &schema.Schema{
							Description: `Network Device Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_mac_address": &schema.Schema{
							Description: `Network Device Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"oper_status": &schema.Schema{
							Description: `Oper Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"peer_stack_member": &schema.Schema{
							Description: `Peer Stack Member`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"peer_stack_port": &schema.Schema{
							Description: `Peer Stack Port`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_channel_id": &schema.Schema{
							Description: `Port Channel Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_mode": &schema.Schema{
							Description: `Port Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_type": &schema.Schema{
							Description: `Port Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rx_discards": &schema.Schema{
							Description: `Rx Discards`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"rx_error": &schema.Schema{
							Description: `Rx Error`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"rx_rate": &schema.Schema{
							Description: `Rx Rate`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"rx_utilization": &schema.Schema{
							Description: `Rx Utilization`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"speed": &schema.Schema{
							Description: `Speed`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"stack_port_type": &schema.Schema{
							Description: `Stack Port Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"tx_discards": &schema.Schema{
							Description: `Tx Discards`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"tx_error": &schema.Schema{
							Description: `Tx Error`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"tx_rate": &schema.Schema{
							Description: `Tx Rate`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"tx_utilization": &schema.Schema{
							Description: `Tx Utilization`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vNetworkDeviceIPAddress, okNetworkDeviceIPAddress := d.GetOk("network_device_ip_address")
	vNetworkDeviceMacAddress, okNetworkDeviceMacAddress := d.GetOk("network_device_mac_address")
	vInterfaceID, okInterfaceID := d.GetOk("interface_id")
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okStartTime, okEndTime, okLimit, okOffset, okSortBy, okOrder, okSiteHierarchy, okSiteHierarchyID, okSiteID, okView, okAttribute, okNetworkDeviceID, okNetworkDeviceIPAddress, okNetworkDeviceMacAddress, okInterfaceID, okInterfaceName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okStartTime, okEndTime, okView, okAttribute}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices")
		queryParams1 := dnacentersdkgo.GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
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
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
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

		response1, restyResp1, err := client.Devices.GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices", err,
				"Failure at GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData")
		vvID := vID.(string)
		queryParams2 := dnacentersdkgo.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataQueryParams{}

		if okStartTime {
			queryParams2.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams2.EndTime = vEndTime.(float64)
		}
		if okView {
			queryParams2.View = vView.(string)
		}
		if okAttribute {
			queryParams2.Attribute = vAttribute.(string)
		}

		response2, restyResp2, err := client.Devices.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData(vvID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData", err,
				"Failure at GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesItems(items *[]dnacentersdkgo.ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["admin_status"] = item.AdminStatus
		respItem["description"] = item.Description
		respItem["duplex_config"] = item.DuplexConfig
		respItem["duplex_oper"] = item.DuplexOper
		respItem["interface_if_index"] = item.InterfaceIfIndex
		respItem["interface_type"] = item.InterfaceType
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_address_list"] = item.IPv6AddressList
		respItem["is_l3_interface"] = boolPtrToString(item.IsL3Interface)
		respItem["is_wan"] = boolPtrToString(item.IsWan)
		respItem["mac_addr"] = item.MacAddr
		respItem["media_type"] = item.MediaType
		respItem["name"] = item.Name
		respItem["oper_status"] = item.OperStatus
		respItem["peer_stack_member"] = item.PeerStackMember
		respItem["peer_stack_port"] = item.PeerStackPort
		respItem["port_channel_id"] = item.PortChannelID
		respItem["port_mode"] = item.PortMode
		respItem["port_type"] = item.PortType
		respItem["rx_discards"] = item.RxDiscards
		respItem["rx_error"] = item.RxError
		respItem["rx_rate"] = item.RxRate
		respItem["rx_utilization"] = item.RxUtilization
		respItem["speed"] = item.Speed
		respItem["stack_port_type"] = item.StackPortType
		respItem["timestamp"] = item.Timestamp
		respItem["tx_discards"] = item.TxDiscards
		respItem["tx_error"] = item.TxError
		respItem["tx_rate"] = item.TxRate
		respItem["tx_utilization"] = item.TxUtilization
		respItem["vlan_id"] = item.VLANID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["network_device_ip_address"] = item.NetworkDeviceIPAddress
		respItem["network_device_mac_address"] = item.NetworkDeviceMacAddress
		respItem["site_name"] = item.SiteName
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataItem(item *dnacentersdkgo.ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["admin_status"] = item.AdminStatus
	respItem["description"] = item.Description
	respItem["duplex_config"] = item.DuplexConfig
	respItem["duplex_oper"] = item.DuplexOper
	respItem["interface_if_index"] = item.InterfaceIfIndex
	respItem["interface_type"] = item.InterfaceType
	respItem["ipv4_address"] = item.IPv4Address
	respItem["ipv6_address_list"] = item.IPv6AddressList
	respItem["is_l3_interface"] = boolPtrToString(item.IsL3Interface)
	respItem["is_wan"] = boolPtrToString(item.IsWan)
	respItem["mac_addr"] = item.MacAddr
	respItem["media_type"] = item.MediaType
	respItem["name"] = item.Name
	respItem["oper_status"] = item.OperStatus
	respItem["peer_stack_member"] = item.PeerStackMember
	respItem["peer_stack_port"] = item.PeerStackPort
	respItem["port_channel_id"] = item.PortChannelID
	respItem["port_mode"] = item.PortMode
	respItem["port_type"] = item.PortType
	respItem["rx_discards"] = item.RxDiscards
	respItem["rx_error"] = item.RxError
	respItem["rx_rate"] = item.RxRate
	respItem["rx_utilization"] = item.RxUtilization
	respItem["speed"] = item.Speed
	respItem["stack_port_type"] = item.StackPortType
	respItem["timestamp"] = item.Timestamp
	respItem["tx_discards"] = item.TxDiscards
	respItem["tx_error"] = item.TxError
	respItem["tx_rate"] = item.TxRate
	respItem["tx_utilization"] = item.TxUtilization
	respItem["vlan_id"] = item.VLANID
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["network_device_ip_address"] = item.NetworkDeviceIPAddress
	respItem["network_device_mac_address"] = item.NetworkDeviceMacAddress
	respItem["site_name"] = item.SiteName
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["site_hierarchy_id"] = item.SiteHierarchyID
	return []map[string]interface{}{
		respItem,
	}
}
