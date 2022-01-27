package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGoldenImageCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Golden Tag image. Set siteId as -1 for Global site.
`,

		ReadContext: dataSourceGoldenImageCreateRead,
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5 
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"site_id": &schema.Schema{
				Description: `SiteId in uuid format. For Global Site "-1" to be used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGoldenImageCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: TagAsGoldenImage")
		request1 := expandRequestGoldenImageCreateTagAsGoldenImage(ctx, "", d)

		response1, restyResp1, err := client.SoftwareImageManagementSwim.TagAsGoldenImage(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing TagAsGoldenImage", err,
				"Failure at TagAsGoldenImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSoftwareImageManagementSwimTagAsGoldenImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting TagAsGoldenImage response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestGoldenImageCreateTagAsGoldenImage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTagAsGoldenImage {
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

func flattenSoftwareImageManagementSwimTagAsGoldenImageItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimTagAsGoldenImageResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
