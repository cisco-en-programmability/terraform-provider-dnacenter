package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPamServerSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on System Settings.

- Creates configuration details of the external IPAM server. You should only create one external IPAM server; delete any
existing external server before creating a new one.

- Deletes configuration details of the external IPAM server.

- Updates configuration details of the external IPAM server.
`,

		CreateContext: resourceIPamServerSettingCreate,
		ReadContext:   resourceIPamServerSettingRead,
		UpdateContext: resourceIPamServerSettingUpdate,
		DeleteContext: resourceIPamServerSettingDelete,
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

						"provider": &schema.Schema{
							Description: `Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"server_name": &schema.Schema{
							Description: `A descriptive name of this external server, used for identification purposes`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"server_url": &schema.Schema{
							Description: `The URL of this external server`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"state": &schema.Schema{
							Description: `State of the the external IPAM.* OK indicates success of most recent periodic communication check with external IPAM.* CRITICAL indicates failure of most recent attempt to communicate with the external IPAM.* SYNCHRONIZING indicates that the process of synchronizing the external IPAM database with the local IPAM database is running and all other IPAM processes will be blocked until the completes.* DISCONNECTED indicates the external IPAM is no longer being used.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"user_name": &schema.Schema{
							Description: `The external IPAM server login username`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"view": &schema.Schema{
							Description: `The view under which pools are created in the external IPAM server.`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"password": &schema.Schema{
							Description: `The password for the external IPAM server login username`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Computed:    true,
						},
						"provider": &schema.Schema{
							Description: `Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"server_name": &schema.Schema{
							Description: `A descriptive name of this external server, used for identification purposes`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"server_url": &schema.Schema{
							Description: `The URL of this external server`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"sync_view": &schema.Schema{
							Description: `Synchronize the IP pools from the local IPAM to this external server`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"user_name": &schema.Schema{
							Description: `The external IPAM server login username`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"view": &schema.Schema{
							Description: `The view under which pools are created in the external IPAM server.`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIPamServerSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestIPamServerSettingCreatesConfigurationDetailsOfTheExternalIPAMServer(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	item2, _, err := client.SystemSettings.RetrievesConfigurationDetailsOfTheExternalIPAMServer()
	if err == nil && item2 != nil {
		return resourceIPamServerSettingRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.SystemSettings.CreatesConfigurationDetailsOfTheExternalIPAMServer(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesConfigurationDetailsOfTheExternalIPAMServer", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesConfigurationDetailsOfTheExternalIPAMServer", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreatesConfigurationDetailsOfTheExternalIPAMServer", err))
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
				"Failure when executing CreatesConfigurationDetailsOfTheExternalIPAMServer", err1))
			return diags
		}
	}
	item3, _, err := client.SystemSettings.RetrievesConfigurationDetailsOfTheExternalIPAMServer()
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreatesConfigurationDetailsOfTheExternalIPAMServer", err,
			"Failure at CreatesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)

	d.SetId(joinResourceID(resourceMap))
	return resourceIPamServerSettingRead(ctx, d, m)
}

func resourceIPamServerSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesConfigurationDetailsOfTheExternalIPAMServer")

		response1, restyResp1, err := client.SystemSettings.RetrievesConfigurationDetailsOfTheExternalIPAMServer()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesConfigurationDetailsOfTheExternalIPAMServer response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceIPamServerSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestIPamServerSettingUpdatesConfigurationDetailsOfTheExternalIPAMServer(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SystemSettings.UpdatesConfigurationDetailsOfTheExternalIPAMServer(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesConfigurationDetailsOfTheExternalIPAMServer", err, restyResp1.String(),
					"Failure at UpdatesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesConfigurationDetailsOfTheExternalIPAMServer", err,
				"Failure at UpdatesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesConfigurationDetailsOfTheExternalIPAMServer", err))
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
					"Failure when executing UpdatesConfigurationDetailsOfTheExternalIPAMServer", err1))
				return diags
			}
		}

	}

	return resourceIPamServerSettingRead(ctx, d, m)
}

func resourceIPamServerSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	response1, restyResp1, err := client.SystemSettings.DeletesConfigurationDetailsOfTheExternalIPAMServer()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesConfigurationDetailsOfTheExternalIPAMServer", err, restyResp1.String(),
				"Failure at DeletesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesConfigurationDetailsOfTheExternalIPAMServer", err,
			"Failure at DeletesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesConfigurationDetailsOfTheExternalIPAMServer", err))
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
				"Failure when executing DeletesConfigurationDetailsOfTheExternalIPAMServer", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestIPamServerSettingCreatesConfigurationDetailsOfTheExternalIPAMServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer {
	request := dnacentersdkgo.RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_url")))) {
		request.ServerURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view")))) {
		request.View = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_view")))) {
		request.SyncView = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPamServerSettingUpdatesConfigurationDetailsOfTheExternalIPAMServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer {
	request := dnacentersdkgo.RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_url")))) {
		request.ServerURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view")))) {
		request.View = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_view")))) {
		request.SyncView = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
