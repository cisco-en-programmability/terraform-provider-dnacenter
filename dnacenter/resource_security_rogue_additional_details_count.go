package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSecurityRogueAdditionalDetailsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- This data source action returns the count for the Rogue Additional Details.
`,

		CreateContext: resourceSecurityRogueAdditionalDetailsCountCreate,
		ReadContext:   resourceSecurityRogueAdditionalDetailsCountRead,
		DeleteContext: resourceSecurityRogueAdditionalDetailsCountDelete,
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
							Description: `This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time
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
							Description: `This information can be fetched from "Get Threat Levels" API
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
							Description: `This information can be fetched from "Get Threat Types" API
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

func resourceSecurityRogueAdditionalDetailsCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSecurityRogueAdditionalDetailsCountRogueAdditionalDetailCount(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.RogueAdditionalDetailCount(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing RogueAdditionalDetailCount", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDevicesRogueAdditionalDetailCountItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RogueAdditionalDetailCount response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSecurityRogueAdditionalDetailsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSecurityRogueAdditionalDetailsCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSecurityRogueAdditionalDetailsCountRogueAdditionalDetailCount(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRogueAdditionalDetailCount {
	request := dnacentersdkgo.RequestDevicesRogueAdditionalDetailCount{}
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

func flattenDevicesRogueAdditionalDetailCountItem(item *dnacentersdkgo.ResponseDevicesRogueAdditionalDetailCount) []map[string]interface{} {
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
