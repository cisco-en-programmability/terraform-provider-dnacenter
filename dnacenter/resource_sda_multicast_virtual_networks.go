package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaMulticastVirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds multicast for virtual networks based on user input.

- Updates multicast configurations for virtual networks based on user input.

- Deletes a multicast configuration for a virtual network based on id.
`,

		CreateContext: resourceSdaMulticastVirtualNetworksCreate,
		ReadContext:   resourceSdaMulticastVirtualNetworksRead,
		UpdateContext: resourceSdaMulticastVirtualNetworksUpdate,
		DeleteContext: resourceSdaMulticastVirtualNetworksDelete,
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

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the multicast configuration.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_pool_name": &schema.Schema{
							Description: `Name of the IP Pool.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4_ssm_ranges": &schema.Schema{
							Description: `IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"multicast_r_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipv4_address": &schema.Schema{
										Description: `IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv4_asm_ranges": &schema.Schema{
										Description: `IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ipv6_address": &schema.Schema{
										Description: `IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv6_asm_ranges": &schema.Schema{
										Description: `IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_default_v4_rp": &schema.Schema{
										Description: `Specifies whether it is a default IPv4 RP.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default_v6_rp": &schema.Schema{
										Description: `Specifies whether it is a default IPv6 RP.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_device_ids": &schema.Schema{
										Description: `IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"rp_device_location": &schema.Schema{
										Description: `Device location of the RP.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"virtual_network_name": &schema.Schema{
							Description: `Name of the virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddMulticastVirtualNetworks`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fabric_id": &schema.Schema{
										Description: `ID of the fabric site this multicast configuration is associated with.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the multicast configuration (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP Pool associated with the fabric site.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv4_ssm_ranges": &schema.Schema{
										Description: `IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"multicast_r_ps": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4_address": &schema.Schema{
													Description: `IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ipv4_asm_ranges": &schema.Schema{
													Description: `IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ipv6_address": &schema.Schema{
													Description: `IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ipv6_asm_ranges": &schema.Schema{
													Description: `IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_default_v4_rp": &schema.Schema{
													Description: `Specifies whether it is a default IPv4 RP.
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"is_default_v6_rp": &schema.Schema{
													Description: `Specifies whether it is a default IPv6 RP.
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"network_device_ids": &schema.Schema{
													Description: `IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"rp_device_location": &schema.Schema{
													Description: `Device location of the RP.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"virtual_network_name": &schema.Schema{
										Description: `Name of the virtual network associated with the fabric site.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

func resourceSdaMulticastVirtualNetworksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworks(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	vName := resourceItem["virtual_network_name"]
	vvName := interfaceToString(vName)

	queryParamImport := dnacentersdkgo.GetMulticastVirtualNetworksQueryParams{}
	queryParamImport.FabricID = vvFabricID
	queryParamImport.VirtualNetworkName = vvName
	item2, err := searchSdaGetMulticastVirtualNetworks(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["fabric_id"] = item2.FabricID
		resourceMap["virtual_network_name"] = item2.VirtualNetworkName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaMulticastVirtualNetworksRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddMulticastVirtualNetworks(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddMulticastVirtualNetworks", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddMulticastVirtualNetworks", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddMulticastVirtualNetworks", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AddMulticastVirtualNetworks", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetMulticastVirtualNetworksQueryParams{}
	queryParamValidate.FabricID = vvFabricID
	queryParamValidate.VirtualNetworkName = vvName
	item3, err := searchSdaGetMulticastVirtualNetworks(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddMulticastVirtualNetworks", err,
			"Failure at AddMulticastVirtualNetworks, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["fabric_id"] = item3.FabricID
	resourceMap["virtual_network_name"] = item3.VirtualNetworkName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaMulticastVirtualNetworksRead(ctx, d, m)
}

func resourceSdaMulticastVirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	vvFabricID := resourceMap["fabric_id"]
	vvName := resourceMap["virtual_network_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMulticastVirtualNetworks")
		queryParams1 := dnacentersdkgo.GetMulticastVirtualNetworksQueryParams{}
		queryParams1.FabricID = vvFabricID
		queryParams1.VirtualNetworkName = vvName
		item1, err := searchSdaGetMulticastVirtualNetworks(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworksResponse{
			*item1,
		}
		vItem1 := flattenSdaGetMulticastVirtualNetworksItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMulticastVirtualNetworks search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaMulticastVirtualNetworksUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworks(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateMulticastVirtualNetworks(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateMulticastVirtualNetworks", err, restyResp1.String(),
					"Failure at UpdateMulticastVirtualNetworks, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateMulticastVirtualNetworks", err,
				"Failure at UpdateMulticastVirtualNetworks, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateMulticastVirtualNetworks", err))
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateMulticastVirtualNetworks", err1))
				return diags
			}
		}

	}

	return resourceSdaMulticastVirtualNetworksRead(ctx, d, m)
}

func resourceSdaMulticastVirtualNetworksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteMulticastVirtualNetworkByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteMulticastVirtualNetworkByID", err, restyResp1.String(),
				"Failure at DeleteMulticastVirtualNetworkByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMulticastVirtualNetworkByID", err,
			"Failure at DeleteMulticastVirtualNetworkByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteMulticastVirtualNetworkByID", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteMulticastVirtualNetworkByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastVirtualNetworks {
	request := dnacentersdkgo.RequestSdaAddMulticastVirtualNetworks{}
	if v := expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworks{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_ssm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_ssm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_ssm_ranges")))) {
		request.IPv4SsmRanges = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_r_ps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_r_ps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_r_ps")))) {
		request.MulticastRPs = expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemMulticastRPsArray(ctx, key+".multicast_r_ps", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemMulticastRPsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworksMulticastRPs {
	request := []dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworksMulticastRPs{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemMulticastRPs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksAddMulticastVirtualNetworksItemMulticastRPs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworksMulticastRPs {
	request := dnacentersdkgo.RequestItemSdaAddMulticastVirtualNetworksMulticastRPs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rp_device_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rp_device_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rp_device_location")))) {
		request.RpDeviceLocation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_v4_rp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_v4_rp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_v4_rp")))) {
		request.IsDefaultV4RP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_v6_rp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_v6_rp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_v6_rp")))) {
		request.IsDefaultV6RP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_ids")))) {
		request.NetworkDeviceIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_asm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_asm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_asm_ranges")))) {
		request.IPv4AsmRanges = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_asm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_asm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_asm_ranges")))) {
		request.IPv6AsmRanges = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateMulticastVirtualNetworks {
	request := dnacentersdkgo.RequestSdaUpdateMulticastVirtualNetworks{}
	if v := expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworks{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_ssm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_ssm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_ssm_ranges")))) {
		request.IPv4SsmRanges = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_r_ps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_r_ps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_r_ps")))) {
		request.MulticastRPs = expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemMulticastRPsArray(ctx, key+".multicast_r_ps", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemMulticastRPsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs {
	request := []dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemMulticastRPs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastVirtualNetworksUpdateMulticastVirtualNetworksItemMulticastRPs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs {
	request := dnacentersdkgo.RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rp_device_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rp_device_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rp_device_location")))) {
		request.RpDeviceLocation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_v4_rp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_v4_rp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_v4_rp")))) {
		request.IsDefaultV4RP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_v6_rp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_v6_rp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_v6_rp")))) {
		request.IsDefaultV6RP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_ids")))) {
		request.NetworkDeviceIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_asm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_asm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_asm_ranges")))) {
		request.IPv4AsmRanges = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_asm_ranges")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_asm_ranges")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_asm_ranges")))) {
		request.IPv6AsmRanges = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetMulticastVirtualNetworks(m interface{}, queryParams dnacentersdkgo.GetMulticastVirtualNetworksQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworksResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworksResponse
	var ite *dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworks
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetMulticastVirtualNetworks(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetMulticastVirtualNetworks(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.VirtualNetworkName != "" {
		ite, _, err = client.Sda.GetMulticastVirtualNetworks(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.VirtualNetworkName == queryParams.VirtualNetworkName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
