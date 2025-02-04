package dnacenter

import (
	"context"

	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceExecuteSuggestedActionsCommands() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Issues.

- This data source action triggers the execution of the suggested actions for an issue, given the Issue Id. It will
return an execution Id. At the completion of the execution, the output of the commands associated with the suggested
actions will be provided
Invoking this API would provide the execution id. Execute the 'Get Business API Execution Details' API with this
execution id, to receive the suggested actions commands output.
`,

		CreateContext: resourceExecuteSuggestedActionsCommandsCreate,
		ReadContext:   resourceExecuteSuggestedActionsCommandsRead,
		DeleteContext: resourceExecuteSuggestedActionsCommandsDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_type": &schema.Schema{
							Description: `Commands provided as part of the suggested actions for an issue can be executed based on issue id. The value here must be issue_id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"entity_value": &schema.Schema{
							Description: `Contains the actual value for the entity type that has been defined
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action_info": &schema.Schema{
										Description: `Actions Info`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"command": &schema.Schema{
										Description: `Command`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"command_output": &schema.Schema{
										Description: `Command Output`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"entity_id": &schema.Schema{
										Description: `Entity Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"hostname": &schema.Schema{
										Description: `Hostname`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"steps_count": &schema.Schema{
										Description: `Steps Count`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"steps_description": &schema.Schema{
										Description: `Steps Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
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

func resourceExecuteSuggestedActionsCommandsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestExecuteSuggestedActionsCommandsExecuteSuggestedActionsCommands(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Issues.ExecuteSuggestedActionsCommands(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ExecuteSuggestedActionsCommands", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenIssuesExecuteSuggestedActionsCommandsItems(response1)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ExecuteSuggestedActionsCommands response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceExecuteSuggestedActionsCommandsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceExecuteSuggestedActionsCommandsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestExecuteSuggestedActionsCommandsExecuteSuggestedActionsCommands(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesExecuteSuggestedActionsCommands {
	request := dnacentersdkgo.RequestIssuesExecuteSuggestedActionsCommands{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".entity_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".entity_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".entity_type")))) {
		request.EntityType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".entity_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".entity_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".entity_value")))) {
		request.EntityValue = interfaceToString(v)
	}
	return &request
}

func flattenIssuesExecuteSuggestedActionsCommandsItems(items *dnacentersdkgo.ResponseIssuesExecuteSuggestedActionsCommands) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["action_info"] = item.ActionInfo
		respItem["steps_count"] = item.StepsCount
		respItem["entity_id"] = item.EntityID
		respItem["hostname"] = item.Hostname
		respItem["steps_description"] = item.StepsDescription
		respItem["command"] = item.Command
		respItem["command_output"] = flattenIssuesExecuteSuggestedActionsCommandsItemsCommandOutput(item.CommandOutput)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesExecuteSuggestedActionsCommandsItemsCommandOutput(item *dnacentersdkgo.ResponseItemIssuesExecuteSuggestedActionsCommandsCommandOutput) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
