package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProjects() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get all matching template projects based on the filters selected.
`,

		ReadContext: dataSourceProjectsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
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

func dataSourceProjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateProjects")
		queryParams1 := dnacentersdkgo.GetTemplateProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateProjects(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplateProjects", err,
				"Failure at GetTemplateProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetTemplateProjectsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateProjects response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateProjectsItems(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjectsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["project_id"] = item.ProjectID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["last_update_time"] = item.LastUpdateTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
