package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Returns the device's CLIs of the ICAP intent. For detailed information about the usage of the API, please refer to the
Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. device id from intent/api/v1/network-device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"preview_activity_id": &schema.Schema{
				Description: `previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"preview_items": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_preview": &schema.Schema{
										Description: `Config Preview`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"config_type": &schema.Schema{
										Description: `Config Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"error_messages": &schema.Schema{
										Description: `Error Messages`,
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
								},
							},
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPreviewActivityID := d.Get("preview_activity_id")
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheDevicesClisOfTheICapintent")
		vvPreviewActivityID := vPreviewActivityID.(string)
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Sensors.RetrievesTheDevicesClisOfTheICapintent(vvPreviewActivityID, vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheDevicesClisOfTheICapintent", err,
				"Failure at RetrievesTheDevicesClisOfTheICapintent, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsRetrievesTheDevicesClisOfTheICapintentItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheDevicesClisOfTheICapintent response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesTheDevicesClisOfTheICapintentItem(item *dnacentersdkgo.ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["preview_items"] = flattenSensorsRetrievesTheDevicesClisOfTheICapintentItemPreviewItems(item.PreviewItems)
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsRetrievesTheDevicesClisOfTheICapintentItemPreviewItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponsePreviewItems) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["config_preview"] = item.ConfigPreview
		respItem["config_type"] = item.ConfigType
		respItem["error_messages"] = item.ErrorMessages
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
