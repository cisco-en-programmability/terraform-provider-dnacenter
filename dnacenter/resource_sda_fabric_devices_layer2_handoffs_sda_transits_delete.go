package dnacenter

import (
	"context"

	"errors"

	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on SDA.

- Deletes layer 3 handoffs with sda transit of a fabric device based on user input.
`,

		CreateContext: resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteCreate,
		ReadContext:   resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteRead,
		DeleteContext: resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteDelete,
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
							Description: `ID of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task status lookup url.
`,
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
						"fabric_id": &schema.Schema{
							Description: `fabricId query parameter. ID of the fabric this device belongs to.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"network_device_id": &schema.Schema{
							Description: `networkDeviceId query parameter. Network device ID of the fabric device.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vFabricID := resourceItem["fabric_id"]

	vNetworkDeviceID := resourceItem["network_device_id"]

	queryParams1 := dnacentersdkgo.DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams{}

	queryParams1.FabricID = vFabricID.(string)

	queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sda.DeleteFabricDeviceLayer3HandoffsWithSdaTransit(&queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricDeviceLayer3HandoffsWithSdaTransit", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricDeviceLayer3HandoffsWithSdaTransit", err))
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
				"Failure when executing DeleteFabricDeviceLayer3HandoffsWithSdaTransit", err1))
			return diags
		}
	}
	vItem1 := flattenSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeleteFabricDeviceLayer3HandoffsWithSdaTransit response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitItem(item *dnacentersdkgo.ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitResponse) []map[string]interface{} {
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
