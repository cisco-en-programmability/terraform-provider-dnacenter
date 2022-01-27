package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceInventoryInsightLinkMismatch() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Find all devices with link mismatch (speed /  vlan)
`,

		ReadContext: dataSourceNetworkDeviceInventoryInsightLinkMismatchRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Description: `category query parameter. Links mismatch category.  Value can be speed-duplex or vlan.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Default value is 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Row Number.  Default value is 1
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order.  Value can be asc or desc.  Default value is asc
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort By
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"avg_update_frequency": &schema.Schema{
							Description: `Average update frequency
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"end_device_host_name": &schema.Schema{
							Description: `End device hostname
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_device_id": &schema.Schema{
							Description: `End device id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_device_ip_address": &schema.Schema{
							Description: `End device ip address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_address": &schema.Schema{
							Description: `End port address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_allowed_vlan_ids": &schema.Schema{
							Description: `End port allowed vlan ids
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_duplex": &schema.Schema{
							Description: `End port duplex
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_id": &schema.Schema{
							Description: `End port id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_mask": &schema.Schema{
							Description: `End port mask
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_name": &schema.Schema{
							Description: `End port name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_native_vlan_id": &schema.Schema{
							Description: `End port native vlan id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_pep_id": &schema.Schema{
							Description: `End port pep id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_port_speed": &schema.Schema{
							Description: `End port speed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance tenant id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Unique instance id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Description: `Last updated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"link_status": &schema.Schema{
							Description: `Link status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"num_updates": &schema.Schema{
							Description: `Number updates
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"start_device_host_name": &schema.Schema{
							Description: `Start device hostname
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_device_id": &schema.Schema{
							Description: `Start device id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_device_ip_address": &schema.Schema{
							Description: `Start device ip address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_address": &schema.Schema{
							Description: `Start port address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_allowed_vlan_ids": &schema.Schema{
							Description: `Start port allowed vlan ids
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_duplex": &schema.Schema{
							Description: `Start port duplex
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_id": &schema.Schema{
							Description: `Start port id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_mask": &schema.Schema{
							Description: `Start port mask
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_name": &schema.Schema{
							Description: `Start port name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_native_vlan_id": &schema.Schema{
							Description: `Start port native vlan id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_pep_id": &schema.Schema{
							Description: `Start port pep id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_port_speed": &schema.Schema{
							Description: `Start port speed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceInventoryInsightLinkMismatchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vCategory := d.Get("category")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: InventoryInsightDeviceLinkMismatchApI")
		vvSiteID := vSiteID.(string)
		queryParams1 := dnacentersdkgo.InventoryInsightDeviceLinkMismatchApIQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		queryParams1.Category = vCategory.(string)

		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Devices.InventoryInsightDeviceLinkMismatchApI(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing InventoryInsightDeviceLinkMismatchApI", err,
				"Failure at InventoryInsightDeviceLinkMismatchApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesInventoryInsightDeviceLinkMismatchApIItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting InventoryInsightDeviceLinkMismatchApI response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesInventoryInsightDeviceLinkMismatchApIItems(items *[]dnacentersdkgo.ResponseDevicesInventoryInsightDeviceLinkMismatchAPIResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["end_port_allowed_vlan_ids"] = item.EndPortAllowedVLANIDs
		respItem["end_port_native_vlan_id"] = item.EndPortNativeVLANID
		respItem["start_port_allowed_vlan_ids"] = item.StartPortAllowedVLANIDs
		respItem["start_port_native_vlan_id"] = item.StartPortNativeVLANID
		respItem["link_status"] = item.LinkStatus
		respItem["end_device_host_name"] = item.EndDeviceHostName
		respItem["end_device_id"] = item.EndDeviceID
		respItem["end_device_ip_address"] = item.EndDeviceIPAddress
		respItem["end_port_address"] = item.EndPortAddress
		respItem["end_port_duplex"] = item.EndPortDuplex
		respItem["end_port_id"] = item.EndPortID
		respItem["end_port_mask"] = item.EndPortMask
		respItem["end_port_name"] = item.EndPortName
		respItem["end_port_pep_id"] = item.EndPortPepID
		respItem["end_port_speed"] = item.EndPortSpeed
		respItem["start_device_host_name"] = item.StartDeviceHostName
		respItem["start_device_id"] = item.StartDeviceID
		respItem["start_device_ip_address"] = item.StartDeviceIPAddress
		respItem["start_port_address"] = item.StartPortAddress
		respItem["start_port_duplex"] = item.StartPortDuplex
		respItem["start_port_id"] = item.StartPortID
		respItem["start_port_mask"] = item.StartPortMask
		respItem["start_port_name"] = item.StartPortName
		respItem["start_port_pep_id"] = item.StartPortPepID
		respItem["start_port_speed"] = item.StartPortSpeed
		respItem["last_updated"] = item.LastUpdated
		respItem["num_updates"] = item.NumUpdates
		respItem["avg_update_frequency"] = item.AvgUpdateFrequency
		respItem["type"] = item.Type
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}
