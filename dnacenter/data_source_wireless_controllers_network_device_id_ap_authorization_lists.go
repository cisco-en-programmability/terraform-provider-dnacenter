package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersNetworkDeviceIDApAuthorizationLists() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get an AP Authorization List details configured for the given provisioned network
device Id. Obtain the network device ID value by using the API GET call '/dna/intent/api/v1/network-device/ip-
address/${ipAddress}'.
`,

		ReadContext: dataSourceWirelessControllersNetworkDeviceIDApAuthorizationListsRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Network Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_authorization_list_name": &schema.Schema{
							Description: `Ap Authorization List Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"local_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac_entries": &schema.Schema{
										Description: `Ap Mac Entries`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ap_serial_number_entries": &schema.Schema{
										Description: `Ap Serial Number Entries`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"remote_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_servers": &schema.Schema{
										Description: `Aaa Servers`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"authorize_ap_with_mac": &schema.Schema{
										Description: `Authorize Ap With Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"authorize_ap_with_serial_number": &schema.Schema{
										Description: `Authorize Ap With Serial Number`,
										// Type:        schema.TypeBool,
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

func dataSourceWirelessControllersNetworkDeviceIDApAuthorizationListsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApAuthorizationListByNetworkDeviceID")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Wireless.GetApAuthorizationListByNetworkDeviceID(vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApAuthorizationListByNetworkDeviceID", err,
				"Failure at GetApAuthorizationListByNetworkDeviceID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetApAuthorizationListByNetworkDeviceIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApAuthorizationListByNetworkDeviceID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApAuthorizationListByNetworkDeviceIDItem(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["ap_authorization_list_name"] = item.ApAuthorizationListName
	respItem["local_authorization"] = flattenWirelessGetApAuthorizationListByNetworkDeviceIDItemLocalAuthorization(item.LocalAuthorization)
	respItem["remote_authorization"] = flattenWirelessGetApAuthorizationListByNetworkDeviceIDItemRemoteAuthorization(item.RemoteAuthorization)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApAuthorizationListByNetworkDeviceIDItemLocalAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseLocalAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ap_mac_entries"] = item.ApMacEntries
	respItem["ap_serial_number_entries"] = item.ApSerialNumberEntries

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApAuthorizationListByNetworkDeviceIDItemRemoteAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseRemoteAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_servers"] = item.AAAServers
	respItem["authorize_ap_with_mac"] = boolPtrToString(item.AuthorizeApWithMac)
	respItem["authorize_ap_with_serial_number"] = boolPtrToString(item.AuthorizeApWithSerialNumber)

	return []map[string]interface{}{
		respItem,
	}

}
