package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaMulticast() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Add multicast in SDA fabric

- Delete multicast from SDA fabric
`,

		CreateContext: resourceSdaMulticastCreate,
		ReadContext:   resourceSdaMulticastRead,
		UpdateContext: resourceSdaMulticastUpdate,
		DeleteContext: resourceSdaMulticastDelete,
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

						"description": &schema.Schema{
							Description: `multicast configuration info retrieved successfully from sda fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"multicast_method": &schema.Schema{
							Description: `Multicast Method
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"multicast_type": &schema.Schema{
							Description: `Multicast Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"multicast_vn_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_rp_ip_address": &schema.Schema{
										Description: `ExternalRpIpAddress
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"internal_rp_ip_address": &schema.Schema{
										Description: `InternalRpIpAddress
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to Fabric Site
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssm_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ssm_group_range": &schema.Schema{
													Description: `SSM group range
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"ssm_wildcard_mask": &schema.Schema{
													Description: `SSM Wildcard Mask 
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to Fabric Site
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"multicast_method": &schema.Schema{
							Description: `Multicast Method
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"multicast_type": &schema.Schema{
							Description: `Multicast Type
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"multicast_vn_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_rp_ip_address": &schema.Schema{
										Description: `ExternalRpIpAddress, required if multicastType is asm_with_external_rp
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"internal_rp_ip_address": &schema.Schema{
										Description: `InternalRpIpAddress, required if multicastType is asm_with_internal_rp
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to Fabric Site
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssm_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ssm_group_range": &schema.Schema{
													Description: `Valid SSM group range ip address(e.g., 230.0.0.0)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ssm_wildcard_mask": &schema.Schema{
													Description: `Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to Fabric Site
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Full path of sda Fabric Site
`,
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

func resourceSdaMulticastCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaMulticastAddMulticastInSdaFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	queryParamImport := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}
	queryParamImport.SiteNameHierarchy = vvSiteNameHierarchy
	item2, _, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParamImport)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaMulticastRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddMulticastInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddMulticastInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddMulticastInSdaFabric", err))
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
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddMulticastInSdaFabric", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}
	queryParamValidate.SiteNameHierarchy = vvSiteNameHierarchy
	item3, _, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddMulticastInSdaFabric", err,
			"Failure at AddMulticastInSdaFabric, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaMulticastRead(ctx, d, m)
}

func resourceSdaMulticastRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMulticastDetailsFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		response1, restyResp1, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetMulticastDetailsFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMulticastDetailsFromSdaFabric response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSdaMulticastUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaMulticastRead(ctx, d, m)
}

func resourceSdaMulticastDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := dnacentersdkgo.DeleteMulticastFromSdaFabricQueryParams{}

	vvSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	queryParamDelete.SiteNameHierarchy = vvSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteMulticastFromSdaFabric(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteMulticastFromSdaFabric", err, restyResp1.String(),
				"Failure at DeleteMulticastFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMulticastFromSdaFabric", err,
			"Failure at DeleteMulticastFromSdaFabric, unexpected response", ""))
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
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteMulticastFromSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaMulticastAddMulticastInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddMulticastInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_method")))) {
		request.MulticastMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_type")))) {
		request.MulticastType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_vn_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_vn_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_vn_info")))) {
		request.MulticastVnInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoArray(ctx, key+".multicast_vn_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo {
	request := []dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo{}
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
		i := expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo {
	request := dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_rp_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_rp_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_rp_ip_address")))) {
		request.InternalRpIPAddress = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_rp_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_rp_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_rp_ip_address")))) {
		request.ExternalRpIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_info")))) {
		request.SsmInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfoArray(ctx, key+".ssm_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo {
	request := []dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo{}
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
		i := expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo {
	request := dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_group_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_group_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_group_range")))) {
		request.SsmGroupRange = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_wildcard_mask")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_wildcard_mask")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_wildcard_mask")))) {
		request.SsmWildcardMask = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
