package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceWirelessPskOverride() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Update/override pass phrase of enterprise SSID
`,

		ReadContext: dataSourceWirelessPskOverrideRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"payload": &schema.Schema{
				Description: `Array of RequestWirelessPSKOverride`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"pass_phrase": &schema.Schema{
							Description: `Pass phrase (create/update)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site": &schema.Schema{
							Description: `site name hierarchy (ex: Global/aaa/zzz/...) 
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssid": &schema.Schema{
							Description: `enterprise ssid name(already created/present)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessPskOverrideRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: PSKOverride")
		request1 := expandRequestWirelessPskOverridePSKOverride(ctx, "", d)

		response1, restyResp1, err := client.Wireless.PSKOverride(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PSKOverride", err,
				"Failure at PSKOverride, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessPSKOverrideItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting PSKOverride response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestWirelessPskOverridePSKOverride(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessPSKOverride {
	request := dnacentersdkgo.RequestWirelessPSKOverride{}
	if v := expandRequestWirelessPskOverridePSKOverrideItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessPskOverridePSKOverrideItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessPSKOverride {
	request := []dnacentersdkgo.RequestItemWirelessPSKOverride{}
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
		i := expandRequestWirelessPskOverridePSKOverrideItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessPskOverridePSKOverrideItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessPSKOverride {
	request := dnacentersdkgo.RequestItemWirelessPSKOverride{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pass_phrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pass_phrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pass_phrase")))) {
		request.PassPhrase = interfaceToString(v)
	}
	return &request
}

func flattenWirelessPSKOverrideItem(item *dnacentersdkgo.ResponseWirelessPSKOverride) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
