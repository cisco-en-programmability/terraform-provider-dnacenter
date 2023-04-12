package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceProviderV2() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Network Settings.

- API to create Service Provider Profile(QOS).

- API to update Service Provider Profile (QoS).
`,

		CreateContext: resourceServiceProviderV2Create,
		ReadContext:   resourceServiceProviderV2Read,
		UpdateContext: resourceServiceProviderV2Update,
		DeleteContext: resourceServiceProviderV2Delete,
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

						"group_uuid": &schema.Schema{
							Description: `Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"key": &schema.Schema{
							Description: `Key`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"namespace": &schema.Schema{
							Description: `Namespace`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"sla_profile_name": &schema.Schema{
										Description: `Sla Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sp_profile_name": &schema.Schema{
										Description: `Sp Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"wan_provider": &schema.Schema{
										Description: `Wan Provider`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeInt,
							Computed:    true,
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

						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"qos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"model": &schema.Schema{
													Description: `Model`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"old_profile_name": &schema.Schema{
													Description: `Old Profile Name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"profile_name": &schema.Schema{
													Description: `Profile Name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"wan_provider": &schema.Schema{
													Description: `Wan Provider`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
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

func resourceServiceProviderV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	// resourceItem := *getResourceItem(d.Get("parameters"))
	vvSpProfileName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.settings"); ok {
			if _, ok := d.GetOk("parameters.0.settings.0"); ok {
				if _, ok := d.GetOk("parameters.0.settings.0.qos"); ok {
					if v, ok := d.GetOk("parameters.0.settings.0.qos.0.profile_name"); ok {
						vvSpProfileName = interfaceToString(v)
					}
				}
			}
		}
	}
	request1 := expandRequestServiceProviderV2CreateSpProfileV2(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	item2, err := searchNetworkSettingsGetServiceProviderDetailsV2(m, vvSpProfileName)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["profile_name"] = vvSpProfileName
		d.SetId(joinResourceID(resourceMap))
		return resourceServiceProviderV2Read(ctx, d, m)
	}
	resp1, restyResp1, err := client.NetworkSettings.CreateSpProfileV2(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSpProfileV2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSpProfileV2", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateSpProfileV2", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateSpProfileV2", err1))
			return diags
		}
	}
	item3, err := searchNetworkSettingsGetServiceProviderDetailsV2(m, vvSpProfileName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateSpProfileV2", err,
			"Failure at CreateSpProfileV2, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["profile_name"] = vvSpProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceServiceProviderV2Read(ctx, d, m)
}

func resourceServiceProviderV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvSpProfileName := resourceMap["profile_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetServiceProviderDetailsV2")

		item1, err := searchNetworkSettingsGetServiceProviderDetailsV2(m, vvSpProfileName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2Response{
			*item1,
		}
		vItem1 := flattenNetworkSettingsGetServiceProviderDetailsV2Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetServiceProviderDetailsV2 search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceServiceProviderV2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestServiceProviderV2UpdateSpProfileV2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateSpProfileV2(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSpProfileV2", err, restyResp1.String(),
					"Failure at UpdateSpProfileV2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSpProfileV2", err,
				"Failure at UpdateSpProfileV2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateSpProfileV2", err))
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
					"Failure when executing UpdateSpProfileV2", err1))
				return diags
			}
		}

	}

	return resourceServiceProviderV2Read(ctx, d, m)
}

func resourceServiceProviderV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete ServiceProviderV2 on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestServiceProviderV2CreateSpProfileV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2 {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2{}
	request.Settings = expandRequestServiceProviderV2CreateSpProfileV2Settings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2CreateSpProfileV2Settings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2Settings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderV2CreateSpProfileV2SettingsQosArray(ctx, key+".qos", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2CreateSpProfileV2SettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2SettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2SettingsQos{}
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
		i := expandRequestServiceProviderV2CreateSpProfileV2SettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2CreateSpProfileV2SettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2SettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileV2SettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2UpdateSpProfileV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2 {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2{}
	request.Settings = expandRequestServiceProviderV2UpdateSpProfileV2Settings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2UpdateSpProfileV2Settings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2Settings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderV2UpdateSpProfileV2SettingsQosArray(ctx, key+".qos", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2UpdateSpProfileV2SettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2SettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2SettingsQos{}
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
		i := expandRequestServiceProviderV2UpdateSpProfileV2SettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestServiceProviderV2UpdateSpProfileV2SettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2SettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileV2SettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".old_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".old_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".old_profile_name")))) {
		request.OldProfileName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsGetServiceProviderDetailsV2(m interface{}, vSProfileName string) (*dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2Response, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2Response
	var ite *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2
	ite, resty, err := client.NetworkSettings.GetServiceProviderDetailsV2()
	log.Printf("[DEBUG] Load data")
	if resty != nil {
		log.Printf("Failure when sarch => %s", resty.String())
	}
	log.Printf("[DEBUG] Load data1")
	if err != nil {
		return foundItem, err
	}
	log.Printf("[DEBUG] Load data2")
	items := ite
	if items == nil {
		return foundItem, err
	}
	log.Printf("[DEBUG] Load data3")
	if items.Response == nil {
		return foundItem, err
	}
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Value != nil {
			for _, item2 := range *item.Value {
				if item2.SpProfileName == vSProfileName {
					log.Printf("[DEBUG] Search finded item")
					var getItem *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2Response
					getItem = &item
					foundItem = getItem
					return foundItem, err
				}
			}
		}
	}
	log.Printf("[DEBUG] Search final")
	return foundItem, err
}
