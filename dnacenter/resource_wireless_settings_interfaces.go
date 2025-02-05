package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessSettingsInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- This resource allows the user to create an interface

- This resource allows the user to delete an interface by ID

- This resource allows the user to update an interface by ID
`,

		CreateContext: resourceWirelessSettingsInterfacesCreate,
		ReadContext:   resourceWirelessSettingsInterfacesRead,
		UpdateContext: resourceWirelessSettingsInterfacesUpdate,
		DeleteContext: resourceWirelessSettingsInterfacesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"id": &schema.Schema{
							Description: `Interface ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Description: `VLAN ID
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. Interface ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Description: `VLAN ID range is 1-4094
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessSettingsInterfacesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsInterfacesCreateInterface(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["interface_name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.Wireless.GetInterfaceByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsInterfacesRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.GetInterfacesQueryParams{}

		response2, err := searchWirelessGetInterfaces(m, queryParamImport, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response2.ID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsInterfacesRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Wireless.CreateInterface(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateInterface", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateInterface", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateInterface", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateInterface", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetInterfacesQueryParams{}
	item3, err := searchWirelessGetInterfaces(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateInterface", err,
			"Failure at CreateInterface, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsInterfacesRead(ctx, d, m)
}

func resourceWirelessSettingsInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInterfaceByID")
		vvID := vID

		response1, restyResp1, err := client.Wireless.GetInterfaceByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenWirelessGetInterfaceByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInterfaces search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessSettingsInterfacesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestWirelessSettingsInterfacesUpdateInterface(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateInterface(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateInterface", err, restyResp1.String(),
					"Failure at UpdateInterface, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateInterface", err,
				"Failure at UpdateInterface, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateInterface", err))
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateInterface", err1))
				return diags
			}
		}

	}

	return resourceWirelessSettingsInterfacesRead(ctx, d, m)
}

func resourceWirelessSettingsInterfacesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Wireless.DeleteInterface(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteInterface", err, restyResp1.String(),
				"Failure at DeleteInterface, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteInterface", err,
			"Failure at DeleteInterface, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteInterface", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteInterface", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessSettingsInterfacesCreateInterface(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateInterface {
	request := dnacentersdkgo.RequestWirelessCreateInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsInterfacesUpdateInterface(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateInterface {
	request := dnacentersdkgo.RequestWirelessUpdateInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetInterfaces(m interface{}, queryParams dnacentersdkgo.GetInterfacesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGetInterfacesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGetInterfacesResponse

	queryParams.Offset = 1
	nResponse, _, err := client.Wireless.GetInterfaces(nil)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.InterfaceName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Wireless.GetInterfaces(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
