package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAnalyticsEndpointsAncPolicyDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on AI Endpoint Analytics.

- Revokes given ANC policy from the endpoint.
`,

		CreateContext: resourceAnalyticsEndpointsAncPolicyDeleteCreate,
		ReadContext:   resourceAnalyticsEndpointsAncPolicyDeleteRead,
		DeleteContext: resourceAnalyticsEndpointsAncPolicyDeleteDelete,
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
					},
				},
			},
		},
	}
}

func resourceAnalyticsEndpointsAncPolicyDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vEpID := resourceItem["ep_id"]

	vvEpID := vEpID.(string)

	// has_unknown_response: True

	response1, err := client.AIEndpointAnalytics.RevokeAncPolicy(vvEpID)

	if err != nil || response1 == nil {
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	//Analizar verificacion.

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RevokeAncPolicy response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceAnalyticsEndpointsAncPolicyDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAnalyticsEndpointsAncPolicyDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
