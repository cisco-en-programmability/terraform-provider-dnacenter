package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiagnosticValidationWorkflows() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- Retrieves the workflows that have been successfully submitted and are currently available. This is sorted by
*submitTime*

- Retrieves workflow details for a workflow id
`,

		ReadContext: dataSourceDiagnosticValidationWorkflowsRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Workflows started before the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Workflow id
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
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"run_status": &schema.Schema{
				Description: `runStatus query parameter. Execution status of the workflow. If the workflow is successfully submitted, runStatus is *PENDING*. If the workflow execution has started, runStatus is *IN_PROGRESS*. If the workflow executed is completed with all validations executed, runStatus is *COMPLETED*. If the workflow execution fails while running validations, runStatus is *FAILED*.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Workflows started after the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Workflow description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Workflow id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Workflow name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"release_version": &schema.Schema{
							Description: `Product version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"run_status": &schema.Schema{
							Description: `Execution status of the workflow. If the workflow is successfully submitted, runStatus will return *PENDING*. If the workflow execution has started, runStatus will return *IN_PROGRESS*. If the workflow executed is completed with all validations executed, runStatus will return *COMPLETED*. If the workflow execution fails while running validations, runStatus will return *FAILED*.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"submit_time": &schema.Schema{
							Description: `Workflow submit time (as milliseconds since UNIX epoch).
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"validation_set_ids": &schema.Schema{
							Description: `List of validation set ids
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"validation_sets_run_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Validation set run finish time (as milliseconds since UNIX epoch).
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Validation set run start time (as milliseconds since UNIX epoch).
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"validation_run_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"validation_id": &schema.Schema{
													Description: `Validation id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"validation_message": &schema.Schema{
													Description: `Validation execution result detail message
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"validation_name": &schema.Schema{
													Description: `Validation name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"validation_status": &schema.Schema{
													Description: `Validation execution result status
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"validation_set_id": &schema.Schema{
										Description: `Validation set id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"validation_status": &schema.Schema{
										Description: `Overall result of the validation set execution. If any of the contained validation execution status is *CRITICAL*, this is marked as *CRITICAL*. Else, if any of the contained validation execution status is *WARNING*, this is marked as *WARNING*. Else, this is marked as *INFORMATION*. This is empty when the workflow is in progress.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version": &schema.Schema{
										Description: `Validation set version
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"validation_status": &schema.Schema{
							Description: `Overall result of the execution of all the validations. If any of the contained validation execution status is *CRITICAL*, this is marked as *CRITICAL*. Else, if any of the contained validation execution status is *WARNING*, this is marked as *WARNING*. Else, this is marked as *INFORMATION*.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Workflow description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Description: `Workflow finish time (as milliseconds since UNIX epoch).
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
							Description: `Workflow name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"run_status": &schema.Schema{
							Description: `Execution status of the workflow. If the workflow is successfully submitted, runStatus will return *PENDING*. If the workflow execution has started, runStatus will return *IN_PROGRESS*. If the workflow executed is completed with all validations executed, runStatus will return *COMPLETED*. If the workflow execution fails while running validations, runStatus will return *FAILED*.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Workflow start time (as milliseconds since UNIX epoch).
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"submit_time": &schema.Schema{
							Description: `Workflow submit time (as milliseconds since UNIX epoch).
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"validation_set_ids": &schema.Schema{
							Description: `List of validation set ids
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"validation_status": &schema.Schema{
							Description: `Overall result of execution of the validation workflow
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

func dataSourceDiagnosticValidationWorkflowsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vRunStatus, okRunStatus := d.GetOk("run_status")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vID, okID := d.GetOk("id")

	method1 := []bool{okStartTime, okEndTime, okRunStatus, okOffset, okLimit}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfValidationWorkflows")
		queryParams1 := dnacentersdkgo.RetrievesTheListOfValidationWorkflowsQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okRunStatus {
			queryParams1.RunStatus = vRunStatus.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.HealthAndPerformance.RetrievesTheListOfValidationWorkflows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfValidationWorkflows", err,
				"Failure at RetrievesTheListOfValidationWorkflows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenHealthAndPerformanceRetrievesTheListOfValidationWorkflowsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfValidationWorkflows response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrievesValidationWorkflowDetails")
		vvID := vID.(string)

		response2, restyResp2, err := client.HealthAndPerformance.RetrievesValidationWorkflowDetails(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesValidationWorkflowDetails", err,
				"Failure at RetrievesValidationWorkflowDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesValidationWorkflowDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceRetrievesTheListOfValidationWorkflowsItems(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["run_status"] = item.RunStatus
		respItem["submit_time"] = item.SubmitTime
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["validation_status"] = item.ValidationStatus
		respItem["validation_set_ids"] = item.ValidationSetIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItem(item *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["run_status"] = item.RunStatus
	respItem["submit_time"] = item.SubmitTime
	respItem["validation_set_ids"] = item.ValidationSetIDs
	respItem["release_version"] = item.ReleaseVersion
	respItem["validation_sets_run_details"] = flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItemValidationSetsRunDetails(item.ValidationSetsRunDetails)
	respItem["validation_status"] = item.ValidationStatus
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItemValidationSetsRunDetails(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["validation_set_id"] = item.ValidationSetID
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["validation_status"] = item.ValidationStatus
		respItem["version"] = item.Version
		respItem["validation_run_details"] = flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItemValidationSetsRunDetailsValidationRunDetails(item.ValidationRunDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItemValidationSetsRunDetailsValidationRunDetails(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetailsValidationRunDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["validation_id"] = item.ValidationID
		respItem["validation_name"] = item.ValidationName
		respItem["validation_message"] = item.ValidationMessage
		respItem["validation_status"] = item.ValidationStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}
