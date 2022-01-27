package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNetworkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to create a network for DHCP and DNS center server settings.
`,

		ReadContext: dataSourceNetworkCreateRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site id to which site details to associate with the network settings.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_and_endpoint_aaa": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address": &schema.Schema{
										Description: `IP address for ISE serve (eg: 1.1.1.4). Mandatory for ISE servers.
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"network": &schema.Schema{
										Description: `IP address for AAA or ISE server (eg: 2.2.2.1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol for AAA or ISE serve (eg: RADIUS)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"servers": &schema.Schema{
										Description: `Server type AAA or ISE server (eg: AAA)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared secret for ISE server. Supported only by ISE servers
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"dhcp_server": &schema.Schema{
							Description: `Dhcp serve Ip (eg: 1.1.1.1)
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dns_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"domain_name": &schema.Schema{
										Description: `Domain name of DHCP (eg; cisco). It can only contain alphanumeric characters or hyphen.
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"primary_ip_address": &schema.Schema{
										Description: `Primary ip address for DHCP (eg: 2.2.2.2). valid range : 1.0.0.0 - 223.255.255.255 
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"secondary_ip_address": &schema.Schema{
										Description: `Secondary ip address for DHCP (eg: 3.3.3.3). valid range : 1.0.0.0 - 223.255.255.255
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"message_of_theday": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"banner_message": &schema.Schema{
										Description: `Massage for banner message (eg; Good day)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"retain_existing_banner": &schema.Schema{
										Description: `Retain existing banner message (eg: "true" or "false")
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"netflowcollector": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address": &schema.Schema{
										Description: `IP address for netflow collector (eg: 3.3.3.1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Description: `Port for netflow collector (eg; 443)
`,
										Type:     schema.TypeFloat,
										Optional: true,
									},
								},
							},
						},
						"network_aaa": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address": &schema.Schema{
										Description: `IP address for AAA and ISE server (eg: 1.1.1.1). Mandatory for ISE servers and for AAA consider this as additional Ip.
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"network": &schema.Schema{
										Description: `IP address for AAA or ISE server (eg: 2.2.2.2). For AAA server consider it as primary IP and For ISE consider as Network
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol for AAA or ISE serve (eg: RADIUS)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"servers": &schema.Schema{
										Description: `Server type for AAA network (eg: AAA)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared secret for ISE server. Supported only by ISE servers
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"ntp_server": &schema.Schema{
							Description: `IP address for NTP server (eg: 1.1.1.2)
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"snmp_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_dnac_ip": &schema.Schema{
										Description: `Configuration dnac ip for snmp server (eg: true)
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"ip_addresses": &schema.Schema{
										Description: `IP address for snmp server (eg: 4.4.4.1)
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"syslog_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_dnac_ip": &schema.Schema{
										Description: `Configuration dnac ip for syslog server (eg: true)
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"ip_addresses": &schema.Schema{
										Description: `IP address for syslog server (eg: 4.4.4.4)
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"timezone": &schema.Schema{
							Description: `Input for time zone (eg: Africa/Abidjan)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateNetwork")
		vvSiteID := vSiteID.(string)
		request1 := expandRequestNetworkCreateCreateNetwork(ctx, "", d)

		headerParams1 := dnacentersdkgo.CreateNetworkHeaderParams{}
		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.CreateNetwork(vvSiteID, request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateNetwork", err,
				"Failure at CreateNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsCreateNetworkItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateNetwork response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNetworkCreateCreateNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetwork {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetwork{}
	request.Settings = expandRequestNetworkCreateCreateNetworkSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server")))) {
		request.DhcpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server")))) {
		request.DNSServer = expandRequestNetworkCreateCreateNetworkSettingsDNSServer(ctx, key+".dns_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".syslog_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".syslog_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".syslog_server")))) {
		request.SyslogServer = expandRequestNetworkCreateCreateNetworkSettingsSyslogServer(ctx, key+".syslog_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_server")))) {
		request.SNMPServer = expandRequestNetworkCreateCreateNetworkSettingsSNMPServer(ctx, key+".snmp_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netflowcollector")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netflowcollector")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netflowcollector")))) {
		request.Netflowcollector = expandRequestNetworkCreateCreateNetworkSettingsNetflowcollector(ctx, key+".netflowcollector.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ntp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ntp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ntp_server")))) {
		request.NtpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timezone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timezone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timezone")))) {
		request.Timezone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".message_of_theday")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".message_of_theday")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".message_of_theday")))) {
		request.MessageOfTheday = expandRequestNetworkCreateCreateNetworkSettingsMessageOfTheday(ctx, key+".message_of_theday.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_aaa")))) {
		request.NetworkAAA = expandRequestNetworkCreateCreateNetworkSettingsNetworkAAA(ctx, key+".network_aaa.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_and_endpoint_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) {
		request.ClientAndEndpointAAA = expandRequestNetworkCreateCreateNetworkSettingsClientAndEndpointAAA(ctx, key+".client_and_endpoint_aaa.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsDNSServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsDNSServer {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsDNSServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain_name")))) {
		request.DomainName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_ip_address")))) {
		request.PrimaryIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_ip_address")))) {
		request.SecondaryIPAddress = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsSyslogServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSyslogServer {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSyslogServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addresses")))) {
		request.IPAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_dnac_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_dnac_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_dnac_ip")))) {
		request.ConfigureDnacIP = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsSNMPServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSNMPServer {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSNMPServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addresses")))) {
		request.IPAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_dnac_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_dnac_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_dnac_ip")))) {
		request.ConfigureDnacIP = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsNetflowcollector(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetflowcollector {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetflowcollector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsMessageOfTheday(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_message")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_message")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_message")))) {
		request.BannerMessage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retain_existing_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retain_existing_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retain_existing_banner")))) {
		request.RetainExistingBanner = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsNetworkAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetworkAAA {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetworkAAA{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network")))) {
		request.Network = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsClientAndEndpointAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network")))) {
		request.Network = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenNetworkSettingsCreateNetworkItem(item *dnacentersdkgo.ResponseNetworkSettingsCreateNetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
