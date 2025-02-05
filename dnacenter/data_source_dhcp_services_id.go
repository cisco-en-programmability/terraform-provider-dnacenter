package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDhcpServicesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieves the details of the DHCP Service matching the given id. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceDhcpServicesIDRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Unique id of the DHCP Service. It is the combination of DHCP Server IP (*serverIp*) and Device UUID (*deviceId*) separated by underscore (*_*). Example: If *serverIp* is *10.76.81.33* and *deviceId* is *6bef213c-19ca-4170-8375-b694e251101c*, then the *id* would be *10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
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

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_name": &schema.Schema{
							Description: `Device Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_hierarchy": &schema.Schema{
							Description: `Device Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_hierarchy_id": &schema.Schema{
							Description: `Device Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_site_id": &schema.Schema{
							Description: `Device Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"discover_offer_latency": &schema.Schema{
							Description: `Discover Offer Latency`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"failed_transactions": &schema.Schema{
							Description: `Failed Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"latency": &schema.Schema{
							Description: `Latency`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"request_acknowledge_latency": &schema.Schema{
							Description: `Request Acknowledge Latency`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"server_ip": &schema.Schema{
							Description: `Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"successful_transactions": &schema.Schema{
							Description: `Successful Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transactions": &schema.Schema{
							Description: `Transactions`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDhcpServicesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService(vvID, &headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService", err,
				"Failure at RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceItem(item *dnacentersdkgo.ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["server_ip"] = item.ServerIP
	respItem["device_id"] = item.DeviceID
	respItem["device_name"] = item.DeviceName
	respItem["device_family"] = item.DeviceFamily
	respItem["device_site_hierarchy"] = item.DeviceSiteHierarchy
	respItem["device_site_id"] = item.DeviceSiteID
	respItem["device_site_hierarchy_id"] = item.DeviceSiteHierarchyID
	respItem["transactions"] = item.Transactions
	respItem["failed_transactions"] = item.FailedTransactions
	respItem["successful_transactions"] = item.SuccessfulTransactions
	respItem["latency"] = item.Latency
	respItem["discover_offer_latency"] = item.DiscoverOfferLatency
	respItem["request_acknowledge_latency"] = item.RequestAcknowledgeLatency
	return []map[string]interface{}{
		respItem,
	}
}
