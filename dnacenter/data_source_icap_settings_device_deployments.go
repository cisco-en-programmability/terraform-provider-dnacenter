package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettingsDeviceDeployments() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves ICAP configuration deployment status(s) per device based on filter criteria. For detailed information about
the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsDeviceDeploymentsRead,
		Schema: map[string]*schema.Schema{
			"deploy_activity_id": &schema.Schema{
				Description: `deployActivityId query parameter. activity from the /deploy task response
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
			"network_device_ids": &schema.Schema{
				Description: `networkDeviceIds query parameter. device ids, retrievable from the id attribute in intent/api/v1/network-device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"config_group_name": &schema.Schema{
							Description: `Config Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"config_group_version": &schema.Schema{
							Description: `Config Group Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"deploy_activity_id": &schema.Schema{
							Description: `Deploy Activity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"error": &schema.Schema{
							Description: `Error`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Computed:    true,
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

func dataSourceIcapSettingsDeviceDeploymentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeployActivityID, okDeployActivityID := d.GetOk("deploy_activity_id")
	vNetworkDeviceIDs, okNetworkDeviceIDs := d.GetOk("network_device_ids")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceDeploymentStatus")
		queryParams1 := dnacentersdkgo.GetDeviceDeploymentStatusQueryParams{}

		if okDeployActivityID {
			queryParams1.DeployActivityID = vDeployActivityID.(string)
		}
		if okNetworkDeviceIDs {
			queryParams1.NetworkDeviceIDs = vNetworkDeviceIDs.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Sensors.GetDeviceDeploymentStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceDeploymentStatus", err,
				"Failure at GetDeviceDeploymentStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsGetDeviceDeploymentStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceDeploymentStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsGetDeviceDeploymentStatusItems(items *[]dnacentersdkgo.ResponseSensorsGetDeviceDeploymentStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["deploy_activity_id"] = item.DeployActivityID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["config_group_name"] = item.ConfigGroupName
		respItem["config_group_version"] = item.ConfigGroupVersion
		respItem["status"] = item.Status
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["error"] = flattenSensorsGetDeviceDeploymentStatusItemsError(item.Error)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsGetDeviceDeploymentStatusItemsError(item *dnacentersdkgo.ResponseSensorsGetDeviceDeploymentStatusResponseError) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
