package dnacenter

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{

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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"client_and_endpoint_aaa": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"servers": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"dhcp_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dns_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"primary_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"secondary_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"message_of_theday": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"banner_message": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"retain_existing_banner": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"netflowcollector": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"network_aaa": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
									"network": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
									"servers": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
								},
							},
						},
						"ntp_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"snmp_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configure_dnac_ip": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
										Optional: true,
									},
									"ip_addresses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
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
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configure_dnac_ip": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
										Optional: true,
									},
									"ip_addresses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"timezone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func constructUpdateNetworkClientAndEndpointAAA(response []interface{}) *dnac.UpdateNetworkRequestSettingsClientAndEndpointAAA {
	result := dnac.UpdateNetworkRequestSettingsClientAndEndpointAAA{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["network"]; ok && v != nil {
			result.Network = v.(string)
		}
		if v, ok := ci["protocol"]; ok && v != nil {
			result.Protocol = v.(string)
		}
		if v, ok := ci["servers"]; ok && v != nil {
			result.Servers = v.(string)
		}
		if v, ok := ci["shared_secret"]; ok && v != nil {
			result.SharedSecret = v.(string)
		}
		return &result
	}
	return nil
}
func constructUpdateNetworkDNSServer(response []interface{}) *dnac.UpdateNetworkRequestSettingsDNSServer {
	result := dnac.UpdateNetworkRequestSettingsDNSServer{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["domain_name"]; ok && v != nil {
			result.DomainName = v.(string)
		}
		if v, ok := ci["primary_ip_address"]; ok && v != nil {
			result.PrimaryIPAddress = v.(string)
		}
		if v, ok := ci["secondary_ip_address"]; ok && v != nil {
			result.SecondaryIPAddress = v.(string)
		}
		return &result
	}
	return nil
}
func constructUpdateNetworkMessageOfTheday(response []interface{}) *dnac.UpdateNetworkRequestSettingsMessageOfTheday {
	var result dnac.UpdateNetworkRequestSettingsMessageOfTheday
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["banner_message"]; ok && v != nil {
			result.BannerMessage = v.(string)
		}
		if v, ok := ci["retain_existing_banner"]; ok && v != nil {
			result.RetainExistingBanner = v.(bool)
		}

		return &result
	}
	return nil
}
func constructUpdateNetworkNetflowcollector(response []interface{}) *dnac.UpdateNetworkRequestSettingsNetflowcollector {
	var result dnac.UpdateNetworkRequestSettingsNetflowcollector
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["port"]; ok && v != nil {
			result.Port = v.(float64)
		}

		return &result
	}
	return nil
}
func constructUpdateNetworkNetworkAAA(response []interface{}) *dnac.UpdateNetworkRequestSettingsNetworkAAA {
	var result dnac.UpdateNetworkRequestSettingsNetworkAAA
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["network"]; ok && v != nil {
			result.Network = v.(string)
		}
		if v, ok := ci["protocol"]; ok && v != nil {
			result.Protocol = v.(string)
		}
		if v, ok := ci["servers"]; ok && v != nil {
			result.Servers = v.(string)
		}
		if v, ok := ci["shared_secret"]; ok && v != nil {
			result.SharedSecret = v.(string)
		}

		return &result
	}
	return nil
}
func constructUpdateNetworkSNMPServer(response []interface{}) *dnac.UpdateNetworkRequestSettingsSNMPServer {
	var result dnac.UpdateNetworkRequestSettingsSNMPServer
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["configure_dnac_ip"]; ok && v != nil {
			result.ConfigureDnacIP = v.(bool)
		}
		if v, ok := ci["ip_addresses"]; ok && v != nil {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}

		return &result
	}
	return nil
}
func constructUpdateNetworkSyslogServer(response []interface{}) *dnac.UpdateNetworkRequestSettingsSyslogServer {
	var result dnac.UpdateNetworkRequestSettingsSyslogServer
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["configure_dnac_ip"]; ok && v != nil {
			result.ConfigureDnacIP = v.(bool)
		}
		if v, ok := ci["ip_addresses"]; ok && v != nil {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}

		return &result
	}
	return nil
}
func constructUpdateNetwork(response []interface{}) *dnac.UpdateNetworkRequest {
	result := dnac.UpdateNetworkRequest{}
	resultSettings := dnac.UpdateNetworkRequestSettings{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["client_and_endpoint_aaa"]; ok && v != nil {
			if w := constructUpdateNetworkClientAndEndpointAAA(v.([]interface{})); w != nil {
				resultSettings.ClientAndEndpointAAA = w
			}
		}
		if v, ok := ci["dhcp_server"]; ok && v != nil {
			resultSettings.DhcpServer = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := ci["dns_server"]; ok && v != nil {
			if w := constructUpdateNetworkDNSServer(v.([]interface{})); w != nil {
				resultSettings.DNSServer = w
			}
		}
		if v, ok := ci["message_of_theday"]; ok && v != nil {
			if w := constructUpdateNetworkMessageOfTheday(v.([]interface{})); w != nil {
				resultSettings.MessageOfTheday = w
			}
		}
		if v, ok := ci["netflowcollector"]; ok && v != nil {
			if w := constructUpdateNetworkNetflowcollector(v.([]interface{})); w != nil {
				resultSettings.Netflowcollector = w
			}
		}
		if v, ok := ci["network_aaa"]; ok && v != nil {
			if w := constructUpdateNetworkNetworkAAA(v.([]interface{})); w != nil {
				resultSettings.NetworkAAA = w
			}
		}
		if v, ok := ci["ntp_server"]; ok && v != nil {
			resultSettings.NtpServer = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := ci["snmp_server"]; ok && v != nil {
			if w := constructUpdateNetworkSNMPServer(v.([]interface{})); w != nil {
				resultSettings.SNMPServer = w
			}
		}
		if v, ok := ci["syslog_server"]; ok && v != nil {
			if w := constructUpdateNetworkSyslogServer(v.([]interface{})); w != nil {
				resultSettings.SyslogServer = w
			}
		}
		if v, ok := ci["timezone"]; ok && v != nil {
			resultSettings.Timezone = v.(string)
		}
	}
	result.Settings = resultSettings
	return &result
}

func constructCreateNetworkClientAndEndpointAAA(response []interface{}) *dnac.CreateNetworkRequestSettingsClientAndEndpointAAA {
	result := dnac.CreateNetworkRequestSettingsClientAndEndpointAAA{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["network"]; ok && v != nil {
			result.Network = v.(string)
		}
		if v, ok := ci["protocol"]; ok && v != nil {
			result.Protocol = v.(string)
		}
		if v, ok := ci["servers"]; ok && v != nil {
			result.Servers = v.(string)
		}
		if v, ok := ci["shared_secret"]; ok && v != nil {
			result.SharedSecret = v.(string)
		}
		return &result
	}
	return nil
}
func constructCreateNetworkDNSServer(response []interface{}) *dnac.CreateNetworkRequestSettingsDNSServer {
	result := dnac.CreateNetworkRequestSettingsDNSServer{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["domain_name"]; ok && v != nil {
			result.DomainName = v.(string)
		}
		if v, ok := ci["primary_ip_address"]; ok && v != nil {
			result.PrimaryIPAddress = v.(string)
		}
		if v, ok := ci["secondary_ip_address"]; ok && v != nil {
			result.SecondaryIPAddress = v.(string)
		}
		return &result
	}
	return nil
}
func constructCreateNetworkMessageOfTheday(response []interface{}) *dnac.CreateNetworkRequestSettingsMessageOfTheday {
	var result dnac.CreateNetworkRequestSettingsMessageOfTheday
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["banner_message"]; ok && v != nil {
			result.BannerMessage = v.(string)
		}
		if v, ok := ci["retain_existing_banner"]; ok && v != nil {
			result.RetainExistingBanner = v.(bool)
		}

		return &result
	}
	return nil
}
func constructCreateNetworkNetflowcollector(response []interface{}) *dnac.CreateNetworkRequestSettingsNetflowcollector {
	var result dnac.CreateNetworkRequestSettingsNetflowcollector
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["port"]; ok && v != nil {
			result.Port = v.(float64)
		}

		return &result
	}
	return nil
}
func constructCreateNetworkNetworkAAA(response []interface{}) *dnac.CreateNetworkRequestSettingsNetworkAAA {
	var result dnac.CreateNetworkRequestSettingsNetworkAAA
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["ip_address"]; ok && v != nil {
			result.IPAddress = v.(string)
		}
		if v, ok := ci["network"]; ok && v != nil {
			result.Network = v.(string)
		}
		if v, ok := ci["protocol"]; ok && v != nil {
			result.Protocol = v.(string)
		}
		if v, ok := ci["servers"]; ok && v != nil {
			result.Servers = v.(string)
		}
		if v, ok := ci["shared_secret"]; ok && v != nil {
			result.SharedSecret = v.(string)
		}

		return &result
	}
	return nil
}
func constructCreateNetworkSNMPServer(response []interface{}) *dnac.CreateNetworkRequestSettingsSNMPServer {
	var result dnac.CreateNetworkRequestSettingsSNMPServer
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["configure_dnac_ip"]; ok && v != nil {
			result.ConfigureDnacIP = v.(bool)
		}
		if v, ok := ci["ip_addresses"]; ok && v != nil {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}

		return &result
	}
	return nil
}
func constructCreateNetworkSyslogServer(response []interface{}) *dnac.CreateNetworkRequestSettingsSyslogServer {
	var result dnac.CreateNetworkRequestSettingsSyslogServer
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["configure_dnac_ip"]; ok && v != nil {
			result.ConfigureDnacIP = v.(bool)
		}
		if v, ok := ci["ip_addresses"]; ok && v != nil {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}

		return &result
	}
	return nil
}
func constructCreateNetwork(response []interface{}) *dnac.CreateNetworkRequest {
	result := dnac.CreateNetworkRequest{}
	resultSettings := dnac.CreateNetworkRequestSettings{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})

		if v, ok := ci["client_and_endpoint_aaa"]; ok && v != nil {
			if w := constructCreateNetworkClientAndEndpointAAA(v.([]interface{})); w != nil {
				resultSettings.ClientAndEndpointAAA = *w
			}
		}
		if v, ok := ci["dhcp_server"]; ok && v != nil {
			resultSettings.DhcpServer = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := ci["dns_server"]; ok && v != nil {
			if w := constructCreateNetworkDNSServer(v.([]interface{})); w != nil {
				resultSettings.DNSServer = *w
			}
		}
		if v, ok := ci["message_of_theday"]; ok && v != nil {
			if w := constructCreateNetworkMessageOfTheday(v.([]interface{})); w != nil {
				resultSettings.MessageOfTheday = *w
			}
		}
		if v, ok := ci["netflowcollector"]; ok && v != nil {
			if w := constructCreateNetworkNetflowcollector(v.([]interface{})); w != nil {
				resultSettings.Netflowcollector = *w
			}
		}
		if v, ok := ci["network_aaa"]; ok && v != nil {
			if w := constructCreateNetworkNetworkAAA(v.([]interface{})); w != nil {
				resultSettings.NetworkAAA = *w
			}
		}
		if v, ok := ci["ntp_server"]; ok && v != nil {
			resultSettings.NtpServer = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := ci["snmp_server"]; ok && v != nil {
			if w := constructCreateNetworkSNMPServer(v.([]interface{})); w != nil {
				resultSettings.SNMPServer = *w
			}
		}
		if v, ok := ci["syslog_server"]; ok && v != nil {
			if w := constructCreateNetworkSyslogServer(v.([]interface{})); w != nil {
				resultSettings.SyslogServer = *w
			}
		}
		if v, ok := ci["timezone"]; ok && v != nil {
			resultSettings.Timezone = v.(string)
		}
	}
	result.Settings = resultSettings
	return &result
}

func networkSimplified(response *dnac.GetNetworkResponse) *dnac.CreateNetworkRequest {
	if response != nil {
		result := dnac.CreateNetworkRequestSettings{}
		for _, item := range response.Response {
			if strings.HasPrefix(item.Key, "aaa.endpoint.server") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})
					if v, ok := valueM["ipAddress"]; ok && v != nil {
						if result.ClientAndEndpointAAA.IPAddress != "" {
							result.ClientAndEndpointAAA.IPAddress += ","
						}
						result.ClientAndEndpointAAA.IPAddress += v.(string)
					}
					if v, ok := valueM["network"]; ok && v != nil {
						if result.ClientAndEndpointAAA.Network != "" {
							result.ClientAndEndpointAAA.Network += ","
						}
						result.ClientAndEndpointAAA.Network += v.(string)
					}
					if v, ok := valueM["protocol"]; ok && v != nil {
						if result.ClientAndEndpointAAA.Protocol != "" {
							result.ClientAndEndpointAAA.Protocol += ","
						}
						result.ClientAndEndpointAAA.Protocol += v.(string)
					}
					if v, ok := valueM["sharedSecret"]; ok && v != nil {
						if result.ClientAndEndpointAAA.SharedSecret != "" {
							result.ClientAndEndpointAAA.SharedSecret += ","
						}
						result.ClientAndEndpointAAA.SharedSecret += v.(string)
					}
					if result.ClientAndEndpointAAA.Servers != "" {
						result.ClientAndEndpointAAA.Servers += ","
					}
					result.ClientAndEndpointAAA.Servers += strings.TrimPrefix(item.Key, "aaa.endpoint.server.")
				}
				continue
			}
			if strings.EqualFold(item.Key, "dhcp.server") {
				result.DhcpServer = convertSliceInterfaceToSliceString(item.Value)
				continue
			}
			if strings.EqualFold(item.Key, "dns.server") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})
					if v, ok := valueM["domainName"]; ok && v != nil {
						result.DNSServer.DomainName = v.(string)
					}
					if v, ok := valueM["primaryIpAddress"]; ok && v != nil {
						result.DNSServer.PrimaryIPAddress = v.(string)
					}
					if v, ok := valueM["secondaryIpAddress"]; ok && v != nil {
						result.DNSServer.SecondaryIPAddress = v.(string)
					}
				}
				continue
			}

			if strings.EqualFold(item.Key, "device.banner") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})

					if v, ok := valueM["bannerMessage"]; ok && v != nil {
						result.MessageOfTheday.BannerMessage = v.(string)
					}
					if v, ok := valueM["retainExistingBanner"]; ok && v != nil {
						result.MessageOfTheday.RetainExistingBanner = v.(bool)
					}
				}
				continue
			}

			if strings.EqualFold(item.Key, "netflow.collector") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})

					if v, ok := valueM["ipAddress"]; ok && v != nil {
						result.Netflowcollector.IPAddress = v.(string)
					}
					if v, ok := valueM["port"]; ok && v != nil {
						if f, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err != nil {
							result.Netflowcollector.Port = f
						}
					}
				}
				continue
			}

			if strings.HasPrefix(item.Key, "aaa.network.server") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})
					if v, ok := valueM["ipAddress"]; ok && v != nil {
						if result.NetworkAAA.IPAddress != "" {
							result.NetworkAAA.IPAddress += ","
						}
						result.NetworkAAA.IPAddress += v.(string)
					}
					if v, ok := valueM["network"]; ok && v != nil {
						if result.NetworkAAA.Network != "" {
							result.NetworkAAA.Network += ","
						}
						result.NetworkAAA.Network += v.(string)
					}
					if v, ok := valueM["protocol"]; ok && v != nil {
						if result.NetworkAAA.Protocol != "" {
							result.NetworkAAA.Protocol += ","
						}
						result.NetworkAAA.Protocol += v.(string)
					}
					if v, ok := valueM["sharedSecret"]; ok && v != nil {
						if result.NetworkAAA.SharedSecret != "" {
							result.NetworkAAA.SharedSecret += ","
						}
						result.NetworkAAA.SharedSecret += v.(string)
					}
					if result.NetworkAAA.Servers != "" {
						result.NetworkAAA.Servers += ","
					}
					result.NetworkAAA.Servers += strings.TrimPrefix(item.Key, "aaa.network.server.")
				}
				continue
			}

			if strings.EqualFold(item.Key, "ntp.server") {
				result.NtpServer = convertSliceInterfaceToSliceString(item.Value)
				continue
			}

			if strings.EqualFold(item.Key, "snmp.trap.receiver") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})

					if v, ok := valueM["configureDnacIP"]; ok && v != nil {
						result.SNMPServer.ConfigureDnacIP = v.(bool)
					}
					if v, ok := valueM["ipAddresses"]; ok && v != nil {
						result.SNMPServer.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
					}
				}
				continue
			}

			if strings.EqualFold(item.Key, "syslog.server") {
				if len(item.Value) > 0 {
					valueM := item.Value[0].(map[string]interface{})

					if v, ok := valueM["configureDnacIP"]; ok && v != nil {
						result.SyslogServer.ConfigureDnacIP = v.(bool)
					}
					if v, ok := valueM["ipAddresses"]; ok && v != nil {
						result.SyslogServer.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
					}
				}
				continue
			}

			if strings.EqualFold(item.Key, "timezone.site") {
				if len(item.Value) > 0 {
					result.Timezone = item.Value[0].(string)
				}
				continue
			}
		}
		return &dnac.CreateNetworkRequest{Settings: result}
	}
	return nil
}

func requiresNetworkUpdate(response *dnac.CreateNetworkRequest, expected *dnac.CreateNetworkRequest) bool {
	log.Printf("requiresNetworkUpdate response %+v", response)
	log.Printf("requiresNetworkUpdate expected %+v", expected)
	cmpResult := response.Settings.Timezone != expected.Settings.Timezone ||
		response.Settings.SyslogServer.ConfigureDnacIP != expected.Settings.SyslogServer.ConfigureDnacIP ||
		!hasSameSliceString(response.Settings.SyslogServer.IPAddresses, expected.Settings.SyslogServer.IPAddresses) ||
		response.Settings.SNMPServer.ConfigureDnacIP != expected.Settings.SNMPServer.ConfigureDnacIP ||
		!hasSameSliceString(response.Settings.SNMPServer.IPAddresses, expected.Settings.SNMPServer.IPAddresses) ||
		!hasSameSliceString(response.Settings.NtpServer, expected.Settings.NtpServer) ||
		response.Settings.NetworkAAA.IPAddress != expected.Settings.NetworkAAA.IPAddress ||
		response.Settings.NetworkAAA.Network != expected.Settings.NetworkAAA.Network ||
		response.Settings.NetworkAAA.Protocol != expected.Settings.NetworkAAA.Protocol ||
		response.Settings.NetworkAAA.Servers != expected.Settings.NetworkAAA.Servers ||
		response.Settings.NetworkAAA.SharedSecret != expected.Settings.NetworkAAA.SharedSecret ||
		response.Settings.Netflowcollector.IPAddress != expected.Settings.Netflowcollector.IPAddress ||
		response.Settings.Netflowcollector.Port != expected.Settings.Netflowcollector.Port ||
		response.Settings.MessageOfTheday.BannerMessage != expected.Settings.MessageOfTheday.BannerMessage ||
		response.Settings.MessageOfTheday.RetainExistingBanner != expected.Settings.MessageOfTheday.RetainExistingBanner ||
		response.Settings.DNSServer.DomainName != expected.Settings.DNSServer.DomainName ||
		response.Settings.DNSServer.PrimaryIPAddress != expected.Settings.DNSServer.PrimaryIPAddress ||
		response.Settings.DNSServer.SecondaryIPAddress != expected.Settings.DNSServer.SecondaryIPAddress ||
		!hasSameSliceString(response.Settings.DhcpServer, expected.Settings.DhcpServer) ||
		response.Settings.ClientAndEndpointAAA.IPAddress != expected.Settings.ClientAndEndpointAAA.IPAddress ||
		response.Settings.ClientAndEndpointAAA.Network != expected.Settings.ClientAndEndpointAAA.Network ||
		response.Settings.ClientAndEndpointAAA.Protocol != expected.Settings.ClientAndEndpointAAA.Protocol ||
		response.Settings.ClientAndEndpointAAA.Servers != expected.Settings.ClientAndEndpointAAA.Servers ||
		response.Settings.ClientAndEndpointAAA.SharedSecret != expected.Settings.ClientAndEndpointAAA.SharedSecret
	log.Printf("requiresNetworkUpdate %+v", cmpResult)
	return cmpResult
}

func resourceNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	items := d.Get("item").([]interface{})
	item := items[0].(map[string]interface{})
	createRequest := constructCreateNetwork(items)

	siteID := item["site_id"].(string)
	searchResponse, _, err := client.NetworkSettings.GetNetwork(&dnac.GetNetworkQueryParams{SiteID: siteID})
	if err == nil && searchResponse != nil {
		networkSimplified := networkSimplified(searchResponse)
		if requiresNetworkUpdate(networkSimplified, createRequest) {
			updateRequest := constructUpdateNetwork(items)

			// Construct payload from resource schema (item)
			log.Printf("updateRequest %+v", updateRequest)
			updateResponse, _, err := client.NetworkSettings.UpdateNetwork(siteID, updateRequest)
			if err != nil {
				return diag.FromErr(err)
			}
			log.Printf("updateResponse %+v", updateResponse)

			// Wait for execution status to complete
			time.Sleep(5 * time.Second)

		}
		d.SetId(siteID)
		resourceNetworkRead(ctx, d, m)
		return diags
	}

	// Construct payload from resource schema (item)
	_, _, err = client.NetworkSettings.CreateNetwork(siteID, createRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse2, _, err := client.NetworkSettings.GetNetwork(&dnac.GetNetworkQueryParams{SiteID: siteID})
	if err != nil {
		return diag.FromErr(err)
	}

	if err == nil && searchResponse2 != nil {
		// Update resource id
		d.SetId(siteID)
		resourceNetworkRead(ctx, d, m)
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to get created network",
	})
	return diags
}

func resourceNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteID := d.Id()

	searchResponse, _, err := client.NetworkSettings.GetNetwork(&dnac.GetNetworkQueryParams{SiteID: siteID})
	if err != nil || searchResponse == nil {
		d.SetId("")
		return diags
	}

	networkSimplified := networkSimplified(searchResponse)
	if networkSimplified == nil {
		d.SetId("")
		return diags
	}

	network := flattenNetworkReadItem(networkSimplified)
	if err := d.Set("item", network); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	siteID := d.Id()

	var diags diag.Diagnostics

	searchResponse, _, err := client.NetworkSettings.GetNetwork(&dnac.GetNetworkQueryParams{SiteID: siteID})
	if err != nil || searchResponse == nil {
		d.SetId("")
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		items := d.Get("item").([]interface{})
		updateRequest := constructUpdateNetwork(items)

		log.Printf("updateRequest %+v", updateRequest)
		updateResponse, _, err := client.NetworkSettings.UpdateNetwork(siteID, updateRequest)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Printf("updateResponse %+v", updateResponse)

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceNetworkRead(ctx, d, m)
}

func resourceNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
