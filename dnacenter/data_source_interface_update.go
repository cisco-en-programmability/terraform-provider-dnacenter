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
func dataSourceInterfaceUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.

- Add/Update Interface description, VLAN membership and change Interface admin status ('UP'/'DOWN') from Request body.
`,

		ReadContext: dataSourceInterfaceUpdateRead,
		Schema: map[string]*schema.Schema{
			"deployment_mode": &schema.Schema{
				Description: `deploymentMode query parameter. Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_uuid": &schema.Schema{
				Description: `interfaceUuid path parameter. Interface ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"task_id": &schema.Schema{
										Description: `Task Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"url": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"required": &schema.Schema{
							Description: `Required`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInterfaceUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInterfaceUUID := d.Get("interface_uuid")
	vDeploymentMode, okDeploymentMode := d.GetOk("deployment_mode")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: UpdateInterfaceDetails")
		vvInterfaceUUID := vInterfaceUUID.(string)
		request1 := expandRequestInterfaceUpdateUpdateInterfaceDetails(ctx, "", d)
		queryParams1 := dnacentersdkgo.UpdateInterfaceDetailsQueryParams{}

		if okDeploymentMode {
			queryParams1.DeploymentMode = vDeploymentMode.(string)
		}

		response1, restyResp1, err := client.Devices.UpdateInterfaceDetails(vvInterfaceUUID, request1, &queryParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateInterfaceDetails", err,
				"Failure at UpdateInterfaceDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesUpdateInterfaceDetailsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateInterfaceDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestInterfaceUpdateUpdateInterfaceDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateInterfaceDetails {
	request := dnacentersdkgo.RequestDevicesUpdateInterfaceDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = expandRequestInterfaceUpdateUpdateInterfaceDetailsDescription(ctx, key+".description.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_status")))) {
		request.AdminStatus = expandRequestInterfaceUpdateUpdateInterfaceDetailsAdminStatus(ctx, key+".admin_status.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = expandRequestInterfaceUpdateUpdateInterfaceDetailsVLANID(ctx, key+".vlan_id.0", d)
	}
	return &request
}

func expandRequestInterfaceUpdateUpdateInterfaceDetailsDescription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsDescription {
	request := dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsDescription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	return &request
}

func expandRequestInterfaceUpdateUpdateInterfaceDetailsAdminStatus(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsAdminStatus {
	request := dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsAdminStatus{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	return &request
}

func expandRequestInterfaceUpdateUpdateInterfaceDetailsVLANID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsVLANID {
	request := dnacentersdkgo.RequestDevicesUpdateInterfaceDetailsVLANID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	return &request
}

func flattenDevicesUpdateInterfaceDetailsItem(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["properties"] = flattenDevicesUpdateInterfaceDetailsItemProperties(item.Properties)
	respItem["required"] = item.Required
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesUpdateInterfaceDetailsItemProperties(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = flattenDevicesUpdateInterfaceDetailsItemPropertiesURL(item.URL)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesUpdateInterfaceDetailsItemPropertiesURL(item *dnacentersdkgo.ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
