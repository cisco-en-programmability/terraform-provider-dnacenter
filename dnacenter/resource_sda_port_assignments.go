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

func resourceSdaPortAssignments() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds port assignments based on user input.

- Updates port assignments based on user input.

- Deletes port assignments based on user input.

- Deletes a port assignment based on id.
`,

		CreateContext: resourceSdaPortAssignmentsCreate,
		ReadContext:   resourceSdaPortAssignmentsRead,
		UpdateContext: resourceSdaPortAssignmentsUpdate,
		DeleteContext: resourceSdaPortAssignmentsDelete,
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

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate template name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"connected_device_type": &schema.Schema{
							Description: `Connected device type of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_vlan_name": &schema.Schema{
							Description: `Data VLAN name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric the device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_description": &schema.Schema{
							Description: `Interface description of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device ID of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_group_name": &schema.Schema{
							Description: `Security group name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"voice_vlan_name": &schema.Schema{
							Description: `Voice VLAN name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddPortAssignments`,
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
									"authenticate_template_name": &schema.Schema{
										Description: `Authenticate template name of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"connected_device_type": &schema.Schema{
										Description: `Connected device type of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"data_vlan_name": &schema.Schema{
										Description: `Data VLAN name of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric the device is assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_description": &schema.Schema{
										Description: `Interface description of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device ID of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scalable_group_name": &schema.Schema{
										Description: `Scalable group name of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"security_group_name": &schema.Schema{
										Description: `Security group name of the port assignment.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"voice_vlan_name": &schema.Schema{
										Description: `Voice VLAN name of the port assignment.
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

func resourceSdaPortAssignmentsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaPortAssignmentsAddPortAssignments(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	vNetworkDeviceID := resourceItem["network_device_id"]
	vvNetworkDeviceID := interfaceToString(vNetworkDeviceID)
	vInterfaceName := resourceItem["interface_name"]
	vvInterfaceName := interfaceToString(vInterfaceName)
	vDataVLANName := resourceItem["data_vlan_name"]
	vvDataVLANName := interfaceToString(vDataVLANName)
	vVoiceVLANName := resourceItem["voice_vlan_name"]
	vvVoiceVLANName := interfaceToString(vVoiceVLANName)

	queryParamImport := dnacentersdkgo.GetPortAssignmentsQueryParams{}
	queryParamImport.FabricID = vvFabricID
	queryParamImport.NetworkDeviceID = vvNetworkDeviceID
	queryParamImport.InterfaceName = vvInterfaceName
	queryParamImport.DataVLANName = vvDataVLANName
	queryParamImport.VoiceVLANName = vvVoiceVLANName
	item2, err := searchSdaGetPortAssignments(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["fabric_id"] = item2.FabricID
		resourceMap["network_device_id"] = item2.NetworkDeviceID
		resourceMap["interface_name"] = item2.InterfaceName
		resourceMap["data_vlan_name"] = item2.DataVLANName
		resourceMap["voice_vlan_name"] = item2.VoiceVLANName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaPortAssignmentsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddPortAssignments(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddPortAssignments", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddPortAssignments", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddPortAssignments", err))
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
				"Failure when executing AddPortAssignments", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetPortAssignmentsQueryParams{}
	queryParamValidate.FabricID = vvFabricID
	queryParamValidate.NetworkDeviceID = vvNetworkDeviceID
	queryParamValidate.InterfaceName = vvInterfaceName
	queryParamValidate.DataVLANName = vvDataVLANName
	queryParamValidate.VoiceVLANName = vvVoiceVLANName
	item3, err := searchSdaGetPortAssignments(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddPortAssignments", err,
			"Failure at AddPortAssignments, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["fabric_id"] = item3.FabricID
	resourceMap["network_device_id"] = item3.NetworkDeviceID
	resourceMap["interface_name"] = item3.InterfaceName
	resourceMap["data_vlan_name"] = item3.DataVLANName
	resourceMap["voice_vlan_name"] = item3.VoiceVLANName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaPortAssignmentsRead(ctx, d, m)
}

func resourceSdaPortAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvFabricID := resourceMap["fabric_id"]
	vvNetworkDeviceID := resourceMap["network_device_id"]
	vvInterfaceName := resourceMap["interface_name"]
	vvDataVLANName := resourceMap["data_vlan_name"]
	vvVoiceVLANName := resourceMap["voice_vlan_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortAssignments")
		queryParams1 := dnacentersdkgo.GetPortAssignmentsQueryParams{}
		queryParams1.FabricID = vvFabricID
		queryParams1.NetworkDeviceID = vvNetworkDeviceID
		queryParams1.InterfaceName = vvInterfaceName
		queryParams1.DataVLANName = vvDataVLANName
		queryParams1.VoiceVLANName = vvVoiceVLANName
		response1, restyResp1, err := client.Sda.GetPortAssignments(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		// Review flatten function used
		vItem1 := flattenSdaGetPortAssignmentsItems(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortAssignments search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaPortAssignmentsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaPortAssignmentsUpdatePortAssignments(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdatePortAssignments(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePortAssignments", err, restyResp1.String(),
					"Failure at UpdatePortAssignments, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePortAssignments", err,
				"Failure at UpdatePortAssignments, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatePortAssignments", err))
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
					"Failure when executing UpdatePortAssignments", err1))
				return diags
			}
		}

	}

	return resourceSdaPortAssignmentsRead(ctx, d, m)
}

func resourceSdaPortAssignmentsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeletePortAssignmentByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePortAssignmentByID", err, restyResp1.String(),
				"Failure at DeletePortAssignmentByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePortAssignmentByID", err,
			"Failure at DeletePortAssignmentByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletePortAssignmentByID", err))
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
				"Failure when executing DeletePortAssignmentByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaPortAssignmentsAddPortAssignments(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddPortAssignments {
	request := dnacentersdkgo.RequestSdaAddPortAssignments{}
	if v := expandRequestSdaPortAssignmentsAddPortAssignmentsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortAssignmentsAddPortAssignmentsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddPortAssignments {
	request := []dnacentersdkgo.RequestItemSdaAddPortAssignments{}
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
		i := expandRequestSdaPortAssignmentsAddPortAssignmentsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortAssignmentsAddPortAssignmentsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddPortAssignments {
	request := dnacentersdkgo.RequestItemSdaAddPortAssignments{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_device_type")))) {
		request.ConnectedDeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_vlan_name")))) {
		request.DataVLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".voice_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".voice_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".voice_vlan_name")))) {
		request.VoiceVLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_group_name")))) {
		request.SecurityGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_description")))) {
		request.InterfaceDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortAssignmentsUpdatePortAssignments(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdatePortAssignments {
	request := dnacentersdkgo.RequestSdaUpdatePortAssignments{}
	if v := expandRequestSdaPortAssignmentsUpdatePortAssignmentsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortAssignmentsUpdatePortAssignmentsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdatePortAssignments {
	request := []dnacentersdkgo.RequestItemSdaUpdatePortAssignments{}
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
		i := expandRequestSdaPortAssignmentsUpdatePortAssignmentsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortAssignmentsUpdatePortAssignmentsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdatePortAssignments {
	request := dnacentersdkgo.RequestItemSdaUpdatePortAssignments{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_device_type")))) {
		request.ConnectedDeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_vlan_name")))) {
		request.DataVLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".voice_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".voice_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".voice_vlan_name")))) {
		request.VoiceVLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_description")))) {
		request.InterfaceDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetPortAssignments(m interface{}, queryParams dnacentersdkgo.GetPortAssignmentsQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetPortAssignmentsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetPortAssignmentsResponse
	var ite *dnacentersdkgo.ResponseSdaGetPortAssignments
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetPortAssignments(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetPortAssignments(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else {
		ite, _, err = client.Sda.GetPortAssignments(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {

			foundItem = &item
			return foundItem, err

		}
		return foundItem, err
	}
}
