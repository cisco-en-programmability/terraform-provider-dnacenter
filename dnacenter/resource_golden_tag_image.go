package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGoldenTagImage() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Software Image Management (SWIM).

- Golden Tag image. Set siteId as -1 for Global site.

- Remove golden tag. Set siteId as -1 for Global site.
`,

		CreateContext: resourceGoldenTagImageCreate,
		ReadContext:   resourceGoldenTagImageRead,
		UpdateContext: resourceGoldenTagImageUpdate,
		DeleteContext: resourceGoldenTagImageDelete,
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

						"device_role": &schema.Schema{
							Description: `Device Role. Possible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_site_id": &schema.Schema{
							Description: `Inherited Site Id. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the uuid of the site from where the tag is inherited. In case the golden tag is inherited from the Global site the value will be "-1".
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_site_name": &schema.Schema{
							Description: `Inherited Site Name. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the name of the site from where the tag is inherited.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tagged_golden": &schema.Schema{
							Description: `Tagged Golden.
`,
							// Type:        schema.TypeBool,
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

						"device_family_identifier": &schema.Schema{
							Description: `Device Family Identifier e.g. : 277696480-283933147, 277696480
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_role": &schema.Schema{
							Description: `Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"image_id": &schema.Schema{
							Description: `imageId in uuid format.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `SiteId in uuid format. For Global Site "-1" to be used.
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

func resourceGoldenTagImageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGoldenTagImageTagAsGoldenImage(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteID, okSiteID := resourceItem["site_id"]
	vvSiteID := ""
	if okSiteID {
		vvSiteID = interfaceToString(vSiteID)
	}

	vDeviceFamilyIDentifier, okDeviceFamilyIDentifier := resourceItem["device_family_identifier"]
	vvDeviceFamilyIDentifier := ""
	if okDeviceFamilyIDentifier {
		vvDeviceFamilyIDentifier = interfaceToString(vDeviceFamilyIDentifier)
	}

	vDeviceRole, okDeviceRole := resourceItem["device_role"]
	vvDeviceRole := ""
	if okDeviceRole {
		vvDeviceRole = interfaceToString(vDeviceRole)
	}

	vImageID, okImageID := resourceItem["image_id"]
	vvImageID := ""
	if okImageID {
		vvImageID = interfaceToString(vImageID)
	}
	if okSiteID && vvSiteID != "" && okDeviceFamilyIDentifier && vvDeviceFamilyIDentifier != "" && okDeviceRole && vvDeviceRole != "" && okImageID && vvImageID != "" {
		getResponse1, _, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["site_id"] = vvSiteID
			resourceMap["device_family_identifier"] = vvDeviceFamilyIDentifier
			resourceMap["device_role"] = vvDeviceRole
			resourceMap["image_id"] = vvImageID
			d.SetId(joinResourceID(resourceMap))
			return resourceGoldenTagImageRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.SoftwareImageManagementSwim.TagAsGoldenImage(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing TagAsGoldenImage", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing TagAsGoldenImage", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	resourceMap["device_family_identifier"] = vvDeviceFamilyIDentifier
	resourceMap["device_role"] = vvDeviceRole
	resourceMap["image_id"] = vvImageID
	d.SetId(joinResourceID(resourceMap))
	return resourceGoldenTagImageRead(ctx, d, m)
}

func resourceGoldenTagImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSiteID := resourceMap["site_id"]

	vDeviceFamilyIDentifier := resourceMap["device_family_identifier"]

	vDeviceRole := resourceMap["device_role"]

	vImageID := resourceMap["image_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGoldenTagStatusOfAnImage")
		vvSiteID := vSiteID
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier
		vvDeviceRole := vDeviceRole
		vvImageID := vImageID

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGoldenTagStatusOfAnImage response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceGoldenTagImageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceGoldenTagImageRead(ctx, d, m)
}

func resourceGoldenTagImageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvSiteID := resourceMap["site_id"]

	vvDeviceFamilyIDentifier := resourceMap["device_family_identifier"]

	vvDeviceRole := resourceMap["device_role"]

	vvImageID := resourceMap["image_id"]
	response1, restyResp1, err := client.SoftwareImageManagementSwim.RemoveGoldenTagForImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing RemoveGoldenTagForImage", err, restyResp1.String(),
				"Failure at RemoveGoldenTagForImage, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RemoveGoldenTagForImage", err,
			"Failure at RemoveGoldenTagForImage, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing RemoveGoldenTagForImage", err))
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
				"Failure when executing RemoveGoldenTagForImage", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGoldenTagImageTagAsGoldenImage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTagAsGoldenImage {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimTagAsGoldenImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_role")))) {
		request.DeviceRole = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_family_identifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_family_identifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_family_identifier")))) {
		request.DeviceFamilyIDentifier = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
