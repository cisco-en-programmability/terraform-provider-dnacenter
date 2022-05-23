package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePlannedAccessPoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Provides a list of Planned Access Points for the Floor it is requested for
`,

		ReadContext: dataSourcePlannedAccessPointsRead,
		Schema: map[string]*schema.Schema{
			"floor_id": &schema.Schema{
				Description: `floorId path parameter. Floor Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"radios": &schema.Schema{
				Description: `radios query parameter. inlcude planned radio details
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"create_date": &schema.Schema{
										Description: `Create Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"domain": &schema.Schema{
										Description: `Domain`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"heirarchy_name": &schema.Schema{
										Description: `Heirarchy Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"macaddress": &schema.Schema{
										Description: `Macaddress`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"source": &schema.Schema{
										Description: `Source`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type_string": &schema.Schema{
										Description: `Type String`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"is_sensor": &schema.Schema{
							Description: `Is Sensor`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `Location`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"position": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"x": &schema.Schema{
										Description: `X`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"y": &schema.Schema{
										Description: `Y`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"z": &schema.Schema{
										Description: `Z`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},

						"radio_count": &schema.Schema{
							Description: `Radio Count`,
							Type:        schema.TypeInt,
							Computed:    true,
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

												"azimuth_angle": &schema.Schema{
													Description: `Azimuth Angle`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"elevation_angle": &schema.Schema{
													Description: `Elevation Angle`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"gain": &schema.Schema{
													Description: `Gain`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"mode": &schema.Schema{
													Description: `Mode`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"channel": &schema.Schema{
													Description: `Channel`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"channel_string": &schema.Schema{
													Description: `Channel String`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"if_mode": &schema.Schema{
													Description: `If Mode`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"if_type_string": &schema.Schema{
													Description: `If Type String`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"if_type_subband": &schema.Schema{
													Description: `If Type Subband`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_uuid": &schema.Schema{
													Description: `Instance Uuid`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"slot_id": &schema.Schema{
													Description: `Slot Id`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},

									"is_sensor": &schema.Schema{
										Description: `Is Sensor`,

										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePlannedAccessPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFloorID := d.Get("floor_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vRadios, okRadios := d.GetOk("radios")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPlannedAccessPointsForFloor")
		vvFloorID := vFloorID.(string)
		queryParams1 := dnacentersdkgo.GetPlannedAccessPointsForFloorQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okRadios {
			queryParams1.Radios = vRadios.(bool)
		}

		response1, restyResp1, err := client.Devices.GetPlannedAccessPointsForFloor(vvFloorID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPlannedAccessPointsForFloor", err,
				"Failure at GetPlannedAccessPointsForFloor, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetPlannedAccessPointsForFloorItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPlannedAccessPointsForFloor response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetPlannedAccessPointsForFloorItems(items *[]dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attributes"] = flattenDevicesGetPlannedAccessPointsForFloorItemsAttributes(item.Attributes)
		respItem["location"] = flattenDevicesGetPlannedAccessPointsForFloorItemsLocation(item.Location)
		respItem["position"] = flattenDevicesGetPlannedAccessPointsForFloorItemsPosition(item.Position)
		respItem["radio_count"] = item.RadioCount
		respItem["radios"] = flattenDevicesGetPlannedAccessPointsForFloorItemsRadios(item.Radios)
		respItem["is_sensor"] = boolPtrToString(item.IsSensor)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetPlannedAccessPointsForFloorItemsAttributes(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["name"] = item.Name
	respItem["type_string"] = item.TypeString
	respItem["domain"] = item.Domain
	respItem["heirarchy_name"] = item.HeirarchyName
	respItem["source"] = item.Source
	respItem["create_date"] = item.CreateDate
	respItem["macaddress"] = flattenDevicesGetPlannedAccessPointsForFloorItemsAttributesMacaddress(item.Macaddress)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsAttributesMacaddress(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributesMacaddress) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsLocation(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsPosition(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition) []map[string]interface{} {
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

func flattenDevicesGetPlannedAccessPointsForFloorItemsRadios(items *[]dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attributes"] = flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributes(item.Attributes)
		respItem["antenna"] = flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAntenna(item.Antenna)
		respItem["is_sensor"] = boolPtrToString(item.IsSensor)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributes(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["slot_id"] = item.SlotID
	respItem["if_type_string"] = item.IfTypeString
	respItem["if_type_subband"] = item.IfTypeSubband
	respItem["channel"] = flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributesChannel(item.Channel)
	respItem["channel_string"] = flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributesChannelString(item.ChannelString)
	respItem["if_mode"] = item.IfMode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributesChannel(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannel) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAttributesChannelString(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannelString) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetPlannedAccessPointsForFloorItemsRadiosAntenna(item *dnacentersdkgo.ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type"] = item.Type
	respItem["mode"] = item.Mode
	respItem["azimuth_angle"] = item.AzimuthAngle
	respItem["elevation_angle"] = item.ElevationAngle
	respItem["gain"] = item.Gain

	return []map[string]interface{}{
		respItem,
	}

}
