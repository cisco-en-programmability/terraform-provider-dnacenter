package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDiagnosticValidationWorkflows() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Health and Performance.

- Submits the workflow for executing the validations for the given validation specifications

- Deletes the workflow for the given id
`,

		CreateContext: resourceDiagnosticValidationWorkflowsCreate,
		ReadContext:   resourceDiagnosticValidationWorkflowsRead,
		UpdateContext: resourceDiagnosticValidationWorkflowsUpdate,
		DeleteContext: resourceDiagnosticValidationWorkflowsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description of the workflow to run
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Workflow id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Description: `Name of the workflow to run. It must be unique.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"validation_set_ids": &schema.Schema{
							Description: `List of validation set ids
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceDiagnosticValidationWorkflowsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDiagnosticValidationWorkflowsSubmitsTheWorkflowForExecutingValidations(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.HealthAndPerformance.RetrievesValidationWorkflowDetails(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = getResponse2.Response.Name
			d.SetId(joinResourceID(resourceMap))
			return resourceDiagnosticValidationWorkflowsRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.RetrievesTheListOfValidationWorkflowsQueryParams{}

		response2, err := searchHealthAndPerformanceRetrievesTheListOfValidationWorkflows(m, queryParamImport, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response2.ID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDiagnosticValidationWorkflowsRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.HealthAndPerformance.SubmitsTheWorkflowForExecutingValidations(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing SubmitsTheWorkflowForExecutingValidations", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing SubmitsTheWorkflowForExecutingValidations", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceDiagnosticValidationWorkflowsRead(ctx, d, m)
}

func resourceDiagnosticValidationWorkflowsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	log.Printf("[DEBUG] Selected method: RetrievesValidationWorkflowDetails")
	vvID := vID

	response1, restyResp1, err := client.HealthAndPerformance.RetrievesValidationWorkflowDetails(vvID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}
	vItem1 := flattenHealthAndPerformanceRetrievesValidationWorkflowDetailsItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheListOfValidationWorkflows search response",
			err))
		return diags
	}

	return diags
}

func resourceDiagnosticValidationWorkflowsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDiagnosticValidationWorkflowsRead(ctx, d, m)
}

func resourceDiagnosticValidationWorkflowsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	restyResp1, err := client.HealthAndPerformance.DeletesAValidationWorkflow(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesAValidationWorkflow", err, restyResp1.String(),
				"Failure at DeletesAValidationWorkflow, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesAValidationWorkflow", err,
			"Failure at DeletesAValidationWorkflow, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDiagnosticValidationWorkflowsSubmitsTheWorkflowForExecutingValidations(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations {
	request := dnacentersdkgo.RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_set_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_set_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_set_ids")))) {
		request.ValidationSetIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchHealthAndPerformanceRetrievesTheListOfValidationWorkflows(m interface{}, queryParams dnacentersdkgo.RetrievesTheListOfValidationWorkflowsQueryParams, vName string) (*dnacentersdkgo.ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsResponse
	// var ite *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflows
	if vName != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.HealthAndPerformance.RetrievesTheListOfValidationWorkflows(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vName == item.Name {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.HealthAndPerformance.RetrievesTheListOfValidationWorkflows(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	}
	return foundItem, err
}
