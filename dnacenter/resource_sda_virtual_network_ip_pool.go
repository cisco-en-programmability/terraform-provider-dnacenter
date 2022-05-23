package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaVirtualNetworkIPPool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete IP Pool from SDA Virtual Network

- Add IP Pool in SDA Virtual Network
`,

		CreateContext: resourceSdaVirtualNetworkIPPoolCreate,
		ReadContext:   resourceSdaVirtualNetworkIPPoolRead,
		UpdateContext: resourceSdaVirtualNetworkIPPoolUpdate,
		DeleteContext: resourceSdaVirtualNetworkIPPoolDelete,
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

						"authentication_policy_name": &schema.Schema{
							Description: `Authentication Policy Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_l2_flooding_enabled": &schema.Schema{
							Description: `Is L2 Flooding Enabled`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"is_this_critical_pool": &schema.Schema{
							Description: `Is This Critical Pool`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"traffic_type": &schema.Schema{
							Description: `Traffic Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name`,
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

						"auto_generate_vlan_name": &schema.Schema{
							Description: `It will auto generate vlanName, if vlanName is empty(applicable for L3  and INFRA_VN)
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name, that is reserved to Fabric Site for (applicable for L3 and INFRA_VN)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_common_pool": &schema.Schema{
							Description: `Common Pool enablement flag(applicable for L3 and L2 and default value is False )
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_ip_directed_broadcast": &schema.Schema{
							Description: `Ip Directed Broadcast enablement flag(applicable for L3 and default value is False )
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_l2_flooding_enabled": &schema.Schema{
							Description: `Layer2 flooding enablement flag(applicable for L3 , L2 and always true for L2 and default value is False )
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_layer2_only": &schema.Schema{
							Description: `Layer2 Only enablement flag and default value is False 
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_this_critical_pool": &schema.Schema{
							Description: `Critical pool enablement(applicable for L3 and default value is False )
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_wireless_pool": &schema.Schema{
							Description: `Wireless Pool enablement flag(applicable for L3  and L2 and default value is False )
`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"pool_type": &schema.Schema{
							Description: `Pool Type (applicable for  INFRA_VN)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name(applicable for L3)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"traffic_type": &schema.Schema{
							Description: `Traffic type(applicable for L3  and L2)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name, that is associated to Fabric Site
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Description: `vlan Id(applicable for L3 , L2 and  INFRA_VN)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_name": &schema.Schema{
							Description: `Vlan name represent the segment name, if empty, vlanName would be auto generated by API
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

func resourceSdaVirtualNetworkIPPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkIPPoolAddIPPoolInSdaVirtualNetwork(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vIPPoolName := resourceItem["ip_pool_name"]
	vVirtualNetworkName := resourceItem["virtual_network_name"]
	vvIPPoolName := interfaceToString(vIPPoolName)
	vvVirtualNetworkName := interfaceToString(vVirtualNetworkName)
	queryParams1 := dnacentersdkgo.GetIPPoolFromSdaVirtualNetworkQueryParams{}

	queryParams1.IPPoolName = vvIPPoolName

	queryParams1.VirtualNetworkName = vvVirtualNetworkName

	getResponse2, _, err := client.Sda.GetIPPoolFromSdaVirtualNetwork(&queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["ip_pool_name"] = vvIPPoolName
		resourceMap["virtual_network_name"] = vvVirtualNetworkName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaVirtualNetworkIPPoolRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Sda.AddIPPoolInSdaVirtualNetwork(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddIPPoolInSdaVirtualNetwork", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddIPPoolInSdaVirtualNetwork", err))
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
				"Failure when executing AddIPPoolInSdaVirtualNetwork", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["ip_pool_name"] = vvIPPoolName
	resourceMap["virtual_network_name"] = vvVirtualNetworkName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaVirtualNetworkIPPoolRead(ctx, d, m)
}

func resourceSdaVirtualNetworkIPPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIPPoolName := resourceMap["ip_pool_name"]
	vVirtualNetworkName := resourceMap["virtual_network_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetIPPoolFromSdaVirtualNetwork")
		queryParams1 := dnacentersdkgo.GetIPPoolFromSdaVirtualNetworkQueryParams{}

		queryParams1.IPPoolName = vIPPoolName

		queryParams1.VirtualNetworkName = vVirtualNetworkName

		response1, restyResp1, err := client.Sda.GetIPPoolFromSdaVirtualNetwork(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetIPPoolFromSdaVirtualNetworkItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPPoolFromSdaVirtualNetwork response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaVirtualNetworkIPPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaVirtualNetworkIPPoolRead(ctx, d, m)
}

func resourceSdaVirtualNetworkIPPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIPPoolName := resourceMap["ip_pool_name"]
	vVirtualNetworkName := resourceMap["virtual_network_name"]

	queryParams1 := dnacentersdkgo.GetIPPoolFromSdaVirtualNetworkQueryParams{}
	queryParams1.IPPoolName = vIPPoolName
	queryParams1.VirtualNetworkName = vVirtualNetworkName
	item, restyResp1, err := client.Sda.GetIPPoolFromSdaVirtualNetwork(&queryParams1)
	if err != nil || item == nil {
		d.SetId("")
		return diags
	}

	// REVIEW: Add getAllItems and search function to get missing params
	queryParams2 := dnacentersdkgo.DeleteIPPoolFromSdaVirtualNetworkQueryParams{}
	queryParams2.IPPoolName = vIPPoolName
	queryParams2.VirtualNetworkName = vVirtualNetworkName
	response1, restyResp1, err := client.Sda.DeleteIPPoolFromSdaVirtualNetwork(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteIPPoolFromSdaVirtualNetwork", err, restyResp1.String(),
				"Failure at DeleteIPPoolFromSdaVirtualNetwork, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteIPPoolFromSdaVirtualNetwork", err,
			"Failure at DeleteIPPoolFromSdaVirtualNetwork, unexpected response", ""))
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
				"Failure when executing DeleteIPPoolFromSdaVirtualNetwork", err))
			return diags
		}
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaVirtualNetworkIPPoolAddIPPoolInSdaVirtualNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddIPPoolInSdaVirtualNetwork {
	request := dnacentersdkgo.RequestSdaAddIPPoolInSdaVirtualNetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_layer2_only")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_layer2_only")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_layer2_only")))) {
		request.IsLayer2Only = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_generate_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_generate_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_generate_vlan_name")))) {
		request.AutoGenerateVLANName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_l2_flooding_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_l2_flooding_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_l2_flooding_enabled")))) {
		request.IsL2FloodingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_this_critical_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_this_critical_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_this_critical_pool")))) {
		request.IsThisCriticalPool = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_wireless_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_wireless_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_wireless_pool")))) {
		request.IsWirelessPool = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_ip_directed_broadcast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_ip_directed_broadcast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_ip_directed_broadcast")))) {
		request.IsIPDirectedBroadcast = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_pool")))) {
		request.IsCommonPool = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pool_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pool_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pool_type")))) {
		request.PoolType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
