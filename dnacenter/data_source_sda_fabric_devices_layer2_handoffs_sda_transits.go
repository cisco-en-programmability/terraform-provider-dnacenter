package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricDevicesLayer2HandoffsSdaTransits() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of layer 3 handoffs with sda transit of fabric devices that match the provided query parameters.
`,

		ReadContext: dataSourceSdaFabricDevicesLayer2HandoffsSdaTransitsRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric this device belongs to.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device ID of the fabric device.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"affinity_id_decider": &schema.Schema{
							Description: `Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"affinity_id_prime": &schema.Schema{
							Description: `Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"connected_to_internet": &schema.Schema{
							Description: `True value for this allows associated site to provide internet access to other sites through sd-access.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_multicast_over_transit_enabled": &schema.Schema{
							Description: `True value for this configures native multicast over multiple sites that are connected to an sd-access transit.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"transit_network_id": &schema.Schema{
							Description: `ID of the transit network of the layer 3 handoff sda transit.
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

func dataSourceSdaFabricDevicesLayer2HandoffsSdaTransitsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevicesLayer3HandoffsWithSdaTransit")
		queryParams1 := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams{}

		queryParams1.FabricID = vFabricID.(string)

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetFabricDevicesLayer3HandoffsWithSdaTransit(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFabricDevicesLayer3HandoffsWithSdaTransit", err,
				"Failure at GetFabricDevicesLayer3HandoffsWithSdaTransit, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetFabricDevicesLayer3HandoffsWithSdaTransitItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevicesLayer3HandoffsWithSdaTransit response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetFabricDevicesLayer3HandoffsWithSdaTransitItems(items *[]dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["fabric_id"] = item.FabricID
		respItem["transit_network_id"] = item.TransitNetworkID
		respItem["affinity_id_prime"] = item.AffinityIDPrime
		respItem["affinity_id_decider"] = item.AffinityIDDecider
		respItem["connected_to_internet"] = boolPtrToString(item.ConnectedToInternet)
		respItem["is_multicast_over_transit_enabled"] = boolPtrToString(item.IsMulticastOverTransitEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}
