package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSwimImportLocal() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions
are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2.
Upload the file to the **file** form data field
`,

		ReadContext: dataSourceSwimImportLocalRead,
		Schema: map[string]*schema.Schema{
			"file_name": &schema.Schema{
				Description: `File name.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_path": &schema.Schema{
				Description: `File absolute path.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"is_third_party": &schema.Schema{
				Description: `isThirdParty query parameter. Third party Image check
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"third_party_application_type": &schema.Schema{
				Description: `thirdPartyApplicationType query parameter. Third Party Application Type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"third_party_image_family": &schema.Schema{
				Description: `thirdPartyImageFamily query parameter. Third Party image family
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"third_party_vendor": &schema.Schema{
				Description: `thirdPartyVendor query parameter. Third Party Vendor
`,
				Type:     schema.TypeString,
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

func dataSourceSwimImportLocalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vIsThirdParty, okIsThirdParty := d.GetOk("is_third_party")
	vThirdPartyVendor, okThirdPartyVendor := d.GetOk("third_party_vendor")
	vThirdPartyImageFamily, okThirdPartyImageFamily := d.GetOk("third_party_image_family")
	vThirdPartyApplicationType, okThirdPartyApplicationType := d.GetOk("third_party_application_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ImportLocalSoftwareImage")
		queryParams1 := dnacentersdkgo.ImportLocalSoftwareImageQueryParams{}

		if okIsThirdParty {
			queryParams1.IsThirdParty = vIsThirdParty.(bool)
		}
		if okThirdPartyVendor {
			queryParams1.ThirdPartyVendor = vThirdPartyVendor.(string)
		}
		if okThirdPartyImageFamily {
			queryParams1.ThirdPartyImageFamily = vThirdPartyImageFamily.(string)
		}
		if okThirdPartyApplicationType {
			queryParams1.ThirdPartyApplicationType = vThirdPartyApplicationType.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ImportLocalSoftwareImage(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportLocalSoftwareImage", err,
				"Failure at ImportLocalSoftwareImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimImportLocalSoftwareImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportLocalSoftwareImage response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimImportLocalSoftwareImageItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimImportLocalSoftwareImageResponse) []map[string]interface{} {
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
