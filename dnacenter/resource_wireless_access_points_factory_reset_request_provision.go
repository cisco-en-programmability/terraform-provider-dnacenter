package dnacenter

import (
	"context"

	"errors"

	"time"

	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessAccessPointsFactoryResetRequestProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action is used to factory reset Access Points. It is supported for maximum 100 Access Points per
request. Factory reset clears all configurations from the Access Points. After factory reset the Access Point may become
unreachable from the currently associated Wireless Controller and may or may not join back the same controller.
`,

		CreateContext: resourceWirelessAccessPointsFactoryResetRequestProvisionCreate,
		ReadContext:   resourceWirelessAccessPointsFactoryResetRequestProvisionRead,
		DeleteContext: resourceWirelessAccessPointsFactoryResetRequestProvisionDelete,
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
						"ap_mac_addresses": &schema.Schema{
							Description: `List of Access Point's Ethernet MAC addresses, set maximum 100 ethernet MAC addresses per request.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"keep_static_ipconfig": &schema.Schema{
							Description: `Set the value of keepStaticIPConfig to false, to clear all configurations from Access Points and set the value of keepStaticIPConfig to true, to clear all configurations from Access Points without clearing static IP configuration.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessAccessPointsFactoryResetRequestProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessAccessPointsFactoryResetRequestProvisionFactoryResetAccessPoints(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.FactoryResetAccessPoints(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing FactoryResetAccessPoints", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing FactoryResetAccessPoints", err))
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
				"Failure when executing FactoryResetAccessPoints", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessFactoryResetAccessPointsItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting FactoryResetAccessPoints response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessAccessPointsFactoryResetRequestProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessAccessPointsFactoryResetRequestProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessAccessPointsFactoryResetRequestProvisionFactoryResetAccessPoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessFactoryResetAccessPoints {
	request := dnacentersdkgo.RequestWirelessFactoryResetAccessPoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".keep_static_ipconfig")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".keep_static_ipconfig")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".keep_static_ipconfig")))) {
		request.KeepStaticIPConfig = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_mac_addresses")))) {
		request.ApMacAddresses = interfaceToSliceString(v)
	}
	return &request
}

func flattenWirelessFactoryResetAccessPointsItem(item *dnacentersdkgo.ResponseWirelessFactoryResetAccessPointsResponse) []map[string]interface{} {
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
