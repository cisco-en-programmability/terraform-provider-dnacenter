package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceVLAN() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns Device Interface VLANs
`,

		ReadContext: dataSourceNetworkDeviceVLANRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"interface_type": &schema.Schema{
				Description: `interfaceType query parameter. Vlan assocaited with sub-interface
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mask": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"number_of_ips": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_number": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"vlan_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceVLANRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInterfaceType, okInterfaceType := d.GetOk("interface_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceInterfaceVLANs")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetDeviceInterfaceVLANsQueryParams{}

		if okInterfaceType {
			queryParams1.InterfaceType = vInterfaceType.(string)
		}

		response1, restyResp1, err := client.Devices.GetDeviceInterfaceVLANs(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceInterfaceVLANs", err,
				"Failure at GetDeviceInterfaceVLANs, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceInterfaceVLANsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceInterfaceVLANs response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceInterfaceVLANsItems(items *[]dnacentersdkgo.ResponseDevicesGetDeviceInterfaceVLANsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface_name"] = item.InterfaceName
		respItem["ip_address"] = item.IPAddress
		respItem["mask"] = item.Mask
		respItem["network_address"] = item.NetworkAddress
		respItem["number_of_ips"] = item.NumberOfIPs
		respItem["prefix"] = item.Prefix
		respItem["vlan_number"] = item.VLANNumber
		respItem["vlan_type"] = item.VLANType
		respItems = append(respItems, respItem)
	}
	return respItems
}
