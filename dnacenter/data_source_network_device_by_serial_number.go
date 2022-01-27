package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceBySerialNumber() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the network device with given serial number
`,

		ReadContext: dataSourceNetworkDeviceBySerialNumberRead,
		Schema: map[string]*schema.Schema{
			"serial_number": &schema.Schema{
				Description: `serialNumber path parameter. Device serial number
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_manager_interface_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"associated_wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_interval": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_status": &schema.Schema{
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

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_status_detail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
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

						"platform_id": &schema.Schema{
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

						"series": &schema.Schema{
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

						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tunnel_udp_port": &schema.Schema{
							Type:     schema.TypeString,
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

						"waas_device_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceBySerialNumberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSerialNumber := d.Get("serial_number")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceBySerialNumber")
		vvSerialNumber := vSerialNumber.(string)

		response1, restyResp1, err := client.Devices.GetDeviceBySerialNumber(vvSerialNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceBySerialNumber", err,
				"Failure at GetDeviceBySerialNumber, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceBySerialNumberItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceBySerialNumber response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceBySerialNumberItem(item *dnacentersdkgo.ResponseDevicesGetDeviceBySerialNumberResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = item.BootDateTime
	respItem["collection_interval"] = item.CollectionInterval
	respItem["collection_status"] = item.CollectionStatus
	respItem["error_code"] = item.ErrorCode
	respItem["error_description"] = item.ErrorDescription
	respItem["family"] = item.Family
	respItem["hostname"] = item.Hostname
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["interface_count"] = item.InterfaceCount
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["last_updated"] = item.LastUpdated
	respItem["line_card_count"] = item.LineCardCount
	respItem["line_card_id"] = item.LineCardID
	respItem["location"] = item.Location
	respItem["location_name"] = item.LocationName
	respItem["mac_address"] = item.MacAddress
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["role"] = item.Role
	respItem["role_source"] = item.RoleSource
	respItem["serial_number"] = item.SerialNumber
	respItem["series"] = item.Series
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["software_type"] = item.SoftwareType
	respItem["software_version"] = item.SoftwareVersion
	respItem["tag_count"] = item.TagCount
	respItem["tunnel_udp_port"] = item.TunnelUDPPort
	respItem["type"] = item.Type
	respItem["up_time"] = item.UpTime
	respItem["waas_device_mode"] = item.WaasDeviceMode
	return []map[string]interface{}{
		respItem,
	}
}
