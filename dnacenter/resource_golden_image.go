package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGoldenImage() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Software Image Management (SWIM).

- Golden Tag image. Set siteId as -1 for Global site.

- Remove golden tag. Set siteId as -1 for Global site.
`,

		CreateContext: resourceGoldenImageCreate,
		ReadContext:   resourceGoldenImageRead,
		UpdateContext: resourceGoldenImageUpdate,
		DeleteContext: resourceGoldenImageDelete,
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

							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_family_identifier": &schema.Schema{
							Description: `Device Family Identifier e.g. : 277696480-283933147, 277696480
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_role": &schema.Schema{
							Description: `Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"image_id": &schema.Schema{
							Description: `imageId in uuid format.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_id": &schema.Schema{
							Description: `SiteId in uuid format. For Global Site "-1" to be used.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGoldenImageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGoldenImageTagAsGoldenImage(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteID, okSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vDeviceFamilyIDentifier, okDeviceFamilyIDentifier := resourceItem["device_family_identifier"]
	vvDeviceFamilyIDentifier := interfaceToString(vDeviceFamilyIDentifier)
	vDeviceRole, okDeviceRole := resourceItem["device_role"]
	vvDeviceRole := interfaceToString(vDeviceRole)
	vImageID, okImageID := resourceItem["image_id"]
	vvImageID := interfaceToString(vImageID)
	if okSiteID && vvSiteID != "" && okDeviceFamilyIDentifier && vvDeviceFamilyIDentifier != "" && okDeviceRole && vvDeviceRole != "" && okImageID && vvImageID != "" {
		getResponse1, _, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)
		if err == nil && getResponse1 != nil && getResponse1.Response != nil && getResponse1.Response.TaggedGolden != nil {
			if *getResponse1.Response.TaggedGolden {
				resourceMap := make(map[string]string)
				resourceMap["site_id"] = vvSiteID
				resourceMap["device_family_identifier"] = vvDeviceFamilyIDentifier
				resourceMap["device_role"] = vvDeviceRole
				resourceMap["image_id"] = vvImageID
				d.SetId(joinResourceID(resourceMap))
				return resourceGoldenImageRead(ctx, d, m)
			}
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
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing TagAsGoldenImage", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing TagAsGoldenImage", err1))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	resourceMap["device_family_identifier"] = vvDeviceFamilyIDentifier
	resourceMap["device_role"] = vvDeviceRole
	resourceMap["image_id"] = vvImageID
	d.SetId(joinResourceID(resourceMap))
	return resourceGoldenImageRead(ctx, d, m)
}

func resourceGoldenImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetGoldenTagStatusOfAnImage")
		vvSiteID := vSiteID
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier
		vvDeviceRole := vDeviceRole
		vvImageID := vImageID

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if !d.IsNewResource() && err != nil || response1 == nil {
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

func resourceGoldenImageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing GoldenImageUpdate", err, "Update method is not supported",
		"Failure at GoldenImageUpdate, unexpected response", ""))

	return diags
}

func resourceGoldenImageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vDeviceFamilyIDentifier := resourceMap["device_family_identifier"]
	vDeviceRole := resourceMap["device_role"]
	vImageID := resourceMap["image_id"]

	selectedMethod := 1
	//var vvID string
	//var vvName string
	if selectedMethod == 1 {
		//vvID = vID
		getResp, _, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vSiteID, vDeviceFamilyIDentifier, vDeviceRole, vImageID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.SoftwareImageManagementSwim.RemoveGoldenTagForImage(vSiteID, vDeviceFamilyIDentifier, vDeviceRole, vImageID)
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
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
func expandRequestGoldenImageTagAsGoldenImage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTagAsGoldenImage {
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
