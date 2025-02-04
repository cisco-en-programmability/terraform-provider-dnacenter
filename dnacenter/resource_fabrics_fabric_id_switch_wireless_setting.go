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

func resourceFabricsFabricIDSwitchWirelessSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Fabric Wireless.

- This resource is used to enable or disable wireless capabilities on switch devices, along with configuring rolling AP
upgrades on the fabric site. Reboot action is required to remove wireless configurations.
`,

		CreateContext: resourceFabricsFabricIDSwitchWirelessSettingCreate,
		ReadContext:   resourceFabricsFabricIDSwitchWirelessSettingRead,
		UpdateContext: resourceFabricsFabricIDSwitchWirelessSettingUpdate,
		DeleteContext: resourceFabricsFabricIDSwitchWirelessSettingDelete,
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

						"enable_wireless": &schema.Schema{
							Description: `Enable Wireless`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Network Device ID of the Wireless Capable Switch
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rolling_ap_upgrade": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_reboot_percentage": &schema.Schema{
										Description: `AP Reboot Percentage. Permissible values - 5, 15, 25
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"enable_rolling_ap_upgrade": &schema.Schema{
										Description: `Enable Rolling Ap Upgrade`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

						"enable_wireless": &schema.Schema{
							Description: `Enable Wireless`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"fabric_id": &schema.Schema{
							Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"id": &schema.Schema{
							Description: `Network Device ID of the wireless capable switch
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rolling_ap_upgrade": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_reboot_percentage": &schema.Schema{
										Description: `AP Reboot Percentage. Permissible values - 5, 15, 25
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"enable_rolling_ap_upgrade": &schema.Schema{
										Description: `Enable Rolling Ap Upgrade`,
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
				},
			},
		},
	}
}

func resourceFabricsFabricIDSwitchWirelessSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	resourceMap["fabric_id"] = interfaceToString(resourceItem["fabric_id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceFabricsFabricIDSwitchWirelessSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSdaWirelessDetailsFromSwitches")
		vvFabricID := vFabricID

		response1, restyResp1, err := client.FabricWireless.GetSdaWirelessDetailsFromSwitches(vvFabricID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		// Review flatten function used
		vItem1 := flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaWirelessDetailsFromSwitches search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceFabricsFabricIDSwitchWirelessSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]

	if d.HasChange("parameters") {
		request1 := expandRequestFabricsFabricIDSwitchWirelessSettingSwitchWirelessSettingAndRollingApUpgradeManagement(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && request1.ID == "" {
			log.Printf("[DEBUG] resty response for update operation => %v", request1)
			diags = append(diags, diagError(
				"Erro expand Request Fabrics", errors.New("error")))
			return diags
		}
		response1, restyResp1, err := client.FabricWireless.SwitchWirelessSettingAndRollingApUpgradeManagement(vFabricID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SwitchWirelessSettingAndRollingApUpgradeManagement", err, restyResp1.String(),
					"Failure at SwitchWirelessSettingAndRollingApUpgradeManagement, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SwitchWirelessSettingAndRollingApUpgradeManagement", err,
				"Failure at SwitchWirelessSettingAndRollingApUpgradeManagement, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SwitchWirelessSettingAndRollingApUpgradeManagement", err))
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
					"Failure when executing SwitchWirelessSettingAndRollingApUpgradeManagement", err1))
				return diags
			}
		}

	}

	return resourceFabricsFabricIDSwitchWirelessSettingRead(ctx, d, m)
}

func resourceFabricsFabricIDSwitchWirelessSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete FabricsFabricIDSwitchWirelessSetting on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestFabricsFabricIDSwitchWirelessSettingSwitchWirelessSettingAndRollingApUpgradeManagement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement {
	request := dnacentersdkgo.RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_wireless")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_wireless")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_wireless")))) {
		request.EnableWireless = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rolling_ap_upgrade")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rolling_ap_upgrade")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rolling_ap_upgrade")))) {
		request.RollingApUpgrade = expandRequestFabricsFabricIDSwitchWirelessSettingSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade(ctx, key+".rolling_ap_upgrade.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFabricsFabricIDSwitchWirelessSettingSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade {
	request := dnacentersdkgo.RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_rolling_ap_upgrade")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_rolling_ap_upgrade")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_rolling_ap_upgrade")))) {
		request.EnableRollingApUpgrade = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_reboot_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_reboot_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_reboot_percentage")))) {
		request.ApRebootPercentage = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesByIDItem(items *[]dnacentersdkgo.ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["enableWireless"] = item.EnableWireless
		respItem["rollingApUpgrade"] = flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesByIDItemRollingApUpgrade(item.RollingApUpgrade)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesByIDItemRollingApUpgrade(item *dnacentersdkgo.ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponseRollingApUpgrade) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enableRollingApUpgrade"] = item.EnableRollingApUpgrade
	respItem["apRebootPercentage"] = item.ApRebootPercentage
	return []map[string]interface{}{
		respItem,
	}

}
