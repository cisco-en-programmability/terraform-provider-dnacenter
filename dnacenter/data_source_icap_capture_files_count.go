package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapCaptureFilesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves the total number of packet capture files matching the specified criteria. For detailed information about the
usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapCaptureFilesCountRead,
		Schema: map[string]*schema.Schema{
			"ap_mac": &schema.Schema{
				Description: `apMac query parameter. The base radio macAddress of the access point
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_mac": &schema.Schema{
				Description: `clientMac query parameter. The macAddress of client
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Capture Type
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapCaptureFilesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vType := d.Get("type")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vApMac, okApMac := d.GetOk("ap_mac")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria")

		headerParams1 := dnacentersdkgo.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams{}

		queryParams1.Type = vType.(string)

		if okClientMac {
			queryParams1.ClientMac = vClientMac.(string)
		}
		if okApMac {
			queryParams1.ApMac = vApMac.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sensors.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria", err,
				"Failure at RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaItem(item *dnacentersdkgo.ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
