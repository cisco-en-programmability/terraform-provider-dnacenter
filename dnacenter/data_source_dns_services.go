package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDNSServices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieves the list of DNS Services and offers basic filtering and sorting capabilities. For detailed information about
the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceDNSServicesRead,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
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
			"server_ip": &schema.Schema{
				Description: `serverIp query parameter. IP Address of the DNS Server. This parameter supports wildcard (***) character -based search. Example: *10.76.81.* or *56.78* or *50.28* Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Field name on which sorting needs to be done
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID Wireless Local Area Network Identifier. This field supports wildcard (***) character-based search. If the field contains the (***) character, please use the /query API for search. Ex: *Alpha* or *Alpha* or *Alpha*
Examples:
*ssid=Alpha* (single ssid requested)
*ssid=Alpha&ssid=Guest* (multiple ssid requested)
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_name": &schema.Schema{
							Description: `Device Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_hierarchy": &schema.Schema{
							Description: `Device Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_hierarchy_id": &schema.Schema{
							Description: `Device Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_id": &schema.Schema{
							Description: `Device Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failed_transactions": &schema.Schema{
							Description: `Failed Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"failures": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failed_transactions": &schema.Schema{
										Description: `Failed Transactions`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_description": &schema.Schema{
										Description: `Failure Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"failure_response_code": &schema.Schema{
										Description: `Failure Response Code`,
										Type:        schema.TypeInt,
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

						"latency": &schema.Schema{
							Description: `Latency`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"server_ip": &schema.Schema{
							Description: `Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"successful_transactions": &schema.Schema{
							Description: `Successful Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transactions": &schema.Schema{
							Description: `Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDNSServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vServerIP, okServerIP := d.GetOk("server_ip")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vDeviceSiteHierarchyID, okDeviceSiteHierarchyID := d.GetOk("device_site_hierarchy_id")
	vDeviceSiteID, okDeviceSiteID := d.GetOk("device_site_id")
	vSSID, okSSID := d.GetOk("ssid")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfDNSServicesForGivenParameters")

		headerParams1 := dnacentersdkgo.RetrievesTheListOfDNSServicesForGivenParametersHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheListOfDNSServicesForGivenParametersQueryParams{}

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
		if okServerIP {
			queryParams1.ServerIP = vServerIP.(string)
		}
		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okDeviceSiteHierarchyID {
			queryParams1.DeviceSiteHierarchyID = vDeviceSiteHierarchyID.(string)
		}
		if okDeviceSiteID {
			queryParams1.DeviceSiteID = vDeviceSiteID.(string)
		}
		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.RetrievesTheListOfDNSServicesForGivenParameters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfDNSServicesForGivenParameters", err,
				"Failure at RetrievesTheListOfDNSServicesForGivenParameters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesRetrievesTheListOfDNSServicesForGivenParametersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfDNSServicesForGivenParameters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrievesTheListOfDNSServicesForGivenParametersItems(items *[]dnacentersdkgo.ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["server_ip"] = item.ServerIP
		respItem["device_id"] = item.DeviceID
		respItem["device_name"] = item.DeviceName
		respItem["device_family"] = item.DeviceFamily
		respItem["device_site_hierarchy"] = item.DeviceSiteHierarchy
		respItem["device_site_id"] = item.DeviceSiteID
		respItem["device_site_hierarchy_id"] = item.DeviceSiteHierarchyID
		respItem["transactions"] = item.Transactions
		respItem["failed_transactions"] = item.FailedTransactions
		respItem["failures"] = flattenDevicesRetrievesTheListOfDNSServicesForGivenParametersItemsFailures(item.Failures)
		respItem["successful_transactions"] = item.SuccessfulTransactions
		respItem["latency"] = item.Latency
		respItem["ssid"] = item.SSID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesRetrievesTheListOfDNSServicesForGivenParametersItemsFailures(items *[]dnacentersdkgo.ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponseFailures) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["failure_response_code"] = item.FailureResponseCode
		respItem["failure_description"] = item.FailureDescription
		respItem["failed_transactions"] = item.FailedTransactions
		respItems = append(respItems, respItem)
	}
	return respItems
}
