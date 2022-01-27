package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryRange() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the discovery by specified range
`,

		ReadContext: dataSourceDiscoveryRangeRead,
		Schema: map[string]*schema.Schema{
			"records_to_return": &schema.Schema{
				Description: `recordsToReturn path parameter. Number of records to return
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_index": &schema.Schema{
				Description: `startIndex path parameter. Start index
`,
				Type:     schema.TypeInt,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"cdp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_condition": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_password_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"global_credential_id_list": &schema.Schema{
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
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
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
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_filter_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auto_cdp": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"lldp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"password_list": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},

						"preferred_mgmt_ipmethod": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol_order": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"snmp_auth_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_ro_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_rw_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"update_mgmt_ip": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_name_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartIndex := d.Get("start_index")
	vRecordsToReturn := d.Get("records_to_return")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDiscoveriesByRange")
		vvStartIndex := vStartIndex.(int)
		vvRecordsToReturn := vRecordsToReturn.(int)

		response1, restyResp1, err := client.Discovery.GetDiscoveriesByRange(vvStartIndex, vvRecordsToReturn)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveriesByRange", err,
				"Failure at GetDiscoveriesByRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetDiscoveriesByRangeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDiscoveriesByRange response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDiscoveriesByRangeItems(items *[]dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDiscoveryGetDiscoveriesByRangeItemsAttributeInfo(item.AttributeInfo)
		respItem["cdp_level"] = item.CdpLevel
		respItem["device_ids"] = item.DeviceIDs
		respItem["discovery_condition"] = item.DiscoveryCondition
		respItem["discovery_status"] = item.DiscoveryStatus
		respItem["discovery_type"] = item.DiscoveryType
		respItem["enable_password_list"] = item.EnablePasswordList
		respItem["global_credential_id_list"] = item.GlobalCredentialIDList
		respItem["http_read_credential"] = flattenDiscoveryGetDiscoveriesByRangeItemsHTTPReadCredential(item.HTTPReadCredential)
		respItem["http_write_credential"] = flattenDiscoveryGetDiscoveriesByRangeItemsHTTPWriteCredential(item.HTTPWriteCredential)
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
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetDiscoveriesByRangeItemsAttributeInfo(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDiscoveryGetDiscoveriesByRangeItemsHTTPReadCredential(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPReadCredential) []map[string]interface{} {
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

func flattenDiscoveryGetDiscoveriesByRangeItemsHTTPWriteCredential(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPWriteCredential) []map[string]interface{} {
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
