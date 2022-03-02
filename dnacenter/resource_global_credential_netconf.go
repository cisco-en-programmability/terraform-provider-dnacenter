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

func resourceGlobalCredentialNetconf() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Updates global netconf credentials

- Adds global netconf credentials

- Deletes global credential for the given ID
`,

		CreateContext: resourceGlobalCredentialNetconfCreate,
		ReadContext:   resourceGlobalCredentialNetconfRead,
		UpdateContext: resourceGlobalCredentialNetconfUpdate,
		DeleteContext: resourceGlobalCredentialNetconfDelete,
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

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"secure": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateNetconfCredentials`,
				Type:        schema.TypeList,
				MaxItems:    1,
				MinItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGlobalCredentialNetconfCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalCredentialNetconfCreateNetconfCredentials(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vNetconfPort := resourceItem["netconf_port"]
	vvNetconfPort := interfaceToString(vNetconfPort)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "NETCONF"
	item, err := searchDiscoveryGetGlobalCredentialsNetConf(m, queryParams1, vvNetconfPort, vvID)
	if item != nil || err != nil {
		resourceMap := make(map[string]string)
		resourceMap["netconf_port"] = vvNetconfPort
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalCredentialNetconfRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Discovery.CreateNetconfCredentials(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetconfCredentials", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetconfCredentials", err))
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
			diags = append(diags, diagError(
				"Failure when executing CreateNetconfCredentials", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["netconf_port"] = vvNetconfPort
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialNetconfRead(ctx, d, m)
}

func resourceGlobalCredentialNetconfRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "NETCONF"
	vNetconfPort := resourceMap["netconf_port"]
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalCredentials")
		queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

		queryParams1.CredentialSubType = vCredentialSubType

		response1, err := searchDiscoveryGetGlobalCredentialsNetConf(m, queryParams1, vNetconfPort, vID)
		if err != nil || response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenDiscoveryGetGlobalCredentialsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalCredentialNetconfUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetconfPort := resourceMap["netconf_port"]
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "NETCONF"
	item, err := searchDiscoveryGetGlobalCredentialsNetConf(m, queryParams1, vNetconfPort, vID)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalCredentials", err,
			"Failure at GetGlobalCredentials, unexpected response", ""))
		return diags
	}

	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestGlobalCredentialNetconfUpdateNetconfCredentials(ctx, "parameters.0", d)
		request1.ID = item.ID
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateNetconfCredentials(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetconfCredentials", err, restyResp1.String(),
					"Failure at UpdateNetconfCredentials, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetconfCredentials", err,
				"Failure at UpdateNetconfCredentials, unexpected response", ""))
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
				diags = append(diags, diagError(
					"Failure when executing UpdateNetconfCredentials", err))
				return diags
			}
		}

	}

	return resourceGlobalCredentialNetconfRead(ctx, d, m)
}

func resourceGlobalCredentialNetconfDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete GlobalCredentialSNMPv2ReadCommunity on Dna Center
	//       Returning empty diags to delete it on Terraform
	// DeleteGlobalCredentialsByID
	client := m.(*dnacentersdkgo.Client)

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vNetconfPort := resourceMap["netconf_port"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "NETCONF"
	item, err := searchDiscoveryGetGlobalCredentialsNetConf(m, queryParams1, vNetconfPort, vID)
	if item == nil && err != nil {
		return resourceGlobalCredentialNetconfRead(ctx, d, m)
	}

	if vID != "" {
		response1, restyResp1, err := client.Discovery.DeleteGlobalCredentialsByID(vID)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteGlobalCredentialsByID", err,
				"Failure at DeleteGlobalCredentialsByID, unexpected response", ""))
			return diags
		}
	}
	return diags
}
func expandRequestGlobalCredentialNetconfCreateNetconfCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateNetconfCredentials {
	request := dnacentersdkgo.RequestDiscoveryCreateNetconfCredentials{}
	if v := expandRequestGlobalCredentialNetconfCreateNetconfCredentialsItemArray(ctx, key, d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialNetconfCreateNetconfCredentialsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateNetconfCredentials {
	request := []dnacentersdkgo.RequestItemDiscoveryCreateNetconfCredentials{}
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
		i := expandRequestGlobalCredentialNetconfCreateNetconfCredentialsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialNetconfCreateNetconfCredentialsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateNetconfCredentials {
	request := dnacentersdkgo.RequestItemDiscoveryCreateNetconfCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialNetconfUpdateNetconfCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateNetconfCredentials {
	request := dnacentersdkgo.RequestDiscoveryUpdateNetconfCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscoveryGetGlobalCredentialsNetConf(m interface{}, queryParams dnacentersdkgo.GetGlobalCredentialsQueryParams, vNetconfPort string, vID string) (*dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
	var ite *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentials
	queryParams.CredentialSubType = "NETCONF"
	ite, _, err = client.Discovery.GetGlobalCredentials(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range *itemsCopy.Response {
		// Call get by _ method and set value to foundItem and return
		if item.ID == vID || item.NetconfPort == vNetconfPort {
			var getItem *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
