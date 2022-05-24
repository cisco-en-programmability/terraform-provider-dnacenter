package dnacenter

import (
	"context"
	"errors"
	"time"

	"reflect"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.
- Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from
Request body.
`,

		CreateContext: resourceInterfaceUpdateCreate,
		ReadContext:   resourceInterfaceUpdateRead,
		DeleteContext: resourceInterfaceUpdateDelete,
		Schema: map[string]*schema.Schema{
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_mode": &schema.Schema{
							Description: `deploymentMode query parameter. Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
			`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"interface_uuid": &schema.Schema{
							Description: `interfaceUuid path parameter. Interface ID
			`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"admin_status": &schema.Schema{
							Description: `Admin Status`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
						"voice_vlan_id": &schema.Schema{
							Description: `Voice Vlan Id`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"task_id": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"url": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"required": &schema.Schema{
							Description: `Required`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceInterfaceUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vInterfaceID := resourceItem["interface_uuid"]
	vvInterfaceID := interfaceToString(vInterfaceID)
	vDeploymentMode, okDeploymentMode := d.GetOk("parameters.0.deployment_mode")
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: UpdateInterfaceDetails")
		request1 := expandRequestInterfaceUpdateUpdateInterfaceDetails(ctx, "parameters.0", d)
		queryParams1 := dnacentersdkgo.UpdateInterfaceDetailsQueryParams{}
		if okDeploymentMode {
			queryParams1.DeploymentMode = vDeploymentMode.(string)
		}
		response1, restyResp1, err := client.Devices.UpdateInterfaceDetails(vvInterfaceID, request1, &queryParams1)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateInterfaceDetails", err,
				"Failure at UpdateInterfaceDetails, unexpected response", ""))
			return diags
		}

		if response1.Response.Properties == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateInterfaceDetails", err))
			return diags
		}
		if response1.Response.Properties.TaskID == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateInterfaceDetails", err))
			return diags
		}
		taskId := response1.Response.Properties.TaskID.Type
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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateInterfaceDetails", err1))
				return diags
			}
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesUpdateInterfaceDetailsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateInterfaceDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags
	}
	return diags
}

func resourceInterfaceUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestInterfaceUpdateUpdateInterfaceDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateInterfaceDetails {
	request := dnacentersdkgo.RequestDevicesUpdateInterfaceDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_status")))) {
		request.AdminStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".voice_vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".voice_vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".voice_vlan_id")))) {
		request.VoiceVLANID = interfaceToIntPtr(v)
	}
	return &request
}

func flattenDevicesUpdateInterfaceDetailsItem(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["properties"] = flattenDevicesUpdateInterfaceDetailsItemProperties(item.Properties)
	respItem["required"] = item.Required
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesUpdateInterfaceDetailsItemProperties(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = flattenDevicesUpdateInterfaceDetailsItemPropertiesTaskID(item.TaskID)
	respItem["url"] = flattenDevicesUpdateInterfaceDetailsItemPropertiesURL(item.URL)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesUpdateInterfaceDetailsItemPropertiesTaskID(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesUpdateInterfaceDetailsItemPropertiesURL(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
