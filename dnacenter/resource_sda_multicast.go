package dnacenter

import (
	"context"
	"log"
	"reflect"
	"time"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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

						"multicast_method": &schema.Schema{
							Description: `Multicast Methods
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
										Description: `External Rp Ip Address, required for muticastType=asm_with_external_rp
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_group_range": &schema.Schema{
										Description: `Valid SSM group range ip address(e.g., 230.0.0.0)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_info": &schema.Schema{
										Description: `Source-specific multicast information, required if muticastType=ssm
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_wildcard_mask": &schema.Schema{
										Description: `Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"muticast_type": &schema.Schema{
							Description: `Muticast Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Full path of sda fabric siteNameHierarchy
`,
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

						"multicast_method": &schema.Schema{
							Description: `Multicast Methods
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"multicast_vn_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_rp_ip_address": &schema.Schema{
										Description: `External Rp Ip Address, required for muticastType=asm_with_external_rp
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssm_group_range": &schema.Schema{
										Description: `Valid SSM group range ip address(e.g., 230.0.0.0)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssm_info": &schema.Schema{
										Description: `Source-specific multicast information, required if muticastType=ssm
`,
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ssm_wildcard_mask": &schema.Schema{
										Description: `Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"muticast_type": &schema.Schema{
							Description: `Muticast Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Full path of sda fabric siteNameHierarchy
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

func resourceSdaMulticastCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaMulticastAddMulticastInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)

	queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}

	queryParams1.SiteNameHierarchy = vvSiteNameHierarchy

	getResponse2, _, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)

	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaMulticastRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Sda.AddMulticastInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddMulticastInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddMulticastInSdaFabric", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
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
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddMulticastInSdaFabric", err))
			return diags
		}
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
		log.Printf("[DEBUG] Selected method 1: GetMulticastDetailsFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		response1, restyResp1, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetMulticastDetailsFromSdaFabric", err,
			// 	"Failure at GetMulticastDetailsFromSdaFabric, unexpected response", ""))
			// return diags
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
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	item, restyResp1, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)
	if err != nil || item == nil {
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetMulticastDetailsFromSDAFabric", err,
		"Failure at GetMulticastDetailsFromSDAFabric, unexpected response", ""))*/
		d.SetId("")
		return diags
	}
	queryParams2 := dnacentersdkgo.DeleteMulticastFromSdaFabricQueryParams{}
	queryParams2.SiteNameHierarchy = vSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteMulticastFromSdaFabric(&queryParams2)
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
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".muticast_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".muticast_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".muticast_type")))) {
		request.MuticastType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_vn_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_vn_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_vn_info")))) {
		request.MulticastVnInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfo(ctx, key+".multicast_vn_info.0", d)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_rp_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_rp_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_rp_ip_address")))) {
		request.ExternalRpIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_info")))) {
		request.SsmInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx, key+".ssm_info.0", d)
	}
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

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo {
	var request dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
