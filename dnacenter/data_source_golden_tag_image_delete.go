package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGoldenTagImageDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Software Image Management (SWIM).

- Remove golden tag. Set siteId as -1 for Global site.
`,

		ReadContext: dataSourceGoldenTagImageDeleteRead,
		Schema: map[string]*schema.Schema{
			"device_family_identifier": &schema.Schema{
				Description: `deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"device_role": &schema.Schema{
				Description: `deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": &schema.Schema{
				Description: `imageId path parameter. Image Id in uuid format.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.
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
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5 
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
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

func dataSourceGoldenTagImageDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vDeviceFamilyIDentifier := d.Get("device_family_identifier")
	vDeviceRole := d.Get("device_role")
	vImageID := d.Get("image_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveGoldenTagForImage")
		vvSiteID := vSiteID.(string)
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier.(string)
		vvDeviceRole := vDeviceRole.(string)
		vvImageID := vImageID.(string)

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RemoveGoldenTagForImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveGoldenTagForImage", err,
				"Failure at RemoveGoldenTagForImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimRemoveGoldenTagForImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveGoldenTagForImage response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRemoveGoldenTagForImageItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
