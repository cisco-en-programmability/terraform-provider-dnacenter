package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeployTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Software Image Management (SWIM).

- Golden Tag image. Set siteId as -1 for Global site.

- Remove golden tag. Set siteId as -1 for Global site.
`,

		CreateContext: resourceDeployTemplateCreate,
		ReadContext:   resourceDeployTemplateRead,
		UpdateContext: resourceDeployTemplateUpdate,
		DeleteContext: resourceDeployTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"force_push_template": &schema.Schema{

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_composite": &schema.Schema{
							Description: `Composite template flag
			`,

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
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
				},
			},
		},
	}
}

func resourceDeployTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: DeployTemplate")
	request1 := expandRequestConfigurationTemplateDeployV2DeployTemplateV2(ctx, "parameters.0", d)

	response1, restyResp1, err := client.ConfigurationTemplates.DeployTemplateV2(request1)

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

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplication", err))
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
			diags = append(diags, diagError(
				"Failure when executing CreateApplication", err))
			return diags
		}
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	d.SetId(getUnixTimeString())

	return resourceDeployTemplateRead(ctx, d, m)
}

func resourceDeployTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func resourceDeployTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDeployTemplateRead(ctx, d, m)
}

func resourceDeployTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
func expandRequestConfigurationTemplateDeployV2DeployTemplateV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2 {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2{}
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
		request.MemberTemplateDeploymentInfo = expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx, key+".member_template_deployment_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_info")))) {
		request.TargetInfo = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoArray(ctx, key+".target_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoParams(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParamsArray(ctx, key+".resource_params", d)
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

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
