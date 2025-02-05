package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesTelemetrySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves telemetry settings for the given site. *null* values indicate that the setting will be inherited from the
parent site.
`,

		ReadContext: dataSourceSitesTelemetrySettingsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id, retrievable from the *id* attribute in */dna/intent/api/v1/sites*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"inherited": &schema.Schema{
				Description: `_inherited query parameter. Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when *false*, *null* values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_visibility": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"collector": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"collector_type": &schema.Schema{
													Description: `Collector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"port": &schema.Schema{
													Description: `Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"enable_on_wired_access_devices": &schema.Schema{
										Description: `Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmp_traps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_trap_servers": &schema.Schema{
										Description: `External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"use_builtin_trap_server": &schema.Schema{
										Description: `Enable this server as a destination server for SNMP traps and messages from your network
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"syslogs": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_syslog_servers": &schema.Schema{
										Description: `External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"use_builtin_syslog_server": &schema.Schema{
										Description: `Enable this server as a destination server for syslog messages.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wired_data_collection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wired_data_collectio": &schema.Schema{
										Description: `Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"wireless_telemetry": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wireless_telemetry": &schema.Schema{
										Description: `Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceSitesTelemetrySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTelemetrySettingsForASite")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.RetrieveTelemetrySettingsForASiteQueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrieveTelemetrySettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTelemetrySettingsForASite", err,
				"Failure at RetrieveTelemetrySettingsForASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTelemetrySettingsForASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItem(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["wired_data_collection"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemWiredDataCollection(item.WiredDataCollection)
	respItem["wireless_telemetry"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemWirelessTelemetry(item.WirelessTelemetry)
	respItem["snmp_traps"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemSNMPTraps(item.SNMPTraps)
	respItem["syslogs"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemSyslogs(item.Syslogs)
	respItem["application_visibility"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemApplicationVisibility(item.ApplicationVisibility)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemWiredDataCollection(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWiredDataCollection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_wired_data_collectio"] = boolPtrToString(item.EnableWiredDataCollectio)
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemWirelessTelemetry(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWirelessTelemetry) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_wireless_telemetry"] = boolPtrToString(item.EnableWirelessTelemetry)
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemSNMPTraps(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSNMPTraps) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["use_builtin_trap_server"] = boolPtrToString(item.UseBuiltinTrapServer)
	respItem["external_trap_servers"] = item.ExternalTrapServers
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemSyslogs(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSyslogs) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["use_builtin_syslog_server"] = boolPtrToString(item.UseBuiltinSyslogServer)
	respItem["external_syslog_servers"] = item.ExternalSyslogServers
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemApplicationVisibility(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibility) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["collector"] = flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemApplicationVisibilityCollector(item.Collector)
	respItem["enable_on_wired_access_devices"] = boolPtrToString(item.EnableOnWiredAccessDevices)
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveTelemetrySettingsForASiteItemApplicationVisibilityCollector(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibilityCollector) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["collector_type"] = item.CollectorType
	respItem["address"] = item.Address
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}
