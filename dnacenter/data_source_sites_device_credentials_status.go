package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesDeviceCredentialsStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Get network devices credentials sync status at a given site.
`,

		ReadContext: dataSourceSitesDeviceCredentialsStatusRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_count": &schema.Schema{
										Description: `Device count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Sync status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v2_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_count": &schema.Schema{
										Description: `Device count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Sync status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v2_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_count": &schema.Schema{
										Description: `Device count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Sync status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_count": &schema.Schema{
										Description: `Device count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Sync status
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

func dataSourceSitesDeviceCredentialsStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDevicesCredentialsSyncStatus")
		vvID := vID.(string)

		response1, restyResp1, err := client.NetworkSettings.GetNetworkDevicesCredentialsSyncStatus(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDevicesCredentialsSyncStatus", err,
				"Failure at GetNetworkDevicesCredentialsSyncStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDevicesCredentialsSyncStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItem(item *dnacentersdkgo.ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli"] = flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemCli(item.Cli)
	respItem["snmp_v2_read"] = flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV2Read(item.SNMPV2Read)
	respItem["snmp_v2_write"] = flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV2Write(item.SNMPV2Write)
	respItem["snmp_v3"] = flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV3(item.SNMPV3)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemCli(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseCli) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_count"] = item.DeviceCount
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV2Read(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Read) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_count"] = item.DeviceCount
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV2Write(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Write) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_count"] = item.DeviceCount
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetNetworkDevicesCredentialsSyncStatusItemSNMPV3(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV3) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_count"] = item.DeviceCount
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
