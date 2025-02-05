package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersManagedApLocationsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves the count of Managed AP locations, including Primary Managed AP Locations, Secondary Managed AP Locations,
and Anchor Managed AP Locations, associated with the specific Wireless Controller.
`,

		ReadContext: dataSourceWirelessControllersManagedApLocationsCountRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"anchor_managed_ap_locations_count": &schema.Schema{
							Description: `The count of the Anchor managed ap  locations.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"primary_managed_ap_locations_count": &schema.Schema{
							Description: `The count of the Primary managed ap locations.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"secondary_managed_ap_locations_count": &schema.Schema{
							Description: `The count of the Secondary managed ap locations.
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

func dataSourceWirelessControllersManagedApLocationsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetManagedApLocationsCountForSpecificWirelessController")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Wireless.GetManagedApLocationsCountForSpecificWirelessController(vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetManagedApLocationsCountForSpecificWirelessController", err,
				"Failure at GetManagedApLocationsCountForSpecificWirelessController, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetManagedApLocationsCountForSpecificWirelessControllerItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetManagedApLocationsCountForSpecificWirelessController response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetManagedApLocationsCountForSpecificWirelessControllerItem(item *dnacentersdkgo.ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["primary_managed_ap_locations_count"] = item.PrimaryManagedApLocationsCount
	respItem["secondary_managed_ap_locations_count"] = item.SecondaryManagedApLocationsCount
	respItem["anchor_managed_ap_locations_count"] = item.AnchorManagedApLocationsCount
	return []map[string]interface{}{
		respItem,
	}
}
