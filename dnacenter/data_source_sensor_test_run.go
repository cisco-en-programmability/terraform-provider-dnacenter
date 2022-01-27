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
func dataSourceSensorTestRun() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.

- Intent API to run a deployed SENSOR test
`,

		ReadContext: dataSourceSensorTestRunRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_name": &schema.Schema{
				Description: `Template Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSensorTestRunRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RunNowSensorTest")
		request1 := expandRequestSensorTestRunRunNowSensorTest(ctx, "", d)

		response1, err := client.Sensors.RunNowSensorTest(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RunNowSensorTest", err,
				"Failure at RunNowSensorTest, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RunNowSensorTest response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSensorTestRunRunNowSensorTest(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRunNowSensorTest {
	request := dnacentersdkgo.RequestSensorsRunNowSensorTest{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_name")))) {
		request.TemplateName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
