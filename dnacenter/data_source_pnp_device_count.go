package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the device count based on filter criteria. This is useful for pagination
`,

		ReadContext: dataSourcePnpDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"cm_state": &schema.Schema{
				Description: `cmState query parameter. Device Connection Manager State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_contact": &schema.Schema{
				Description: `lastContact query parameter. Device Has Contacted lastContact > 0
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Device Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"onb_state": &schema.Schema{
				Description: `onbState query parameter. Device Onboarding State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pid": &schema.Schema{
				Description: `pid query parameter. Device ProductId
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_id": &schema.Schema{
				Description: `projectId query parameter. Device Project Id
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_name": &schema.Schema{
				Description: `projectName query parameter. Device Project Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Device Serial Number
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"smart_account_id": &schema.Schema{
				Description: `smartAccountId query parameter. Device Smart Account
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": &schema.Schema{
				Description: `source query parameter. Device Source
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Description: `state query parameter. Device State
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_account_id": &schema.Schema{
				Description: `virtualAccountId query parameter. Device Virtual Account
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_id": &schema.Schema{
				Description: `workflowId query parameter. Device Workflow Id
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_name": &schema.Schema{
				Description: `workflowName query parameter. Device Workflow Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vState, okState := d.GetOk("state")
	vOnbState, okOnbState := d.GetOk("onb_state")
	vCmState, okCmState := d.GetOk("cm_state")
	vName, okName := d.GetOk("name")
	vPid, okPid := d.GetOk("pid")
	vSource, okSource := d.GetOk("source")
	vProjectID, okProjectID := d.GetOk("project_id")
	vWorkflowID, okWorkflowID := d.GetOk("workflow_id")
	vProjectName, okProjectName := d.GetOk("project_name")
	vWorkflowName, okWorkflowName := d.GetOk("workflow_name")
	vSmartAccountID, okSmartAccountID := d.GetOk("smart_account_id")
	vVirtualAccountID, okVirtualAccountID := d.GetOk("virtual_account_id")
	vLastContact, okLastContact := d.GetOk("last_contact")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceCount")
		queryParams1 := dnacentersdkgo.GetDeviceCountQueryParams{}

		if okSerialNumber {
			queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
		}
		if okState {
			queryParams1.State = interfaceToSliceString(vState)
		}
		if okOnbState {
			queryParams1.OnbState = interfaceToSliceString(vOnbState)
		}
		if okCmState {
			queryParams1.CmState = interfaceToSliceString(vCmState)
		}
		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}
		if okPid {
			queryParams1.Pid = interfaceToSliceString(vPid)
		}
		if okSource {
			queryParams1.Source = interfaceToSliceString(vSource)
		}
		if okProjectID {
			queryParams1.ProjectID = interfaceToSliceString(vProjectID)
		}
		if okWorkflowID {
			queryParams1.WorkflowID = interfaceToSliceString(vWorkflowID)
		}
		if okProjectName {
			queryParams1.ProjectName = interfaceToSliceString(vProjectName)
		}
		if okWorkflowName {
			queryParams1.WorkflowName = interfaceToSliceString(vWorkflowName)
		}
		if okSmartAccountID {
			queryParams1.SmartAccountID = interfaceToSliceString(vSmartAccountID)
		}
		if okVirtualAccountID {
			queryParams1.VirtualAccountID = interfaceToSliceString(vVirtualAccountID)
		}
		if okLastContact {
			queryParams1.LastContact = vLastContact.(bool)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetDeviceCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceCount", err,
				"Failure at GetDeviceCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpGetDeviceCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetDeviceCountItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetDeviceCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
