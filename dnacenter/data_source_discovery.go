package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscovery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoveryRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Description: `Deprecated
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"cdp_level": &schema.Schema{
							Description: `CDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"device_ids": &schema.Schema{
							Description: `Ids of the devices discovered in a discovery
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_condition": &schema.Schema{
							Description: `To indicate the discovery status. Available options: Complete or In Progress
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_status": &schema.Schema{
							Description: `Status of the discovery. Available options are: Active, Inactive, Edit
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_type": &schema.Schema{
							Description: `Type of the discovery. Available types are: 'Single', 'Range', 'CDP', 'LLDP', 'CIDR'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_password_list": &schema.Schema{
							Description: `Enable Password of the devices to be discovered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"global_credential_id_list": &schema.Schema{
							Description: `List of global credential ids to be used
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Credential Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Credential Tenant Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Credential Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Credential Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Credential Tenant Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Credential Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Unique Discovery Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address_list": &schema.Schema{
							Description: `List of IP address of the devices to be discovered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_filter_list": &schema.Schema{
							Description: `IP addresses of the devices to be filtered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auto_cdp": &schema.Schema{
							Description: `Flag to mention if CDP discovery or not
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"lldp_level": &schema.Schema{
							Description: `LLDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name for the discovery
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_port": &schema.Schema{
							Description: `Netconf port on the device. Netconf will need valid sshv2 credentials for it to work
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"num_devices": &schema.Schema{
							Description: `Number of devices discovered in the discovery
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"parent_discovery_id": &schema.Schema{
							Description: `Parent Discovery Id from which the discovery was initiated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"password_list": &schema.Schema{
							Description: `Password of the devices to be discovered
`,
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},

						"preferred_mgmt_ipmethod": &schema.Schema{
							Description: `Preferred management IP method. Available options are 'None' and 'UseLoopBack'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol_order": &schema.Schema{
							Description: `Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"retry_count": &schema.Schema{
							Description: `Number of times to try establishing connection to device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"snmp_auth_passphrase": &schema.Schema{
							Description: `Auth passphrase for SNMP
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_auth_protocol": &schema.Schema{
							Description: `SNMP auth protocol. SHA' or 'MD5'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_mode": &schema.Schema{
							Description: `Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_priv_passphrase": &schema.Schema{
							Description: `Passphrase for SNMP privacy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_priv_protocol": &schema.Schema{
							Description: `SNMP privacy protocol. 'AES128'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_ro_community": &schema.Schema{
							Description: `SNMP RO community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_ro_community_desc": &schema.Schema{
							Description: `Description for SNMP RO community
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_rw_community": &schema.Schema{
							Description: `SNMP RW community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_rw_community_desc": &schema.Schema{
							Description: `Description for SNMP RW community
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_user_name": &schema.Schema{
							Description: `SNMP username of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"time_out": &schema.Schema{
							Description: `Time to wait for device response.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"update_mgmt_ip": &schema.Schema{
							Description: `Updates Maganement IP if multiple IPs are available for a device. If set to true, when a device is rediscovered with a different IP, the management IP is updated. Default value is false
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_name_list": &schema.Schema{
							Description: `Username of the devices to be discovered
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

func dataSourceDiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDiscoveryByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Discovery.GetDiscoveryByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryGetDiscoveryByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDiscoveryByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDiscoveryByIDItem(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveryByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_info"] = flattenDiscoveryGetDiscoveryByIDItemAttributeInfo(item.AttributeInfo)
	respItem["cdp_level"] = item.CdpLevel
	respItem["device_ids"] = item.DeviceIDs
	respItem["discovery_condition"] = item.DiscoveryCondition
	respItem["discovery_status"] = item.DiscoveryStatus
	respItem["discovery_type"] = item.DiscoveryType
	respItem["enable_password_list"] = item.EnablePasswordList
	respItem["global_credential_id_list"] = item.GlobalCredentialIDList
	respItem["http_read_credential"] = flattenDiscoveryGetDiscoveryByIDItemHTTPReadCredential(item.HTTPReadCredential)
	respItem["http_write_credential"] = flattenDiscoveryGetDiscoveryByIDItemHTTPWriteCredential(item.HTTPWriteCredential)
	respItem["id"] = item.ID
	respItem["ip_address_list"] = item.IPAddressList
	respItem["ip_filter_list"] = item.IPFilterList
	respItem["is_auto_cdp"] = boolPtrToString(item.IsAutoCdp)
	respItem["lldp_level"] = item.LldpLevel
	respItem["name"] = item.Name
	respItem["netconf_port"] = item.NetconfPort
	respItem["num_devices"] = item.NumDevices
	respItem["parent_discovery_id"] = item.ParentDiscoveryID
	respItem["password_list"] = item.PasswordList
	respItem["preferred_mgmt_ipmethod"] = item.PreferredMgmtIPMethod
	respItem["protocol_order"] = item.ProtocolOrder
	respItem["retry_count"] = item.RetryCount
	respItem["snmp_auth_passphrase"] = item.SNMPAuthPassphrase
	respItem["snmp_auth_protocol"] = item.SNMPAuthProtocol
	respItem["snmp_mode"] = item.SNMPMode
	respItem["snmp_priv_passphrase"] = item.SNMPPrivPassphrase
	respItem["snmp_priv_protocol"] = item.SNMPPrivProtocol
	respItem["snmp_ro_community"] = item.SNMPRoCommunity
	respItem["snmp_ro_community_desc"] = item.SNMPRoCommunityDesc
	respItem["snmp_rw_community"] = item.SNMPRwCommunity
	respItem["snmp_rw_community_desc"] = item.SNMPRwCommunityDesc
	respItem["snmp_user_name"] = item.SNMPUserName
	respItem["time_out"] = item.TimeOut
	respItem["update_mgmt_ip"] = boolPtrToString(item.UpdateMgmtIP)
	respItem["user_name_list"] = item.UserNameList
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDiscoveryGetDiscoveryByIDItemAttributeInfo(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveryByIDResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDiscoveryGetDiscoveryByIDItemHTTPReadCredential(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveryByIDResponseHTTPReadCredential) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["comments"] = item.Comments
	respItem["credential_type"] = item.CredentialType
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["password"] = item.Password
	respItem["port"] = item.Port
	respItem["secure"] = boolPtrToString(item.Secure)
	respItem["username"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDiscoveryGetDiscoveryByIDItemHTTPWriteCredential(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveryByIDResponseHTTPWriteCredential) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["comments"] = item.Comments
	respItem["credential_type"] = item.CredentialType
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["password"] = item.Password
	respItem["port"] = item.Port
	respItem["secure"] = boolPtrToString(item.Secure)
	respItem["username"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}
