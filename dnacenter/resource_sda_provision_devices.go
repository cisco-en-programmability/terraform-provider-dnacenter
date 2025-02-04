package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaProvisionDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Delete provisioned devices based on query parameters.

- Provisions network devices to respective Sites based on user input.

- Re-provisions network devices to the site based on the user input.

- Deletes provisioned device based on Id.
`,

		CreateContext: resourceSdaProvisionDevicesCreate,
		ReadContext:   resourceSdaProvisionDevicesRead,
		UpdateContext: resourceSdaProvisionDevicesUpdate,
		DeleteContext: resourceSdaProvisionDevicesDelete,
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
							Description: `ID of the provisioned device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_device_id": &schema.Schema{
							Description: `ID of the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `ID of the site this device is provisioned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaProvisionDevices`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `ID of the provisioned device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `ID of network device to be provisioned.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"site_id": &schema.Schema{
										Description: `ID of the site this network device needs to be provisioned.
`,
										Type:     schema.TypeString,
										Optional: true,
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

func resourceSdaProvisionDevicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaProvisionDevicesProvisionDevices(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vNetworkID := resourceItem["network_device_id"]
	vvNetworkDeviceID := interfaceToString(vNetworkID)
	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	queryParamImport := dnacentersdkgo.GetProvisionedDevicesQueryParams{}
	queryParamImport.NetworkDeviceID = vvNetworkDeviceID
	queryParamImport.SiteID = vvSiteID
	item2, err := searchSdaGetProvisionedDevices(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["network_device_id"] = item2.NetworkDeviceID
		resourceMap["site_id"] = item2.SiteID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaProvisionDevicesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.ProvisionDevices(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ProvisionDevices", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ProvisionDevices", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ProvisionDevices", err))
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
				"Failure when executing ProvisionDevices", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetProvisionedDevicesQueryParams{}
	queryParamValidate.NetworkDeviceID = vvNetworkDeviceID
	queryParamValidate.SiteID = vvSiteID
	item3, err := searchSdaGetProvisionedDevices(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ProvisionDevices", err,
			"Failure at ProvisionDevices, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["network_device_id"] = item3.NetworkDeviceID
	resourceMap["site_id"] = item3.SiteID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaProvisionDevicesRead(ctx, d, m)
}

func resourceSdaProvisionDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	vvNetworkDeviceID := resourceMap["network_device_id"]
	vvSiteID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProvisionedDevices")
		queryParams1 := dnacentersdkgo.GetProvisionedDevicesQueryParams{}
		queryParams1.NetworkDeviceID = vvNetworkDeviceID
		queryParams1.SiteID = vvSiteID
		item1, err := searchSdaGetProvisionedDevices(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseSdaGetProvisionedDevicesResponse{
			*item1,
		}

		vItem1 := flattenSdaGetProvisionedDevicesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisionedDevices search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaProvisionDevicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaProvisionDevicesReProvisionDevices(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.ReProvisionDevices(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ReProvisionDevices", err, restyResp1.String(),
					"Failure at ReProvisionDevices, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReProvisionDevices", err,
				"Failure at ReProvisionDevices, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing ReProvisionDevices", err))
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
					"Failure when executing ReProvisionDevices", err1))
				return diags
			}
		}

	}

	return resourceSdaProvisionDevicesRead(ctx, d, m)
}

func resourceSdaProvisionDevicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteProvisionedDeviceByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteProvisionedDeviceByID", err, restyResp1.String(),
				"Failure at DeleteProvisionedDeviceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteProvisionedDeviceByID", err,
			"Failure at DeleteProvisionedDeviceByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteProvisionedDeviceByID", err))
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
				"Failure when executing DeleteProvisionedDeviceByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaProvisionDevicesProvisionDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaProvisionDevices {
	request := dnacentersdkgo.RequestSdaProvisionDevices{}
	if v := expandRequestSdaProvisionDevicesProvisionDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDevicesProvisionDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaProvisionDevices {
	request := []dnacentersdkgo.RequestItemSdaProvisionDevices{}
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
		i := expandRequestSdaProvisionDevicesProvisionDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDevicesProvisionDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaProvisionDevices {
	request := dnacentersdkgo.RequestItemSdaProvisionDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDevicesReProvisionDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaReProvisionDevices {
	request := dnacentersdkgo.RequestSdaReProvisionDevices{}
	if v := expandRequestSdaProvisionDevicesReProvisionDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDevicesReProvisionDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaReProvisionDevices {
	request := []dnacentersdkgo.RequestItemSdaReProvisionDevices{}
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
		i := expandRequestSdaProvisionDevicesReProvisionDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDevicesReProvisionDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaReProvisionDevices {
	request := dnacentersdkgo.RequestItemSdaReProvisionDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetProvisionedDevices(m interface{}, queryParams dnacentersdkgo.GetProvisionedDevicesQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetProvisionedDevicesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetProvisionedDevicesResponse
	var ite *dnacentersdkgo.ResponseSdaGetProvisionedDevices
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetProvisionedDevices(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetProvisionedDevices(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.NetworkDeviceID != "" {
		ite, _, err = client.Sda.GetProvisionedDevices(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.NetworkDeviceID == queryParams.NetworkDeviceID {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
