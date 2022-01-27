package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoldenTagImageDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Get golden tag status of an image. Set siteId as -1 for Global site.
`,

		ReadContext: dataSourceGoldenTagImageDetailsRead,
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

						"device_role": &schema.Schema{
							Description: `Device Role. Possible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_id": &schema.Schema{
							Description: `Inherited Site Id. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the uuid of the site from where the tag is inherited. In case the golden tag is inherited from the Global site the value will be "-1".
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_name": &schema.Schema{
							Description: `Inherited Site Name. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the name of the site from where the tag is inherited.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tagged_golden": &schema.Schema{
							Description: `Tagged Golden.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGoldenTagImageDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vDeviceFamilyIDentifier := d.Get("device_family_identifier")
	vDeviceRole := d.Get("device_role")
	vImageID := d.Get("image_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGoldenTagStatusOfAnImage")
		vvSiteID := vSiteID.(string)
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier.(string)
		vvDeviceRole := vDeviceRole.(string)
		vvImageID := vImageID.(string)

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGoldenTagStatusOfAnImage", err,
				"Failure at GetGoldenTagStatusOfAnImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGoldenTagStatusOfAnImage response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_role"] = item.DeviceRole
	respItem["tagged_golden"] = boolPtrToString(item.TaggedGolden)
	respItem["inherited_site_name"] = item.InheritedSiteName
	respItem["inherited_site_id"] = item.InheritedSiteID
	return []map[string]interface{}{
		respItem,
	}
}
