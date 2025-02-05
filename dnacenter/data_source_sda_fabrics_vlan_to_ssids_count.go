package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricsVLANToSSIDsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Return the count of all the fabric site which has SSID to IP Pool mapping
`,

		ReadContext: dataSourceSdaFabricsVLANToSSIDsCountRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Return the count of all the fabric site which has SSID to IP Pool mapping
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

func dataSourceSdaFabricsVLANToSSIDsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping")

		response1, restyResp1, err := client.FabricWireless.ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping", err,
				"Failure at ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingItem(item *dnacentersdkgo.ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
