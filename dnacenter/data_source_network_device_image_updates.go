package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceImageUpdates() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns the list of network device image updates based on the given filter criteria
`,

		ReadContext: dataSourceNetworkDeviceImageUpdatesRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Image update started before the given time (as milliseconds since UNIX epoch)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Description: `hostName query parameter. Host name of the network device for the image update. Supports case-insensitive partial search
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Update id which is unique for each network device under the parentId
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_name": &schema.Schema{
				Description: `imageName query parameter. Software image name for the update
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
			"management_address": &schema.Schema{
				Description: `managementAddress query parameter. Management address of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device id
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
			"parent_id": &schema.Schema{
				Description: `parentId query parameter. Updates that have this parent id
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
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Image update started after the given time (as milliseconds since UNIX epoch)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Status of the image update. Available values : FAILURE, SUCCESS, IN_PROGRESS, PENDING
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"end_time": &schema.Schema{
							Description: `Image update end time (as milliseconds since UNIX epoch)
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"host_name": &schema.Schema{
							Description: `Host name of the network device for the image update
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique identifier for the image update
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_address": &schema.Schema{
							Description: `Management address of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent identifier for the image update
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Image update started after the given time (as milliseconds since UNIX epoch)
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the image update
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type of the image update
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"update_image_version": &schema.Schema{
							Description: `Software image version
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

func dataSourceNetworkDeviceImageUpdatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vParentID, okParentID := d.GetOk("parent_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vStatus, okStatus := d.GetOk("status")
	vImageName, okImageName := d.GetOk("image_name")
	vHostName, okHostName := d.GetOk("host_name")
	vManagementAddress, okManagementAddress := d.GetOk("management_address")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceImageUpdates")
		queryParams1 := dnacentersdkgo.GetNetworkDeviceImageUpdatesQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okParentID {
			queryParams1.ParentID = vParentID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okImageName {
			queryParams1.ImageName = vImageName.(string)
		}
		if okHostName {
			queryParams1.HostName = vHostName.(string)
		}
		if okManagementAddress {
			queryParams1.ManagementAddress = vManagementAddress.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetNetworkDeviceImageUpdates(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDeviceImageUpdates", err,
				"Failure at GetNetworkDeviceImageUpdates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceImageUpdates response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["parent_id"] = item.ParentID
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["status"] = item.Status
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["management_address"] = item.ManagementAddress
		respItem["host_name"] = item.HostName
		respItem["update_image_version"] = item.UpdateImageVersion
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}
