package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get default authentication profile from SDA Fabric
`,

		ReadContext: dataSourceSdaFabricAuthenticationProfileRead,
		Schema: map[string]*schema.Schema{
			"authenticate_template_name": &schema.Schema{
				Description: `authenticateTemplateName query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name_hierarchy": &schema.Schema{
				Description: `siteNameHierarchy query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_id": &schema.Schema{
							Description: `Authenticate Template Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaFabricAuthenticationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteNameHierarchy := d.Get("site_name_hierarchy")
	vAuthenticateTemplateName, okAuthenticateTemplateName := d.GetOk("authenticate_template_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDefaultAuthenticationProfileFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		if okAuthenticateTemplateName {
			queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName.(string)
		}

		response1, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDefaultAuthenticationProfileFromSdaFabric", err,
				"Failure at GetDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetDefaultAuthenticationProfileFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDefaultAuthenticationProfileFromSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetDefaultAuthenticationProfileFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["authenticate_template_name"] = item.AuthenticateTemplateName
	respItem["authenticate_template_id"] = item.AuthenticateTemplateID
	return []map[string]interface{}{
		respItem,
	}
}
