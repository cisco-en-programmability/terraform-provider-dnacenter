package dnacenter

import (
	"context"
	"fmt"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func discoveryHTTPCredentialParam() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"credential_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDiscovery() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDiscoveryCreate,
		ReadContext:   resourceDiscoveryRead,
		UpdateContext: resourceDiscoveryUpdate,
		DeleteContext: resourceDiscoveryDelete,
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
				MaxItems: 1,
				Required: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cdp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"discovery_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"enable_password_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"global_credential_id_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,

							Elem: discoveryHTTPCredentialParam(),
						},
						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,

							Elem: discoveryHTTPCredentialParam(),
						},
						"ip_address_list": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"lldp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"no_add_new_device": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preferred_mgmt_ip_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol_order": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"re_discovery": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"retry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Optional: true,
							Type:     schema.TypeString,
						},
						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"update_mgmt_ip": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"user_name_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"discovery_condition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_auto_cdp": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func constructStartDiscoveryRequest(discovery map[string]interface{}) *dnac.StartDiscoveryRequest {
	discoveryRequest := dnac.StartDiscoveryRequest{}
	discoveryRequest.DiscoveryType = discovery["discovery_type"].(string)
	discoveryRequest.IPAddressList = discovery["ip_address_list"].(string)
	discoveryRequest.Name = discovery["name"].(string)

	if v, ok := discovery["cdp_level"]; ok {
		discoveryRequest.CdpLevel = v.(int)
	}
	if v, ok := discovery["enable_password_list"]; ok {
		discoveryRequest.EnablePasswordList = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := discovery["global_credential_id_list"]; ok {
		discoveryRequest.GlobalCredentialIDList = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := discovery["http_read_credential"]; ok {
		httpCredentials := v.([]interface{})
		if len(httpCredentials) > 0 {
			hC := httpCredentials[0]
			httpCredential := hC.(map[string]interface{})
			if v, ok := httpCredential["comments"]; ok {
				discoveryRequest.HTTPReadCredential.Comments = v.(string)
			}
			if v, ok := httpCredential["credential_type"]; ok {
				discoveryRequest.HTTPReadCredential.CredentialType = v.(string)
			}
			if v, ok := httpCredential["description"]; ok {
				discoveryRequest.HTTPReadCredential.Description = v.(string)
			}
			if v, ok := httpCredential["id"]; ok {
				discoveryRequest.HTTPReadCredential.ID = v.(string)
			}
			if v, ok := httpCredential["instance_tenant_id"]; ok {
				discoveryRequest.HTTPReadCredential.InstanceTenantID = v.(string)
			}
			if v, ok := httpCredential["instance_uuid"]; ok {
				discoveryRequest.HTTPReadCredential.InstanceUUID = v.(string)
			}
			if v, ok := httpCredential["password"]; ok {
				discoveryRequest.HTTPReadCredential.Password = v.(string)
			}
			if v, ok := httpCredential["port"]; ok {
				discoveryRequest.HTTPReadCredential.Port = v.(int)
			}
			if v, ok := httpCredential["secure"]; ok {
				discoveryRequest.HTTPReadCredential.Secure = v.(bool)
			}
			if v, ok := httpCredential["username"]; ok {
				discoveryRequest.HTTPReadCredential.Username = v.(string)
			}
		}
	}
	if v, ok := discovery["http_write_credential"]; ok {
		httpCredentials := v.([]interface{})
		if len(httpCredentials) > 0 {
			hC := httpCredentials[0]
			httpCredential := hC.(map[string]interface{})
			if v, ok := httpCredential["comments"]; ok {
				discoveryRequest.HTTPWriteCredential.Comments = v.(string)
			}
			if v, ok := httpCredential["credential_type"]; ok {
				discoveryRequest.HTTPWriteCredential.CredentialType = v.(string)
			}
			if v, ok := httpCredential["description"]; ok {
				discoveryRequest.HTTPWriteCredential.Description = v.(string)
			}
			if v, ok := httpCredential["id"]; ok {
				discoveryRequest.HTTPWriteCredential.ID = v.(string)
			}
			if v, ok := httpCredential["instance_tenant_id"]; ok {
				discoveryRequest.HTTPWriteCredential.InstanceTenantID = v.(string)
			}
			if v, ok := httpCredential["instance_uuid"]; ok {
				discoveryRequest.HTTPWriteCredential.InstanceUUID = v.(string)
			}
			if v, ok := httpCredential["password"]; ok {
				discoveryRequest.HTTPWriteCredential.Password = v.(string)
			}
			if v, ok := httpCredential["port"]; ok {
				discoveryRequest.HTTPWriteCredential.Port = v.(int)
			}
			if v, ok := httpCredential["secure"]; ok {
				discoveryRequest.HTTPWriteCredential.Secure = v.(bool)
			}
			if v, ok := httpCredential["username"]; ok {
				discoveryRequest.HTTPWriteCredential.Username = v.(string)
			}
		}
	}
	if v, ok := discovery["ip_filter_list"]; ok {
		discoveryRequest.IPFilterList = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := discovery["lldp_level"]; ok {
		discoveryRequest.LldpLevel = v.(int)
	}
	if v, ok := discovery["netconf_port"]; ok {
		discoveryRequest.NetconfPort = v.(string)
	}
	if v, ok := discovery["no_add_new_device"]; ok {
		discoveryRequest.NoAddNewDevice = v.(bool)
	}
	if v, ok := discovery["parent_discovery_id"]; ok {
		discoveryRequest.ParentDiscoveryID = v.(string)
	}
	if v, ok := discovery["password_list"]; ok {
		discoveryRequest.PasswordList = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := discovery["preferred_mgmt_ip_method"]; ok {
		discoveryRequest.PreferredMgmtIPMethod = v.(string)
	}
	if v, ok := discovery["protocol_order"]; ok {
		discoveryRequest.ProtocolOrder = v.(string)
	}
	if v, ok := discovery["re_discovery"]; ok {
		discoveryRequest.ReDiscovery = v.(bool)
	}
	if v, ok := discovery["retry"]; ok {
		discoveryRequest.Retry = v.(int)
	}
	if v, ok := discovery["snmp_auth_passphrase"]; ok {
		discoveryRequest.SNMPAuthPassphrase = v.(string)
	}
	if v, ok := discovery["snmp_auth_protocol"]; ok {
		discoveryRequest.SNMPAuthProtocol = v.(string)
	}
	if v, ok := discovery["snmp_mode"]; ok {
		discoveryRequest.SNMPMode = v.(string)
	}
	if v, ok := discovery["snmp_priv_passphrase"]; ok {
		discoveryRequest.SNMPPrivPassphrase = v.(string)
	}
	if v, ok := discovery["snmp_priv_protocol"]; ok {
		discoveryRequest.SNMPPrivProtocol = v.(string)
	}
	if v, ok := discovery["snmp_ro_community"]; ok {
		discoveryRequest.SNMPROCommunity = v.(string)
	}
	if v, ok := discovery["snmp_ro_community_desc"]; ok {
		discoveryRequest.SNMPROCommunityDesc = v.(string)
	}
	if v, ok := discovery["snmp_rw_community"]; ok {
		discoveryRequest.SNMPRWCommunity = v.(string)
	}
	if v, ok := discovery["snmp_rw_community_desc"]; ok {
		discoveryRequest.SNMPRWCommunityDesc = v.(string)
	}
	if v, ok := discovery["snmp_user_name"]; ok {
		discoveryRequest.SNMPUserName = v.(string)
	}
	if v, ok := discovery["snmp_version"]; ok {
		discoveryRequest.SNMPUserName = v.(string)
	}
	if v, ok := discovery["timeout"]; ok {
		discoveryRequest.Timeout = v.(int)
	}
	if v, ok := discovery["update_mgmt_ip"]; ok {
		discoveryRequest.UpdateMgmtIP = v.(bool)
	}
	if v, ok := discovery["user_name_list"]; ok {
		discoveryRequest.UserNameList = convertSliceInterfaceToSliceString(v.([]interface{}))
	}

	return &discoveryRequest
}

func constructUpdateDiscoveryRequest(prevID string, discovery map[string]interface{}) *dnac.UpdatesAnExistingDiscoveryBySpecifiedIDRequest {
	discoveryRequest := dnac.UpdatesAnExistingDiscoveryBySpecifiedIDRequest{}

	if v, ok := discovery["cdp_level"]; ok {
		discoveryRequest.CdpLevel = v.(int)
	}
	if v, ok := discovery["device_ids"]; ok {
		discoveryRequest.DeviceIDs = v.(string)
	}
	if v, ok := discovery["discovery_condition"]; ok {
		discoveryRequest.DiscoveryCondition = v.(string)
	}
	if v, ok := discovery["enable_password_list"]; ok {
		discoveryRequest.EnablePasswordList = strings.Join(convertSliceInterfaceToSliceString(v.([]interface{})), ",")
	}
	if v, ok := discovery["global_credential_id_list"]; ok {
		discoveryRequest.DeviceIDs = strings.Join(convertSliceInterfaceToSliceString(v.([]interface{})), ",")
	}
	if v, ok := discovery["ip_filter_list"]; ok {
		discoveryRequest.IPFilterList = strings.Join(convertSliceInterfaceToSliceString(v.([]interface{})), ",")
	}
	if v, ok := discovery["is_auto_cdp"]; ok {
		discoveryRequest.IsAutoCdp = v.(bool)
	}
	if v, ok := discovery["lldp_level"]; ok {
		discoveryRequest.LldpLevel = v.(int)
	}
	if v, ok := discovery["netconf_port"]; ok {
		discoveryRequest.NetconfPort = v.(string)
	}
	if v, ok := discovery["parent_discovery_id"]; ok {
		discoveryRequest.ParentDiscoveryID = v.(string)
	}
	if v, ok := discovery["password_list"]; ok {
		discoveryRequest.PasswordList = strings.Join(convertSliceInterfaceToSliceString(v.([]interface{})), ",")
	}
	if v, ok := discovery["preferred_mgmt_ip_method"]; ok {
		discoveryRequest.PreferredMgmtIPMethod = v.(string)
	}
	if v, ok := discovery["protocol_order"]; ok {
		discoveryRequest.ProtocolOrder = v.(string)
	}
	if v, ok := discovery["retry"]; ok {
		discoveryRequest.RetryCount = v.(int)
	}
	if v, ok := discovery["snmp_auth_passphrase"]; ok {
		discoveryRequest.SNMPAuthPassphrase = v.(string)
	}
	if v, ok := discovery["snmp_auth_protocol"]; ok {
		discoveryRequest.SNMPAuthProtocol = v.(string)
	}
	if v, ok := discovery["snmp_mode"]; ok {
		discoveryRequest.SNMPMode = v.(string)
	}
	if v, ok := discovery["snmp_priv_passphrase"]; ok {
		discoveryRequest.SNMPPrivPassphrase = v.(string)
	}
	if v, ok := discovery["snmp_priv_protocol"]; ok {
		discoveryRequest.SNMPPrivProtocol = v.(string)
	}
	if v, ok := discovery["snmp_ro_community"]; ok {
		discoveryRequest.SNMPRoCommunity = v.(string)
	}
	if v, ok := discovery["snmp_ro_community_desc"]; ok {
		discoveryRequest.SNMPRoCommunityDesc = v.(string)
	}
	if v, ok := discovery["snmp_rw_community"]; ok {
		discoveryRequest.SNMPRwCommunity = v.(string)
	}
	if v, ok := discovery["snmp_rw_community_desc"]; ok {
		discoveryRequest.SNMPRwCommunityDesc = v.(string)
	}
	if v, ok := discovery["snmp_user_name"]; ok {
		discoveryRequest.SNMPUserName = v.(string)
	}
	if v, ok := discovery["timeout"]; ok {
		discoveryRequest.TimeOut = v.(int)
	}
	if v, ok := discovery["update_mgmt_ip"]; ok {
		discoveryRequest.UpdateMgmtIP = v.(bool)
	}
	if v, ok := discovery["user_name_list"]; ok {
		discoveryRequest.UserNameList = strings.Join(convertSliceInterfaceToSliceString(v.([]interface{})), ",")
	}

	if v, ok := discovery["http_read_credential"]; ok {
		httpCredentials := v.([]interface{})
		if len(httpCredentials) > 0 {
			hC := httpCredentials[0]
			httpCredential := hC.(map[string]interface{})
			if v, ok := httpCredential["comments"]; ok {
				discoveryRequest.HTTPReadCredential.Comments = v.(string)
			}
			if v, ok := httpCredential["credential_type"]; ok {
				discoveryRequest.HTTPReadCredential.CredentialType = v.(string)
			}
			if v, ok := httpCredential["description"]; ok {
				discoveryRequest.HTTPReadCredential.Description = v.(string)
			}
			if v, ok := httpCredential["id"]; ok {
				discoveryRequest.HTTPReadCredential.ID = v.(string)
			}
			if v, ok := httpCredential["instance_tenant_id"]; ok {
				discoveryRequest.HTTPReadCredential.InstanceTenantID = v.(string)
			}
			if v, ok := httpCredential["instance_uuid"]; ok {
				discoveryRequest.HTTPReadCredential.InstanceUUID = v.(string)
			}
			if v, ok := httpCredential["password"]; ok {
				discoveryRequest.HTTPReadCredential.Password = v.(string)
			}
			if v, ok := httpCredential["port"]; ok {
				discoveryRequest.HTTPReadCredential.Port = v.(int)
			}
			if v, ok := httpCredential["secure"]; ok {
				discoveryRequest.HTTPReadCredential.Secure = v.(bool)
			}
			if v, ok := httpCredential["username"]; ok {
				discoveryRequest.HTTPReadCredential.Username = v.(string)
			}
		}
	}
	if v, ok := discovery["http_write_credential"]; ok {
		httpCredentials := v.([]interface{})
		if len(httpCredentials) > 0 {
			hC := httpCredentials[0]
			httpCredential := hC.(map[string]interface{})
			if v, ok := httpCredential["comments"]; ok {
				discoveryRequest.HTTPWriteCredential.Comments = v.(string)
			}
			if v, ok := httpCredential["credential_type"]; ok {
				discoveryRequest.HTTPWriteCredential.CredentialType = v.(string)
			}
			if v, ok := httpCredential["description"]; ok {
				discoveryRequest.HTTPWriteCredential.Description = v.(string)
			}
			if v, ok := httpCredential["id"]; ok {
				discoveryRequest.HTTPWriteCredential.ID = v.(string)
			}
			if v, ok := httpCredential["instance_tenant_id"]; ok {
				discoveryRequest.HTTPWriteCredential.InstanceTenantID = v.(string)
			}
			if v, ok := httpCredential["instance_uuid"]; ok {
				discoveryRequest.HTTPWriteCredential.InstanceUUID = v.(string)
			}
			if v, ok := httpCredential["password"]; ok {
				discoveryRequest.HTTPWriteCredential.Password = v.(string)
			}
			if v, ok := httpCredential["port"]; ok {
				discoveryRequest.HTTPWriteCredential.Port = v.(int)
			}
			if v, ok := httpCredential["secure"]; ok {
				discoveryRequest.HTTPWriteCredential.Secure = v.(bool)
			}
			if v, ok := httpCredential["username"]; ok {
				discoveryRequest.HTTPWriteCredential.Username = v.(string)
			}
		}
	}
	discoveryRequest.ID = prevID
	discoveryRequest.DiscoveryStatus = discovery["discovery_status"].(string) // Apparently this is required for update
	discoveryRequest.DiscoveryType = discovery["discovery_type"].(string)
	discoveryRequest.IPAddressList = discovery["ip_address_list"].(string)
	discoveryRequest.Name = discovery["name"].(string)
	return &discoveryRequest
}

func resourceDiscoveryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	discovery := item.(map[string]interface{})

	// Check if element already exists
	v1, ok1 := discovery["id"]
	var userGaveID bool
	var prevID string
	if !userGaveID && ok1 {
		prevID = v1.(string)
		userGaveID = true
	}

	if userGaveID {
		searchResponse, _, err := client.Discovery.GetDiscoveryByID(prevID)
		if err == nil && searchResponse != nil {
			// It already exists on DNAC. Update resource id
			d.SetId(prevID)

			_, canUpdate := discovery["discovery_status"]
			if !canUpdate {
				// Can not update on DNAC, then just retrieve data for Tf
				resourceDiscoveryRead(ctx, d, m)
				return diags
			}
			// Construct payload from resource schema (item)
			discoveryRequest := constructUpdateDiscoveryRequest(prevID, discovery)
			// Call function to update tag resource
			response, _, err := client.Discovery.UpdatesAnExistingDiscoveryBySpecifiedID(discoveryRequest)
			if err != nil {
				return diag.FromErr(err)
			}

			// Call function to check task
			taskID := response.Response.TaskID
			taskResponse, _, err := client.Task.GetTaskByID(taskID)
			if err != nil {
				return diag.FromErr(err)
			}

			// Check if task was completed successfully
			if taskResponse.Response.IsError {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create updated discovery",
					Detail:   taskResponse.Response.FailureReason,
				})
				return diags
			}

			// Update resource data (on Terraform and DNAC), other data is unknown
			resourceDiscoveryRead(ctx, d, m)
			return diags
		}
	}

	// Construct payload from resource schema (item)
	discoveryRequest := constructStartDiscoveryRequest(discovery)

	// Call function to create tag resource
	response, _, err := client.Discovery.StartDiscovery(discoveryRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(10 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if task was completed successfully
	if taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to start discovery",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	idFound := false

	if !idFound && taskResponse.Response.Data != "" {
		d.SetId(taskResponse.Response.Data)
		idFound = true
	}
	if !idFound && taskResponse.Response.Progress != "" {
		d.SetId(taskResponse.Response.Progress)
		idFound = true
	}
	if idFound {
		resourceDiscoveryRead(ctx, d, m)
		return diags
	}
	return diag.FromErr(fmt.Errorf("Unable to retrieve id of object created"))
}

func resourceDiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of tag.id)
	discoveryID := d.Id()

	// Call function to read tag.id
	response, _, err := client.Discovery.GetDiscoveryByID(discoveryID)
	if err != nil || response == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	discovery := flattenDiscoveryReadItem(response)
	if err := d.Set("item", discovery); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDiscoveryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	discoveryID := d.Id()
	searchResponse, _, err := client.Discovery.GetDiscoveryByID(discoveryID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		discovery := item.(map[string]interface{})

		_, canUpdate := discovery["discovery_status"]
		if !canUpdate {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update discovery",
				Detail:   "item.discovery_status parameter is missing",
			})
			return diags
		}

		// Construct payload from resource schema (item)
		discoveryRequest := constructUpdateDiscoveryRequest(discoveryID, discovery)
		// Call function to update tag resource
		response, _, err := client.Discovery.UpdatesAnExistingDiscoveryBySpecifiedID(discoveryRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Call function to check task
		taskID := response.Response.TaskID
		taskResponse, _, err := client.Task.GetTaskByID(taskID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Check if task was completed successfully
		if taskResponse.Response.IsError {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update discovery",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	discoveryID := d.Id()

	searchResponse, _, err := client.Discovery.GetDiscoveryByID(discoveryID)
	if err != nil || searchResponse == nil {
		return diags
	}

	// Call function to delete resource
	response, _, err := client.Discovery.DeleteDiscoveryByID(discoveryID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResponse != nil && taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete discovery",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	return diags
}
