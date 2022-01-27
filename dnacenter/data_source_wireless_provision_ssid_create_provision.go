package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceWirelessProvisionSSIDCreateProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Creates SSID, updates the SSID to the corresponding site profiles and provision it to the devices matching the given
sites
`,

		ReadContext: dataSourceWirelessProvisionSSIDCreateProvisionRead,
		Schema: map[string]*schema.Schema{
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"enable_fabric": &schema.Schema{
				Description: `Enable SSID for Fabric
`,
				// Type:        schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"flex_connect": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_flex_connect": &schema.Schema{
							Description: `Enable Flex Connect
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"local_to_vlan": &schema.Schema{
							Description: `Local To Vlan (range is 1 to 4094)
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
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
			"managed_aplocations": &schema.Schema{
				Description: `Managed AP Locations (Enter entire Site(s) hierarchy)
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ssid_details": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_broadcast_ssi_d": &schema.Schema{
							Description: `Enable Broadcast SSID
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_fast_lane": &schema.Schema{
							Description: `Enable Fast Lane
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_mac_filtering": &schema.Schema{
							Description: `Enable MAC Filtering
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `SSID Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"passphrase": &schema.Schema{
							Description: `Pass Phrase ( Only applicable for SSID with PERSONAL auth type )
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_policy": &schema.Schema{
							Description: `Radio Policy. Allowed values are 'Dual band operation (2.4GHz and 5GHz)', 'Dual band operation with band select', '5GHz only', '2.4GHz only'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_level": &schema.Schema{
							Description: `Security Level(For guest SSID OPEN/WEB_AUTH, For Enterprise SSID ENTERPRISE/PERSONAL/OPEN)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"traffic_type": &schema.Schema{
							Description: `Traffic Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"web_auth_url": &schema.Schema{
							Description: `Web Auth URL
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ssid_type": &schema.Schema{
				Description: `SSID Type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceWirelessProvisionSSIDCreateProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateAndProvisionSSID")
		request1 := expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSID(ctx, "", d)

		headerParams1 := dnacentersdkgo.CreateAndProvisionSSIDHeaderParams{}

		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.Wireless.CreateAndProvisionSSID(request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateAndProvisionSSID", err,
				"Failure at CreateAndProvisionSSID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessCreateAndProvisionSSIDItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateAndProvisionSSID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateAndProvisionSSID {
	request := dnacentersdkgo.RequestWirelessCreateAndProvisionSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_aplocations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_aplocations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".managed_aplocations")))) {
		request.ManagedApLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSIDSSIDDetails(ctx, key+".ssid_details.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_type")))) {
		request.SSIDType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSIDFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSIDSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateAndProvisionSSIDSSIDDetails {
	request := dnacentersdkgo.RequestWirelessCreateAndProvisionSSIDSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_level")))) {
		request.SecurityLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fast_lane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fast_lane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fast_lane")))) {
		request.EnableFastLane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) {
		request.EnableBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_policy")))) {
		request.RadioPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_mac_filtering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_mac_filtering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_mac_filtering")))) {
		request.EnableMacFiltering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".web_auth_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".web_auth_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".web_auth_url")))) {
		request.WebAuthURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProvisionSSIDCreateProvisionCreateAndProvisionSSIDFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateAndProvisionSSIDFlexConnect {
	request := dnacentersdkgo.RequestWirelessCreateAndProvisionSSIDFlexConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_flex_connect")))) {
		request.EnableFlexConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_to_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_to_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_to_vlan")))) {
		request.LocalToVLAN = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenWirelessCreateAndProvisionSSIDItem(item *dnacentersdkgo.ResponseWirelessCreateAndProvisionSSID) []map[string]interface{} {
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
