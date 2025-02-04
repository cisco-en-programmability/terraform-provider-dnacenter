package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettingsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves the count of deployed ICAP configurations while supporting basic filtering. For detailed information about
the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsCountRead,
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
			"wlc_id": &schema.Schema{
				Description: `wlcId query parameter. The wireless controller device's UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSettingsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vCaptureType, okCaptureType := d.GetOk("capture_type")
	vCaptureStatus := d.Get("capture_status")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vAPID, okAPID := d.GetOk("apid")
	vWlcID, okWlcID := d.GetOk("wlc_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering")
		queryParams1 := dnacentersdkgo.RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams{}

		if okCaptureType {
			queryParams1.CaptureType = vCaptureType.(string)
		}
		queryParams1.CaptureStatus = vCaptureStatus.(string)

		if okClientMac {
			queryParams1.ClientMac = vClientMac.(string)
		}
		if okAPID {
			queryParams1.APID = vAPID.(string)
		}
		if okWlcID {
			queryParams1.WlcID = vWlcID.(string)
		}

		response1, restyResp1, err := client.Sensors.RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering", err,
				"Failure at RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringItem(item *dnacentersdkgo.ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
