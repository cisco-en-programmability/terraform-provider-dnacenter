package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Updates fabric devices based on user input.

- Deletes fabric devices based on user input.

- Adds fabric devices based on user input.

- Deletes a fabric device based on id.
`,

		CreateContext: resourceSdaFabricDevicesCreate,
		ReadContext:   resourceSdaFabricDevicesRead,
		UpdateContext: resourceSdaFabricDevicesUpdate,
		DeleteContext: resourceSdaFabricDevicesDelete,
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

						"border_device_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"border_types": &schema.Schema{
										Description: `List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"layer3_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"border_priority": &schema.Schema{
													Description: `Border priority of the fabric border device.  A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"import_external_routes": &schema.Schema{
													Description: `Import external routes value of the fabric border device.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_default_exit": &schema.Schema{
													Description: `Is default exit value of the fabric border device.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"local_autonomous_system_number": &schema.Schema{
													Description: `BGP Local autonomous system number of the fabric border device.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"prepend_autonomous_system_count": &schema.Schema{
													Description: `Prepend autonomous system count of the fabric border device.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"device_roles": &schema.Schema{
							Description: `List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric of this fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddFabricDevices`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"border_device_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"border_types": &schema.Schema{
													Description: `List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"layer3_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"border_priority": &schema.Schema{
																Description: `Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"import_external_routes": &schema.Schema{
																Description: `Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane.
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"is_default_exit": &schema.Schema{
																Description: `Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes.
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"local_autonomous_system_number": &schema.Schema{
																Description: `BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295].
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"prepend_autonomous_system_count": &schema.Schema{
																Description: `Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"device_roles": &schema.Schema{
										Description: `List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE].
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric of this fabric device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the fabric device. (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device ID of the fabric device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

func resourceSdaFabricDevicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaFabricDevicesAddFabricDevices(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vNetworkDevice := resourceItem["network_device"]
	vvNetworkDevice := interfaceToString(vNetworkDevice)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	queryParamImport := dnacentersdkgo.GetFabricDevicesQueryParams{}
	queryParamImport.FabricID = vvFabricID
	queryParamImport.NetworkDeviceID = vvNetworkDevice
	item2, err := searchSdaGetFabricDevices(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["fabric_id"] = item2.FabricID
		resourceMap["network_device"] = item2.NetworkDeviceID
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricDevicesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddFabricDevices(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabricDevices", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevices", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevices", err))
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
				"Failure when executing AddFabricDevices", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetFabricDevicesQueryParams{}
	queryParamValidate.FabricID = vvFabricID
	item3, err := searchSdaGetFabricDevices(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddFabricDevices", err,
			"Failure at AddFabricDevices, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["fabric_id"] = item3.FabricID
	resourceMap["network_device"] = item3.NetworkDeviceID
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricDevicesRead(ctx, d, m)
}

func resourceSdaFabricDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]
	vvID := resourceMap["id"]
	vNetworkDevice := resourceMap["network_device"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevices")
		queryParams1 := dnacentersdkgo.GetFabricDevicesQueryParams{}

		queryParams1.FabricID = vFabricID
		queryParams1.NetworkDeviceID = vNetworkDevice

		item1, err := searchSdaGetFabricDevices(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}

		items := []dnacentersdkgo.ResponseSdaGetFabricDevicesResponse{
			*item1,
		}
		// Review flatten function used
		vItem1 := flattenSdaGetFabricDevicesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevices search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricDevicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricDevicesUpdateFabricDevices(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateFabricDevices(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFabricDevices", err, restyResp1.String(),
					"Failure at UpdateFabricDevices, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFabricDevices", err,
				"Failure at UpdateFabricDevices, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateFabricDevices", err))
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
					"Failure when executing UpdateFabricDevices", err1))
				return diags
			}
		}

	}

	return resourceSdaFabricDevicesRead(ctx, d, m)
}

func resourceSdaFabricDevicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Sda.DeleteFabricDeviceByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFabricDeviceByID", err, restyResp1.String(),
				"Failure at DeleteFabricDeviceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFabricDeviceByID", err,
			"Failure at DeleteFabricDeviceByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricDeviceByID", err))
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
				"Failure when executing DeleteFabricDeviceByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricDevicesAddFabricDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddFabricDevices {
	request := dnacentersdkgo.RequestSdaAddFabricDevices{}
	if v := expandRequestSdaFabricDevicesAddFabricDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesAddFabricDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddFabricDevices {
	request := []dnacentersdkgo.RequestItemSdaAddFabricDevices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricDevicesAddFabricDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesAddFabricDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricDevices {
	request := dnacentersdkgo.RequestItemSdaAddFabricDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_roles")))) {
		request.DeviceRoles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_device_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_device_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_device_settings")))) {
		request.BorderDeviceSettings = expandRequestSdaFabricDevicesAddFabricDevicesItemBorderDeviceSettings(ctx, key+".border_device_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesAddFabricDevicesItemBorderDeviceSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricDevicesBorderDeviceSettings {
	request := dnacentersdkgo.RequestItemSdaAddFabricDevicesBorderDeviceSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_types")))) {
		request.BorderTypes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3_settings")))) {
		request.Layer3Settings = expandRequestSdaFabricDevicesAddFabricDevicesItemBorderDeviceSettingsLayer3Settings(ctx, key+".layer3_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesAddFabricDevicesItemBorderDeviceSettingsLayer3Settings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricDevicesBorderDeviceSettingsLayer3Settings {
	request := dnacentersdkgo.RequestItemSdaAddFabricDevicesBorderDeviceSettingsLayer3Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_autonomous_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_autonomous_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_autonomous_system_number")))) {
		request.LocalAutonomousSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_exit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_exit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_exit")))) {
		request.IsDefaultExit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".import_external_routes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".import_external_routes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".import_external_routes")))) {
		request.ImportExternalRoutes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_priority")))) {
		request.BorderPriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prepend_autonomous_system_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prepend_autonomous_system_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prepend_autonomous_system_count")))) {
		request.PrependAutonomousSystemCount = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesUpdateFabricDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateFabricDevices {
	request := dnacentersdkgo.RequestSdaUpdateFabricDevices{}
	if v := expandRequestSdaFabricDevicesUpdateFabricDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesUpdateFabricDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateFabricDevices {
	request := []dnacentersdkgo.RequestItemSdaUpdateFabricDevices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricDevicesUpdateFabricDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesUpdateFabricDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricDevices {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_roles")))) {
		request.DeviceRoles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_device_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_device_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_device_settings")))) {
		request.BorderDeviceSettings = expandRequestSdaFabricDevicesUpdateFabricDevicesItemBorderDeviceSettings(ctx, key+".border_device_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesUpdateFabricDevicesItemBorderDeviceSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricDevicesBorderDeviceSettings {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricDevicesBorderDeviceSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_types")))) {
		request.BorderTypes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3_settings")))) {
		request.Layer3Settings = expandRequestSdaFabricDevicesUpdateFabricDevicesItemBorderDeviceSettingsLayer3Settings(ctx, key+".layer3_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesUpdateFabricDevicesItemBorderDeviceSettingsLayer3Settings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricDevicesBorderDeviceSettingsLayer3Settings {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricDevicesBorderDeviceSettingsLayer3Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_autonomous_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_autonomous_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_autonomous_system_number")))) {
		request.LocalAutonomousSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_exit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_exit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_exit")))) {
		request.IsDefaultExit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".import_external_routes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".import_external_routes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".import_external_routes")))) {
		request.ImportExternalRoutes = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_priority")))) {
		request.BorderPriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prepend_autonomous_system_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prepend_autonomous_system_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prepend_autonomous_system_count")))) {
		request.PrependAutonomousSystemCount = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetFabricDevices(m interface{}, queryParams dnacentersdkgo.GetFabricDevicesQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetFabricDevicesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetFabricDevicesResponse
	var ite *dnacentersdkgo.ResponseSdaGetFabricDevices

	ite, _, err = client.Sda.GetFabricDevices(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err
	}
	itemsCopy := *ite.Response
	if itemsCopy == nil {
		return foundItem, err
	}
	for _, item := range itemsCopy {
		if item.NetworkDeviceID == queryParams.NetworkDeviceID {
			foundItem = &item
			return foundItem, err
		}
	}

	return foundItem, err
}
