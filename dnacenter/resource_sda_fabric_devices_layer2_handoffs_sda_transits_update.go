package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on SDA.

- Updates layer 3 handoffs with sda transit of fabric devices based on user input.
`,

		CreateContext: resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateCreate,
		ReadContext:   resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateRead,
		DeleteContext: resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateDelete,
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
						"payload": &schema.Schema{
							Description: `Array of RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"affinity_id_decider": &schema.Schema{
										Description: `Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"affinity_id_prime": &schema.Schema{
										Description: `Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"connected_to_internet": &schema.Schema{
										Description: `Set this true to allow associated site to provide internet access to other sites through sd-access.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this device is assigned to. (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"is_multicast_over_transit_enabled": &schema.Schema{
										Description: `Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device ID of the fabric device. (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"transit_network_id": &schema.Schema{
										Description: `ID of the transit network of the layer 3 handoff sda transit. (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransit(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sda.UpdateFabricDevicesLayer3HandoffsWithSdaTransit(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateFabricDevicesLayer3HandoffsWithSdaTransit", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UpdateFabricDevicesLayer3HandoffsWithSdaTransit", err))
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
				"Failure when executing UpdateFabricDevicesLayer3HandoffsWithSdaTransit", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateFabricDevicesLayer3HandoffsWithSdaTransit response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransit(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit {
	request := dnacentersdkgo.RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit{}
	if v := expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransitItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransitItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit {
	request := []dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit{}
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
		i := expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransitItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsSdaTransitsUpdateUpdateFabricDevicesLayer3HandoffsWithSdaTransitItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_network_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_network_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_network_id")))) {
		request.TransitNetworkID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".affinity_id_prime")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".affinity_id_prime")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".affinity_id_prime")))) {
		request.AffinityIDPrime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".affinity_id_decider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".affinity_id_decider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".affinity_id_decider")))) {
		request.AffinityIDDecider = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_to_internet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_to_internet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_to_internet")))) {
		request.ConnectedToInternet = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_multicast_over_transit_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) {
		request.IsMulticastOverTransitEnabled = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitItem(item *dnacentersdkgo.ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitResponse) []map[string]interface{} {
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
