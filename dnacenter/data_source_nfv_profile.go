package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNfvProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- API to get NFV network profile.
`,

		ReadContext: dataSourceNfvProfileRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. ID of network profile to retrieve.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of profile to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of network profile to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_networks": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connection_type": &schema.Schema{
													Description: `Type of network connection from custom network (eg: lan)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"network_name": &schema.Schema{
													Description: `name of custom network (eg: cust-1)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"services_to_connect": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"service_name": &schema.Schema{
																Description: `Name of service to be connected to the custom network (eg: router-1)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"vlan_id": &schema.Schema{
													Description: `Vlan id for the custom network(eg: 4000)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vlan_mode": &schema.Schema{
													Description: `Vlan network mode (eg Access or Trunk)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"custom_template": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_type": &schema.Schema{
													Description: `Type of the device(eg: Cisco 5400 Enterprise Network Compute System)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"template": &schema.Schema{
													Description: `Name of the template(eg NFVIS template)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"template_type": &schema.Schema{
													Description: `Name of the template to which template is associated (eg: Cloud DayN Templates)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"device_tag": &schema.Schema{
										Description: `Device Tag name(eg: dev1)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_type": &schema.Schema{
										Description: `Name of the device used in creating nfv profile(eg: Cisco 5400 Enterprise Network Compute System)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"direct_internet_access_for_firewall": &schema.Schema{
										Description: `Direct internet access value should be boolean (eg: false)
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"service_provider_profile": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connect": &schema.Schema{
													Description: `Connection of service provider and device value should be boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"connect_default_gateway_on_wan": &schema.Schema{
													Description: `Default gateway connect value as boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"link_type": &schema.Schema{
													Description: `Name of connection type(eg: GigabitEthernet) 
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"service_provider": &schema.Schema{
													Description: `Name of the service provider(eg: Airtel)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"services": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"firewall_mode": &schema.Schema{
													Description: `Mode of firewall (eg: routed, transparent)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"image_name": &schema.Schema{
													Description: `Service image name (eg: isrv-universalk9.16.12.01a.tar.gz)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"profile_type": &schema.Schema{
													Description: `Profile type of service (eg: ISRv-mini)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"service_name": &schema.Schema{
													Description: `Name of service (eg: router-1)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"service_type": &schema.Schema{
													Description: `Service type (eg: ISRV)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"v_nic_mapping": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"assign_ip_address_to_network": &schema.Schema{
																Description: `Assign ip address to network (eg: true)
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"network_type": &schema.Schema{
																Description: `Type of connection (eg:  wan, lan or internal)
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

									"vlan_for_l2": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"vlan_description": &schema.Schema{
													Description: `Vlan description(eg. Access 4018)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vlan_id": &schema.Schema{
													Description: `Vlan id(eg.4018)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vlan_type": &schema.Schema{
													Description: `Vlan type(eg. Access or Trunk)
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

						"id": &schema.Schema{
							Description: `Id of nfv created nfv profile
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `Name of the profile to create NFV profile( eg: Nfvis_profile)
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

func dataSourceNfvProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNfvProfile")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetNfvProfileQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.SiteDesign.GetNfvProfile(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNfvProfile", err,
				"Failure at GetNfvProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignGetNfvProfileItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNfvProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetNfvProfileItems(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_name"] = item.ProfileName
		respItem["id"] = item.ID
		respItem["device"] = flattenSiteDesignGetNfvProfileItemsDevice(item.Device)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDevice(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_type"] = item.DeviceType
		respItem["device_tag"] = item.DeviceTag
		respItem["service_provider_profile"] = flattenSiteDesignGetNfvProfileItemsDeviceServiceProviderProfile(item.ServiceProviderProfile)
		respItem["direct_internet_access_for_firewall"] = boolPtrToString(item.DirectInternetAccessForFirewall)
		respItem["services"] = flattenSiteDesignGetNfvProfileItemsDeviceServices(item.Services)
		respItem["custom_networks"] = flattenSiteDesignGetNfvProfileItemsDeviceCustomNetworks(item.CustomNetworks)
		respItem["vlan_for_l2"] = flattenSiteDesignGetNfvProfileItemsDeviceVLANForL2(item.VLANForL2)
		respItem["custom_template"] = flattenSiteDesignGetNfvProfileItemsDeviceCustomTemplate(item.CustomTemplate)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceServiceProviderProfile(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceServiceProviderProfile) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["link_type"] = item.LinkType
		respItem["connect"] = boolPtrToString(item.Connect)
		respItem["connect_default_gateway_on_wan"] = boolPtrToString(item.ConnectDefaultGatewayOnWan)
		respItem["service_provider"] = item.ServiceProvider
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceServices(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceServices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["service_type"] = item.ServiceType
		respItem["profile_type"] = item.ProfileType
		respItem["service_name"] = item.ServiceName
		respItem["image_name"] = item.ImageName
		respItem["v_nic_mapping"] = flattenSiteDesignGetNfvProfileItemsDeviceServicesVNicMapping(item.VNicMapping)
		respItem["firewall_mode"] = item.FirewallMode
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceServicesVNicMapping(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceServicesVnicMapping) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["network_type"] = item.NetworkType
		respItem["assign_ip_address_to_network"] = boolPtrToString(item.AssignIPAddressToNetwork)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceCustomNetworks(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["network_name"] = item.NetworkName
		respItem["services_to_connect"] = flattenSiteDesignGetNfvProfileItemsDeviceCustomNetworksServicesToConnect(item.ServicesToConnect)
		respItem["connection_type"] = item.ConnectionType
		respItem["vlan_mode"] = item.VLANMode
		respItem["vlan_id"] = item.VLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceCustomNetworksServicesToConnect(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworksServicesToConnect) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["service_name"] = item.ServiceName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceVLANForL2(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceVLANForL2) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["vlan_type"] = item.VLANType
		respItem["vlan_id"] = item.VLANID
		respItem["vlan_description"] = item.VLANDescription
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetNfvProfileItemsDeviceCustomTemplate(items *[]dnacentersdkgo.ResponseSiteDesignGetNfvProfileResponseDeviceCustomTemplate) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_type"] = item.DeviceType
		respItem["template"] = item.Template
		respItem["template_type"] = item.TemplateType
		respItems = append(respItems, respItem)
	}
	return respItems
}
