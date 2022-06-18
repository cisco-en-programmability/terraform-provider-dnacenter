package dnacenter

import (
	"context"
	"errors"
	"time"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceInterfaceOperationCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the
future more possible operations will be added to this API
`,

		CreateContext: resourceInterfaceOperationCreateCreate,
		ReadContext:   resourceInterfaceOperationCreateRead,
		DeleteContext: resourceInterfaceOperationCreateDelete,
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
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"deployment_mode": &schema.Schema{
							Description: `deploymentMode query parameter. Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"interface_uuid": &schema.Schema{
							Description: `interfaceUuid path parameter. Interface Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"operation": &schema.Schema{
							Description: `Operation`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"payload": &schema.Schema{
							Description: `Payload`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceOperationCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vInterfaceUUID := resourceItem["interface_uuid"]
	vDeploymentMode, okDeploymentMode := resourceItem["deployment_mode"]

	vvInterfaceUUID := vInterfaceUUID.(string)
	request1 := expandRequestInterfaceOperationCreateClearMacAddressTable(ctx, "parameters.0", d)
	queryParams1 := dnacentersdkgo.ClearMacAddressTableQueryParams{}

	if okDeploymentMode {
		queryParams1.DeploymentMode = vDeploymentMode.(string)
	}

	response1, restyResp1, err := client.Devices.ClearMacAddressTable(vvInterfaceUUID, request1, &queryParams1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ClearMacAddressTable", err,
			"Failure at ClearMacAddressTable, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ClearMacAddressTable", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ClearMacAddressTable", err1))
			return diags
		}
	}

	vItem1 := flattenDevicesClearMacAddressTableItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClearMacAddressTable response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceInterfaceOperationCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceInterfaceOperationCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestInterfaceOperationCreateClearMacAddressTable(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesClearMacAddressTable {
	request := dnacentersdkgo.RequestDevicesClearMacAddressTable{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request.Payload = expandRequestInterfaceOperationCreateClearMacAddressTablePayload(ctx, key+".payload.0", d)
	}
	return &request
}

func expandRequestInterfaceOperationCreateClearMacAddressTablePayload(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesClearMacAddressTablePayload {
	var request dnacentersdkgo.RequestDevicesClearMacAddressTablePayload
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenDevicesClearMacAddressTableItem(item *dnacentersdkgo.ResponseDevicesClearMacAddressTableResponse) []map[string]interface{} {
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
