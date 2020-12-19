package dnacenter

import (
	"context"
	"strings"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaSDAFabricBorderDeviceNetworkWideSettingsIP() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"padded_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func schemaSDAFabricBorderDeviceNetworkWideSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aaa": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cmx": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deploy_pending": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     schemaSDAFabricBorderDeviceNetworkWideSettingsIP(),
						},
					},
				},
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     schemaSDAFabricBorderDeviceNetworkWideSettingsIP(),
						},
					},
				},
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ldap": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"native_vlan": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"netflow": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ntp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snmp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"syslogs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func schemaSDAFabricBorderDeviceDeviceSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connected_to": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"deploy_pending": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ext_connectivity_settings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deploy_pending": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_domain_protocol_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_version": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interface_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"l2_handoff": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"l3_handoff": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deploy_pending": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_version": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"local_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"virtual_network": &schema.Schema{
										Type:     schema.TypeList, //GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id_ref": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"vlan_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"policy_propagation_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"policy_sgt_tag": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"external_connectivity_ip_pool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_domain_routing_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"internal_domain_protocol_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"node_type": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"storage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func schemaSDAFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"akc_settings_cfs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"auth_entity_class": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_entity_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cfs_change_info": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custom_provisions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deploy_pending": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployed": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"device_interface_info": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"device_settings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     schemaSDAFabricBorderDeviceDeviceSettings(),
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_seeded": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_stale": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_update_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"managed_sites": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_device_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_wide_settings": &schema.Schema{
				Type:     schema.TypeList, // GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettings,
				Computed: true,
				Elem:     schemaSDAFabricBorderDeviceNetworkWideSettings(),
			},
			"other_device": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"provisioning_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"save_wan_connectivity_details_only": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"transit_networks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id_ref": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_network": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"wlan": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSDAFabricBorderDevice() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricBorderDeviceCreate,
		ReadContext:   resourceSDAFabricBorderDeviceRead,
		UpdateContext: resourceSDAFabricBorderDeviceUpdate,
		DeleteContext: resourceSDAFabricBorderDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			// For other get elements
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Required: true, //REVIEW:.
				// MinItems: 1,
				Elem: schemaSDAFabricBorderDevice(),
			},

			"device_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"site_name_hierarchy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_domain_routing_protocol_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW:.
			},
			"external_connectivity_ip_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW:.
			},
			"internal_autonomous_system_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW:.
			},
			"border_session_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW:.
			},
			"connected_to_internet": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true, //REVIEW:.
			},
			"external_connectivity_settings": &schema.Schema{
				Type:     schema.TypeList,
				Required: true, //REVIEW:.
				// MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true, //REVIEW.
						},
						"external_autonomous_system_number": &schema.Schema{
							Type:     schema.TypeString,
							Required: true, //REVIEW.
						},
						"l3_handoff": &schema.Schema{
							Type:     schema.TypeList,
							Required: true, //REVIEW:.
							// MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_network": &schema.Schema{
										Type:     schema.TypeList,
										Required: true, //REVIEW:.
										// MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"virtual_network_name": &schema.Schema{
													Type:     schema.TypeString,
													Required: true, //REVIEW.
												},
											},
										},
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

func constructCreateSDAFabricBorderDeviceExternalConnectivitySettingsL3HandoffVirtualNetwork(response []interface{}) *dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork {
	result := dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork{}
	if len(response) > 0 {
		item := response[0]
		ci := item.(map[string]interface{})
		if v, ok := ci["virtual_network_name"]; ok {
			result.VirtualNetworkName = v.(string)
		}
		return &result
	}
	return nil
}

func constructCreateSDAFabricBorderDeviceExternalConnectivitySettingsL3Handoff(response []interface{}) *[]dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff {
	var result []dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff
	for _, item := range response {
		ci := item.(map[string]interface{})
		cItem := dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff{}
		if v, ok := ci["virtual_network"]; ok && v != nil {
			if w := constructCreateSDAFabricBorderDeviceExternalConnectivitySettingsL3HandoffVirtualNetwork(v.([]interface{})); w != nil {
				cItem.VirtualNetwork = *w
			}
		}
		result = append(result, cItem)
	}
	return &result
}

func constructCreateSDAFabricBorderDeviceExternalConnectivitySettings(response []interface{}) *[]dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings {
	var result []dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings
	for _, item := range response {
		ci := item.(map[string]interface{})
		cItem := dnac.AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings{}
		if v, ok := ci["external_autonomous_system_number"]; ok {
			cItem.ExternalAutonomouSystemNumber = v.(string)
		}
		if v, ok := ci["interface_name"]; ok {
			cItem.InterfaceName = v.(string)
		}
		if v, ok := ci["l3_handoff"]; ok && v != nil {
			if w := constructCreateSDAFabricBorderDeviceExternalConnectivitySettingsL3Handoff(v.([]interface{})); w != nil {
				cItem.L3Handoff = *w
			}
		}
		result = append(result, cItem)
	}
	return &result
}

func resourceSDAFabricBorderDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	deviceIPAddress := d.Get("device_ip_address").(string)
	siteNameHierarchy := d.Get("site_name_hierarchy").(string)

	searchResponse, _, err := client.SDA.GetsBorderDeviceDetailFromSDAFabric(&dnac.GetsBorderDeviceDetailFromSDAFabricQueryParams{
		DeviceIPAddress: deviceIPAddress,
	})
	if err == nil && searchResponse != nil {
		if "success" == searchResponse.Status {
			// Update resource id
			d.SetId(strings.Join([]string{deviceIPAddress, siteNameHierarchy}, "_/_"))
			resourceSDAFabricBorderDeviceRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddsBorderDeviceInSDAFabricRequest
	var request dnac.AddsBorderDeviceInSDAFabricRequest
	request.DeviceManagementIPAddress = deviceIPAddress
	request.SiteNameHierarchy = siteNameHierarchy
	if v, ok := d.GetOk("external_domain_routing_protocol_name"); ok {
		request.ExternalDomainRoutingProtocolName = v.(string)
	}
	if v, ok := d.GetOk("external_connectivity_ip_pool_name"); ok {
		request.ExternalConnectivityIPPoolName = v.(string)
	}
	if v, ok := d.GetOk("internal_autonomous_system_number"); ok {
		request.InternalAutonomouSystemNumber = v.(string)
	}
	if v, ok := d.GetOk("border_session_type"); ok {
		request.BorderSessionType = v.(string)
	}
	if v, ok := d.GetOk("connected_to_internet"); ok {
		request.ConnectedToInternet = v.(bool)
	}
	if v, ok := d.GetOk("external_connectivity_settings"); ok && v != nil {
		if w := constructCreateSDAFabricBorderDeviceExternalConnectivitySettings(v.([]interface{})); w != nil {
			request.ExternalConnectivitySettings = *w
		}
	}

	requests = append(requests, request)
	_, _, err = client.SDA.AddsBorderDeviceInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(strings.Join([]string{deviceIPAddress, siteNameHierarchy}, "_/_"))
	resourceSDAFabricBorderDeviceRead(ctx, d, m)
	return diags
}

func resourceSDAFabricBorderDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	deviceIPAddress, siteNameHierarchy := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetsBorderDeviceDetailFromSDAFabric(&dnac.GetsBorderDeviceDetailFromSDAFabricQueryParams{
		DeviceIPAddress: deviceIPAddress,
	})
	if err != nil || searchResponse == nil {
		d.SetId("")
		return diags
	}
	if "success" != searchResponse.Status {
		d.SetId("")
		return diags
	}

	if err := d.Set("device_ip_address", deviceIPAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("site_name_hierarchy", siteNameHierarchy); err != nil {
		return diag.FromErr(err)
	}
	// REVIEW:.
	item := flattenSDAFabricBorderDevice(searchResponse)
	if err := d.Set("item", item); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricBorderDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSDAFabricBorderDeviceRead(ctx, d, m)
}

func resourceSDAFabricBorderDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
	/*
		client := m.(*dnac.Client)
		resourceIDs := strings.Split(d.Id(), "_/_")
		deviceIPAddress := resourceIDs[0]

		searchResponse, _, err := client.SDA.GetsBorderDeviceDetailFromSDAFabric(&dnac.GetsBorderDeviceDetailFromSDAFabricQueryParams{
			DeviceIPAddress: deviceIPAddress,
		})
		if err != nil || searchResponse == nil {
			return diags
		}
		if "success" != searchResponse.Status {
			return diags
		}

		// Call function to delete resource
		deleteRequest := dnac.DeletesBorderDeviceFromSDAFabricRequest{}
		_, _, err = client.SDA.DeleteSiteFromSDAFabric(&dnac.DeletesBorderDeviceFromSDAFabricQueryParams{
			DeviceIPAddress: deviceIPAddress,
		}, &deleteRequest{})
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		searchResponse, _, err = client.SDA.GetsBorderDeviceDetailFromSDAFabric(&dnac.GetsBorderDeviceDetailFromSDAFabricQueryParams{
			DeviceIPAddress: deviceIPAddress,
		})

		if err == nil && searchResponse != nil {
			// Check if element already exists
			if "success" == searchResponse.Status {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to delete SDA fabric border device",
				})
			}
			return diags
		}

		return diags
	*/
}
