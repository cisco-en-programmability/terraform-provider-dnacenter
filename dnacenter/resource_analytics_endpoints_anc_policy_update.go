package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAnalyticsEndpointsAncPolicyUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on AI Endpoint Analytics.

- Applies given ANC policy to the endpoint.
`,

		CreateContext: resourceAnalyticsEndpointsAncPolicyUpdateCreate,
		ReadContext:   resourceAnalyticsEndpointsAncPolicyUpdateRead,
		DeleteContext: resourceAnalyticsEndpointsAncPolicyUpdateDelete,
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
						"ep_id": &schema.Schema{
							Description: `epId path parameter. Unique identifier for the endpoint.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"anc_policy": &schema.Schema{
							Description: `ANC policy name.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"granular_anc_policy": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name of the granular ANC policy.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"nas_ip_address": &schema.Schema{
										Description: `IP address of the network device where endpoint is attached.
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

func resourceAnalyticsEndpointsAncPolicyUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vEpID := resourceItem["ep_id"]

	vvEpID := vEpID.(string)
	request1 := expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicy(ctx, "parameters.0", d)

	// has_unknown_response: True

	response1, err := client.AIEndpointAnalytics.ApplyAncPolicy(vvEpID, request1)

	if err != nil || response1 == nil {
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	//Analizar verificacion.

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApplyAncPolicy response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceAnalyticsEndpointsAncPolicyUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAnalyticsEndpointsAncPolicyUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicy {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anc_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anc_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anc_policy")))) {
		request.AncPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".granular_anc_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".granular_anc_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".granular_anc_policy")))) {
		request.GranularAncPolicy = expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicyGranularAncPolicyArray(ctx, key+".granular_anc_policy", d)
	}
	return &request
}

func expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicyGranularAncPolicyArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy {
	request := []dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy{}
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
		i := expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicyGranularAncPolicy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAnalyticsEndpointsAncPolicyUpdateApplyAncPolicyGranularAncPolicy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nas_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nas_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nas_ip_address")))) {
		request.NasIPAddress = interfaceToString(v)
	}
	return &request
}
