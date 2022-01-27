package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceConfigurationTemplateClone() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- API to clone template
`,

		ReadContext: dataSourceConfigurationTemplateCloneRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name path parameter. Template name to clone template(Name should be different than existing template name within same project)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId path parameter. UUID of the project in which the template needs to be created
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"template_id": &schema.Schema{
				Description: `templateId path parameter. UUID of the template to clone it
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
		},
	}
}

func dataSourceConfigurationTemplateCloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName := d.Get("name")
	vTemplateID := d.Get("template_id")
	vProjectID := d.Get("project_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreatesACloneOfTheGivenTemplate")
		vvName := vName.(string)
		vvTemplateID := vTemplateID.(string)
		vvProjectID := vProjectID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.CreatesACloneOfTheGivenTemplate(vvName, vvTemplateID, vvProjectID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreatesACloneOfTheGivenTemplate", err,
				"Failure at CreatesACloneOfTheGivenTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesCreatesACloneOfTheGivenTemplateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreatesACloneOfTheGivenTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesCreatesACloneOfTheGivenTemplateItem(item *dnacentersdkgo.ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateResponse) []map[string]interface{} {
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
