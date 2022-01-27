package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceConfigurationTemplateImportProject() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- Imports the Projects provided in the DTO
`,

		ReadContext: dataSourceConfigurationTemplateImportProjectRead,
		Schema: map[string]*schema.Schema{
			"do_version": &schema.Schema{
				Description: `doVersion query parameter. If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
`,
				Type:     schema.TypeBool,
				Optional: true,
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

func dataSourceConfigurationTemplateImportProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDoVersion, okDoVersion := d.GetOk("do_version")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportsTheProjectsProvided")
		queryParams1 := dnacentersdkgo.ImportsTheProjectsProvidedQueryParams{}

		if okDoVersion {
			queryParams1.DoVersion = vDoVersion.(bool)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.ImportsTheProjectsProvided(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportsTheProjectsProvided", err,
				"Failure at ImportsTheProjectsProvided, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesImportsTheProjectsProvidedItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportsTheProjectsProvided response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesImportsTheProjectsProvidedItem(item *dnacentersdkgo.ResponseConfigurationTemplatesImportsTheProjectsProvidedResponse) []map[string]interface{} {
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
