package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceReplacements() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Replacement.

- Retrieve the list of device replacements with replacement details. Filters can be applied based on faulty device name,
faulty device platform, faulty device serial number, replacement device platform, replacement device serial number,
device replacement status, device family.
`,

		ReadContext: dataSourceNetworkDeviceReplacementsRead,
		Schema: map[string]*schema.Schema{
			"family": &schema.Schema{
				Description: `family query parameter. Faulty device family.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"faulty_device_name": &schema.Schema{
				Description: `faultyDeviceName query parameter. Faulty device name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"faulty_device_platform": &schema.Schema{
				Description: `faultyDevicePlatform query parameter. Faulty device platform.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"faulty_device_serial_number": &schema.Schema{
				Description: `faultyDeviceSerialNumber query parameter. Faulty device serial number.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Maximum value can be 500.
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
			"replacement_device_platform": &schema.Schema{
				Description: `replacementDevicePlatform query parameter. Replacement device platform.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_device_serial_number": &schema.Schema{
				Description: `replacementDeviceSerialNumber query parameter. Replacement device serial number.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_status": &schema.Schema{
				Description: `replacementStatus query parameter. Device replacement status. Available values : MARKED_FOR_REPLACEMENT, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED, READY_FOR_REPLACEMENT, REPLACEMENT_SCHEDULED, REPLACEMENT_IN_PROGRESS, REPLACED, ERROR. Replacement status: 'MARKED_FOR_REPLACEMENT' The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' Device replacement is currently in progress. 'REPLACED' Device replacement was successful. 'ERROR' Device replacement has failed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by. Available values : id, creationTime, family, faultyDeviceId, fautyDeviceName, faultyDevicePlatform, faultyDeviceSerialNumber, replacementDevicePlatform, replacementDeviceSerialNumber, replacementTime.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Whether ascending or descending order should be used to sort the response. Available values : ASC, DESC
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
							Description: `Time of marking the device for replacement in Unix epoch time in milliseconds
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
							Description: `Faulty device id
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

						"neighbor_device_id": &schema.Schema{
							Description: `Unique identifier of the neighbor device to create the DHCP server
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
							Description: `Device Replacement status. 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replacement_time": &schema.Schema{
							Description: `The Unix epoch time in milliseconds at which the device was replaced successfully
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"workflow": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Completion time of the workflow in Unix epoch time in milliseconds
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Workflow id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the workflow
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Start time of the workflow in Unix epoch time in milliseconds
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"steps": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_time": &schema.Schema{
													Description: `Completion time of the workflow step in Unix epoch time in milliseconds
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Workflow step name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"start_time": &schema.Schema{
													Description: `Start time of the workflow step in Unix epoch time in milliseconds
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"status": &schema.Schema{
													Description: `Workflow step status. 'INIT' - Workflow step has not started execution. 'RUNNING' - Workflow step is currently in progress. 'SUCCESS' - Workflow step completed successfully. 'FAILED' - Workflow step completed with failure. 'ABORTED' - Workflow step aborted execution due to failure of the previous step. 'TIMEOUT' - Workflow step timedout to complete execution.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"status_message": &schema.Schema{
													Description: `Detailed status message for the step
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"workflow_status": &schema.Schema{
										Description: `Workflow status. 'RUNNING' - Workflow is currently in progress. 'SUCCESS' - Workflow completed successfully. 'FAILED' - Workflow completed with failure.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceReplacementsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFamily, okFamily := d.GetOk("family")
	vFaultyDeviceName, okFaultyDeviceName := d.GetOk("faulty_device_name")
	vFaultyDevicePlatform, okFaultyDevicePlatform := d.GetOk("faulty_device_platform")
	vFaultyDeviceSerialNumber, okFaultyDeviceSerialNumber := d.GetOk("faulty_device_serial_number")
	vReplacementDevicePlatform, okReplacementDevicePlatform := d.GetOk("replacement_device_platform")
	vReplacementDeviceSerialNumber, okReplacementDeviceSerialNumber := d.GetOk("replacement_device_serial_number")
	vReplacementStatus, okReplacementStatus := d.GetOk("replacement_status")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vSortOrder, okSortOrder := d.GetOk("sort_order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheStatusOfAllTheDeviceReplacementWorkflows")
		queryParams1 := dnacentersdkgo.RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams{}

		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okFaultyDeviceName {
			queryParams1.FaultyDeviceName = vFaultyDeviceName.(string)
		}
		if okFaultyDevicePlatform {
			queryParams1.FaultyDevicePlatform = vFaultyDevicePlatform.(string)
		}
		if okFaultyDeviceSerialNumber {
			queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber.(string)
		}
		if okReplacementDevicePlatform {
			queryParams1.ReplacementDevicePlatform = vReplacementDevicePlatform.(string)
		}
		if okReplacementDeviceSerialNumber {
			queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber.(string)
		}
		if okReplacementStatus {
			queryParams1.ReplacementStatus = vReplacementStatus.(string)
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
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.DeviceReplacement.RetrieveTheStatusOfAllTheDeviceReplacementWorkflows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheStatusOfAllTheDeviceReplacementWorkflows", err,
				"Failure at RetrieveTheStatusOfAllTheDeviceReplacementWorkflows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheStatusOfAllTheDeviceReplacementWorkflows response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItems(items *[]dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponse) []map[string]interface{} {
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
		respItem["neighbor_device_id"] = item.NeighborDeviceID
		respItem["replacement_device_platform"] = item.ReplacementDevicePlatform
		respItem["replacement_device_serial_number"] = item.ReplacementDeviceSerialNumber
		respItem["replacement_status"] = item.ReplacementStatus
		respItem["replacement_time"] = item.ReplacementTime
		respItem["workflow"] = flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItemsWorkflow(item.Workflow)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItemsWorkflow(item *dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["workflow_status"] = item.WorkflowStatus
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["steps"] = flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItemsWorkflowSteps(item.Steps)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsItemsWorkflowSteps(items *[]dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflowSteps) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["status"] = item.Status
		respItem["status_message"] = item.StatusMessage
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
