package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Deploy authentication template in SDA Fabric

- Update default authentication profile in SDA Fabric

- Add default authentication profile in SDA Fabric
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

						"authenticate_template_id": &schema.Schema{
							Description: `Authenticate Template Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name. Allowed values are 'No Authentication ', 'Open Authentication', 'Closed Authentication', 'Low Impact'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_order": &schema.Schema{
							Description: `Authentication Order. Allowed values are 'dot1x ', 'mac'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `In a network that includes both devices that support and devices that do not support IEEE 802.1X, MAB can be deployed as a fallback, or complementary, mechanism to IEEE 802.1X. If the network does not have any IEEE 802.1X-capable devices, MAB can be deployed as a standalone authentication mechanism (e.g. [3-120])
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"number_of_hosts": &schema.Schema{
							Description: `Number of hosts specifies the number of data hosts that can be connected to a port. With Single selected, you can have only one data client  on the port. With Unlimited selected, you can have multiple data clients and one voice client on the port
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy should be a valid fabric site name hierarchy. e.g Global/USA/San Jose
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"wake_on_lan": &schema.Schema{
							Description: `The IEEE 802.1X Wake on LAN (WoL) Support feature allows dormant systems to be powered up when the  switch receives a specific Ethernet frame. You can use this feature in cases when hosts on power save and needs to receive a  magic packet to turn them on. This feature works on a per subnet basis and send the subnet broadcast to all hosts in the subnet
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
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

	request1 := expandRequestSdaFabricAuthenticationProfileDeployAuthenticationTemplateInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	vAuthenticateTemplateName := resourceItem["authenticate_template_name"]
	vvAuthenticateTemplateName := interfaceToString(vAuthenticateTemplateName)

	queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParams1.SiteNameHierarchy = vvSiteNameHierarchy
	queryParams1.AuthenticateTemplateName = vvAuthenticateTemplateName
	getResponse2, _, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)
	if err == nil && getResponse2 != nil && getResponse2.SiteNameHierarchy != "" {
		resourceMap := make(map[string]string)
		resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
		resourceMap["authenticate_template_name"] = vvAuthenticateTemplateName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Sda.DeployAuthenticationTemplateInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing DeployAuthenticationTemplateInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing DeployAuthenticationTemplateInSdaFabric", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
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
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeployAuthenticationTemplateInSdaFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
	resourceMap["authenticate_template_name"] = vvAuthenticateTemplateName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSdaFabricAuthenticationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy, okAuthenticateTemplateName := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDefaultAuthenticationProfileFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		if okAuthenticateTemplateName {
			queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
		}

		response1, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDefaultAuthenticationProfileFromSdaFabric", err,
			// 	"Failure at GetDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
			// return diags
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

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
	item, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

	if err != nil || item == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDefaultAuthenticationProfileFromSdaFabric", err,
			"Failure at GetDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
		return diags
	}

	vvName := item.SiteNameHierarchy
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabric(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
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
			response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
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
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
	item, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

	if err != nil || item == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetDefaultAuthenticationProfileFromSdaFabric", err,
		"Failure at GetDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))*/
		d.SetId("")
		return diags
	}

	queryParams2 := dnacentersdkgo.DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams{}
	queryParams2.SiteNameHierarchy = item.SiteNameHierarchy
	response1, restyResp1, err := client.Sda.DeleteDefaultAuthenticationProfileFromSdaFabric(&queryParams2)
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
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
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
func expandRequestSdaFabricAuthenticationProfileDeployAuthenticationTemplateInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaDeployAuthenticationTemplateInSdaFabric {
	request := dnacentersdkgo.RequestSdaDeployAuthenticationTemplateInSdaFabric{}
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
