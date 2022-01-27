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
func dataSourceThreatDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- The details for the Rogue and aWIPS threats
`,

		ReadContext: dataSourceThreatDetailRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `End Time`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"is_new_threat": &schema.Schema{
				Description: `Is New Threat`,
				// Type:        schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_name": &schema.Schema{
							Description: `Ap Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"threat_level": &schema.Schema{
							Description: `Threat Level`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"threat_type": &schema.Schema{
							Description: `Threat Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_time": &schema.Schema{
							Description: `Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"vendor": &schema.Schema{
							Description: `Vendor`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"limit": &schema.Schema{
				Description: `Limit`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `Offset`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"site_id": &schema.Schema{
				Description: `Site Id`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_time": &schema.Schema{
				Description: `Start Time`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"threat_level": &schema.Schema{
				Description: `Threat Level`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"threat_type": &schema.Schema{
				Description: `Threat Type`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceThreatDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ThreatDetails")
		request1 := expandRequestThreatDetailThreatDetails(ctx, "", d)

		response1, restyResp1, err := client.Devices.ThreatDetails(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ThreatDetails", err,
				"Failure at ThreatDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesThreatDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ThreatDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestThreatDetailThreatDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesThreatDetails {
	request := dnacentersdkgo.RequestDevicesThreatDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_new_threat")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_new_threat")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_new_threat")))) {
		request.IsNewThreat = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDevicesThreatDetailsItems(items *[]dnacentersdkgo.ResponseDevicesThreatDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mac_address"] = item.MacAddress
		respItem["updated_time"] = item.UpdatedTime
		respItem["vendor"] = item.Vendor
		respItem["threat_type"] = item.ThreatType
		respItem["threat_level"] = item.ThreatLevel
		respItem["ap_name"] = item.ApName
		respItem["ssid"] = item.SSID
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItems = append(respItems, respItem)
	}
	return respItems
}
