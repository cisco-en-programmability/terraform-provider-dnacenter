package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProjectsProjectID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get a template project by the project's ID.
`,

		ReadContext: dataSourceProjectsProjectIDRead,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Description: `projectId path parameter. The id of the project to get, retrieveable from *GET /dna/intent/api/v1/projects*
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Timestamp of when the project was updated or modified
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_id": &schema.Schema{
							Description: `UUID of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProjectsProjectIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProjectID := d.Get("project_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateProject")
		vvProjectID := vProjectID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateProject(vvProjectID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplateProject", err,
				"Failure at GetTemplateProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesGetTemplateProjectItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateProject response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateProjectItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjectResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["project_id"] = item.ProjectID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["last_update_time"] = item.LastUpdateTime
	return []map[string]interface{}{
		respItem,
	}
}
