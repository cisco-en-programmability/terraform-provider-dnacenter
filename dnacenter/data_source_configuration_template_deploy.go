package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceConfigurationTemplateDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- API to deploy a template.
`,

		ReadContext: dataSourceConfigurationTemplateDeployRead,
		Schema: map[string]*schema.Schema{
			"force_push_template": &schema.Schema{
				// Type:     schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"is_composite": &schema.Schema{
				Description: `Composite template flag
`,
				// Type:        schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deployment_id": &schema.Schema{
							Description: `UUID of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_name": &schema.Schema{
							Description: `Name of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"detailed_status_message": &schema.Schema{
										Description: `Device detailed status message
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_id": &schema.Schema{
										Description: `UUID of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"duration": &schema.Schema{
										Description: `Total duration of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_time": &schema.Schema{
										Description: `EndTime of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": &schema.Schema{
										Description: `Identifier of device based on the target type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `Device IPAddress
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Description: `StartTime of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Current status of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_type": &schema.Schema{
										Description: `Target type of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"duration": &schema.Schema{
							Description: `Total deployment duration
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_time": &schema.Schema{
							Description: `Deployment end time
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Description: `Deployment start time
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Current status of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_message": &schema.Schema{
							Description: `Status message of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Description: `Name of template deployed
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_version": &schema.Schema{
							Description: `Version of template deployed
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"main_template_id": &schema.Schema{
				Description: `Main template UUID of versioned template
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"member_template_deployment_info": &schema.Schema{
				Description: `memberTemplateDeploymentInfo
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"host_name": &schema.Schema{
							Description: `Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `UUID of target is required if targetType is MANAGED_DEVICE_UUID
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"params": &schema.Schema{
							Description: `Template params/values to be provisioned
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_params": &schema.Schema{
							Description: `Resource params to be provisioned
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Target type of device
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"versioned_template_id": &schema.Schema{
							Description: `Versioned templateUUID to be provisioned
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"template_id": &schema.Schema{
				Description: `UUID of template to be provisioned
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceConfigurationTemplateDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeployTemplate")
		request1 := expandRequestConfigurationTemplateDeployDeployTemplate(ctx, "", d)

		response1, restyResp1, err := client.ConfigurationTemplates.DeployTemplate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployTemplate", err,
				"Failure at DeployTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesDeployTemplateItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestConfigurationTemplateDeployDeployTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_push_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_push_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_push_template")))) {
		request.ForcePushTemplate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_composite")))) {
		request.IsComposite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".main_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".main_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".main_template_id")))) {
		request.MainTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_template_deployment_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_template_deployment_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_template_deployment_info")))) {
		request.MemberTemplateDeploymentInfo = expandRequestConfigurationTemplateDeployDeployTemplateMemberTemplateDeploymentInfoArray(ctx, key+".member_template_deployment_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_info")))) {
		request.TargetInfo = expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoArray(ctx, key+".target_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateMemberTemplateDeploymentInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateDeployDeployTemplateMemberTemplateDeploymentInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateMemberTemplateDeploymentInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfo{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateDeployDeployTemplateTargetInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateTargetInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfo {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoParams(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoResourceParamsArray(ctx, key+".resource_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".versioned_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".versioned_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".versioned_template_id")))) {
		request.VersionedTemplateID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoResourceParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoResourceParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployDeployTemplateTargetInfoResourceParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenConfigurationTemplatesDeployTemplateItem(item *dnacentersdkgo.ResponseConfigurationTemplatesDeployTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["deployment_name"] = item.DeploymentName
	respItem["devices"] = flattenConfigurationTemplatesDeployTemplateItemDevices(item.Devices)
	respItem["duration"] = item.Duration
	respItem["end_time"] = item.EndTime
	respItem["project_name"] = item.ProjectName
	respItem["start_time"] = item.StartTime
	respItem["status"] = item.Status
	respItem["status_message"] = item.StatusMessage
	respItem["template_name"] = item.TemplateName
	respItem["template_version"] = item.TemplateVersion
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesDeployTemplateItemDevices(items *[]dnacentersdkgo.ResponseConfigurationTemplatesDeployTemplateDevices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["detailed_status_message"] = item.DetailedStatusMessage
		respItem["device_id"] = item.DeviceID
		respItem["duration"] = item.Duration
		respItem["end_time"] = item.EndTime
		respItem["identifier"] = item.IDentifier
		respItem["ip_address"] = item.IPAddress
		respItem["name"] = item.Name
		respItem["start_time"] = item.StartTime
		respItem["status"] = item.Status
		respItem["target_type"] = item.TargetType
		respItems = append(respItems, respItem)
	}
	return respItems
}
