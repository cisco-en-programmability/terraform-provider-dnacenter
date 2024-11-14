package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersWirelessMobilityGroups() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieve all configured mobility groups if no Network Device Id is provided as a query parameter. If a Network Device
Id is given and a mobility group is configured for it, return the configured details; otherwise, return the default
values from the device.
`,

		ReadContext: dataSourceWirelessControllersWirelessMobilityGroupsRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Employ this query parameter to obtain the details of the Mobility Group corresponding to the provided networkDeviceId. Obtain the network device ID value by using the API GET call /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data_link_encryption": &schema.Schema{
							Description: `A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dtls_high_cipher": &schema.Schema{
							Description: `DTLS High Cipher.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `Device mobility MAC Address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip": &schema.Schema{
							Description: `Self device wireless Management IP.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mobility_group_name": &schema.Schema{
							Description: `Self device Group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mobility_peers": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_link_encryption": &schema.Schema{
										Description: `A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. 
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_series": &schema.Schema{
										Description: `Peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"hash_key": &schema.Schema{
										Description: `SSC hash string must be 40 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"member_mac_address": &schema.Schema{
										Description: `Peer device mobility MAC Address.  Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mobility_group_name": &schema.Schema{
										Description: `Peer device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"peer_ip": &schema.Schema{
										Description: `This indicates public IP address.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"peer_network_device_id": &schema.Schema{
										Description: `Peer device Id. The possible values are UNKNOWN or valid UUID of Network device ID. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device ID.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"private_ip_address": &schema.Schema{
										Description: `This indicates private/management IP address.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Possible values are - Control and Data Path Down, Data Path Down, Control Path Down, UP.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"network_device_id": &schema.Schema{
							Description: `Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
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

func dataSourceWirelessControllersWirelessMobilityGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllMobilityGroups")
		queryParams1 := dnacentersdkgo.GetAllMobilityGroupsQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}

		response1, restyResp1, err := client.Wireless.GetAllMobilityGroups(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllMobilityGroups", err,
				"Failure at GetAllMobilityGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetAllMobilityGroupsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllMobilityGroups response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAllMobilityGroupsItems(items *[]dnacentersdkgo.ResponseWirelessGetAllMobilityGroupsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mobility_group_name"] = item.MobilityGroupName
		respItem["mac_address"] = item.MacAddress
		respItem["management_ip"] = item.ManagementIP
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["dtls_high_cipher"] = boolPtrToString(item.DtlsHighCipher)
		respItem["data_link_encryption"] = boolPtrToString(item.DataLinkEncryption)
		respItem["mobility_peers"] = flattenWirelessGetAllMobilityGroupsItemsMobilityPeers(item.MobilityPeers)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAllMobilityGroupsItemsMobilityPeers(items *[]dnacentersdkgo.ResponseWirelessGetAllMobilityGroupsResponseMobilityPeers) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mobility_group_name"] = item.MobilityGroupName
		respItem["peer_network_device_id"] = item.PeerNetworkDeviceID
		respItem["member_mac_address"] = item.MemberMacAddress
		respItem["device_series"] = item.DeviceSeries
		respItem["data_link_encryption"] = boolPtrToString(item.DataLinkEncryption)
		respItem["hash_key"] = item.HashKey
		respItem["status"] = item.Status
		respItem["peer_ip"] = item.PeerIP
		respItem["private_ip_address"] = item.PrivateIPAddress
		respItems = append(respItems, respItem)
	}
	return respItems
}
