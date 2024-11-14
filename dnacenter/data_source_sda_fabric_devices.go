package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of fabric devices that match the provided query parameters.
`,

		ReadContext: dataSourceSdaFabricDevicesRead,
		Schema: map[string]*schema.Schema{
			"device_roles": &schema.Schema{
				Description: `deviceRoles query parameter. Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
`,
				Type:     schema.TypeString,
				Optional: true,
			},
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

						"border_device_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"border_types": &schema.Schema{
										Description: `List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"layer3_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"border_priority": &schema.Schema{
													Description: `Border priority of the fabric border device.  A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"import_external_routes": &schema.Schema{
													Description: `Import external routes value of the fabric border device.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"is_default_exit": &schema.Schema{
													Description: `Is default exit value of the fabric border device.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"local_autonomous_system_number": &schema.Schema{
													Description: `BGP Local autonomous system number of the fabric border device.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"prepend_autonomous_system_count": &schema.Schema{
													Description: `Prepend autonomous system count of the fabric border device.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"device_roles": &schema.Schema{
							Description: `List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric of this fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device ID of the fabric device.
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

func dataSourceSdaFabricDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vDeviceRoles, okDeviceRoles := d.GetOk("device_roles")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevices")
		queryParams1 := dnacentersdkgo.GetFabricDevicesQueryParams{}

		queryParams1.FabricID = vFabricID.(string)

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okDeviceRoles {
			queryParams1.DeviceRoles = vDeviceRoles.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetFabricDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFabricDevices", err,
				"Failure at GetFabricDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetFabricDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetFabricDevicesItems(items *[]dnacentersdkgo.ResponseSdaGetFabricDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["fabric_id"] = item.FabricID
		respItem["device_roles"] = item.DeviceRoles
		respItem["border_device_settings"] = flattenSdaGetFabricDevicesItemsBorderDeviceSettings(item.BorderDeviceSettings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetFabricDevicesItemsBorderDeviceSettings(item *dnacentersdkgo.ResponseSdaGetFabricDevicesResponseBorderDeviceSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["border_types"] = item.BorderTypes
	respItem["layer3_settings"] = flattenSdaGetFabricDevicesItemsBorderDeviceSettingsLayer3Settings(item.Layer3Settings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetFabricDevicesItemsBorderDeviceSettingsLayer3Settings(item *dnacentersdkgo.ResponseSdaGetFabricDevicesResponseBorderDeviceSettingsLayer3Settings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["local_autonomous_system_number"] = item.LocalAutonomousSystemNumber
	respItem["is_default_exit"] = boolPtrToString(item.IsDefaultExit)
	respItem["import_external_routes"] = boolPtrToString(item.ImportExternalRoutes)
	respItem["border_priority"] = item.BorderPriority
	respItem["prepend_autonomous_system_count"] = item.PrependAutonomousSystemCount

	return []map[string]interface{}{
		respItem,
	}

}
