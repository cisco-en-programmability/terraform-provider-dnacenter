package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSecurityThreatsDetailsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- The details count for the Rogue and aWIPS threats
`,

		CreateContext: resourceSecurityThreatsDetailsCountCreate,
		ReadContext:   resourceSecurityThreatsDetailsCountRead,
		DeleteContext: resourceSecurityThreatsDetailsCountDelete,
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

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
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
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"is_new_threat": &schema.Schema{
							Description: `Is New Threat`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"limit": &schema.Schema{
							Description: `Limit`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"offset": &schema.Schema{
							Description: `Offset`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"threat_level": &schema.Schema{
							Description: `Threat Level`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"threat_type": &schema.Schema{
							Description: `Threat Type`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
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

func resourceSecurityThreatsDetailsCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSecurityThreatsDetailsCountThreatDetailCount(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.ThreatDetailCount(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ThreatDetailCount", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDevicesThreatDetailCountItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ThreatDetailCount response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSecurityThreatsDetailsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSecurityThreatsDetailsCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSecurityThreatsDetailsCountThreatDetailCount(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesThreatDetailCount {
	request := dnacentersdkgo.RequestDevicesThreatDetailCount{}
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
	return &request
}

func flattenDevicesThreatDetailCountItem(item *dnacentersdkgo.ResponseDevicesThreatDetailCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}