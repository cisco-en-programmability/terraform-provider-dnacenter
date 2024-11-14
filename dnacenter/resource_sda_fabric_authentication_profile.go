package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Add default authentication template in SDA Fabric

- Update default authentication profile in SDA Fabric

- Delete default authentication profile in SDA Fabric
`,

		CreateContext: resourceSdaFabricAuthenticationProfileCreate,
		ReadContext:   resourceSdaFabricAuthenticationProfileRead,
		UpdateContext: resourceSdaFabricAuthenticationProfileUpdate,
		DeleteContext: resourceSdaFabricAuthenticationProfileDelete,
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

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication_order": &schema.Schema{
							Description: `Authentication Order
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Authenticate Template info reterieved successfully in sda fabric site
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `Dot1x To Mab Fallback Timeout
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"number_of_hosts": &schema.Schema{
							Description: `Number Of Hosts
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"wake_on_lan": &schema.Schema{
							Description: `Wake On Lan
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddDefaultAuthenticationTemplateInSDAFabric`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication_order": &schema.Schema{
							Description: `Authentication Order
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `Dot1x To MabFallback Timeout( Allowed range is [3-120])
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"number_of_hosts": &schema.Schema{
							Description: `Number Of Hosts
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wake_on_lan": &schema.Schema{
							Description: `Wake On Lan
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricAuthenticationProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabric(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	queryParamImport := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParamImport.SiteNameHierarchy = vvSiteNameHierarchy
	item2, _, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_name_hierarchy"] = item2.SiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddDefaultAuthenticationTemplateInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err))
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
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParamValidate.SiteNameHierarchy = vvSiteNameHierarchy
	item3, _, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err,
			"Failure at AddDefaultAuthenticationTemplateInSdaFabric, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["site_name_hierarchy"] = item3.SiteNameHierarchy

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSdaFabricAuthenticationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDefaultAuthenticationProfileFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetDefaultAuthenticationProfileFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDefaultAuthenticationProfileFromSdaFabric response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSdaFabricAuthenticationProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabric(ctx, "parameters", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.UpdateDefaultAuthenticationProfileInSdaFabric(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDefaultAuthenticationProfileInSdaFabric", err, restyResp1.String(),
					"Failure at UpdateDefaultAuthenticationProfileInSdaFabric, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDefaultAuthenticationProfileInSdaFabric", err,
				"Failure at UpdateDefaultAuthenticationProfileInSdaFabric, unexpected response", ""))
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
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
			for response2.Status == "IN_PROGRESS" {
				time.Sleep(10 * time.Second)
				response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response2 == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if response2.Status == "FAILURE" {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing UpdateDefaultAuthenticationProfileInSdaFabric", err))
				return diags
			}
		}

	}

	return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSdaFabricAuthenticationProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := dnacentersdkgo.DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams{}

	vvSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	queryParamDelete.SiteNameHierarchy = vvSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteDefaultAuthenticationProfileFromSdaFabric(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDefaultAuthenticationProfileFromSdaFabric", err, restyResp1.String(),
				"Failure at DeleteDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDefaultAuthenticationProfileFromSdaFabric", err,
			"Failure at DeleteDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
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
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteDefaultAuthenticationProfileFromSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddDefaultAuthenticationTemplateInSdaFabric{}
	if v := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := []dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric{}
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
		i := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := dnacentersdkgo.RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
	if v := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := []dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
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
		i := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_order")))) {
		request.AuthenticationOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot1x_to_mab_fallback_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) {
		request.Dot1XToMabFallbackTimeout = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wake_on_lan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wake_on_lan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wake_on_lan")))) {
		request.WakeOnLan = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_hosts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_hosts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_hosts")))) {
		request.NumberOfHosts = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
