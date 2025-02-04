package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityRogueWirelessContainmentStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Intent API to check the wireless rogue access point containment status. The response includes all the details like
containment status, contained by WLC, containment status of each BSSID etc. This data source also includes the
information of strongest detecting WLC for this rogue access point.
`,

		ReadContext: dataSourceSecurityRogueWirelessContainmentStatusRead,
		Schema: map[string]*schema.Schema{
			"mac_address": &schema.Schema{
				Description: `macAddress path parameter. MAC Address of the Wireless Rogue AP
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bssid_containment_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"contained_by_wlc_ip": &schema.Schema{
										Description: `Contained By Wlc Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"containment_status": &schema.Schema{
										Description: `Containment Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"is_adhoc": &schema.Schema{
										Description: `Is Adhoc`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_type": &schema.Schema{
										Description: `Radio Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"classification": &schema.Schema{
							Description: `Classification`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"contained_by_wlc_ip": &schema.Schema{
							Description: `Contained By Wlc Ip`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"containment_status": &schema.Schema{
							Description: `Containment Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_seen": &schema.Schema{
							Description: `Last Seen`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_task_detail": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"initiated_on_bssid": &schema.Schema{
										Description: `Initiated On Bssid`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"initiated_on_wlc_ip": &schema.Schema{
										Description: `Initiated On Wlc Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"task_id": &schema.Schema{
										Description: `Task Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"task_start_time": &schema.Schema{
										Description: `Task Start Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"task_state": &schema.Schema{
										Description: `Task State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"task_type": &schema.Schema{
										Description: `Task Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"strongest_detecting_wlc_ip": &schema.Schema{
							Description: `Strongest Detecting Wlc Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityRogueWirelessContainmentStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vMacAddress := d.Get("mac_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: WirelessRogueApContainmentStatus")
		vvMacAddress := vMacAddress.(string)

		response1, restyResp1, err := client.Devices.WirelessRogueApContainmentStatus(vvMacAddress)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 WirelessRogueApContainmentStatus", err,
				"Failure at WirelessRogueApContainmentStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesWirelessRogueApContainmentStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting WirelessRogueApContainmentStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesWirelessRogueApContainmentStatusItems(items *[]dnacentersdkgo.ResponseDevicesWirelessRogueApContainmentStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mac_address"] = item.MacAddress
		respItem["type"] = item.Type
		respItem["classification"] = item.Classification
		respItem["containment_status"] = item.ContainmentStatus
		respItem["contained_by_wlc_ip"] = item.ContainedByWlcIP
		respItem["last_seen"] = item.LastSeen
		respItem["strongest_detecting_wlc_ip"] = item.StrongestDetectingWlcIP
		respItem["last_task_detail"] = flattenDevicesWirelessRogueApContainmentStatusItemsLastTaskDetail(item.LastTaskDetail)
		respItem["bssid_containment_status"] = flattenDevicesWirelessRogueApContainmentStatusItemsBssidContainmentStatus(item.BssidContainmentStatus)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesWirelessRogueApContainmentStatusItemsLastTaskDetail(item *dnacentersdkgo.ResponseDevicesWirelessRogueApContainmentStatusResponseLastTaskDetail) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["task_type"] = item.TaskType
	respItem["task_state"] = item.TaskState
	respItem["task_start_time"] = item.TaskStartTime
	respItem["initiated_on_wlc_ip"] = item.InitiatedOnWlcIP
	respItem["initiated_on_bssid"] = item.InitiatedOnBssid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesWirelessRogueApContainmentStatusItemsBssidContainmentStatus(items *[]dnacentersdkgo.ResponseDevicesWirelessRogueApContainmentStatusResponseBssidContainmentStatus) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bssid"] = item.Bssid
		respItem["ssid"] = item.SSID
		respItem["radio_type"] = item.RadioType
		respItem["containment_status"] = item.ContainmentStatus
		respItem["contained_by_wlc_ip"] = item.ContainedByWlcIP
		respItem["is_adhoc"] = boolPtrToString(item.IsAdhoc)
		respItems = append(respItems, respItem)
	}
	return respItems
}
