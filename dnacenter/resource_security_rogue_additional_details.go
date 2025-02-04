package dnacenter

import (
	"context"

	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSecurityRogueAdditionalDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- This data source action provides additional information of the rogue threats with details at BSSID level. The
additional information includes Switch Port details in case of Rogue on Wire, first time when the rogue is seen in the
network etc.
`,

		CreateContext: resourceSecurityRogueAdditionalDetailsCreate,
		ReadContext:   resourceSecurityRogueAdditionalDetailsRead,
		DeleteContext: resourceSecurityRogueAdditionalDetailsDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time": &schema.Schema{
							Description: `This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_name": &schema.Schema{
										Description: `Detecting AP Name
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"channel_number": &schema.Schema{
										Description: `Channel Number on which the Rogue is detected
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"containment": &schema.Schema{
										Description: `Containment Status of the Rogue
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"controller_ip": &schema.Schema{
										Description: `IP Address of the Controller detecting this Rogue
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"controller_name": &schema.Schema{
										Description: `Name of the Controller detecting this Rogue
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"created_time": &schema.Schema{
										Description: `First time when the Rogue is seen in the network
`,
										Type:     schema.TypeInt,
										ForceNew: true,
										Computed: true,
									},
									"detecting_apmac": &schema.Schema{
										Description: `MAC Address of the Detecting AP
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"encryption": &schema.Schema{
										Description: `Security status of the Rogue SSID
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC Address of the Rogue BSSID
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"mld_mac_address": &schema.Schema{
										Description: `MLD MAC Address of the Rogue BSSID, this is applicable only for Wi-Fi 7 Rogues
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"port_description": &schema.Schema{
										Description: `Port information of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"radio_type": &schema.Schema{
										Description: `Radio Type on which Rogue is detected
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"site_name_hierarchy": &schema.Schema{
										Description: `Site Hierarchy of the Rogue
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Description: `Rogue SSID
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"switch_ip": &schema.Schema{
										Description: `IP Address of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"switch_name": &schema.Schema{
										Description: `Name of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"threat_level": &schema.Schema{
										Description: `Level of the Rogue Threat
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"threat_type": &schema.Schema{
										Description: `Type of the Rogue Threat
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"updated_time": &schema.Schema{
										Description: `Last time when the Rogue is seen in the network
`,
										Type:     schema.TypeInt,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"limit": &schema.Schema{
							Description: `The maximum number of entries to return. Default value is 1000
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"offset": &schema.Schema{
							Description: `The offset of the first item in the collection to return. Default value is 1
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `Filter Rogues by location. Site IDs information can be fetched from "Get Site" API
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"start_time": &schema.Schema{
							Description: `This is the epoch start time in milliseconds from which data need to be fetched. Default value is 24 hours earlier to endTime
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"threat_level": &schema.Schema{
							Description: `Filter Rogues by Threat Level. Threat Level information can be fetched from "Get Threat Levels" API
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"threat_type": &schema.Schema{
							Description: `Filter Rogues by Threat Type. Threat Type information can be fetched from "Get Threat Types" API
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceSecurityRogueAdditionalDetailsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSecurityRogueAdditionalDetailsRogueAdditionalDetails(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.RogueAdditionalDetails(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing RogueAdditionalDetails", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesRogueAdditionalDetailsItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RogueAdditionalDetails response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSecurityRogueAdditionalDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSecurityRogueAdditionalDetailsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSecurityRogueAdditionalDetailsRogueAdditionalDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRogueAdditionalDetails {
	request := dnacentersdkgo.RequestDevicesRogueAdditionalDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".threat_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".threat_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".threat_level")))) {
		request.ThreatLevel = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".threat_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".threat_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".threat_type")))) {
		request.ThreatType = interfaceToSliceString(v)
	}
	return &request
}

func flattenDevicesRogueAdditionalDetailsItems(items *[]dnacentersdkgo.ResponseDevicesRogueAdditionalDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mac_address"] = item.MacAddress
		respItem["mld_mac_address"] = item.MldMacAddress
		respItem["updated_time"] = item.UpdatedTime
		respItem["created_time"] = item.CreatedTime
		respItem["threat_type"] = item.ThreatType
		respItem["threat_level"] = item.ThreatLevel
		respItem["ap_name"] = item.ApName
		respItem["detecting_apmac"] = item.DetectingApMac
		respItem["ssid"] = item.SSID
		respItem["containment"] = item.Containment
		respItem["radio_type"] = item.RadioType
		respItem["controller_ip"] = item.ControllerIP
		respItem["controller_name"] = item.ControllerName
		respItem["channel_number"] = item.ChannelNumber
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["encryption"] = item.Encryption
		respItem["switch_ip"] = item.SwitchIP
		respItem["switch_name"] = item.SwitchName
		respItem["port_description"] = item.PortDescription
		respItems = append(respItems, respItem)
	}
	return respItems
}
