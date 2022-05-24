package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSiteAssign() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Sites.

- Assigns list of devices to a site
`,

		CreateContext: resourceSiteAssignCreate,
		ReadContext:   resourceSiteAssignRead,
		UpdateContext: resourceSiteAssignUpdate,
		DeleteContext: resourceSiteAssignDelete,
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

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"response": &schema.Schema{
										Description: `Response`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"site": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"response": &schema.Schema{
										Description: `Response`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"version": &schema.Schema{
										Description: `Version`,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": &schema.Schema{
										Description: `Device ip (eg: 10.104.240.64)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id to which site the device to assign
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

func resourceSiteAssignCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteAssignAssignDeviceToSite(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteID, okSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	if okSiteID && vvSiteID != "" {
		getResponse1, _, err := client.Sites.GetMembership(vvSiteID, nil)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["site_id"] = vvSiteID
			d.SetId(joinResourceID(resourceMap))
			return resourceSiteAssignRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Sites.AssignDeviceToSite(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AssignDeviceToSite", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AssignDeviceToSite", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteAssignRead(ctx, d, m)
}

func resourceSiteAssignRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]
	vDeviceFamily := resourceMap["device_family"]
	vSerialNumber := resourceMap["serial_number"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetMembership")
		vvSiteID := vSiteID
		queryParams1 := dnacentersdkgo.GetMembershipQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset
		}
		if okLimit {
			queryParams1.Limit = vLimit
		}
		if okDeviceFamily {
			queryParams1.DeviceFamily = vDeviceFamily
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber
		}

		response1, restyResp1, err := client.Sites.GetMembership(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMembership", err,
				"Failure at GetMembership, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetMembershipItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMembership response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSiteAssignUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSiteAssignRead(ctx, d, m)
}

func resourceSiteAssignDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SiteAssign on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSiteAssignAssignDeviceToSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDeviceToSite {
	request := dnacentersdkgo.RequestSitesAssignDeviceToSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestSiteAssignAssignDeviceToSiteDeviceArray(ctx, key+".device", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteAssignAssignDeviceToSiteDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice {
	request := []dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSiteAssignAssignDeviceToSiteDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteAssignAssignDeviceToSiteDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice {
	request := dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip")))) {
		request.IP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
