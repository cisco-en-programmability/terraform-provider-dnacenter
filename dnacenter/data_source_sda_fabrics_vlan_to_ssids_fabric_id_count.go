package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricsVLANToSSIDsFabricIDCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Returns the count of VLANs mapped to SSIDs in a Fabric Site. The 'fabricId' represents the Fabric ID of a particular
Fabric Site.
`,

		ReadContext: dataSourceSdaFabricsVLANToSSIDsFabricIDCountRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Returns the count of VLANs mapped to SSIDs in a Fabric Site
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

func dataSourceSdaFabricsVLANToSSIDsFabricIDCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite")
		vvFabricID := vFabricID.(string)

		response1, restyResp1, err := client.FabricWireless.ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite(vvFabricID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite", err,
				"Failure at ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteItem(item *dnacentersdkgo.ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
