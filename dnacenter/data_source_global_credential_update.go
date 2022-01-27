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
func dataSourceGlobalCredentialUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Discovery.

- Update global credential for network devices in site(s)
`,

		ReadContext: dataSourceGlobalCredentialUpdateRead,
		Schema: map[string]*schema.Schema{
			"global_credential_id": &schema.Schema{
				Description: `globalCredentialId path parameter. Global credential Uuid
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"site_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGlobalCredentialUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vGlobalCredentialID := d.Get("global_credential_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateGlobalCredentials")
		vvGlobalCredentialID := vGlobalCredentialID.(string)
		request1 := expandRequestGlobalCredentialUpdateUpdateGlobalCredentials(ctx, "", d)

		response1, restyResp1, err := client.Discovery.UpdateGlobalCredentials(vvGlobalCredentialID, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalCredentials", err,
				"Failure at UpdateGlobalCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryUpdateGlobalCredentialsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateGlobalCredentials response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGlobalCredentialUpdateUpdateGlobalCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentials {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_uuids")))) {
		request.SiteUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDiscoveryUpdateGlobalCredentialsItem(item *dnacentersdkgo.ResponseDiscoveryUpdateGlobalCredentialsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
