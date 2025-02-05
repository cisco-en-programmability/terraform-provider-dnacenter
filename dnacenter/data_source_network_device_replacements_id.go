package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceReplacementsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Replacement.

- Fetches the status of the device replacement workflow for a given device replacement *id*. Invoke the API
*/dna/intent/api/v1/networkDeviceReplacements* to *GET* the list of all device replacements and use the *id* field data
as input to this API.
`,

		ReadContext: dataSourceNetworkDeviceReplacementsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Instance UUID of the device replacement
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceNetworkDeviceReplacementsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice")
		vvID := vID.(string)

		response1, restyResp1, err := client.DeviceReplacement.RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice", err,
				"Failure at RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItem(item *dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["workflow"] = flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItemWorkflow(item.Workflow)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItemWorkflow(item *dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["workflow_status"] = item.WorkflowStatus
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["steps"] = flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItemWorkflowSteps(item.Steps)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceItemWorkflowSteps(items *[]dnacentersdkgo.ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflowSteps) []map[string]interface{} {
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
