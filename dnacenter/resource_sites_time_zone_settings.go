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

func resourceSitesTimeZoneSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Network Settings.

- Set time zone settings for a site; *null* values indicate that the setting will be inherited from the parent site;
empty objects (*{}*) indicate that the settings is unset.
`,

		CreateContext: resourceSitesTimeZoneSettingsCreate,
		ReadContext:   resourceSitesTimeZoneSettingsRead,
		UpdateContext: resourceSitesTimeZoneSettingsUpdate,
		DeleteContext: resourceSitesTimeZoneSettingsDelete,
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

						"time_zone": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"identifier": &schema.Schema{
										Description: `Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example : GMT
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

						"id": &schema.Schema{
							Description: `id path parameter. Site Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"time_zone": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"identifier": &schema.Schema{
										Description: `Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example: GMT
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

func resourceSitesTimeZoneSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSitesTimeZoneSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTimeZoneSettingsForASite")
		vvID := vID
		queryParams1 := dnacentersdkgo.RetrieveTimeZoneSettingsForASiteQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrieveTimeZoneSettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveTimeZoneSettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTimeZoneSettingsForASite response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSitesTimeZoneSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		request1 := expandRequestSitesTimeZoneSettingsSetTimeZoneForASite(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.SetTimeZoneForASite(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetTimeZoneForASite", err, restyResp1.String(),
					"Failure at SetTimeZoneForASite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetTimeZoneForASite", err,
				"Failure at SetTimeZoneForASite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SetTimeZoneForASite", err))
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
					"Failure when executing SetTimeZoneForASite", err1))
				return diags
			}
		}

	}

	return resourceSitesTimeZoneSettingsRead(ctx, d, m)
}

func resourceSitesTimeZoneSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SitesTimeZoneSettings", err, "Delete method is not supported",
		"Failure at SitesTimeZoneSettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestSitesTimeZoneSettingsSetTimeZoneForASite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetTimeZoneForASite {
	request := dnacentersdkgo.RequestNetworkSettingsSetTimeZoneForASite{}
	request.TimeZone = expandRequestSitesTimeZoneSettingsSetTimeZoneForASiteTimeZone(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTimeZoneSettingsSetTimeZoneForASiteTimeZone(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsSetTimeZoneForASiteTimeZone {
	request := dnacentersdkgo.RequestNetworkSettingsSetTimeZoneForASiteTimeZone{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identifier")))) {
		request.IDentifier = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
