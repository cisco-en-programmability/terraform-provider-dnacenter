package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSensorTestDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.
		- Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID
	
`,

		CreateContext: resourceSensorTestDeleteCreate,
		ReadContext:   resourceSensorTestDeleteRead,
		DeleteContext: resourceSensorTestDeleteDelete,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_name": &schema.Schema{
							Description: `templateName query parameter.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSensorTestDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	vTemplateName := d.Get("parameters.0.template_name").(string)
	log.Printf("[DEBUG] Selected method 1: DeleteSensorTest")
	queryParams1 := dnacentersdkgo.DeleteSensorTestQueryParams{}
	queryParams1.TemplateName = vTemplateName

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
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceSensorTestDeleteRead(ctx, d, m)
}

func resourceSensorTestDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceSensorTestDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
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
