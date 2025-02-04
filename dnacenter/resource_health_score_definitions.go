package dnacenter

import (
	"context"
	"errors"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthScoreDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Devices.

- Update health threshold, include status of overall health status.
And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information
about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		CreateContext: resourceHealthScoreDefinitionsCreate,
		ReadContext:   resourceHealthScoreDefinitionsRead,
		UpdateContext: resourceHealthScoreDefinitionsUpdate,
		DeleteContext: resourceHealthScoreDefinitionsDelete,
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
						"device_family": &schema.Schema{
							Description: `Device Family`,
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
						"include_for_overall_health": &schema.Schema{
							Description: `Include For Overall Health`,
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
						"synchronize_to_issue_threshold": &schema.Schema{
							Description: `Synchronize To Issue Threshold`,
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
							Description: `id path parameter. Health score definition id.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"include_for_overall_health": &schema.Schema{
							Description: `Include For Overall Health`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"synchronize_to_issue_threshold": &schema.Schema{
							Description: `Synchronize To Issue Threshold`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"threshold_value": &schema.Schema{
							Description: `Thresehold Value`,
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

func resourceHealthScoreDefinitionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceHealthScoreDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetHealthScoreDefinitionForTheGivenID")
		vvID := vID

		response1, restyResp1, err := client.Devices.GetHealthScoreDefinitionForTheGivenID(vvID, nil)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceHealthScoreDefinitionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestHealthScoreDefinitionsUpdateHealthScoreDefinitionForTheGivenID(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Devices.UpdateHealthScoreDefinitionForTheGivenID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateHealthScoreDefinitionForTheGivenID", err, restyResp1.String(),
					"Failure at UpdateHealthScoreDefinitionForTheGivenID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateHealthScoreDefinitionForTheGivenID", err,
				"Failure at UpdateHealthScoreDefinitionForTheGivenID, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceHealthScoreDefinitionsRead(ctx, d, m)
}

func resourceHealthScoreDefinitionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing HealthScoreDefinitions", err, "Delete method is not supported",
		"Failure at HealthScoreDefinitionsDelete, unexpected response", ""))
	return diags
}
func expandRequestHealthScoreDefinitionsUpdateHealthScoreDefinitionForTheGivenID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateHealthScoreDefinitionForTheGivenID {
	request := dnacentersdkgo.RequestDevicesUpdateHealthScoreDefinitionForTheGivenID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_for_overall_health")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_for_overall_health")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_for_overall_health")))) {
		request.IncludeForOverallHealth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".threshold_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".threshold_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".threshold_value")))) {
		request.ThresholdValue = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".synchronize_to_issue_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".synchronize_to_issue_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".synchronize_to_issue_threshold")))) {
		request.SynchronizeToIssueThreshold = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
