package dnacenter

import (
	"context"
	"errors"
	"log"
	"reflect"
	"time"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSitesAAASettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Network Settings.

- Set AAA settings for a site; *null* values indicate that the settings will be inherited from the parent site; empty
objects (*{}*) indicate that the settings is unset.
`,

		CreateContext: resourceSitesAAASettingsCreate,
		ReadContext:   resourceSitesAAASettingsRead,
		UpdateContext: resourceSitesAAASettingsUpdate,
		DeleteContext: resourceSitesAAASettingsDelete,
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

						"aaa_client": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"pan": &schema.Schema{
										Description: `Administration Node. Required for ISE.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"aaa_network": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"pan": &schema.Schema{
										Description: `Administration Node. Required for ISE.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
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

						"aaa_client": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"pan": &schema.Schema{
										Description: `Administration Node.  Required for ISE.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"aaa_network": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"pan": &schema.Schema{
										Description: `Administration Node. Required for ISE.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Site Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSitesAAASettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSitesAAASettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveAAASettingsForASite")
		vvID := vID
		queryParams1 := dnacentersdkgo.RetrieveAAASettingsForASiteQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrieveAAASettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveAAASettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveAAASettingsForASite response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSitesAAASettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSitesAAASettingsSetAAASettingsForASite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.SetAAASettingsForASite(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetAAASettingsForASite", err, restyResp1.String(),
					"Failure at SetAAASettingsForASite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetAAASettingsForASite", err,
				"Failure at SetAAASettingsForASite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SetAAASettingsForASite", err))
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
					"Failure when executing SetAAASettingsForASite", err1))
				return diags
			}
		}

	}

	return resourceSitesAAASettingsRead(ctx, d, m)
}

func resourceSitesAAASettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SitesAAASettings", err, "Delete method is not supported",
		"Failure at SitesAAASettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestSitesAAASettingsSetAAASettingsForASite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASite {
	request := dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_network")))) {
		request.AAANetwork = expandRequestSitesAAASettingsSetAAASettingsForASiteAAANetwork(ctx, key+".aaa_network.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_client")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_client")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_client")))) {
		request.AAAClient = expandRequestSitesAAASettingsSetAAASettingsForASiteAAAClient(ctx, key+".aaa_client.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesAAASettingsSetAAASettingsForASiteAAANetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASiteAAANetwork {
	request := dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASiteAAANetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_type")))) {
		request.ServerType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pan")))) {
		request.Pan = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_server_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_server_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_server_ip")))) {
		request.PrimaryServerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_server_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_server_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_server_ip")))) {
		request.SecondaryServerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesAAASettingsSetAAASettingsForASiteAAAClient(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASiteAAAClient {
	request := dnacentersdkgo.RequestNetworkSettingsSetAAASettingsForASiteAAAClient{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_type")))) {
		request.ServerType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pan")))) {
		request.Pan = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_server_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_server_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_server_ip")))) {
		request.PrimaryServerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_server_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_server_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_server_ip")))) {
		request.SecondaryServerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
