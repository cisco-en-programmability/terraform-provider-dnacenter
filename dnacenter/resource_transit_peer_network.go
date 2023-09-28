package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTransitPeerNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete Transit Peer Network from SD-Access

- Add Transit Peer Network in SD-Access
`,

		CreateContext: resourceTransitPeerNetworkCreate,
		ReadContext:   resourceTransitPeerNetworkRead,
		UpdateContext: resourceTransitPeerNetworkUpdate,
		DeleteContext: resourceTransitPeerNetworkDelete,
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
							Description: `Transit Peer network info retrieved successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autonomous_system_number": &schema.Schema{
										Description: `Autonomous System Number  
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"routing_protocol_name": &schema.Schema{
										Description: `Routing Protocol Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sda_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"transit_control_plane_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_management_ip_address": &schema.Schema{
													Description: `Device Management Ip Address 
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"site_name_hierarchy": &schema.Schema{
													Description: `Site Name Hierarchy 
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
						"status": &schema.Schema{
							Description: `status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"transit_peer_network_name": &schema.Schema{
							Description: `Transit Peer Network Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"transit_peer_network_type": &schema.Schema{
							Description: `Transit Peer Network Type
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

						"ip_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autonomous_system_number": &schema.Schema{
										Description: `Autonomous System Number  (e.g.,1-65535)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"routing_protocol_name": &schema.Schema{
										Description: `Routing Protocol Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"sda_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"transit_control_plane_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_management_ip_address": &schema.Schema{
													Description: `Device Management Ip Address of provisioned device
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"site_name_hierarchy": &schema.Schema{
													Description: `Site Name Hierarchy where device is provisioned
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
						"transit_peer_network_name": &schema.Schema{
							Description: `Transit Peer Network Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transit_peer_network_type": &schema.Schema{
							Description: `Transit Peer Network Type
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

func resourceTransitPeerNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTransitPeerNetworkAddTransitPeerNetwork(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTransitPeerNetworkName := resourceItem["transit_peer_network_name"]
	vvTransitPeerNetworkName := interfaceToString(vTransitPeerNetworkName)
	queryParamImport := dnacentersdkgo.GetTransitPeerNetworkInfoQueryParams{}
	queryParamImport.TransitPeerNetworkName = vvTransitPeerNetworkName
	item2, _, err := client.Sda.GetTransitPeerNetworkInfo(&queryParamImport)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["transit_peer_network_name"] = item2.TransitPeerNetworkName
		d.SetId(joinResourceID(resourceMap))
		return resourceTransitPeerNetworkRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddTransitPeerNetwork(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddTransitPeerNetwork", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddTransitPeerNetwork", err))
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
				"Failure when executing AddTransitPeerNetwork", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetTransitPeerNetworkInfoQueryParams{}
	queryParamValidate.TransitPeerNetworkName = vvTransitPeerNetworkName
	item3, _, err := client.Sda.GetTransitPeerNetworkInfo(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddTransitPeerNetwork", err,
			"Failure at AddTransitPeerNetwork, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["transit_peer_network_name"] = item3.TransitPeerNetworkName

	d.SetId(joinResourceID(resourceMap))
	return resourceTransitPeerNetworkRead(ctx, d, m)
}

func resourceTransitPeerNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTransitPeerNetworkName := resourceMap["transit_peer_network_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransitPeerNetworkInfo")
		queryParams1 := dnacentersdkgo.GetTransitPeerNetworkInfoQueryParams{}

		queryParams1.TransitPeerNetworkName = vTransitPeerNetworkName

		response1, restyResp1, err := client.Sda.GetTransitPeerNetworkInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetTransitPeerNetworkInfoItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransitPeerNetworkInfo response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceTransitPeerNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing TransitPeerNetworkUpdate", err, "Update method is not supported",
		"Failure at TransitPeerNetworkUpdate, unexpected response", ""))

	return diags
}

func resourceTransitPeerNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := dnacentersdkgo.DeleteTransitPeerNetworkQueryParams{}

	vvTransitPeerNetworkName := resourceMap["transit_peer_network_name"]
	queryParamDelete.TransitPeerNetworkName = vvTransitPeerNetworkName

	response1, restyResp1, err := client.Sda.DeleteTransitPeerNetwork(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTransitPeerNetwork", err, restyResp1.String(),
				"Failure at DeleteTransitPeerNetwork, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTransitPeerNetwork", err,
			"Failure at DeleteTransitPeerNetwork, unexpected response", ""))
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
				"Failure when executing DeleteTransitPeerNetwork", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTransitPeerNetworkAddTransitPeerNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddTransitPeerNetwork {
	request := dnacentersdkgo.RequestSdaAddTransitPeerNetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_peer_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_peer_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_peer_network_name")))) {
		request.TransitPeerNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_peer_network_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_peer_network_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_peer_network_type")))) {
		request.TransitPeerNetworkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_transit_settings")))) {
		request.IPTransitSettings = expandRequestTransitPeerNetworkAddTransitPeerNetworkIPTransitSettings(ctx, key+".ip_transit_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sda_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sda_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sda_transit_settings")))) {
		request.SdaTransitSettings = expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettings(ctx, key+".sda_transit_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTransitPeerNetworkAddTransitPeerNetworkIPTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddTransitPeerNetworkIPTransitSettings {
	request := dnacentersdkgo.RequestSdaAddTransitPeerNetworkIPTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".routing_protocol_name")))) {
		request.RoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".autonomous_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".autonomous_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".autonomous_system_number")))) {
		request.AutonomousSystemNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettings {
	request := dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_control_plane_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_control_plane_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_control_plane_settings")))) {
		request.TransitControlPlaneSettings = expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettingsArray(ctx, key+".transit_control_plane_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings {
	request := []dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings{}
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
		i := expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTransitPeerNetworkAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings {
	request := dnacentersdkgo.RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
