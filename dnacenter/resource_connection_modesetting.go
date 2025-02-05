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

func resourceConnectionModesetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Licenses.

- Update Cisco Smart Software Manager (CSSM) connection mode for the system.
`,

		CreateContext: resourceConnectionModesettingCreate,
		ReadContext:   resourceConnectionModesettingRead,
		UpdateContext: resourceConnectionModesettingUpdate,
		DeleteContext: resourceConnectionModesettingDelete,
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

						"connection_mode": &schema.Schema{
							Description: `The CSSM connection modes of Catalyst Center are DIRECT, ON_PREMISE and SMART_PROXY
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_id": &schema.Schema{
										Description: `On-premise CSSM client id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"on_premise_host": &schema.Schema{
										Description: `On-premise host
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"smart_account_name": &schema.Schema{
										Description: `On-premise CSSM local smart account name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

						"connection_mode": &schema.Schema{
							Description: `The CSSM connection modes of Catalyst Center are DIRECT, ON_PREMISE and SMART_PROXY.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_id": &schema.Schema{
										Description: `On-premise CSSM client id
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"client_secret": &schema.Schema{
										Description: `On-premise CSSM client secret
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"on_premise_host": &schema.Schema{
										Description: `On-premise CSSM hostname or IP address
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smart_account_name": &schema.Schema{
										Description: `On-premise CSSM local smart account name
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

func resourceConnectionModesettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceConnectionModesettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesCSSMConnectionMode")

		response1, restyResp1, err := client.Licenses.RetrievesCSSMConnectionMode()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesRetrievesCSSMConnectionModeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesCSSMConnectionMode response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceConnectionModesettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestConnectionModesettingUpdateCSSMConnectionMode(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Licenses.UpdateCSSMConnectionMode(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateCSSMConnectionMode", err, restyResp1.String(),
					"Failure at UpdateCSSMConnectionMode, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateCSSMConnectionMode", err,
				"Failure at UpdateCSSMConnectionMode, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateCSSMConnectionMode", err))
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
					"Failure when executing UpdateCSSMConnectionMode", err1))
				return diags
			}
		}

	}

	return resourceConnectionModesettingRead(ctx, d, m)
}

func resourceConnectionModesettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete ConnectionModesetting on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestConnectionModesettingUpdateCSSMConnectionMode(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLicensesUpdateCSSMConnectionMode {
	request := dnacentersdkgo.RequestLicensesUpdateCSSMConnectionMode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_mode")))) {
		request.ConnectionMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameters")))) {
		request.Parameters = expandRequestConnectionModesettingUpdateCSSMConnectionModeParameters(ctx, key+".parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConnectionModesettingUpdateCSSMConnectionModeParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLicensesUpdateCSSMConnectionModeParameters {
	request := dnacentersdkgo.RequestLicensesUpdateCSSMConnectionModeParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".on_premise_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".on_premise_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".on_premise_host")))) {
		request.OnPremiseHost = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_name")))) {
		request.SmartAccountName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_id")))) {
		request.ClientID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_secret")))) {
		request.ClientSecret = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
