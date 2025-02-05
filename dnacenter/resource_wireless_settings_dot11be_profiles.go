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

func resourceWirelessSettingsDot11BeProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- This resource allows the user to create a 802.11be Profile.Catalyst Center will push this profile to device's
"default-dot11be-profile”.Also please note , 802.11be Profile is supported only on IOS-XE controllers since device
version 17.15

- This resource allows the user to delete a 802.11be Profile,if the 802.11be Profile is not mapped to any Wireless
Network Profile

- This resource allows the user to update a 802.11be Profile
`,

		CreateContext: resourceWirelessSettingsDot11BeProfilesCreate,
		ReadContext:   resourceWirelessSettingsDot11BeProfilesRead,
		UpdateContext: resourceWirelessSettingsDot11BeProfilesUpdate,
		DeleteContext: resourceWirelessSettingsDot11BeProfilesDelete,
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

						"default": &schema.Schema{
							Description: `Is 802.11be Profile marked as default in System . (Read only field)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `802.11be Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"mu_mimo_down_link": &schema.Schema{
							Description: `MU-MIMO Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"mu_mimo_up_link": &schema.Schema{
							Description: `MU-MIMO Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ofdma_down_link": &schema.Schema{
							Description: `OFDMA Downlink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ofdma_multi_ru": &schema.Schema{
							Description: `OFDMA Multi-RU
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ofdma_up_link": &schema.Schema{
							Description: `OFDMA Uplink
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Description: `802.11be Profile Name
`,
							Type:     schema.TypeString,
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
							Description: `id path parameter. 802.11be Profile ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"mu_mimo_down_link": &schema.Schema{
							Description: `MU-MIMO Downlink (Default: false)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"mu_mimo_up_link": &schema.Schema{
							Description: `MU-MIMO Uplink (Default: false)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ofdma_down_link": &schema.Schema{
							Description: `OFDMA Downlink (Default: true)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ofdma_multi_ru": &schema.Schema{
							Description: `OFDMA Multi-RU (Default: false)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ofdma_up_link": &schema.Schema{
							Description: `OFDMA Uplink (Default: true)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"profile_name": &schema.Schema{
							Description: `802.11be Profile Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessSettingsDot11BeProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsDot11BeProfilesCreateA80211BeProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["profile_name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.Wireless.Get80211BeProfileByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsDot11BeProfilesRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.Get80211BeProfilesQueryParams{}

		response2, err := searchWirelessGetAll80211BeProfiles(m, queryParamImport, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response2.ID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsDot11BeProfilesRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Wireless.CreateA80211BeProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateA80211BeProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateA80211BeProfile", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateA80211BeProfile", err))
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
				"Failure when executing CreateA80211BeProfile", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.Get80211BeProfilesQueryParams{}
	item3, err := searchWirelessGetAll80211BeProfiles(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateA80211BeProfile", err,
			"Failure at CreateA80211BeProfile, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsDot11BeProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsDot11BeProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: Get80211BeProfileByID")
		vvID := vID

		response1, restyResp1, err := client.Wireless.Get80211BeProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		// Review flatten function used
		vItem1 := flattenWirelessGet80211BeProfileByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAll80211BeProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessSettingsDot11BeProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestWirelessSettingsDot11BeProfilesUpdate80211BeProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.Update80211BeProfile(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing Update80211BeProfile", err, restyResp1.String(),
					"Failure at Update80211BeProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Update80211BeProfile", err,
				"Failure at Update80211BeProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing Update80211BeProfile", err))
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
					"Failure when executing Update80211BeProfile", err1))
				return diags
			}
		}

	}

	return resourceWirelessSettingsDot11BeProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsDot11BeProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Wireless.DeleteA80211BeProfile(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteA80211BeProfile", err, restyResp1.String(),
				"Failure at DeleteA80211BeProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteA80211BeProfile", err,
			"Failure at DeleteA80211BeProfile, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteA80211BeProfile", err))
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
				"Failure when executing DeleteA80211BeProfile", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessSettingsDot11BeProfilesCreateA80211BeProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateA80211BeProfile {
	request := dnacentersdkgo.RequestWirelessCreateA80211BeProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_multi_ru")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) {
		request.OfdmaMultiRu = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsDot11BeProfilesUpdate80211BeProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdate80211BeProfile {
	request := dnacentersdkgo.RequestWirelessUpdate80211BeProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_multi_ru")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) {
		request.OfdmaMultiRu = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetAll80211BeProfiles(m interface{}, queryParams dnacentersdkgo.Get80211BeProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGet80211BeProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGet80211BeProfilesResponse
	// var ite *dnacentersdkgo.ResponseWirelessGetAll80211BeProfiles

	queryParams.Offset = 1
	nResponse, _, err := client.Wireless.Get80211BeProfiles(nil)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.ProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Wireless.Get80211BeProfiles(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
