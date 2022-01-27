package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDiscovery() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Discovery.

- Stops all the discoveries and removes them

- Stops or starts an existing discovery

- Initiates discovery with the given parameters

- Stops the discovery for the given Discovery ID and removes it. Discovery ID can be obtained using the "Get Discoveries
by range" API.
`,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cdp_level": &schema.Schema{
							Description: `CDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
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
						"discovery_type": &schema.Schema{
							Description: `Type of Discovery. 'SINGLE', 'RANGE', 'MULTI RANGE', 'CDP', 'LLDP'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_password_list": &schema.Schema{
							Description: `Enable Password of the devices to be discovered

ERROR: Different types for param enablePasswordList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"global_credential_id_list": &schema.Schema{
							Description: `Global Credential Ids to be used for discovery
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
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
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
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
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address_list": &schema.Schema{
							Description: `IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_filter_list": &schema.Schema{
							Description: `IP Addresses of the devices to be filtered out during discovery

ERROR: Different types for param ipFilterList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"is_auto_cdp": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"lldp_level": &schema.Schema{
							Description: `LLDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of the discovery
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"netconf_port": &schema.Schema{
							Description: `Netconf Port. It will need valid SSH credentials to work
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_list": &schema.Schema{
							Description: `Password of the devices to be discovered

ERROR: Different types for param passwordList schema.TypeList schema.TypeString`,
							Type:      schema.TypeList,
							Optional:  true,
							Sensitive: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preferred_mgmt_ipmethod": &schema.Schema{
							Description: `Preferred Management IP Method.'None' or 'UseLoopBack'. Default is 'None'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol_order": &schema.Schema{
							Description: `Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"retry": &schema.Schema{
							Description: `Number of times to try establishing connection to device
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Description: `Auth Pass phrase for SNMP
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Description: `SNMP auth protocol. SHA' or 'MD5'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Description: `Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Description: `Pass phrase for SNMP privacy
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Description: `SNMP privacy protocol. 'DES' or 'AES128'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Description: `Snmp RO community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Description: `Description for Snmp RO community
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Description: `Snmp RW community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Description: `Description for Snmp RW community
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_user_name": &schema.Schema{
							Description: `SNMP username of the device
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_version": &schema.Schema{
							Description: `Version of SNMP. v2 or v3
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"timeout": &schema.Schema{
							Description: `Time to wait for device response in seconds
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"update_mgmt_ip": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"user_name_list": &schema.Schema{
							Description: `Username of the devices to be discovered

ERROR: Different types for param userNameList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceDiscoveryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDiscoveryStartDiscovery(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.Discovery.GetDiscoveryByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDiscoveryRead(ctx, d, m)
		}
	}

	if okName && vvName != "" {
		getResponse1, err := searchDiscovery(m, vvName)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDiscoveryRead(ctx, d, m)
		}
	}

	resp1, restyResp1, err := client.Discovery.StartDiscovery(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing StartDiscovery", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing StartDiscovery", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing StartDiscovery", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			diags = append(diags, diagError(
				"Failure when executing CreateDiscovery", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vName := resourceMap["name"]

	if vID != "" {
		log.Printf("[DEBUG] Selected method 1: GetDiscoveryByID")
		vvID := vID

		response1, restyResp1, err := client.Discovery.GetDiscoveryByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDiscoveryByID", err,
			// 	"Failure at GetDiscoveryByID, unexpected response", ""))
			// return diags
			d.SetId("")
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
		return diags

	}

	if vName != "" {
		response1, err := searchDiscovery(m, vName)
		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDiscoveryByID", err,
			// 	"Failure at GetDiscoveryByID, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		vID = response1.ID
		response2, restyResp1, err := client.Discovery.GetDiscoveryByID(vID)

		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}
		vItem1 := flattenDiscoveryGetDiscoveryByIDItem(response2.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDiscoveryByID response",
				err))
			return diags
		}
		return diags
	}
	return diags
}

func resourceDiscoveryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vName := resourceMap["name"]

	var vvID string
	var vvName string
	if vID != "" {
		vvID = vID
		getResp, _, err := client.Discovery.GetDiscoveryByID(vvID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}
	} else if vName != "" {
		getResp, err := searchDiscovery(m, vName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}
		vvID = getResp.ID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil && request1.ID == "" {
			request1.ID = vvID
		}
		response1, restyResp1, err := client.Discovery.UpdatesAnExistingDiscoveryBySpecifiedID(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAnExistingDiscoveryBySpecifiedID", err, restyResp1.String(),
					"Failure at UpdatesAnExistingDiscoveryBySpecifiedID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAnExistingDiscoveryBySpecifiedID", err,
				"Failure at UpdatesAnExistingDiscoveryBySpecifiedID, unexpected response", ""))
			return diags
		}
		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesAnExistingDiscoveryBySpecifiedID", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				diags = append(diags, diagError(
					"Failure when executing UpdateDiscovery", err))
				return diags
			}
		}
	}

	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vName := resourceMap["name"]

	var vvID string
	if vID != "" {
		vvID = vID
		getResp, _, err := client.Discovery.GetDiscoveryByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	} else if vName != "" {
		getResp, err := searchDiscovery(m, vName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		if getResp != nil && getResp.ID == "" {
			vvID = getResp.ID
		}
	}
	response1, restyResp1, err := client.Discovery.DeleteDiscoveryByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDiscoveryByID", err, restyResp1.String(),
				"Failure at DeleteDiscoveryByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDiscoveryByID", err,
			"Failure at DeleteDiscoveryByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDiscoveryStartDiscovery(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscovery {
	request := dnacentersdkgo.RequestDiscoveryStartDiscovery{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cdp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cdp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cdp_level")))) {
		request.CdpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_type")))) {
		request.DiscoveryType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_list")))) {
		request.EnablePasswordList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_credential_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_credential_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_credential_id_list")))) {
		request.GlobalCredentialIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read_credential")))) {
		request.HTTPReadCredential = expandRequestDiscoveryStartDiscoveryHTTPReadCredential(ctx, key+".http_read_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write_credential")))) {
		request.HTTPWriteCredential = expandRequestDiscoveryStartDiscoveryHTTPWriteCredential(ctx, key+".http_write_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_list")))) {
		request.IPAddressList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_filter_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_filter_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_filter_list")))) {
		request.IPFilterList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lldp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lldp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lldp_level")))) {
		request.LldpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_list")))) {
		request.PasswordList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_mgmt_ipmethod")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) {
		request.PreferredMgmtIPMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol_order")))) {
		request.ProtocolOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retry")))) {
		request.Retry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) {
		request.SNMPROCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) {
		request.SNMPRWCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_list")))) {
		request.UserNameList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryStartDiscoveryHTTPReadCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPReadCredential {
	request := dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPReadCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryStartDiscoveryHTTPWriteCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPWriteCredential {
	request := dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPWriteCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_info")))) {
		value := d.Get(fixKeyAccess(key + ".attribute_info"))
		if len(interfaceToString(value)) > 0 {
			request.AttributeInfo = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo(ctx, key+".attribute_info", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cdp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cdp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cdp_level")))) {
		request.CdpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ids")))) {
		request.DeviceIDs = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_condition")))) {
		request.DiscoveryCondition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_status")))) {
		request.DiscoveryStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_type")))) {
		request.DiscoveryType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_list")))) {
		request.EnablePasswordList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_credential_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_credential_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_credential_id_list")))) {
		request.GlobalCredentialIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read_credential")))) {
		request.HTTPReadCredential = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential(ctx, key+".http_read_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write_credential")))) {
		request.HTTPWriteCredential = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential(ctx, key+".http_write_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_list")))) {
		request.IPAddressList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_filter_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_filter_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_filter_list")))) {
		request.IPFilterList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auto_cdp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auto_cdp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auto_cdp")))) {
		request.IsAutoCdp = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lldp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lldp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lldp_level")))) {
		request.LldpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_devices")))) {
		request.NumDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_discovery_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_discovery_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_discovery_id")))) {
		request.ParentDiscoveryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_list")))) {
		request.PasswordList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_mgmt_ipmethod")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) {
		request.PreferredMgmtIPMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol_order")))) {
		request.ProtocolOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retry_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retry_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retry_count")))) {
		request.RetryCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPRoCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) {
		request.SNMPRoCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRwCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) {
		request.SNMPRwCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_out")))) {
		request.TimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_mgmt_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_mgmt_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_mgmt_ip")))) {
		request.UpdateMgmtIP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_list")))) {
		request.UserNameList = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo {
	var request dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscovery(m interface{}, vName string) (*dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDiscoveryGetDiscoveriesByRangeResponse
	if vName != "" {
		totalDiscovery, _, err := client.Discovery.GetCountOfAllDiscoveryJobs()
		if err != nil || totalDiscovery == nil {
			return foundItem, err
		}

		if err != nil || totalDiscovery.Response == nil || *totalDiscovery.Response < 1 {
			return foundItem, err
		}

		response, _, err := client.Discovery.GetDiscoveriesByRange(1, *totalDiscovery.Response)
		if err != nil || response == nil {
			return foundItem, err
		}

		for _, item := range *response.Response {
			if vName == item.Name {
				return &item, err
			}
		}
	}
	return foundItem, err
}
