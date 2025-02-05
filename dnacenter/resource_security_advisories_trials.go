package dnacenter

import (
	"context"

	"errors"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityAdvisoriesTrials() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Compliance.

- Creates a trial for security advisories detection on network devices. The consent to connect agreement must have been
accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.
`,

		CreateContext: resourceSecurityAdvisoriesTrialsCreate,
		ReadContext:   resourceSecurityAdvisoriesTrialsRead,
		DeleteContext: resourceSecurityAdvisoriesTrialsDelete,
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

						"active": &schema.Schema{
							Description: `Indicates if the trial is active
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"contract_level": &schema.Schema{
							Description: `Contract level for which trial was created. this was used in older versions and exists only for compatibility.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_time": &schema.Schema{
							Description: `Trial end time; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"feature": &schema.Schema{
							Description: `Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"seconds_remaining_to_expiry": &schema.Schema{
							Description: `Seconds remaining in the trial before it expires. for expired trials this will be 0.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"seconds_since_expired": &schema.Schema{
							Description: `Seconds elapsed after the trial has expired. for active trials this will be 0.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Description: `Trial start time; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
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

func resourceSecurityAdvisoriesTrialsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	item2, _, err := client.Compliance.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices()
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		d.SetId(joinResourceID(resourceMap))
		return resourceSecurityAdvisoriesTrialsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Compliance.CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices()
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices", err1))
			return diags
		}
	}
	item3, _, err := client.Compliance.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices()
	if err != nil || item3 != nil {
		resourceMap := make(map[string]string)
		d.SetId(joinResourceID(resourceMap))
		return resourceSecurityAdvisoriesTrialsRead(ctx, d, m)
	}
	item4, _, err := client.Compliance.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices()
	if err != nil || item4 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices", err,
			"Failure at CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)

	d.SetId(joinResourceID(resourceMap))
	return resourceSecurityAdvisoriesTrialsRead(ctx, d, m)
}

func resourceSecurityAdvisoriesTrialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices")

		response1, restyResp1, err := client.Compliance.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSecurityAdvisoriesTrialsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SecurityAdvisoriesTrials on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
