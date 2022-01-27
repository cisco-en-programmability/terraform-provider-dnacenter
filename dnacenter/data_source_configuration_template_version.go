package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplateVersion() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get all the versions of template by its id
`,

		ReadContext: dataSourceConfigurationTemplateVersionRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Description: `templateId path parameter. templateId(UUID) to get list of versioned templates
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"versions_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of version template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_comment": &schema.Schema{
										Description: `Version comment
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_time": &schema.Schema{
										Description: `Template version time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceConfigurationTemplateVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsAllTheVersionsOfAGivenTemplate")
		vvTemplateID := vTemplateID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.GetsAllTheVersionsOfAGivenTemplate(vvTemplateID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsAllTheVersionsOfAGivenTemplate", err,
				"Failure at GetsAllTheVersionsOfAGivenTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAllTheVersionsOfAGivenTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateItems(items *dnacentersdkgo.ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["name"] = item.Name
		respItem["project_id"] = item.ProjectID
		respItem["project_name"] = item.ProjectName
		respItem["template_id"] = item.TemplateID
		respItem["versions_info"] = flattenConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateItemsVersionsInfo(item.VersionsInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateItemsVersionsInfo(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateVersionsInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["author"] = item.Author
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["version"] = item.Version
		respItem["version_comment"] = item.VersionComment
		respItem["version_time"] = item.VersionTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
