package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapCaptureFiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Lists the ICAP packet capture (pcap) files matching the specified criteria. For detailed information about the usage
of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapCaptureFilesRead,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_mac": &schema.Schema{
							Description: `Ap Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_mac": &schema.Schema{
							Description: `Client Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"file_creation_timestamp": &schema.Schema{
							Description: `File Creation Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"file_name": &schema.Schema{
							Description: `File Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"file_size": &schema.Schema{
							Description: `File Size`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_updated_timestamp": &schema.Schema{
							Description: `Last Updated Timestamp`,
							Type:        schema.TypeInt,
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
		},
	}
}

func dataSourceIcapCaptureFilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vType := d.Get("type")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vApMac, okApMac := d.GetOk("ap_mac")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ListsICapPacketCaptureFilesMatchingSpecifiedCriteria")

		headerParams1 := dnacentersdkgo.ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams{}
		queryParams1 := dnacentersdkgo.ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams{}

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
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sensors.ListsICapPacketCaptureFilesMatchingSpecifiedCriteria(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ListsICapPacketCaptureFilesMatchingSpecifiedCriteria", err,
				"Failure at ListsICapPacketCaptureFilesMatchingSpecifiedCriteria, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ListsICapPacketCaptureFilesMatchingSpecifiedCriteria response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaItems(items *[]dnacentersdkgo.ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["file_name"] = item.FileName
		respItem["file_size"] = item.FileSize
		respItem["type"] = item.Type
		respItem["client_mac"] = item.ClientMac
		respItem["ap_mac"] = item.ApMac
		respItem["file_creation_timestamp"] = item.FileCreationTimestamp
		respItem["last_updated_timestamp"] = item.LastUpdatedTimestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}
