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
func resourceSecurityThreatsSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- The Threat Summary for the Rogues and aWIPS
`,

		CreateContext: resourceSecurityThreatsSummaryCreate,
		ReadContext:   resourceSecurityThreatsSummaryRead,
		DeleteContext: resourceSecurityThreatsSummaryDelete,
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
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"threat_data": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"threat_count": &schema.Schema{
													Description: `Threat Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"threat_level": &schema.Schema{
													Description: `Threat Level`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"threat_type": &schema.Schema{
													Description: `Threat Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
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

func resourceSecurityThreatsSummaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSecurityThreatsSummaryThreatSummary(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.ThreatSummary(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ThreatSummary", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesThreatSummaryItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ThreatSummary response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSecurityThreatsSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSecurityThreatsSummaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSecurityThreatsSummaryThreatSummary(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesThreatSummary {
	request := dnacentersdkgo.RequestDevicesThreatSummary{}
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
	return &request
}

func flattenDevicesThreatSummaryItems(items *[]dnacentersdkgo.ResponseDevicesThreatSummaryResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["threat_data"] = flattenDevicesThreatSummaryItemsThreatData(item.ThreatData)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesThreatSummaryItemsThreatData(items *[]dnacentersdkgo.ResponseDevicesThreatSummaryResponseThreatData) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["threat_type"] = item.ThreatType
		respItem["threat_level"] = item.ThreatLevel
		respItem["threat_count"] = item.ThreatCount
		respItems = append(respItems, respItem)
	}
	return respItems
}
