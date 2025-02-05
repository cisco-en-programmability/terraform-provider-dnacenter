package dnacenter

import (
	"context"
	"errors"
	"log"
	"reflect"
	"time"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSitesDNSSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Network Settings.

- Set DNS settings for a site; *null* values indicate that the setting will be inherited from the parent site; empty
objects (*{}*) indicate that the settings is unset.
`,

		CreateContext: resourceSitesDNSSettingsCreate,
		ReadContext:   resourceSitesDNSSettingsRead,
		UpdateContext: resourceSitesDNSSettingsUpdate,
		DeleteContext: resourceSitesDNSSettingsDelete,
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

						"dns": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_servers": &schema.Schema{
										Description: `DNS servers for hostname resolution.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"domain_name": &schema.Schema{
										Description: `Network's domain name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
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

						"dns": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_servers": &schema.Schema{
										Description: `DNS servers for hostname resolution.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"domain_name": &schema.Schema{
										Description: `Network's domain name. Example : myCompnay.com
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

func resourceSitesDNSSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceSitesDNSSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveDNSSettingsForASite")
		vvID := vID
		queryParams1 := dnacentersdkgo.RetrieveDNSSettingsForASiteQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrieveDNSSettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveDNSSettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveDNSSettingsForASite response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSitesDNSSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSitesDNSSettingsSetDNSSettingsForASite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.SetDNSSettingsForASite(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetDNSSettingsForASite", err, restyResp1.String(),
					"Failure at SetDNSSettingsForASite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetDNSSettingsForASite", err,
				"Failure at SetDNSSettingsForASite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SetDNSSettingsForASite", err))
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
					"Failure when executing SetDNSSettingsForASite", err1))
				return diags
			}
		}

	}

	return resourceSitesDNSSettingsRead(ctx, d, m)
}

func resourceSitesDNSSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SitesDNSSettings", err, "Delete method is not supported",
		"Failure at SitesDNSSettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestSitesDNSSettingsSetDNSSettingsForASite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetDNSSettingsForASite {
	request := dnacentersdkgo.RequestNetworkSettingsSetDNSSettingsForASite{}
	request.DNS = expandRequestSitesDNSSettingsSetDNSSettingsForASiteDNS(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesDNSSettingsSetDNSSettingsForASiteDNS(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetDNSSettingsForASiteDNS {
	request := dnacentersdkgo.RequestNetworkSettingsSetDNSSettingsForASiteDNS{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain_name")))) {
		request.DomainName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_servers")))) {
		request.DNSServers = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
