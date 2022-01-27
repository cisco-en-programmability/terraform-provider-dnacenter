package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalPool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to update global pool

- API to create global pool.

- API to delete global IP pool.
`,

		CreateContext: resourceGlobalPoolCreate,
		ReadContext:   resourceGlobalPoolRead,
		UpdateContext: resourceGlobalPoolUpdate,
		DeleteContext: resourceGlobalPoolDelete,
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

						"client_options": &schema.Schema{
							Description: `Client Options`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"configure_external_dhcp": &schema.Schema{
							Description: `Configure External Dhcp`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"context_key": &schema.Schema{
										Description: `Context Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"context_value": &schema.Schema{
										Description: `Context Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"owner": &schema.Schema{
										Description: `Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"create_time": &schema.Schema{
							Description: `Create Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"dhcp_server_ips": &schema.Schema{
							Description: `Dhcp Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"dns_server_ips": &schema.Schema{
							Description: `Dns Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"gateways": &schema.Schema{
							Description: `Gateways`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_cidr": &schema.Schema{
							Description: `Ip Pool Cidr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"overlapping": &schema.Schema{
							Description: `Overlapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"owner": &schema.Schema{
							Description: `Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_uuid": &schema.Schema{
							Description: `Parent Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"shared": &schema.Schema{
							Description: `Shared`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"total_ip_address_count": &schema.Schema{
							Description: `Total Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"used_ip_address_count": &schema.Schema{
							Description: `Used Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"used_percentage": &schema.Schema{
							Description: `Used Percentage`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"id": &schema.Schema{
							Description: `id path parameter. global pool id
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ippool": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address_space": &schema.Schema{
													Description: `Ip Address Space. Allowed values are IPv6 or IPv4.`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dhcp_server_ips": &schema.Schema{
													Description: `Dhcp Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"dns_server_ips": &schema.Schema{
													Description: `Dns Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"gateway": &schema.Schema{
													Description: `Gateway`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ip_pool_cidr": &schema.Schema{
													Description: `Ip Pool Cidr`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ip_pool_name": &schema.Schema{
													Description: `Ip Pool Name`,
													Type:        schema.TypeString,
													Required:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Optional:    true,
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

func resourceGlobalPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolCreate")
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestGlobalPoolCreateGlobalPool(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resourceItem := *getResourceItem(d.Get("parameters"))
	vvID := interfaceToString(resourceItem["id"])
	vvIpPoolName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.settings"); ok {
			if _, ok := d.GetOk("parameters.0.settings.0"); ok {
				if _, ok := d.GetOk("parameters.0.settings.0.ippool"); ok {
					if _, ok := d.GetOk("parameters.0.settings.0.ippool.0"); ok {
						if v, ok := d.GetOk("parameters.0.settings.0.ippool.0.ip_pool_name"); ok {
							vvIpPoolName = interfaceToString(v)
						}
					}
				}
			}
		}
	}
	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

	response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vvID, vvIpPoolName)
	if response1 != nil {
		log.Printf("[DEBUG] searchNetworkSettingsGetGlobalPool result %v", responseInterfaceToString(*response1))
	}
	if err == nil && (response1 != nil && len(*response1) > 0) {
		resourceMap := make(map[string]string)
		resourceMap["ip_pool_name"] = vvIpPoolName
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalPoolRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.NetworkSettings.CreateGlobalPool(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGlobalPool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGlobalPool", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateGlobalPool", err,
				"Failure at CreateGlobalPool execution", bapiError))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["ip_pool_name"] = vvIpPoolName
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolRead")

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalPool")
		queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

		response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
		if err != nil || response1 == nil || len(*response1) <= 0 {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetGlobalPool", err,
			// 	"Failure at GetGlobalPool, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetGlobalPoolItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalPool search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolUpdate")
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
	if err != nil || item == nil || len(*item) <= 0 {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalPool", err,
			"Failure at GetGlobalPool, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		request1 := expandRequestGlobalPoolUpdateGlobalPool(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if item != nil && len(*item) > 0 {
			if request1 != nil && request1.Settings != nil && request1.Settings.IPpool != nil && len(*request1.Settings.IPpool) > 0 {
				found := *item
				req := *request1.Settings.IPpool
				req[0].ID = found[0].ID
			}
		}
		response1, restyResp1, err := client.NetworkSettings.UpdateGlobalPool(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGlobalPool", err, restyResp1.String(),
					"Failure at UpdateGlobalPool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalPool", err,
				"Failure at UpdateGlobalPool, unexpected response", ""))
			return diags
		}
		executionId := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionId)
		if executionId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetBusinessAPIExecutionDetails", err,
					"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
				return diags
			}
			for response2.Status == "IN_PROGRESS" {
				time.Sleep(10 * time.Second)
				response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response2 == nil {
					if restyResp1 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if response2.Status == "FAILURE" {
				bapiError := response2.BapiError
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing UpdateGlobalPool", err,
					"Failure at UpdateGlobalPool execution", bapiError))
				return diags
			}
		}
	}

	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolDelete")

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalPool", err,
			"Failure at GetGlobalPool, unexpected response", ""))
		return diags
	}
	if item == nil || len(*item) == 0 {
		return diags
	}
	if vID == "" && item != nil && len(*item) > 0 {
		vID = (*item)[0].ID
	}

	response1, restyResp1, err := client.NetworkSettings.DeleteGlobalIPPool(vID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGlobalIPPool", err, restyResp1.String(),
				"Failure at DeleteGlobalIPPool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGlobalIPPool", err,
			"Failure at DeleteGlobalIPPool, unexpected response", ""))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalPool", err,
				"Failure at UpdateGlobalPool execution", bapiError))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGlobalPoolCreateGlobalPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPool {
	log.Println("[DEBUG] call expandRequestGlobalPoolCreateGlobalPool")
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".settings")))) {
		request.Settings = expandRequestGlobalPoolCreateGlobalPoolSettings(ctx, key+".settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettings {
	log.Println("[DEBUG] call expandRequestGlobalPoolCreateGlobalPoolSettings")
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolCreateGlobalPoolSettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool {
	log.Println("[DEBUG] call expandRequestGlobalPoolCreateGlobalPoolSettingsIPpoolArray")
	request := []dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool{}
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
		i := expandRequestGlobalPoolCreateGlobalPoolSettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool {
	log.Println("[DEBUG] call expandRequestGlobalPoolCreateGlobalPoolSettingsIPpool")
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_cidr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_cidr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_cidr")))) {
		request.IPPoolCidr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_space")))) {
		request.IPAddressSpace = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPool {
	log.Println("[DEBUG] call expandRequestGlobalPoolUpdateGlobalPool")
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".settings")))) {
		request.Settings = expandRequestGlobalPoolUpdateGlobalPoolSettings(ctx, key+".settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettings {
	log.Println("[DEBUG] call expandRequestGlobalPoolUpdateGlobalPoolSettings")
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool {
	log.Println("[DEBUG] call expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpoolArray")
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool{}
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
		i := expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool {
	log.Println("[DEBUG] call expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpool")
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchNetworkSettingsGetGlobalPool(m interface{}, queryParams dnacentersdkgo.GetGlobalPoolQueryParams, vID string, vIPPoolName string) (*[]dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItems []dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse
	offset := 1
	queryParams.Offset = strconv.Itoa(offset)

	nResponse, _, err := client.NetworkSettings.GetGlobalPool(&queryParams)
	if err != nil {
		log.Printf("[DEBUG] GetGlobalPool error %s", err.Error())
		return nil, err
	}
	if nResponse == nil {
		return nil, err
	}
	if nResponse.Response == nil {
		return nil, err
	}
	maxPageSize := len(*nResponse.Response)
	//maxPageSize := 10
	for nResponse != nil && nResponse.Response != nil && len(*nResponse.Response) > 0 {
		for _, item := range *nResponse.Response {
			if vIPPoolName == item.IPPoolName {
				foundItems = append(foundItems, item)
				return &foundItems, err
			}
			if vID == item.ID {
				foundItems = append(foundItems, item)
				return &foundItems, err
			}
		}

		queryParams.Limit = strconv.Itoa(maxPageSize)
		offset += maxPageSize
		queryParams.Offset = strconv.Itoa(offset)
		time.Sleep(15 * time.Second)
		nResponse, _, err = client.NetworkSettings.GetGlobalPool(&queryParams)
	}
	return &foundItems, err
}
