package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesAssignedToSiteID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Get site assigned network device. The items in the list are arranged in an order that corresponds with their internal
identifiers.
`,

		ReadContext: dataSourceNetworkDevicesAssignedToSiteIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Network Device Id.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_id": &schema.Schema{
							Description: `Site assigned network device Id.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id where device has been assigned.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site name hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_type": &schema.Schema{
							Description: `Type of the site where device has been assigned.
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

func dataSourceNetworkDevicesAssignedToSiteIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAssignedNetworkDevice")
		vvID := vID.(string)

		response1, restyResp1, err := client.SiteDesign.GetSiteAssignedNetworkDevice(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteAssignedNetworkDevice", err,
				"Failure at GetSiteAssignedNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetSiteAssignedNetworkDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAssignedNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetSiteAssignedNetworkDeviceItem(item *dnacentersdkgo.ResponseSiteDesignGetSiteAssignedNetworkDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_id"] = item.DeviceID
	respItem["site_id"] = item.SiteID
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["site_type"] = item.SiteType
	return []map[string]interface{}{
		respItem,
	}
}
