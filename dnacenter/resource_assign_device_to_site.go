package dnacenter

import (
	"context"

	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAssignDeviceToSite() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Assigns unassigned devices to a site. This data source action does not move assigned devices to other sites.
`,

		CreateContext: resourceAssignDeviceToSiteCreate,
		ReadContext:   resourceAssignDeviceToSiteRead,
		DeleteContext: resourceAssignDeviceToSiteDelete,
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

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site Id where device(s) needs to be assigned
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": &schema.Schema{
										Description: `Device IP. It can be either IPv4 or IPv6. IPV4 e.g., 10.104.240.64. IPV6 e.g., 2001:420:284:2004:4:181:500:183
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceAssignDeviceToSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vSiteID := resourceItem["site_id"]

	vRunsync := resourceItem["runsync"]

	vTimeout := resourceItem["timeout"]

	vPersistbapioutput := resourceItem["persistbapioutput"]

	vvSiteID := vSiteID.(string)
	request1 := expandRequestAssignDeviceToSiteAssignDevicesToSite(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.AssignDevicesToSiteHeaderParams{}

	headerParams1.Runsync = vRunsync.(string)

	headerParams1.Timeout = vTimeout.(string)

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sites.AssignDevicesToSite(vvSiteID, request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignDevicesToSite", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AssignDevicesToSite", err,
				"Failure at AssignDevicesToSite execution", bapiError))
			return diags
		}
	}

	vItem1 := flattenSitesAssignDevicesToSiteItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignDevicesToSite response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceAssignDeviceToSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssignDeviceToSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssignDeviceToSiteAssignDevicesToSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDevicesToSite {
	request := dnacentersdkgo.RequestSitesAssignDevicesToSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestAssignDeviceToSiteAssignDevicesToSiteDeviceArray(ctx, key+".device", d)
	}
	return &request
}

func expandRequestAssignDeviceToSiteAssignDevicesToSiteDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesAssignDevicesToSiteDevice {
	request := []dnacentersdkgo.RequestSitesAssignDevicesToSiteDevice{}
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
		i := expandRequestAssignDeviceToSiteAssignDevicesToSiteDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssignDeviceToSiteAssignDevicesToSiteDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDevicesToSiteDevice {
	request := dnacentersdkgo.RequestSitesAssignDevicesToSiteDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip")))) {
		request.IP = interfaceToString(v)
	}
	return &request
}

func flattenSitesAssignDevicesToSiteItem(item *dnacentersdkgo.ResponseSitesAssignDevicesToSite) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
