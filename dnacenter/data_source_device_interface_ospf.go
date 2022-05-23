package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceInterfaceOspf() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the interfaces that has OSPF enabled
`,

		ReadContext: dataSourceDeviceInterfaceOspfRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"class_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"duplex": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"if_index": &schema.Schema{
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

						"interface_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_mask": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"isis_support": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"media_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"native_vlan_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ospf_support": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"pid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_no": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"speed": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"voice_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceInterfaceOspfRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetOspfInterfaces")

		response1, restyResp1, err := client.Devices.GetOspfInterfaces()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetOspfInterfaces", err,
				"Failure at GetOspfInterfaces, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetOspfInterfacesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetOspfInterfaces response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetOspfInterfacesItems(items *[]dnacentersdkgo.ResponseDevicesGetOspfInterfacesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["admin_status"] = item.AdminStatus
		respItem["class_name"] = item.ClassName
		respItem["description"] = item.Description
		respItem["device_id"] = item.DeviceID
		respItem["duplex"] = item.Duplex
		respItem["id"] = item.ID
		respItem["if_index"] = item.IfIndex
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["interface_type"] = item.InterfaceType
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv4_mask"] = item.IPv4Mask
		respItem["isis_support"] = item.IsisSupport
		respItem["last_updated"] = item.LastUpdated
		respItem["mac_address"] = item.MacAddress
		respItem["mapped_physical_interface_id"] = item.MappedPhysicalInterfaceID
		respItem["mapped_physical_interface_name"] = item.MappedPhysicalInterfaceName
		respItem["media_type"] = item.MediaType
		respItem["native_vlan_id"] = item.NativeVLANID
		respItem["ospf_support"] = item.OspfSupport
		respItem["pid"] = item.Pid
		respItem["port_mode"] = item.PortMode
		respItem["port_name"] = item.PortName
		respItem["port_type"] = item.PortType
		respItem["serial_no"] = item.SerialNo
		respItem["series"] = item.Series
		respItem["speed"] = item.Speed
		respItem["status"] = item.Status
		respItem["vlan_id"] = item.VLANID
		respItem["voice_vlan"] = item.VoiceVLAN
		respItems = append(respItems, respItem)
	}
	return respItems
}
