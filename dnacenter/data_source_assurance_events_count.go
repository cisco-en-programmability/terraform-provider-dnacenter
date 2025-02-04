package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceEventsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support
Documentation' section to understand which fields are supported. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceEventsCountRead,
		Schema: map[string]*schema.Schema{
			"ap_mac": &schema.Schema{
				Description: `apMac query parameter. MAC address of the access point. This parameter is applicable for *Unified AP* and *Wireless Client* events.
This field supports wildcard (***) character-based search. Ex: **50:0F** or *50:0F** or **50:0F*
Examples:
*apMac=50:0F:80:0F:F7:E0* (single apMac requested)
*apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0* (multiple apMac requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_mac": &schema.Schema{
				Description: `clientMac query parameter. MAC address of the client. This parameter is applicable for *Wired Client* and *Wireless Client* events.
This field supports wildcard (***) character-based search. Ex: **66:2B** or *66:2B** or **66:2B*
Examples:
*clientMac=66:2B:B8:D2:01:56* (single clientMac requested)
*clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89* (multiple clientMac requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_family": &schema.Schema{
				Description: `deviceFamily query parameter. Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing *Routers* along with *Wireless Client* or *Unified AP* is not supported. Examples:
*deviceFamily=Switches and Hubs* (single deviceFamily requested)
*deviceFamily=Switches and Hubs&deviceFamily=Routers* (multiple deviceFamily requested)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *endTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"message_type": &schema.Schema{
				Description: `messageType query parameter. Message type for the event.
Examples:
*messageType=Syslog* (single messageType requested)
*messageType=Trap&messageType=Syslog* (multiple messageType requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. The list of Network Device Uuids. (Ex. *6bef213c-19ca-4170-8375-b694e251101c*)
Examples:
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c* (single networkDeviceId requested)
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0* (multiple networkDeviceId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_name": &schema.Schema{
				Description: `networkDeviceName query parameter. Network device name. This parameter is applicable for network device related families. This field supports wildcard (***) character-based search. Ex: **Branch** or *Branch** or **Branch* Examples:
*networkDeviceName=Branch-3-Gateway* (single networkDeviceName requested)
*networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch* (multiple networkDeviceName requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and *Wired Client* events.
| Value | Severity    | | ----| ----------| | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        |
Examples:
*severity=0* (single severity requested)
*severity=0&severity=1* (multiple severity requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. **uuid*, *uuid, uuid**
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
Examples:
*?siteId=id1* (single siteId requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple siteId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *startTime* is not provided, API will default to current time minus 24 hours.
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

func dataSourceAssuranceEventsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceFamily := d.Get("device_family")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vMessageType, okMessageType := d.GetOk("message_type")
	vSeverity, okSeverity := d.GetOk("severity")
	vSiteID, okSiteID := d.GetOk("site_id")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vNetworkDeviceName, okNetworkDeviceName := d.GetOk("network_device_name")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vApMac, okApMac := d.GetOk("ap_mac")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountTheNumberOfEvents")

		headerParams1 := dnacentersdkgo.CountTheNumberOfEventsHeaderParams{}
		queryParams1 := dnacentersdkgo.CountTheNumberOfEventsQueryParams{}

		queryParams1.DeviceFamily = vDeviceFamily.(string)

		if okStartTime {
			queryParams1.StartTime = vStartTime.(string)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(string)
		}
		if okMessageType {
			queryParams1.MessageType = vMessageType.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okNetworkDeviceName {
			queryParams1.NetworkDeviceName = vNetworkDeviceName.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okApMac {
			queryParams1.ApMac = vApMac.(string)
		}
		if okClientMac {
			queryParams1.ClientMac = vClientMac.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.CountTheNumberOfEvents(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountTheNumberOfEvents", err,
				"Failure at CountTheNumberOfEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesCountTheNumberOfEventsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountTheNumberOfEvents response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesCountTheNumberOfEventsItem(item *dnacentersdkgo.ResponseDevicesCountTheNumberOfEventsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
