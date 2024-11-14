package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAnalyticsEndpoints() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on AI Endpoint Analytics.

- Register a new endpoint in the system.

- Update attributes of a registered endpoint.

- Deletes the endpoint for the given unique identifier 'epId'.
`,

		CreateContext: resourceAnalyticsEndpointsCreate,
		ReadContext:   resourceAnalyticsEndpointsRead,
		UpdateContext: resourceAnalyticsEndpointsUpdate,
		DeleteContext: resourceAnalyticsEndpointsDelete,
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

						"anc_policy": &schema.Schema{
							Description: `ANC policy currently applied to the endpoint in ISE.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": &schema.Schema{
							Description: `Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},
						"device_type": &schema.Schema{
							Description: `Type of the device represented by this endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"duid": &schema.Schema{
							Description: `Unique DUID.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"granular_anc_policy": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name of the granular ANC policy.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"nas_ip_address": &schema.Schema{
										Description: `IP address of the network device where endpoint is attached.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hardware_manufacturer": &schema.Schema{
							Description: `Hardware manufacturer for the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"hardware_model": &schema.Schema{
							Description: `Hardware model of the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `Unique identifier for the endpoint.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_probe_collection_timestamp": &schema.Schema{
							Description: `Last probe collection timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Description: `MAC address of the endpoint.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"operating_system": &schema.Schema{
							Description: `Operating system of the endpoint.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"random_mac": &schema.Schema{
							Description: `Flag to indicate whether MAC address is a randomized one or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"registered": &schema.Schema{
							Description: `Flag to indicate whether this is a manually registered endpoint or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"trust_data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ai_spoofing_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to AI spoofing.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_method": &schema.Schema{
										Description: `Authentication method.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"changed_profile_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to changing profile labels.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"concurrent_mac_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to concurrent MAC address.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_blocklist_detected": &schema.Schema{
										Description: `Flag to fetch endpoints hitting IP blocklist or not.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"nat_trust_level": &schema.Schema{
										Description: `Trust level of the endpoint due to NAT access.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"posture_status": &schema.Schema{
										Description: `Posture status.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"trust_score": &schema.Schema{
										Description: `Overall trust score of the endpoint.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"unauth_port_detected": &schema.Schema{
										Description: `Flag to fetch endpoints exposing unauthorized ports or not.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"weak_cred_detected": &schema.Schema{
										Description: `Flag to fetch endpoints having weak credentials or not.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

						"device_type": &schema.Schema{
							Description: `Type of the device represented by this endpoint.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_id": &schema.Schema{
							Description: `epId path parameter. Unique identifier for the endpoint.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"hardware_manufacturer": &schema.Schema{
							Description: `Hardware manufacturer for the endpoint.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hardware_model": &schema.Schema{
							Description: `Hardware model of the endpoint.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Description: `MAC address of the endpoint.
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

func resourceAnalyticsEndpointsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAnalyticsEndpointsRegisterAnEndpoint(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vEpID, okEpID := resourceItem["ep_id"]
	vvEpID := interfaceToString(vEpID)
	vMacAdress := resourceItem["mac_address"]
	vvMacAdress := interfaceToString(vMacAdress)
	if okEpID && vvEpID != "" {
		getResponse2, _, err := client.AIEndpointAnalytics.GetEndpointDetails(vvEpID, nil)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["ep_id"] = vvEpID
			d.SetId(joinResourceID(resourceMap))
			return resourceAnalyticsEndpointsRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.QueryTheEndpointsQueryParams{
			MacAddress: vvMacAdress,
		}

		response2, _, err := client.AIEndpointAnalytics.QueryTheEndpoints(&queryParamImport)

		if response2 != nil && err == nil && *response2.TotalResults > 0 {
			items := *response2.Items
			item := items[0]
			resourceMap := make(map[string]string)
			resourceMap["ep_id"] = item.ID
			d.SetId(joinResourceID(resourceMap))
			return resourceAnalyticsEndpointsRead(ctx, d, m)
		}
	}
	restyResp1, err := client.AIEndpointAnalytics.RegisterAnEndpoint(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing RegisterAnEndpoint", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing RegisterAnEndpoint", err))
		return diags
	}

	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.QueryTheEndpointsQueryParams{
		MacAddress: vvMacAdress,
	}
	item3, _, err := client.AIEndpointAnalytics.QueryTheEndpoints(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RegisterAnEndpoint", err,
			"Failure at RegisterAnEndpoint, unexpected response", ""))
		return diags
	}
	if item3 != nil && err == nil && *item3.TotalResults > 0 {
		items := *item3.Items
		item := items[0]
		resourceMap := make(map[string]string)
		resourceMap["ep_id"] = item.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceAnalyticsEndpointsRead(ctx, d, m)
	} else {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RegisterAnEndpoint", err,
			"Failure at RegisterAnEndpoint, unexpected response", ""))
		return diags
	}
}

func resourceAnalyticsEndpointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["ep_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEndpointDetails")
		vvEpID := vvID
		queryParams1 := dnacentersdkgo.GetEndpointDetailsQueryParams{}

		// has_unknown_response: None

		response1, restyResp1, err := client.AIEndpointAnalytics.GetEndpointDetails(vvEpID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem2 := flattenAIEndpointAnalyticsGetEndpointDetailsItem(response1)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointDetails response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAnalyticsEndpointsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["ep_id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestAnalyticsEndpointsUpdateARegisteredEndpoint(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		restyResp1, err := client.AIEndpointAnalytics.UpdateARegisteredEndpoint(vvID, request1)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateARegisteredEndpoint", err, restyResp1.String(),
					"Failure at UpdateARegisteredEndpoint, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateARegisteredEndpoint", err,
				"Failure at UpdateARegisteredEndpoint, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceAnalyticsEndpointsRead(ctx, d, m)
}

func resourceAnalyticsEndpointsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["ep_id"]

	restyResp1, err := client.AIEndpointAnalytics.DeleteAnEndpoint(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAnEndpoint", err, restyResp1.String(),
				"Failure at DeleteAnEndpoint, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAnEndpoint", err,
			"Failure at DeleteAnEndpoint, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAnalyticsEndpointsRegisterAnEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsRegisterAnEndpoint {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsRegisterAnEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAnalyticsEndpointsUpdateARegisteredEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsUpdateARegisteredEndpoint {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsUpdateARegisteredEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
