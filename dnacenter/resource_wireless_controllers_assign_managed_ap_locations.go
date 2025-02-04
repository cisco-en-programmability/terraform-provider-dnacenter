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
func resourceWirelessControllersAssignManagedApLocations() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action allows user to assign Managed AP Locations for WLC by device ID. The payload should always be
a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and
deletion.
`,

		CreateContext: resourceWirelessControllersAssignManagedApLocationsCreate,
		ReadContext:   resourceWirelessControllersAssignManagedApLocationsRead,
		DeleteContext: resourceWirelessControllersAssignManagedApLocationsDelete,
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
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task URL
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
						"device_id": &schema.Schema{
							Description: `deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"primary_managed_aplocations_site_ids": &schema.Schema{
							Description: `Site IDs of Primary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"secondary_managed_aplocations_site_ids": &schema.Schema{
							Description: `Site IDs of Secondary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
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

func resourceWirelessControllersAssignManagedApLocationsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDeviceID := resourceItem["device_id"]

	vvDeviceID := vDeviceID.(string)
	request1 := expandRequestWirelessControllersAssignManagedApLocationsAssignManagedApLocationsForWLC(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.AssignManagedApLocationsForWLC(vvDeviceID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignManagedApLocationsForWLC", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignManagedAPLocationsForWLC", err))
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
				"Failure when executing AssignManagedAPLocationsForWLC", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessAssignManagedApLocationsForWLCItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignManagedApLocationsForWLC response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessControllersAssignManagedApLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessControllersAssignManagedApLocationsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessControllersAssignManagedApLocationsAssignManagedApLocationsForWLC(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessAssignManagedApLocationsForWLC {
	request := dnacentersdkgo.RequestWirelessAssignManagedApLocationsForWLC{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_managed_aplocations_site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_managed_aplocations_site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_managed_aplocations_site_ids")))) {
		request.PrimaryManagedApLocationsSiteIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_managed_aplocations_site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_managed_aplocations_site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_managed_aplocations_site_ids")))) {
		request.SecondaryManagedApLocationsSiteIDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenWirelessAssignManagedApLocationsForWLCItem(item *dnacentersdkgo.ResponseWirelessAssignManagedApLocationsForWLCResponse) []map[string]interface{} {
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
