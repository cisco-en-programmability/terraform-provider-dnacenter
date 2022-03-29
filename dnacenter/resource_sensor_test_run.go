package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSensorTestRun() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.
- Intent API to run a deployed SENSOR test
`,

		CreateContext: resourceSensorTestRunCreate,
		ReadContext:   resourceSensorTestRunRead,
		DeleteContext: resourceSensorTestRunDelete,

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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_name": &schema.Schema{
							Description: `Template Name`,
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSensorTestRunCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestSensorTestRunRunNowSensorTest(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, err := client.Sensors.RunNowSensorTest(request1)
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

func resourceSensorTestRunRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSensorTestRunUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSensorTestRunRead(ctx, d, m)
}

func resourceSensorTestRunDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
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
