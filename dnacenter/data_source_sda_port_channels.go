package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPortChannels() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of port channels that match the provided query parameters.
`,

		ReadContext: dataSourceSdaPortChannelsRead,
		Schema: map[string]*schema.Schema{
			"connected_device_type": &schema.Schema{
				Description: `connectedDeviceType query parameter. Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the device is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. ID of the network device.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"port_channel_name": &schema.Schema{
				Description: `portChannelName query parameter. Name of the port channel.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_device_type": &schema.Schema{
							Description: `Connected device type of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric the device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_names": &schema.Schema{
							Description: `Interface names of this port channel.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"network_device_id": &schema.Schema{
							Description: `ID of the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_channel_name": &schema.Schema{
							Description: `Name of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol": &schema.Schema{
							Description: `Protocol of the port channel.
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

func dataSourceSdaPortChannelsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vPortChannelName, okPortChannelName := d.GetOk("port_channel_name")
	vConnectedDeviceType, okConnectedDeviceType := d.GetOk("connected_device_type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortChannels")
		queryParams1 := dnacentersdkgo.GetPortChannelsQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okPortChannelName {
			queryParams1.PortChannelName = vPortChannelName.(string)
		}
		if okConnectedDeviceType {
			queryParams1.ConnectedDeviceType = vConnectedDeviceType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetPortChannels(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPortChannels", err,
				"Failure at GetPortChannels, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetPortChannelsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortChannels response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPortChannelsItems(items *[]dnacentersdkgo.ResponseSdaGetPortChannelsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["port_channel_name"] = item.PortChannelName
		respItem["interface_names"] = item.InterfaceNames
		respItem["connected_device_type"] = item.ConnectedDeviceType
		respItem["protocol"] = item.Protocol
		respItem["description"] = item.Description
		respItems = append(respItems, respItem)
	}
	return respItems
}
