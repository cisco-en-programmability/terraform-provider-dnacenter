package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceImageUpdatesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Returns the count of network device image updates based on the given filter criteria
`,

		ReadContext: dataSourceNetworkDeviceImageUpdatesCountRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Image update started before the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Description: `hostName query parameter. Host name of the network device for the image update. Supports case-insensitive partial search.
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
			"parent_id": &schema.Schema{
				Description: `parentId query parameter. Updates that have this parent id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Image update started after the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Status of the image update. Available values: FAILURE, SUCCESS, IN_PROGRESS, PENDING
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Reports a count, for example, a total count of records in a given resource.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceImageUpdatesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountOfNetworkDeviceImageUpdates")
		queryParams1 := dnacentersdkgo.CountOfNetworkDeviceImageUpdatesQueryParams{}

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

		response1, restyResp1, err := client.SoftwareImageManagementSwim.CountOfNetworkDeviceImageUpdates(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountOfNetworkDeviceImageUpdates", err,
				"Failure at CountOfNetworkDeviceImageUpdates, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountOfNetworkDeviceImageUpdates response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
