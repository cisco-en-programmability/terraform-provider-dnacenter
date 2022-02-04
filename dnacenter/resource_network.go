package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Network Settings.

- API to create a network for DHCP and DNS center server settings.

- API to update network for DHCP and DNS center server settings.
`,

		CreateContext: resourceNetworkCreate,
		ReadContext:   resourceNetworkRead,
		UpdateContext: resourceNetworkUpdate,
		DeleteContext: resourceNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_and_endpoint_aaa": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id to update the network settings which is associated with the site
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkCreateNetwork(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteID, okSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	resp1, restyResp1, err := client.NetworkSettings.CreateNetwork(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetwork", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetwork", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkRead(ctx, d, m)
}

func resourceNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetwork")
		queryParams1 := dnacentersdkgo.GetNetworkQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID
		}

		response1, restyResp1, err := client.NetworkSettings.GetNetwork(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetwork", err,
				"Failure at GetNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenNetworkSettingsGetNetworkItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetwork search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]

	queryParams1 := dnacentersdkgo.GetNetworkQueryParams
	queryParams1.SiteID = vSiteID
	item, err := searchNetworkSettingsGetNetwork(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetNetwork", err,
			"Failure at GetNetwork, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestNetworkUpdateNetwork(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateNetwork(vvSiteID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetwork", err, restyResp1.String(),
					"Failure at UpdateNetwork, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetwork", err,
				"Failure at UpdateNetwork, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkRead(ctx, d, m)
}

func resourceNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete Network on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNetworkCreateNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetwork {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetwork{}
	request.Settings = expandRequestNetworkCreateNetworkSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateNetworkSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server")))) {
		request.DhcpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server")))) {
		request.DNSServer = expandRequestNetworkCreateNetworkSettingsDNSServer(ctx, key+".dns_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".syslog_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".syslog_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".syslog_server")))) {
		request.SyslogServer = expandRequestNetworkCreateNetworkSettingsSyslogServer(ctx, key+".syslog_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_server")))) {
		request.SNMPServer = expandRequestNetworkCreateNetworkSettingsSNMPServer(ctx, key+".snmp_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netflowcollector")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netflowcollector")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netflowcollector")))) {
		request.Netflowcollector = expandRequestNetworkCreateNetworkSettingsNetflowcollector(ctx, key+".netflowcollector.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ntp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ntp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ntp_server")))) {
		request.NtpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timezone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timezone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timezone")))) {
		request.Timezone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".message_of_theday")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".message_of_theday")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".message_of_theday")))) {
		request.MessageOfTheday = expandRequestNetworkCreateNetworkSettingsMessageOfTheday(ctx, key+".message_of_theday.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_aaa")))) {
		request.NetworkAAA = expandRequestNetworkCreateNetworkSettingsNetworkAAA(ctx, key+".network_aaa.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_and_endpoint_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) {
		request.ClientAndEndpointAAA = expandRequestNetworkCreateNetworkSettingsClientAndEndpointAAA(ctx, key+".client_and_endpoint_aaa.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkCreateNetworkSettingsDNSServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsDNSServer {
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

func expandRequestNetworkCreateNetworkSettingsSyslogServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSyslogServer {
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

func expandRequestNetworkCreateNetworkSettingsSNMPServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSNMPServer {
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

func expandRequestNetworkCreateNetworkSettingsNetflowcollector(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetflowcollector {
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

func expandRequestNetworkCreateNetworkSettingsMessageOfTheday(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday {
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

func expandRequestNetworkCreateNetworkSettingsNetworkAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetworkAAA {
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

func expandRequestNetworkCreateNetworkSettingsClientAndEndpointAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA {
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

func expandRequestNetworkUpdateNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetwork {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetwork{}
	request.Settings = expandRequestNetworkUpdateNetworkSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkUpdateNetworkSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server")))) {
		request.DhcpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server")))) {
		request.DNSServer = expandRequestNetworkUpdateNetworkSettingsDNSServer(ctx, key+".dns_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".syslog_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".syslog_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".syslog_server")))) {
		request.SyslogServer = expandRequestNetworkUpdateNetworkSettingsSyslogServer(ctx, key+".syslog_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_server")))) {
		request.SNMPServer = expandRequestNetworkUpdateNetworkSettingsSNMPServer(ctx, key+".snmp_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netflowcollector")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netflowcollector")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netflowcollector")))) {
		request.Netflowcollector = expandRequestNetworkUpdateNetworkSettingsNetflowcollector(ctx, key+".netflowcollector.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ntp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ntp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ntp_server")))) {
		request.NtpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timezone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timezone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timezone")))) {
		request.Timezone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".message_of_theday")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".message_of_theday")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".message_of_theday")))) {
		request.MessageOfTheday = expandRequestNetworkUpdateNetworkSettingsMessageOfTheday(ctx, key+".message_of_theday.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_aaa")))) {
		request.NetworkAAA = expandRequestNetworkUpdateNetworkSettingsNetworkAAA(ctx, key+".network_aaa.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_and_endpoint_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) {
		request.ClientAndEndpointAAA = expandRequestNetworkUpdateNetworkSettingsClientAndEndpointAAA(ctx, key+".client_and_endpoint_aaa.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkUpdateNetworkSettingsDNSServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsDNSServer {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsDNSServer{}
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

func expandRequestNetworkUpdateNetworkSettingsSyslogServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsSyslogServer {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsSyslogServer{}
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

func expandRequestNetworkUpdateNetworkSettingsSNMPServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsSNMPServer {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsSNMPServer{}
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

func expandRequestNetworkUpdateNetworkSettingsNetflowcollector(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsNetflowcollector {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsNetflowcollector{}
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

func expandRequestNetworkUpdateNetworkSettingsMessageOfTheday(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsMessageOfTheday {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsMessageOfTheday{}
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

func expandRequestNetworkUpdateNetworkSettingsNetworkAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsNetworkAAA {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsNetworkAAA{}
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

func expandRequestNetworkUpdateNetworkSettingsClientAndEndpointAAA(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsClientAndEndpointAAA {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateNetworkSettingsClientAndEndpointAAA{}
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

func searchNetworkSettingsGetNetwork(m interface{}, queryParams dnacentersdkgo.GetNetworkQueryParams) (*dnacentersdkgo.ResponseItemNetworkSettingsGetNetwork, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemNetworkSettingsGetNetwork
	var ite *dnacentersdkgo.ResponseNetworkSettingsGetNetwork
	ite, _, err = client.NetworkSettings.GetNetwork(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemNetworkSettingsGetNetwork
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
