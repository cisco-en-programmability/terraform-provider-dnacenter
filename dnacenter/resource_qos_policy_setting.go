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

func resourceQosPolicySetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Application Policy.

- API to update the application QoS policy setting.
`,

		CreateContext: resourceQosPolicySettingCreate,
		ReadContext:   resourceQosPolicySettingRead,
		UpdateContext: resourceQosPolicySettingUpdate,
		DeleteContext: resourceQosPolicySettingDelete,
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

						"deploy_by_default_on_wired_devices": &schema.Schema{
							Description: `Flag to indicate whether QoS policy should be deployed automatically on wired network device when it is provisioned. This would be only applicable for cases where the network device is assigned to a site where a QoS policy has been configured.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
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

						"deploy_by_default_on_wired_devices": &schema.Schema{
							Description: `Flag to indicate whether QoS policy should be deployed automatically on wired network device when it is provisioned. This would be only applicable for cases where the network device is assigned to a site where a QoS policy has been configured.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceQosPolicySettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceQosPolicySettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheApplicationQoSPolicySetting")

		response1, restyResp1, err := client.ApplicationPolicy.RetrievesTheApplicationQoSPolicySetting()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyRetrievesTheApplicationQoSPolicySettingItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheApplicationQoSPolicySetting response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceQosPolicySettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestQosPolicySettingUpdatesTheApplicationQoSPolicySetting(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		restyResp1, err := client.ApplicationPolicy.UpdatesTheApplicationQoSPolicySetting(request1)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesTheApplicationQoSPolicySetting", err, restyResp1.String(),
					"Failure at UpdatesTheApplicationQoSPolicySetting, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesTheApplicationQoSPolicySetting", err,
				"Failure at UpdatesTheApplicationQoSPolicySetting, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceQosPolicySettingRead(ctx, d, m)
}

func resourceQosPolicySettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete QosPolicySetting on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestQosPolicySettingUpdatesTheApplicationQoSPolicySetting(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdatesTheApplicationQoSPolicySetting {
	request := dnacentersdkgo.RequestApplicationPolicyUpdatesTheApplicationQoSPolicySetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_by_default_on_wired_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_by_default_on_wired_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_by_default_on_wired_devices")))) {
		request.DeployByDefaultOnWiredDevices = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
