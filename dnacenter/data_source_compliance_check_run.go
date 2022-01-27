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
func dataSourceComplianceCheckRun() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Compliance.

- Run compliance check for device(s).
`,

		ReadContext: dataSourceComplianceCheckRunRead,
		Schema: map[string]*schema.Schema{
			"categories": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"device_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"trigger_full": &schema.Schema{
				// Type:     schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
		},
	}
}

func dataSourceComplianceCheckRunRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RunCompliance")
		request1 := expandRequestComplianceCheckRunRunCompliance(ctx, "", d)

		response1, restyResp1, err := client.Compliance.RunCompliance(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RunCompliance", err,
				"Failure at RunCompliance, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceRunComplianceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RunCompliance response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestComplianceCheckRunRunCompliance(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestComplianceRunCompliance {
	request := dnacentersdkgo.RequestComplianceRunCompliance{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trigger_full")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trigger_full")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trigger_full")))) {
		request.TriggerFull = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenComplianceRunComplianceItem(item *dnacentersdkgo.ResponseComplianceRunComplianceResponse) []map[string]interface{} {
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
