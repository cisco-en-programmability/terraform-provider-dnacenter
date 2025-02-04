package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloorsFloorIDAccessPointPositions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieve all Access Points positions assigned for a specific floor
`,

		ReadContext: dataSourceFloorsFloorIDAccessPointPositionsRead,
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
				Description: `macAddress query parameter. Access Point mac address.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": &schema.Schema{
				Description: `model query parameter. Access Point model.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Access Point name.
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
				Description: `type query parameter. Access Point type.
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
							Description: `Access Point Id (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `Access Point MAC address (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"model": &schema.Schema{
							Description: `Access Point model (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Access Point Name (readonly)
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
										Description: `Access Point X coordinate in feet
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"y": &schema.Schema{
										Description: `Access Point Y coordinate in feet
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"z": &schema.Schema{
										Description: `Access Point Z coordinate in feet
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
													Description: `Antenna type for this Access Point. Use */dna/intent/api/v1/maps/supported-access-points* to find supported Antennas for a particualr Access Point model.
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
										Description: `Channel to be used by the Access Point (readonly). The value gets updated only every 24 hours
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
										Description: `Transmit power for the channel in Decibel milliwatts (dBm) (readonly). The value gets updated only every 24 hours
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Description: `Access Point type (readonly)
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

func dataSourceFloorsFloorIDAccessPointPositionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFloorID := d.Get("floor_id")
	vName, okName := d.GetOk("name")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vType, okType := d.GetOk("type")
	vModel, okModel := d.GetOk("model")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointsPositionsV2")
		vvFloorID := vFloorID.(string)
		queryParams1 := dnacentersdkgo.GetAccessPointsPositionsV2QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okModel {
			queryParams1.Model = vModel.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SiteDesign.GetAccessPointsPositionsV2(vvFloorID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointsPositionsV2", err,
				"Failure at GetAccessPointsPositionsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignGetAccessPointsPositionsV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointsPositionsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetAccessPointsPositionsV2Items(items *[]dnacentersdkgo.ResponseSiteDesignGetAccessPointsPositionsV2Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["mac_address"] = item.MacAddress
		respItem["model"] = item.Model
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["position"] = flattenSiteDesignGetAccessPointsPositionsV2ItemsPosition(item.Position)
		respItem["radios"] = flattenSiteDesignGetAccessPointsPositionsV2ItemsRadios(item.Radios)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetAccessPointsPositionsV2ItemsPosition(item *dnacentersdkgo.ResponseSiteDesignGetAccessPointsPositionsV2ResponsePosition) []map[string]interface{} {
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

func flattenSiteDesignGetAccessPointsPositionsV2ItemsRadios(items *[]dnacentersdkgo.ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadios) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["bands"] = item.Bands
		respItem["channel"] = item.Channel
		respItem["tx_power"] = item.TxPower
		respItem["antenna"] = flattenSiteDesignGetAccessPointsPositionsV2ItemsRadiosAntenna(item.Antenna)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetAccessPointsPositionsV2ItemsRadiosAntenna(item *dnacentersdkgo.ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadiosAntenna) []map[string]interface{} {
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
