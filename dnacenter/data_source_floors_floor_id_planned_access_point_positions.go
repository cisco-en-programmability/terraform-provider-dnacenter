package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloorsFloorIDPlannedAccessPointPositions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieve all Planned Access Points Positions designated for a specific floor
`,

		ReadContext: dataSourceFloorsFloorIDPlannedAccessPointPositionsRead,
		Schema: map[string]*schema.Schema{
			"floor_id": &schema.Schema{
				Description: `floorId path parameter. Floor Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Planned Access Point mac address.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Planned Access Point name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Minimum: 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Planned Access Point type.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Planned Access Point Id (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `Planned Access Point MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Planned Access Point Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"position": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"x": &schema.Schema{
										Description: `Planned Access Point X coordinate in feet
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"y": &schema.Schema{
										Description: `Planned Access Point Y coordinate in feet
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"z": &schema.Schema{
										Description: `Planned Access Point Z coordinate in feet
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"radios": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"antenna": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"azimuth": &schema.Schema{
													Description: `Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"elevation": &schema.Schema{
													Description: `Elevation of the antenna. The elevation range is from -90 through 90
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Antenna type for this Planned Access Point. Use */dna/intent/api/v1/maps/supported-access-points* to find supported Antennas for a particualr Planned Access Point type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"bands": &schema.Schema{
										Description: `Radio frequencies in GHz (readonly). Radio frequencies are expected to be 2.4, 5, and 6. MinItems: 1; MaxItems: 3
`,
										Type:     schema.TypeList,
										Computed: true,
									},

									"channel": &schema.Schema{
										Description: `Channel to be used by the Planned Access Point. In the context of a Planned Access Point, the channel have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Radio Id (readonly)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tx_power": &schema.Schema{
										Description: `Transmit power for the channel in Decibel milliwatts (dBm). In the context of a Planned Access Point, the txPower have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Description: `Planned Access Point type. Use *dna/intent/api/v1/maps/supported-access-points* to find the supported models
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

func dataSourceFloorsFloorIDPlannedAccessPointPositionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFloorID := d.Get("floor_id")
	vName, okName := d.GetOk("name")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPlannedAccessPointsPositionsV2")
		vvFloorID := vFloorID.(string)
		queryParams1 := dnacentersdkgo.GetPlannedAccessPointsPositionsV2QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SiteDesign.GetPlannedAccessPointsPositionsV2(vvFloorID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPlannedAccessPointsPositionsV2", err,
				"Failure at GetPlannedAccessPointsPositionsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignGetPlannedAccessPointsPositionsV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPlannedAccessPointsPositionsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetPlannedAccessPointsPositionsV2Items(items *[]dnacentersdkgo.ResponseSiteDesignGetPlannedAccessPointsPositionsV2Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["mac_address"] = item.MacAddress
		respItem["type"] = item.Type
		respItem["position"] = flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsPosition(item.Position)
		respItem["radios"] = flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsRadios(item.Radios)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsPosition(item *dnacentersdkgo.ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponsePosition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["x"] = item.X
	respItem["y"] = item.Y
	respItem["z"] = item.Z

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsRadios(items *[]dnacentersdkgo.ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadios) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["channel"] = item.Channel
		respItem["tx_power"] = item.TxPower
		respItem["antenna"] = flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsRadiosAntenna(item.Antenna)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetPlannedAccessPointsPositionsV2ItemsRadiosAntenna(item *dnacentersdkgo.ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadiosAntenna) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["azimuth"] = item.Azimuth
	respItem["elevation"] = item.Elevation

	return []map[string]interface{}{
		respItem,
	}

}
