package dnacenter

import (
	"context"
	"errors"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Issues.

- Update issue trigger threshold, priority for the given id.
Also enable or disable issue trigger for the given id. For detailed information about the usage of the API, please refer
to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		CreateContext: resourceSystemIssueDefinitionsCreate,
		ReadContext:   resourceSystemIssueDefinitionsRead,
		UpdateContext: resourceSystemIssueDefinitionsUpdate,
		DeleteContext: resourceSystemIssueDefinitionsDelete,
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

						"category_name": &schema.Schema{
							Description: `Category Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"default_priority": &schema.Schema{
							Description: `Default Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"definition_status": &schema.Schema{
							Description: `Definition Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"issue_enabled": &schema.Schema{
							Description: `Issue Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"profile_id": &schema.Schema{
							Description: `Profile Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"synchronize_to_health_threshold": &schema.Schema{
							Description: `Synchronize To Health Threshold`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"threshold_value": &schema.Schema{
							Description: `Threshold Value`,
							Type:        schema.TypeFloat,
							Computed:    true,
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

						"id": &schema.Schema{
							Description: `id path parameter. Issue trigger definition id.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"issue_enabled": &schema.Schema{
							Description: `Issue Enabled`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"synchronize_to_health_threshold": &schema.Schema{
							Description: `Synchronize To Health Threshold`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"threshold_value": &schema.Schema{
							Description: `Threshold Value`,
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemIssueDefinitionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSystemIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIssueTriggerDefinitionForGivenID")
		vvID := vID

		response1, restyResp1, err := client.Issues.GetIssueTriggerDefinitionForGivenID(vvID, nil)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		// Review flatten function used
		vItem1 := flattenIssuesGetIssueTriggerDefinitionForGivenIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsAllIssueTriggerDefinitionsForGivenFilters search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSystemIssueDefinitionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSystemIssueDefinitionsIssueTriggerDefinitionUpdate(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Issues.IssueTriggerDefinitionUpdate(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing IssueTriggerDefinitionUpdate", err, restyResp1.String(),
					"Failure at IssueTriggerDefinitionUpdate, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing IssueTriggerDefinitionUpdate", err,
				"Failure at IssueTriggerDefinitionUpdate, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceSystemIssueDefinitionsRead(ctx, d, m)
}

func resourceSystemIssueDefinitionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SystemIssueDefinitions", err, "Delete method is not supported",
		"Failure at SystemIssueDefinitionsDelete, unexpected response", ""))
	return diags
}
func expandRequestSystemIssueDefinitionsIssueTriggerDefinitionUpdate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesIssueTriggerDefinitionUpdate {
	request := dnacentersdkgo.RequestIssuesIssueTriggerDefinitionUpdate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".synchronize_to_health_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".synchronize_to_health_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".synchronize_to_health_threshold")))) {
		request.SynchronizeToHealthThreshold = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issue_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issue_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issue_enabled")))) {
		request.IssueEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".threshold_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".threshold_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".threshold_value")))) {
		request.ThresholdValue = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
