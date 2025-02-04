package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAnalyticsCmdbEndpoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on AI Endpoint Analytics.

- Processes incoming CMDB endpoints data and imports the same in AI Endpoint Analytics.
`,

		CreateContext: resourceAnalyticsCmdbEndpointsCreate,
		ReadContext:   resourceAnalyticsCmdbEndpointsRead,
		DeleteContext: resourceAnalyticsCmdbEndpointsDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
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
						"payload": &schema.Schema{
							Description: `Array of RequestAIEndpointAnalyticsProcessCMDBEndpoints`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"asset_tag": &schema.Schema{
										Description: `Asset tag.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"department": &schema.Schema{
										Description: `Department that asset belongs to.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of the asset.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"last_update_timestamp": &schema.Schema{
										Description: `Last update timestamp in epoch milliseconds.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location": &schema.Schema{
										Description: `Location of the asset.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address of the endpoint.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"managed_by": &schema.Schema{
										Description: `Asset managed by.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"model": &schema.Schema{
										Description: `Asset model.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"model_category": &schema.Schema{
										Description: `Category of the model.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"serial_number": &schema.Schema{
										Description: `Serial number of the endpoint.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceAnalyticsCmdbEndpointsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpoints(ctx, "parameters.0", d)

	// has_unknown_response: True

	response1, err := client.AIEndpointAnalytics.ProcessCmdbEndpoints(request1)

	if err != nil || response1 == nil {
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	//Analizar verificacion.

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ProcessCmdbEndpoints response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceAnalyticsCmdbEndpointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAnalyticsCmdbEndpointsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsProcessCmdbEndpoints {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsProcessCmdbEndpoints{}
	if v := expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpointsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpointsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemAIEndpointAnalyticsProcessCmdbEndpoints {
	request := []dnacentersdkgo.RequestItemAIEndpointAnalyticsProcessCmdbEndpoints{}
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
		i := expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpointsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAnalyticsCmdbEndpointsProcessCmdbEndpointsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemAIEndpointAnalyticsProcessCmdbEndpoints {
	request := dnacentersdkgo.RequestItemAIEndpointAnalyticsProcessCmdbEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".asset_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".asset_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".asset_tag")))) {
		request.AssetTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_category")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_category")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_category")))) {
		request.ModelCategory = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".department")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".department")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".department")))) {
		request.Department = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".managed_by")))) {
		request.ManagedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_timestamp")))) {
		request.LastUpdateTimestamp = interfaceToIntPtr(v)
	}
	return &request
}
