package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSensorDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Sensors.

- Intent API to delete an existing SENSOR test template
`,

		ReadContext: dataSourceSensorDeleteRead,
		Schema: map[string]*schema.Schema{
			"template_name": &schema.Schema{
				Description: `templateName query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"template_name": &schema.Schema{
							Description: `Template Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSensorDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateName, okTemplateName := d.GetOk("template_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteSensorTest")
		queryParams1 := dnacentersdkgo.DeleteSensorTestQueryParams{}

		if okTemplateName {
			queryParams1.TemplateName = vTemplateName.(string)
		}

		response1, restyResp1, err := client.Sensors.DeleteSensorTest(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteSensorTest", err,
				"Failure at DeleteSensorTest, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsDeleteSensorTestItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteSensorTest response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsDeleteSensorTestItem(item *dnacentersdkgo.ResponseSensorsDeleteSensorTestResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["template_name"] = item.TemplateName
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}
