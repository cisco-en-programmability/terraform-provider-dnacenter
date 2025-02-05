package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesTemplateIDVersionsVersionID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get a template's version by the version ID.
`,

		ReadContext: dataSourceTemplatesTemplateIDVersionsVersionIDRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Description: `templateId path parameter. The id of the template to get versions of, retrieveable from *GET /dna/intent/api/v1/templates*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"version_id": &schema.Schema{
				Description: `versionId path parameter. The id of the versioned template to get versions of, retrieveable from *GET /dna/intent/api/v1/templates/{id}/versions*
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"composite_template": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_policy": &schema.Schema{
										Description: `Policy to handle failure only applicable for composite templates  CONTINUE_ON_ERROR: If a composed template fails while deploying a device, continue deploying the next composed template  ABORT_TARGET_ON_ERROR: If a composed template fails while deploying to a device, abort the subsequent composed templates to that device if there any remaining
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Timestamp of when the template was updated or modified
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"products": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Family name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_name": &schema.Schema{
													Description: `Name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_series": &schema.Schema{
													Description: `Series name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"project_id": &schema.Schema{
										Description: `Id of the project
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_family": &schema.Schema{
										Description: `Software Family`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"template_id": &schema.Schema{
										Description: `The id of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `The type of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"regular_template": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"language": &schema.Schema{
										Description: `Language of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Timestamp of when the template was updated or modified
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"products": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Family name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_name": &schema.Schema{
													Description: `Name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_series": &schema.Schema{
													Description: `Series name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"project_id": &schema.Schema{
										Description: `Id of the project
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_family": &schema.Schema{
										Description: `Software Family`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"template_content": &schema.Schema{
										Description: `Template content (uses LF styling for line-breaks)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"template_id": &schema.Schema{
										Description: `The id of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `The type of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `The version number of this version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version_id": &schema.Schema{
							Description: `The id of this version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version_time": &schema.Schema{
							Description: `Time at which this version was committed
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplatesTemplateIDVersionsVersionIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")
	vVersionID := d.Get("version_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateVersion")
		vvTemplateID := vTemplateID.(string)
		vvVersionID := vVersionID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateVersion(vvTemplateID, vvVersionID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplateVersion", err,
				"Failure at GetTemplateVersion, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesGetTemplateVersionItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateVersion response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateVersionItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version_id"] = item.VersionID
	respItem["version"] = item.Version
	respItem["version_time"] = item.VersionTime
	respItem["regular_template"] = flattenConfigurationTemplatesGetTemplateVersionItemRegularTemplate(item.RegularTemplate)
	respItem["composite_template"] = flattenConfigurationTemplatesGetTemplateVersionItemCompositeTemplate(item.CompositeTemplate)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetTemplateVersionItemRegularTemplate(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["template_id"] = item.TemplateID
	respItem["name"] = item.Name
	respItem["project_id"] = item.ProjectID
	respItem["description"] = item.Description
	respItem["software_family"] = item.SoftwareFamily
	respItem["author"] = item.Author
	respItem["products"] = flattenConfigurationTemplatesGetTemplateVersionItemRegularTemplateProducts(item.Products)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["type"] = item.Type
	respItem["language"] = item.Language
	respItem["template_content"] = item.TemplateContent

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplateVersionItemRegularTemplateProducts(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplateProducts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_name"] = item.ProductName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplateVersionItemCompositeTemplate(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["template_id"] = item.TemplateID
	respItem["name"] = item.Name
	respItem["project_id"] = item.ProjectID
	respItem["description"] = item.Description
	respItem["software_family"] = item.SoftwareFamily
	respItem["author"] = item.Author
	respItem["products"] = flattenConfigurationTemplatesGetTemplateVersionItemCompositeTemplateProducts(item.Products)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["type"] = item.Type
	respItem["failure_policy"] = item.FailurePolicy

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplateVersionItemCompositeTemplateProducts(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplateProducts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_name"] = item.ProductName
		respItems = append(respItems, respItem)
	}
	return respItems
}
