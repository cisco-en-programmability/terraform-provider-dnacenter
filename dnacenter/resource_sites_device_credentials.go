package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSitesDeviceCredentials() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Network Settings.

- Updates device credential settings for a site; *null* values indicate that the setting will be inherited from the
parent site; empty objects (*{}*) indicate that the credential is unset, and that no credential of that type will be
used for the site.
`,

		CreateContext: resourceSitesDeviceCredentialsCreate,
		ReadContext:   resourceSitesDeviceCredentialsRead,
		UpdateContext: resourceSitesDeviceCredentialsUpdate,
		DeleteContext: resourceSitesDeviceCredentialsDelete,
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

						"cli_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"http_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"http_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmpv2c_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmpv2c_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmpv3_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
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

						"cli_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"http_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"http_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Site Id, retrievable from the *id* attribute in */dna/intent/api/v1/sites*
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"snmpv2c_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"snmpv2c_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"snmpv3_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
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

func resourceSitesDeviceCredentialsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSitesDeviceCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCredentialSettingsForASite")
		vvID := vID
		queryParams1 := dnacentersdkgo.GetDeviceCredentialSettingsForASiteQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.GetDeviceCredentialSettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCredentialSettingsForASite response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSitesDeviceCredentialsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateDeviceCredentialSettingsForASite(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceCredentialSettingsForASite", err, restyResp1.String(),
					"Failure at UpdateDeviceCredentialSettingsForASite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceCredentialSettingsForASite", err,
				"Failure at UpdateDeviceCredentialSettingsForASite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateDeviceCredentialSettingsForASite", err))
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
					"Failure when executing UpdateDeviceCredentialSettingsForASite", err1))
				return diags
			}
		}

	}

	return resourceSitesDeviceCredentialsRead(ctx, d, m)
}

func resourceSitesDeviceCredentialsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SitesDeviceCredentials", err, "Delete method is not supported",
		"Failure at SitesDeviceCredentialsDelete, unexpected response", ""))
	return diags
}
func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASite {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credentials_id")))) {
		request.CliCredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteCliCredentialsID(ctx, key+".cli_credentials_id.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmpv2c_read_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmpv2c_read_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmpv2c_read_credentials_id")))) {
		request.SNMPv2CReadCredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID(ctx, key+".snmpv2c_read_credentials_id.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmpv2c_write_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmpv2c_write_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmpv2c_write_credentials_id")))) {
		request.SNMPv2CWriteCredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID(ctx, key+".snmpv2c_write_credentials_id.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmpv3_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmpv3_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmpv3_credentials_id")))) {
		request.SNMPv3CredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID(ctx, key+".snmpv3_credentials_id.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read_credentials_id")))) {
		request.HTTPReadCredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID(ctx, key+".http_read_credentials_id.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write_credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write_credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write_credentials_id")))) {
		request.HTTPWriteCredentialsID = expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID(ctx, key+".http_write_credentials_id.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteCliCredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteCliCredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteCliCredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDeviceCredentialsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credentials_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credentials_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credentials_id")))) {
		request.CredentialsID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
