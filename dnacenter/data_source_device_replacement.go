package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceReplacement() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Replacement.

- Get list of replacement devices with replacement details and it can filter replacement devices based on Faulty Device
Name,Faulty Device Platform, Replacement Device Platform, Faulty Device Serial Number,Replacement Device Serial Number,
Device Replacement status, Product Family.
`,

		ReadContext: dataSourceDeviceReplacementRead,
		Schema: map[string]*schema.Schema{
			"family": &schema.Schema{
				Description: `family query parameter. List of families[Routers, Switches and Hubs, AP]
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"faulty_device_name": &schema.Schema{
				Description: `faultyDeviceName query parameter. Faulty Device Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"faulty_device_platform": &schema.Schema{
				Description: `faultyDevicePlatform query parameter. Faulty Device Platform
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"faulty_device_serial_number": &schema.Schema{
				Description: `faultyDeviceSerialNumber query parameter. Faulty Device Serial Number
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"replacement_device_platform": &schema.Schema{
				Description: `replacementDevicePlatform query parameter. Replacement Device Platform
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_device_serial_number": &schema.Schema{
				Description: `replacementDeviceSerialNumber query parameter. Replacement Device Serial Number
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_status": &schema.Schema{
				Description: `replacementStatus query parameter. Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. SortBy this field. SortBy is mandatory when order is used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Order on displayName[ASC,DESC]
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_time": &schema.Schema{
							Description: `Date and time of marking the device for replacement
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"family": &schema.Schema{
							Description: `Faulty device family
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"faulty_device_id": &schema.Schema{
							Description: `Unique identifier of the faulty device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"faulty_device_name": &schema.Schema{
							Description: `Faulty device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"faulty_device_platform": &schema.Schema{
							Description: `Faulty device platform
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"faulty_device_serial_number": &schema.Schema{
							Description: `Faulty device serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique identifier of the device replacement resource
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"neighbour_device_id": &schema.Schema{
							Description: `Unique identifier of the neighbor device to create the DHCP server
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_readiness_task_id": &schema.Schema{
							Description: `Unique identifier of network readiness task
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"readinesscheck_task_id": &schema.Schema{
							Description: `Unique identifier of the readiness check task for the replacement device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replacement_device_platform": &schema.Schema{
							Description: `Replacement device platform
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replacement_device_serial_number": &schema.Schema{
							Description: `Replacement device serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replacement_status": &schema.Schema{
							Description: `Device Replacement status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replacement_time": &schema.Schema{
							Description: `Date and time of device replacement
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"workflow_failed_step": &schema.Schema{
							Description: `Step in which the device replacement failed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"workflow_id": &schema.Schema{
							Description: `Unique identifier of the device replacement workflow
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

func dataSourceDeviceReplacementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFaultyDeviceName, okFaultyDeviceName := d.GetOk("faulty_device_name")
	vFaultyDevicePlatform, okFaultyDevicePlatform := d.GetOk("faulty_device_platform")
	vReplacementDevicePlatform, okReplacementDevicePlatform := d.GetOk("replacement_device_platform")
	vFaultyDeviceSerialNumber, okFaultyDeviceSerialNumber := d.GetOk("faulty_device_serial_number")
	vReplacementDeviceSerialNumber, okReplacementDeviceSerialNumber := d.GetOk("replacement_device_serial_number")
	vReplacementStatus, okReplacementStatus := d.GetOk("replacement_status")
	vFamily, okFamily := d.GetOk("family")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnListOfReplacementDevicesWithReplacementDetails")
		queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

		if okFaultyDeviceName {
			queryParams1.FaultyDeviceName = vFaultyDeviceName.(string)
		}
		if okFaultyDevicePlatform {
			queryParams1.FaultyDevicePlatform = vFaultyDevicePlatform.(string)
		}
		if okReplacementDevicePlatform {
			queryParams1.ReplacementDevicePlatform = vReplacementDevicePlatform.(string)
		}
		if okFaultyDeviceSerialNumber {
			queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber.(string)
		}
		if okReplacementDeviceSerialNumber {
			queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber.(string)
		}
		if okReplacementStatus {
			queryParams1.ReplacementStatus = interfaceToSliceString(vReplacementStatus)
		}
		if okFamily {
			queryParams1.Family = interfaceToSliceString(vFamily)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		response1, restyResp1, err := client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnListOfReplacementDevicesWithReplacementDetails", err,
				"Failure at ReturnListOfReplacementDevicesWithReplacementDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnListOfReplacementDevicesWithReplacementDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsItems(items *[]dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["creation_time"] = item.CreationTime
		respItem["family"] = item.Family
		respItem["faulty_device_id"] = item.FaultyDeviceID
		respItem["faulty_device_name"] = item.FaultyDeviceName
		respItem["faulty_device_platform"] = item.FaultyDevicePlatform
		respItem["faulty_device_serial_number"] = item.FaultyDeviceSerialNumber
		respItem["id"] = item.ID
		respItem["neighbour_device_id"] = item.NeighbourDeviceID
		respItem["network_readiness_task_id"] = item.NetworkReadinessTaskID
		respItem["replacement_device_platform"] = item.ReplacementDevicePlatform
		respItem["replacement_device_serial_number"] = item.ReplacementDeviceSerialNumber
		respItem["replacement_status"] = item.ReplacementStatus
		respItem["replacement_time"] = item.ReplacementTime
		respItem["workflow_id"] = item.WorkflowID
		respItem["workflow_failed_step"] = item.WorkflowFailedStep
		respItem["readinesscheck_task_id"] = item.ReadinesscheckTaskID
		respItems = append(respItems, respItem)
	}
	return respItems
}
