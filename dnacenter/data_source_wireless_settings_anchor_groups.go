package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsAnchorGroups() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get AnchorGroups that captured in wireless settings design.
`,

		ReadContext: dataSourceWirelessSettingsAnchorGroupsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"anchor_group_name": &schema.Schema{
							Description: `Anchor Group Name. Max length is 32 characters
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Anchor Profile unique ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mobility_anchors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"anchor_priority": &schema.Schema{
										Description: `This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_name": &schema.Schema{
										Description: `Peer Host Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_address": &schema.Schema{
										Description: `This indicates Mobility public IP address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Description: `Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"managed_anchor_wlc": &schema.Schema{
										Description: `This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mobility_group_name": &schema.Schema{
										Description: `Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"peer_device_type": &schema.Schema{
										Description: `Indicates peer device mobility belongs to AireOS or IOS-XE family.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"private_ip": &schema.Schema{
										Description: `This indicates private management IP address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessSettingsAnchorGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnchorGroups")

		response1, restyResp1, err := client.Wireless.GetAnchorGroups()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAnchorGroups", err,
				"Failure at GetAnchorGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAnchorGroupsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnchorGroups response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAnchorGroupsItem(item *dnacentersdkgo.ResponseWirelessGetAnchorGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["anchor_group_name"] = item.AnchorGroupName
	respItem["mobility_anchors"] = flattenWirelessGetAnchorGroupsItemMobilityAnchors(item.MobilityAnchors)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetAnchorGroupsItemMobilityAnchors(items *[]dnacentersdkgo.ResponseWirelessGetAnchorGroupsMobilityAnchors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_name"] = item.DeviceName
		respItem["ip_address"] = item.IPAddress
		respItem["anchor_priority"] = item.AnchorPriority
		respItem["managed_anchor_wlc"] = boolPtrToString(item.ManagedAnchorWlc)
		respItem["peer_device_type"] = item.PeerDeviceType
		respItem["mac_address"] = item.MacAddress
		respItem["mobility_group_name"] = item.MobilityGroupName
		respItem["private_ip"] = item.PrivateIP
		respItems = append(respItems, respItem)
	}
	return respItems
}
