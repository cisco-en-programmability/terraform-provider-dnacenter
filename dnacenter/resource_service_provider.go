package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceProvider() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to create service provider profile(QOS).

- API to update SP profile
`,

		CreateContext: resourceServiceProviderCreate,
		ReadContext:   resourceServiceProviderRead,
		UpdateContext: resourceServiceProviderUpdate,
		DeleteContext: resourceServiceProviderDelete,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"qos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"model": &schema.Schema{
													Description: `Model`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"profile_name": &schema.Schema{
													Description: `Profile Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"wan_provider": &schema.Schema{
													Description: `Wan Provider`,
													Type:        schema.TypeString,
													Optional:    true,
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

func resourceServiceProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestServiceProviderCreateSpProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

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

	item, err := searchNetworkSettingsGetServiceProviderDetails(m, vvSpProfileName)
	if err == nil && item != nil {
		resourceMap := make(map[string]string)
		resourceMap["profile_name"] = vvSpProfileName
		d.SetId(joinResourceID(resourceMap))
		return resourceServiceProviderRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.NetworkSettings.CreateSpProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSpProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSpProfile", err))
		return diags
	}

	executionId := resp1.ExecutionID
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
				"Failure when executing CreateSpProfile", err,
				"Failure at CreateSpProfile execution", bapiError))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["profile_name"] = vvSpProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceServiceProviderRead(ctx, d, m)
}

func resourceServiceProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSPProfleName := resourceMap["profile_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetServiceProviderDetails")

		response1, err := searchNetworkSettingsGetServiceProviderDetails(m, vSPProfleName)

		if err != nil || response1 == nil {
			if err != nil {
				log.Printf("[DEBUG] Error => %s", err.Error())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenNetworkSettingsGetServiceProviderDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetServiceProviderDetails search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceServiceProviderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	isProfileNameChange := false
	newProfileName := ""
	vSpProfileName := resourceMap["profile_name"]
	item, err := searchNetworkSettingsGetServiceProviderDetails(m, vSpProfileName)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetServiceProviderDetails", err,
			"Failure at yGetApplications, unexpected response", ""))
		return diags
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vSpProfileName)
		request1 := expandRequestServiceProviderUpdateSpProfile(ctx, "parameters.0", d)
		newQos := *request1.Settings.Qos
		if d.HasChange("parameters.0.settings.0.qos.0.profile_name") {
			old, _ := d.GetChange("parameters.0.settings.0.qos.0.profile_name")
			isProfileNameChange = true
			newQos[0].OldProfileName = interfaceToString(old)
			request1.Settings.Qos = &newQos
			newProfileName = newQos[0].ProfileName
		}
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateSpProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSpProfile", err, restyResp1.String(),
					"Failure at UpdateSpProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSpProfile", err,
				"Failure at UpdateSpProfile, unexpected response", ""))
			return diags
		}
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
					"Failure when executing UpdateSpProfile", err,
					"Failure at UpdateSpProfile execution", bapiError))
				return diags
			}
		}
	}
	if isProfileNameChange {
		resourceMap := make(map[string]string)
		resourceMap["profile_name"] = newProfileName
		d.SetId(joinResourceID(resourceMap))
	}
	return resourceServiceProviderRead(ctx, d, m)
}

func resourceServiceProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSpProfileName := resourceMap["profile_name"]
	item, err := searchNetworkSettingsGetServiceProviderDetails(m, vSpProfileName)
	if err != nil || item == nil {
		d.SetId("")
		return diags
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteSpProfile")

		response1, restyResp1, err := client.NetworkSettings.DeleteSpProfile(vSpProfileName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteSpProfile", err,
				"Failure at DeleteSpProfile, unexpected response", ""))
			return diags
		}

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
					"Failure when executing DeleteSpProfile", err,
					"Failure at DeleteSpProfile execution", bapiError))
				return diags
			}
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		d.SetId("")
		return diags

	}
	return diags
}
func expandRequestServiceProviderCreateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfile{}
	request.Settings = expandRequestServiceProviderCreateSpProfileSettings(ctx, key+".settings.0", d)

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderCreateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
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
		i := expandRequestServiceProviderCreateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile{}
	request.Settings = expandRequestServiceProviderUpdateSpProfileSettings(ctx, key+".settings.0", d)

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderUpdateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
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
		i := expandRequestServiceProviderUpdateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}

	return &request
}

func searchNetworkSettingsGetServiceProviderDetails(m interface{}, vSProfileName string) (*dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsResponse, error) {
	log.Printf("[DEBUG] Search")
	log.Printf("[DEBUG] Search sp profile name =>%s", vSProfileName)
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsResponse
	var ite *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetails
	ite, resty, err := client.NetworkSettings.GetServiceProviderDetails()
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
					var getItem *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsResponse
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
