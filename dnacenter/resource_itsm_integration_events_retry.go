package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceItsmIntegrationEventsRetry() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on ITSM.
		- Allows retry of multiple failed ITSM event instances. The retry request payload can be given as a list of strings:
		["instance1","instance2","instance3",..] A minimum of one instance Id is mandatory. The list of failed event instance
		Ids can be retrieved using the 'Get Failed ITSM Events' API in the 'instanceId' attribute.
	
`,

		CreateContext: resourceItsmIntegrationEventsRetryCreate,
		ReadContext:   resourceItsmIntegrationEventsRetryRead,
		DeleteContext: resourceItsmIntegrationEventsRetryDelete,

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

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
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
						"payload": &schema.Schema{
							Description: `Array of RequestItsmRetryIntegrationEvents`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceItsmIntegrationEventsRetryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: RetryIntegrationEvents")
	request1 := expandRequestItsmIntegrationEventsRetryRetryIntegrationEvents(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Itsm.RetryIntegrationEvents(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RetryIntegrationEvents", err,
			"Failure at RetryIntegrationEvents, unexpected response", ""))
		return diags
	}

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetryIntegrationEvents", err,
				"Failure at RetryIntegrationEvents execution", bapiError))
			return diags
		}
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenItsmRetryIntegrationEventsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetryIntegrationEvents response",
			err))
		return diags
	}
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceItsmIntegrationEventsRetryRead(ctx, d, m)
}

func resourceItsmIntegrationEventsRetryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceItsmIntegrationEventsRetryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestItsmIntegrationEventsRetryRetryIntegrationEvents(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItsmRetryIntegrationEvents {
	request := dnacentersdkgo.RequestItsmRetryIntegrationEvents{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenItsmRetryIntegrationEventsItem(item *dnacentersdkgo.ResponseItsmRetryIntegrationEvents) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
