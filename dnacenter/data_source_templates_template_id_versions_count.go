package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesTemplateIDVersionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get the count of a template's version information.
`,

		ReadContext: dataSourceTemplatesTemplateIDVersionsCountRead,
		Schema: map[string]*schema.Schema{
			"latest_version": &schema.Schema{
				Description: `latestVersion query parameter. Filter response to only include the latest version of a template
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"template_id": &schema.Schema{
				Description: `templateId path parameter. The id of the template to get versions of, retrieveable from *GET /dna/intent/api/v1/templates*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"version_number": &schema.Schema{
				Description: `versionNumber query parameter. Filter response to only get the template version that matches this version number
`,
				Type:     schema.TypeInt,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplatesTemplateIDVersionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")
	vVersionNumber, okVersionNumber := d.GetOk("version_number")
	vLatestVersion, okLatestVersion := d.GetOk("latest_version")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateVersionsCount")
		vvTemplateID := vTemplateID.(string)
		queryParams1 := dnacentersdkgo.GetTemplateVersionsCountQueryParams{}

		if okVersionNumber {
			queryParams1.VersionNumber = vVersionNumber.(int)
		}
		if okLatestVersion {
			queryParams1.LatestVersion = vLatestVersion.(bool)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateVersionsCount(vvTemplateID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplateVersionsCount", err,
				"Failure at GetTemplateVersionsCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesGetTemplateVersionsCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateVersionsCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateVersionsCountItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionsCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
