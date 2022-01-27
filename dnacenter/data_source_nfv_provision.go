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

// dataSourceAction
func dataSourceNfvProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Design and Provision single/multi NFV device with given site/area/building/floor .
`,

		ReadContext: dataSourceNfvProvisionRead,
		Schema: map[string]*schema.Schema{
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
			"runsync": &schema.Schema{
				Description: `__runsync header parameter. Enable this parameter to execute the API and return a response synchronously
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"timeout": &schema.Schema{
				Description: `__timeout header parameter. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
			`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"provisioning": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_networks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address_pool": &schema.Schema{
													Description: `IP address pool of sub pool (eg: 175.175.140.1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of custom network (eg: cust-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"port": &schema.Schema{
													Description: `Port for custom network (eg: 443)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"device_serial_number": &schema.Schema{
										Description: `Serial number of device (eg: FGL210710QY)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip": &schema.Schema{
										Description: `IP address of the device (eg: 172.20.126.90)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"service_providers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"service_provider": &schema.Schema{
													Description: `Name of the service provider (eg: Airtel)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"wan_interface": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth": &schema.Schema{
																Description: `Bandwidth limit (eg: 100)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"gateway": &schema.Schema{
																Description: `Gateway (eg: 175.175.190.1)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"interface_name": &schema.Schema{
																Description: `Name of the interface (eg: GE0-0)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"ip_address": &schema.Schema{
																Description: `IP address (eg: 175.175.190.205)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"subnetmask": &schema.Schema{
																Description: `Subnet mask (eg: 255.255.255.0)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"services": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_password_hash": &schema.Schema{
													Description: `Admin password hash
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"central_manager_ip": &schema.Schema{
													Description: `WAAS Package needs to be installed to populate Central Manager IP automatically.
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"central_registration_key": &schema.Schema{
													Description: `Central registration key 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"common_key": &schema.Schema{
													Description: `Common key 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"disk": &schema.Schema{
													Description: `Name of disk type (eg: internal)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"mode": &schema.Schema{
													Description: `Mode of firewall (eg: transparent)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"system_ip": &schema.Schema{
													Description: `System IP 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Type of service (eg: ISR)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"sub_pools": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"gateway": &schema.Schema{
													Description: `IP address for gate way (eg: 175.175.140.1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"ip_subnet": &schema.Schema{
													Description: `IP pool cidir (eg: 175.175.140.0)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the ip sub pool (eg; Lan-65)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_pool_name": &schema.Schema{
													Description: `Name of parent pool (global pool name)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Tyep of ip sub pool (eg: Lan)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"tag_name": &schema.Schema{
										Description: `Name of device tag (eg: dev1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"template_param": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"asav": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"var1": &schema.Schema{
																Description: `Variable for asav template (eg: "test":"Hello asav")
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"nfvis": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"var1": &schema.Schema{
																Description: `Variable for nfvis template (eg: "test":"Hello nfvis")
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"vlan": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Vlan id(e: .4018)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"interfaces": &schema.Schema{
													Description: `Interface (eg: GigabitEathernet1/0)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"network": &schema.Schema{
													Description: `Network name to connect (eg: lan-net)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Vlan type(eg. Access or Trunk)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"site": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"area": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name of the area (eg: Area1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of the area to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"building": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `Address of the building to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"latitude": &schema.Schema{
													Description: `Latitude coordinate of the building (eg:37.338)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"longitude": &schema.Schema{
													Description: `Longitude coordinate of the building (eg:-121.832)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the building (eg: building1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Address of the building to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"floor": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"height": &schema.Schema{
													Description: `Height of the floor (eg: 15)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"length": &schema.Schema{
													Description: `Length of the floor (eg: 100)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the floor (eg:floor-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"parent_name": &schema.Schema{
													Description: `Parent name of the floor to be created
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"rf_model": &schema.Schema{
													Description: `Type of floor (eg: Cubes And Walled Offices)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"width": &schema.Schema{
													Description: `Width of the floor (eg:100)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
											},
										},
									},
									"site_profile_name": &schema.Schema{
										Description: `Name of site profile to be provision with device 
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"site_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_networks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connection_type": &schema.Schema{
													Description: `Type of network connection from custom network (eg: lan)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of custom network (eg: cust-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"network_mode": &schema.Schema{
													Description: `Network mode (eg Access or Trunk)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"services_to_connect": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"service": &schema.Schema{
																Description: `Name of service to be connected to the custom network (eg: router-1)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"vlan": &schema.Schema{
													Description: `Vlan id for the custom network(eg: 4000)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"custom_services": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"application_type": &schema.Schema{
													Description: `Application type of custom service (eg: LINUX)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"image_name": &schema.Schema{
													Description: `Image name of custom service (eg: redhat7.tar.gz.tar.gz)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of custom service (eg: LINUX-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"profile": &schema.Schema{
													Description: `Profile type of service (eg: rhel7-medium)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"topology": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"assign_ip": &schema.Schema{
																Description: `Assign ip to network (eg: true)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"name": &schema.Schema{
																Description: `Name of connection from custom service(eg: wan-net)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"type": &schema.Schema{
																Description: `Type of connection from custom service (eg:  wan, lan or internal)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"custom_template": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_type": &schema.Schema{
													Description: `Type of the device(eg: NFVIS)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"template": &schema.Schema{
													Description: `Name of the template(eg NFVIS template)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"device_type": &schema.Schema{
										Description: `Name of the device used in creating nfv profile(eg: ENCS5400)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"dia": &schema.Schema{
										Description: `Direct internet access value should be boolean (eg: false)
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"service_providers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connect": &schema.Schema{
													Description: `Connection of service provider and device value should be boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"default_gateway": &schema.Schema{
													Description: `Default gateway connect value as boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"link_type": &schema.Schema{
													Description: `Name of connection type(eg: GigabitEthernet) 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"service_provider": &schema.Schema{
													Description: `Name of the service provider(eg: Airtel)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"services": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"image_name": &schema.Schema{
													Description: `Name of image (eg: isrv-universalk9.16.06.02.tar.gz)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"mode": &schema.Schema{
													Description: `Mode of firewall (eg: routed, transparent)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of the service (eg: isrv) 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"profile": &schema.Schema{
													Description: `Profile type of service (eg: ISRv-mini)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"topology": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"assign_ip": &schema.Schema{
																Description: `Assign ip address to network (eg: true)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"name": &schema.Schema{
																Description: `Name of connection (eg: wan-net)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"type": &schema.Schema{
																Description: `Type of connection (eg:  wan, lan or internal)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"type": &schema.Schema{
													Description: `Service type (eg: ISRV)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"tag_name": &schema.Schema{
										Description: `Device Tag name(eg: dev1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"vlan": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Vlan id(eg.4018)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Vlan type(eg. Access or Trunk)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"site_profile_name": &schema.Schema{
							Description: `Name of the profile to create site profile profile( eg: profile-1)
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

func dataSourceNfvProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vRunsync, okRunsync := d.GetOk("runsync")
	vTimeout, okTimeout := d.GetOk("timeout")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ProvisionNfv")
		request1 := expandRequestNfvProvisionProvisionNfv(ctx, "", d)

		headerParams1 := dnacentersdkgo.ProvisionNfvHeaderParams{}

		if okRunsync {
			headerParams1.Runsync = vRunsync.(string)
		}

		if okTimeout {
			headerParams1.Timeout = vTimeout.(string)
		}

		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.SiteDesign.ProvisionNfv(request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ProvisionNfv", err,
				"Failure at ProvisionNfv, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignProvisionNfvItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ProvisionNfv response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNfvProvisionProvisionNfv(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfv {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfv{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_profile")))) {
		request.SiteProfile = expandRequestNfvProvisionProvisionNfvSiteProfileArray(ctx, key+".site_profile", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provisioning")))) {
		request.Provisioning = expandRequestNfvProvisionProvisionNfvProvisioningArray(ctx, key+".provisioning", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfile {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfile{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfile(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfile {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_profile_name")))) {
		request.SiteProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceArray(ctx, key+".device", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDevice {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDevice{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDevice {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag_name")))) {
		request.TagName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_providers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_providers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_providers")))) {
		request.ServiceProviders = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServiceProvidersArray(ctx, key+".service_providers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dia")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dia")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dia")))) {
		request.Dia = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServicesArray(ctx, key+".services", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_services")))) {
		request.CustomServices = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServicesArray(ctx, key+".custom_services", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_networks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_networks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_networks")))) {
		request.CustomNetworks = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksArray(ctx, key+".custom_networks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan")))) {
		request.VLAN = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceVLANArray(ctx, key+".vlan", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_template")))) {
		request.CustomTemplate = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomTemplateArray(ctx, key+".custom_template", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServiceProvidersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServiceProviders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServiceProviders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_provider")))) {
		request.ServiceProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link_type")))) {
		request.LinkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connect")))) {
		request.Connect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_gateway")))) {
		request.DefaultGateway = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServices {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServices {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode")))) {
		request.Mode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_name")))) {
		request.ImageName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".topology")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".topology")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".topology")))) {
		request.Topology = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServicesTopology(ctx, key+".topology.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceServicesTopology(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServicesTopology {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceServicesTopology{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_ip")))) {
		request.AssignIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".topology")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".topology")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".topology")))) {
		request.Topology = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServicesTopology(ctx, key+".topology.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_name")))) {
		request.ImageName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomServicesTopology(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServicesTopology {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServicesTopology{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_ip")))) {
		request.AssignIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services_to_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services_to_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services_to_connect")))) {
		request.ServicesToConnect = expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnectArray(ctx, key+".services_to_connect", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_type")))) {
		request.ConnectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_mode")))) {
		request.NetworkMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan")))) {
		request.VLAN = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnectArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service")))) {
		request.Service = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceVLANArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceVLAN(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceVLAN(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomTemplateArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomTemplate(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvSiteProfileDeviceCustomTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template")))) {
		request.Template = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioning {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioning{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioning(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioning(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioning {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioning{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestNfvProvisionProvisionNfvProvisioningSite(ctx, key+".site.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestNfvProvisionProvisionNfvProvisioningDeviceArray(ctx, key+".device", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSite {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_profile_name")))) {
		request.SiteProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".area")))) {
		request.Area = expandRequestNfvProvisionProvisionNfvProvisioningSiteArea(ctx, key+".area.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".building")))) {
		request.Building = expandRequestNfvProvisionProvisionNfvProvisioningSiteBuilding(ctx, key+".building.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor")))) {
		request.Floor = expandRequestNfvProvisionProvisionNfvProvisioningSiteFloor(ctx, key+".floor.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningSiteArea(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteArea {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteArea{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningSiteBuilding(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteBuilding {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteBuilding{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningSiteFloor(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteFloor {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningSiteFloor{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_name")))) {
		request.ParentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_model")))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".width")))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".length")))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".height")))) {
		request.Height = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDevice {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDevice{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDevice {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip")))) {
		request.IP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_serial_number")))) {
		request.DeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag_name")))) {
		request.TagName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_providers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_providers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_providers")))) {
		request.ServiceProviders = expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProvidersArray(ctx, key+".service_providers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = expandRequestNfvProvisionProvisionNfvProvisioningDeviceServicesArray(ctx, key+".services", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan")))) {
		request.VLAN = expandRequestNfvProvisionProvisionNfvProvisioningDeviceVLANArray(ctx, key+".vlan", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sub_pools")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sub_pools")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sub_pools")))) {
		request.SubPools = expandRequestNfvProvisionProvisionNfvProvisioningDeviceSubPoolsArray(ctx, key+".sub_pools", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_networks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_networks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_networks")))) {
		request.CustomNetworks = expandRequestNfvProvisionProvisionNfvProvisioningDeviceCustomNetworksArray(ctx, key+".custom_networks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_param")))) {
		request.TemplateParam = expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParam(ctx, key+".template_param.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProvidersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProviders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProviders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_provider")))) {
		request.ServiceProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_interface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_interface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_interface")))) {
		request.WanInterface = expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProvidersWanInterface(ctx, key+".wan_interface.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceServiceProvidersWanInterface(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProvidersWanInterface {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServiceProvidersWanInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnetmask")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnetmask")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnetmask")))) {
		request.Subnetmask = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth")))) {
		request.Bandwidth = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceServicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServices {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDeviceServices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceServices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServices {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceServices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode")))) {
		request.Mode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_ip")))) {
		request.SystemIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".central_manager_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".central_manager_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".central_manager_ip")))) {
		request.CentralManagerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".central_registration_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".central_registration_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".central_registration_key")))) {
		request.CentralRegistrationKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".common_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".common_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".common_key")))) {
		request.CommonKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password_hash")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password_hash")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password_hash")))) {
		request.AdminPasswordHash = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".disk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".disk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".disk")))) {
		request.Disk = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceVLANArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceVLAN {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceVLAN{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDeviceVLAN(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceVLAN(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceVLAN {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceVLAN{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network")))) {
		request.Network = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceSubPoolsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceSubPools {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceSubPools{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDeviceSubPools(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceSubPools(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceSubPools {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceSubPools{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_subnet")))) {
		request.IPSubnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_pool_name")))) {
		request.ParentPoolName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceCustomNetworksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks {
	request := []dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNfvProvisionProvisionNfvProvisioningDeviceCustomNetworks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceCustomNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_pool")))) {
		request.IPAddressPool = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParam(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParam {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParam{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nfvis")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nfvis")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nfvis")))) {
		request.Nfvis = expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParamNfvis(ctx, key+".nfvis.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".asav")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".asav")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".asav")))) {
		request.Asav = expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParamAsav(ctx, key+".asav.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParamNfvis(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamNfvis {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamNfvis{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".var1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".var1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".var1")))) {
		request.Var1 = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProvisionProvisionNfvProvisioningDeviceTemplateParamAsav(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamAsav {
	request := dnacentersdkgo.RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamAsav{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".var1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".var1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".var1")))) {
		request.Var1 = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSiteDesignProvisionNfvItem(item *dnacentersdkgo.ResponseSiteDesignProvisionNfv) []map[string]interface{} {
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
