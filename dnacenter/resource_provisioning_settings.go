package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProvisioningSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on System Settings.

- Sets provisioning settings
`,

		CreateContext: resourceProvisioningSettingsCreate,
		ReadContext:   resourceProvisioningSettingsRead,
		UpdateContext: resourceProvisioningSettingsUpdate,
		DeleteContext: resourceProvisioningSettingsDelete,
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

						"require_itsm_approval": &schema.Schema{
							Description: `If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"require_preview": &schema.Schema{
							Description: `If require preview is enabled, the device configurations must be reviewed before deploying them
`,
							// Type:        schema.TypeBool,
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

						"require_itsm_approval": &schema.Schema{
							Description: `If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"require_preview": &schema.Schema{
							Description: `If require preview is enabled, the device configurations must be reviewed before deploying them
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceProvisioningSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceProvisioningSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProvisioningSettings")

		response1, restyResp1, err := client.SystemSettings.GetProvisioningSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsGetProvisioningSettingsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisioningSettings response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceProvisioningSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestProvisioningSettingsSetProvisioningSettings(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SystemSettings.SetProvisioningSettings(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetProvisioningSettings", err, restyResp1.String(),
					"Failure at SetProvisioningSettings, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetProvisioningSettings", err,
				"Failure at SetProvisioningSettings, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SetProvisioningSettings", err))
			return diags
		}
		taskId := response1.Response.TaskID
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
					"Failure when executing SetProvisioningSettings", err1))
				return diags
			}
		}

	}

	return resourceProvisioningSettingsRead(ctx, d, m)
}

func resourceProvisioningSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing ProvisioningSettings", err, "Delete method is not supported",
		"Failure at ProvisioningSettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestProvisioningSettingsSetProvisioningSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSystemSettingsSetProvisioningSettings {
	request := dnacentersdkgo.RequestSystemSettingsSetProvisioningSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_itsm_approval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_itsm_approval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_itsm_approval")))) {
		request.RequireItsmApproval = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_preview")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_preview")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_preview")))) {
		request.RequirePreview = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
