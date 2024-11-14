package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricDevicesLayer2HandoffsIPTransitsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.
`,

		ReadContext: dataSourceSdaFabricDevicesLayer2HandoffsIPTransitsCountRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric this device belongs to.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device ID of the fabric device.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Number of fabric device layer 3 handoffs with IP transit.
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

func dataSourceSdaFabricDevicesLayer2HandoffsIPTransitsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevicesLayer3HandoffsWithIPTransitCount")
		queryParams1 := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitCountQueryParams{}

		queryParams1.FabricID = vFabricID.(string)

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}

		response1, restyResp1, err := client.Sda.GetFabricDevicesLayer3HandoffsWithIPTransitCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFabricDevicesLayer3HandoffsWithIPTransitCount", err,
				"Failure at GetFabricDevicesLayer3HandoffsWithIPTransitCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevicesLayer3HandoffsWithIPTransitCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountItem(item *dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
