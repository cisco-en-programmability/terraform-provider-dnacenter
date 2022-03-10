package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImageDistribution() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).
	- Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it
	can be distributed
`,

		CreateContext: resourceImageDistributionCreate,
		ReadContext:   resourceImageDistributionRead,
		DeleteContext: resourceImageDistributionDelete,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"image_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceImageDistributionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: ClaimADeviceToASite")
	request1 := expandRequestSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx, "parameters.0", d)

	response1, restyResp1, err := client.SoftwareImageManagementSwim.TriggerSoftwareImageDistribution(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing TriggerSoftwareImageDistribution", err,
			"Failure at TriggerSoftwareImageDistribution, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing TriggerSoftwareImageDistribution", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil || response2.Response == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if !strings.Contains(response2.Response.Progress, "Starting") && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			err1 := errors.New(response2.Response.Progress)
			diags = append(diags, diagError(
				"Failure when executing TriggerSoftwareImageDistribution", err1))
			return diags
		}
		for strings.Contains(response2.Response.Progress, "Starting") {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err = client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil || response2.Response == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.Progress)
				err1 := errors.New(response2.Response.Progress)
				diags = append(diags, diagError(
					"Failure when executing TriggerSoftwareImageDistribution", err1))
				return diags
			}
		}
		if *response2.Response.IsError {
			log.Printf("[DEBUG] Error %s", response2.Response.Progress)
			err1 := errors.New(response2.Response.Progress)
			diags = append(diags, diagError(
				"Failure when executing TriggerSoftwareImageDistribution", err1))
			return diags
		}
	}

	vItem1 := flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TriggerSoftwareImageDistribution response",
			err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceImageDistributionRead(ctx, d, m)
}

func resourceImageDistributionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceImageDistributionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
func expandRequestSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
	if v := expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := []dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
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
		i := expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuid")))) {
		request.DeviceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_uuid")))) {
		request.ImageUUID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse) []map[string]interface{} {
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
