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
func dataSourceSwimTriggerActivation() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Activates a software image on a given device. Software image must be present in the device flash
`,

		ReadContext: dataSourceSwimTriggerActivationRead,
		Schema: map[string]*schema.Schema{
			"client_type": &schema.Schema{
				Description: `Client-Type header parameter. Client-type (Optional)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_url": &schema.Schema{
				Description: `Client-Url header parameter. Client-url (Optional)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_validate": &schema.Schema{
				Description: `scheduleValidate query parameter. scheduleValidate, validates data before schedule (Optional)
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"payload": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate_lower_image_version": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"device_upgrade_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"distribute_if_needed": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"image_uuid_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"smu_image_uuid_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSwimTriggerActivationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vScheduleValidate, okScheduleValidate := d.GetOk("schedule_validate")
	vClientType, okClientType := d.GetOk("client_type")
	vClientURL, okClientURL := d.GetOk("client_url")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: TriggerSoftwareImageActivation")
		request1 := expandRequestSwimTriggerActivationTriggerSoftwareImageActivation(ctx, "", d)

		headerParams1 := dnacentersdkgo.TriggerSoftwareImageActivationHeaderParams{}
		queryParams1 := dnacentersdkgo.TriggerSoftwareImageActivationQueryParams{}

		if okScheduleValidate {
			queryParams1.ScheduleValidate = vScheduleValidate.(bool)
		}

		if okClientType {
			headerParams1.ClientType = vClientType.(string)
		}

		if okClientURL {
			headerParams1.ClientURL = vClientURL.(string)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.TriggerSoftwareImageActivation(request1, &headerParams1, &queryParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing TriggerSoftwareImageActivation", err,
				"Failure at TriggerSoftwareImageActivation, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimTriggerSoftwareImageActivationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting TriggerSoftwareImageActivation response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSwimTriggerActivationTriggerSoftwareImageActivation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation{}
	if v := expandRequestSwimTriggerActivationTriggerSoftwareImageActivationItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimTriggerActivationTriggerSoftwareImageActivationItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation {
	request := []dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation{}
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
		i := expandRequestSwimTriggerActivationTriggerSoftwareImageActivationItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimTriggerActivationTriggerSoftwareImageActivationItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation {
	request := dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".activate_lower_image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".activate_lower_image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".activate_lower_image_version")))) {
		request.ActivateLowerImageVersion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_upgrade_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_upgrade_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_upgrade_mode")))) {
		request.DeviceUpgradeMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuid")))) {
		request.DeviceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".distribute_if_needed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".distribute_if_needed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".distribute_if_needed")))) {
		request.DistributeIfNeeded = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_uuid_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_uuid_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_uuid_list")))) {
		request.ImageUUIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smu_image_uuid_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smu_image_uuid_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smu_image_uuid_list")))) {
		request.SmuImageUUIDList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSoftwareImageManagementSwimTriggerSoftwareImageActivationItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
