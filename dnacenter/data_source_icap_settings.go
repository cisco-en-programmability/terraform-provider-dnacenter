package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves deployed ICAP configurations while supporting basic filtering. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsRead,
		Schema: map[string]*schema.Schema{
			"apid": &schema.Schema{
				Description: `apId query parameter. The AP device's UUID.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"capture_status": &schema.Schema{
				Description: `captureStatus query parameter. Catalyst Center ICAP status
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"capture_type": &schema.Schema{
				Description: `captureType query parameter. Catalyst Center ICAP type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_mac": &schema.Schema{
				Description: `clientMac query parameter. The client device MAC address in ICAP configuration
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"wlc_id": &schema.Schema{
				Description: `wlcId query parameter. The wireless controller device's UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"activity_id": &schema.Schema{
							Description: `Activity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"apid": &schema.Schema{
							Description: `Ap Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"capture_type": &schema.Schema{
							Description: `Capture Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_mac": &schema.Schema{
							Description: `Client Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"create_time": &schema.Schema{
							Description: `Create Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"deployed_id": &schema.Schema{
							Description: `Deployed Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"disable_activity_id": &schema.Schema{
							Description: `Disable Activity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duration_in_mins": &schema.Schema{
							Description: `Duration In Mins`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ota_band": &schema.Schema{
							Description: `Ota Band`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ota_channel": &schema.Schema{
							Description: `Ota Channel`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"ota_channel_width": &schema.Schema{
							Description: `Ota Channel Width`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"slots": &schema.Schema{
							Description: `Slots`,
							Type:        schema.TypeList,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wlc_id": &schema.Schema{
							Description: `Wlc Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vCaptureStatus := d.Get("capture_status")
	vCaptureType, okCaptureType := d.GetOk("capture_type")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vAPID, okAPID := d.GetOk("apid")
	vWlcID, okWlcID := d.GetOk("wlc_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering")
		queryParams1 := dnacentersdkgo.RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams{}

		queryParams1.CaptureStatus = vCaptureStatus.(string)

		if okCaptureType {
			queryParams1.CaptureType = vCaptureType.(string)
		}
		if okClientMac {
			queryParams1.ClientMac = vClientMac.(string)
		}
		if okAPID {
			queryParams1.APID = vAPID.(string)
		}
		if okWlcID {
			queryParams1.WlcID = vWlcID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sensors.RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering", err,
				"Failure at RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["slots"] = item.Slots
		respItem["ota_band"] = item.OtaBand
		respItem["ota_channel"] = item.OtaChannel
		respItem["ota_channel_width"] = item.OtaChannelWidth
		respItem["id"] = item.ID
		respItem["deployed_id"] = item.DeployedID
		respItem["disable_activity_id"] = item.DisableActivityID
		respItem["activity_id"] = item.ActivityID
		respItem["capture_type"] = item.CaptureType
		respItem["apid"] = item.APID
		respItem["wlc_id"] = item.WlcID
		respItem["client_mac"] = item.ClientMac
		respItem["create_time"] = item.CreateTime
		respItem["end_time"] = item.EndTime
		respItem["duration_in_mins"] = item.DurationInMins
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
