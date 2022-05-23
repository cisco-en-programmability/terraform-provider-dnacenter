package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryDeviceRange() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the network devices discovered for the given discovery and for the given range. The maximum number of records
that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoveryDeviceRangeRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"records_to_return": &schema.Schema{
				Description: `recordsToReturn path parameter. Number of records to return
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_index": &schema.Schema{
				Description: `startIndex path parameter. Start index
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"task_id": &schema.Schema{
				Description: `taskId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"anchor_wlc_for_ap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_model_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"avg_update_frequency": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"cli_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"duplicate_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"http_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"image_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ingress_queue_config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"location_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"memory_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"num_updates": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ping_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_range": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"qos_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"role_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_contact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"up_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlc_ap_device_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryDeviceRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vStartIndex := d.Get("start_index")
	vRecordsToReturn := d.Get("records_to_return")
	vTaskID, okTaskID := d.GetOk("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDiscoveredDevicesByRange")
		vvID := vID.(string)
		vvStartIndex := vStartIndex.(int)
		vvRecordsToReturn := vRecordsToReturn.(int)
		queryParams1 := dnacentersdkgo.GetDiscoveredDevicesByRangeQueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}

		response1, restyResp1, err := client.Discovery.GetDiscoveredDevicesByRange(vvID, vvStartIndex, vvRecordsToReturn, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveredDevicesByRange", err,
				"Failure at GetDiscoveredDevicesByRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetDiscoveredDevicesByRangeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDiscoveredDevicesByRange response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDiscoveredDevicesByRangeItems(items *[]dnacentersdkgo.ResponseDiscoveryGetDiscoveredDevicesByRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["anchor_wlc_for_ap"] = item.AnchorWlcForAp
		respItem["auth_model_id"] = item.AuthModelID
		respItem["avg_update_frequency"] = item.AvgUpdateFrequency
		respItem["boot_date_time"] = item.BootDateTime
		respItem["cli_status"] = item.CliStatus
		respItem["duplicate_device_id"] = item.DuplicateDeviceID
		respItem["error_code"] = item.ErrorCode
		respItem["error_description"] = item.ErrorDescription
		respItem["family"] = item.Family
		respItem["hostname"] = item.Hostname
		respItem["http_status"] = item.HTTPStatus
		respItem["id"] = item.ID
		respItem["image_name"] = item.ImageName
		respItem["ingress_queue_config"] = item.IngressQueueConfig
		respItem["interface_count"] = item.InterfaceCount
		respItem["inventory_collection_status"] = item.InventoryCollectionStatus
		respItem["inventory_reachability_status"] = item.InventoryReachabilityStatus
		respItem["last_updated"] = item.LastUpdated
		respItem["line_card_count"] = item.LineCardCount
		respItem["line_card_id"] = item.LineCardID
		respItem["location"] = item.Location
		respItem["location_name"] = item.LocationName
		respItem["mac_address"] = item.MacAddress
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["memory_size"] = item.MemorySize
		respItem["netconf_status"] = item.NetconfStatus
		respItem["num_updates"] = item.NumUpdates
		respItem["ping_status"] = item.PingStatus
		respItem["platform_id"] = item.PlatformID
		respItem["port_range"] = item.PortRange
		respItem["qos_status"] = item.QosStatus
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["serial_number"] = item.SerialNumber
		respItem["snmp_contact"] = item.SNMPContact
		respItem["snmp_location"] = item.SNMPLocation
		respItem["snmp_status"] = item.SNMPStatus
		respItem["software_version"] = item.SoftwareVersion
		respItem["tag"] = item.Tag
		respItem["tag_count"] = item.TagCount
		respItem["type"] = item.Type
		respItem["up_time"] = item.UpTime
		respItem["vendor"] = item.Vendor
		respItem["wlc_ap_device_status"] = item.WlcApDeviceStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}
