package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapCaptureFilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves details of a specific ICAP packet capture file. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapCaptureFilesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.
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

func dataSourceIcapCaptureFilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesDetailsOfASpecificICapPacketCaptureFile")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.RetrievesDetailsOfASpecificICapPacketCaptureFileHeaderParams{}

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sensors.RetrievesDetailsOfASpecificICapPacketCaptureFile(vvID, &headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesDetailsOfASpecificICapPacketCaptureFile", err,
				"Failure at RetrievesDetailsOfASpecificICapPacketCaptureFile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesDetailsOfASpecificICapPacketCaptureFile response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileItem(item *dnacentersdkgo.ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["file_name"] = item.FileName
	respItem["file_size"] = item.FileSize
	respItem["type"] = item.Type
	respItem["client_mac"] = item.ClientMac
	respItem["ap_mac"] = item.ApMac
	respItem["file_creation_timestamp"] = item.FileCreationTimestamp
	respItem["last_updated_timestamp"] = item.LastUpdatedTimestamp
	return []map[string]interface{}{
		respItem,
	}
}
