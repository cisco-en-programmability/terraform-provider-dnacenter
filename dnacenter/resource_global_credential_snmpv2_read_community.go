package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalCredentialSNMPv2ReadCommunity() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Updates global SNMP read community

- Adds global SNMP read community

- Deletes global credential for the given ID
`,

		CreateContext: resourceGlobalCredentialSNMPv2ReadCommunityCreate,
		ReadContext:   resourceGlobalCredentialSNMPv2ReadCommunityRead,
		UpdateContext: resourceGlobalCredentialSNMPv2ReadCommunityUpdate,
		DeleteContext: resourceGlobalCredentialSNMPv2ReadCommunityDelete,
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

						"read_community": &schema.Schema{
							Type:     schema.TypeString,
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

						"write_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateSNMPReadCommunity`,
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"comments": &schema.Schema{
							Description: `Comments to identify the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"credential_type": &schema.Schema{
							Description: `Credential type to identify the application that uses the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Name/Description of the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_community": &schema.Schema{
							Description: `SNMP read community
`,
							Type:     schema.TypeString,
							Optional: true,
							// ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceGlobalCredentialSNMPv2ReadCommunityCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunity(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vReadCommunity := resourceItem["read_community"]
	vvReadCommunity := interfaceToString(vReadCommunity)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}
	// resourceID := d.Id()
	// if resourceID != "" {
	// 	log.Printf("[DEBUG] ResourceID => %s", resourceID)
	// 	resourceMap := separateResourceID(resourceID)
	// 	vvID = resourceMap["id"]
	// }

	queryParams1.CredentialSubType = "SNMPV2_READ_COMMUNITY"
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Read(m, queryParams1, vvID)
	if item != nil || err != nil {
		resourceMap := make(map[string]string)
		resourceMap["read_community"] = vvReadCommunity
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalCredentialSNMPv2ReadCommunityRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Discovery.CreateSNMPReadCommunity(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSNMPReadCommunity", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSNMPReadCommunity", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateSNMPReadCommunity", err1))
			return diags
		}
		vvID = response2.Response.Progress
	}
	resourceMap := make(map[string]string)
	resourceMap["read_community"] = vvReadCommunity
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialSNMPv2ReadCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2ReadCommunityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "SNMPV2_READ_COMMUNITY"
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalCredentials")
		queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

		queryParams1.CredentialSubType = vCredentialSubType

		response1, err := searchDiscoveryGetGlobalCredentialsSmpv2Read(m, queryParams1, vID)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}
		if response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		items := []dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse{
			*response1,
		}
		vItem1 := flattenDiscoveryGetGlobalCredentialsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalCredentialSNMPv2ReadCommunityUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "SNMPV2_READ_COMMUNITY"
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}
	queryParams1.CredentialSubType = vCredentialSubType
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Read(m, queryParams1, vID)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalCredentials", err,
			"Failure at GetGlobalCredentials, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] ReadCommunity=> %s", item.ReadCommunity)

	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestGlobalCredentialSNMPv2ReadCommunityUpdateSNMPReadCommunity(ctx, "parameters.0", d)
		request1.InstanceUUID = item.InstanceUUID
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateSNMPReadCommunity(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSNMPReadCommunity", err, restyResp1.String(),
					"Failure at UpdateSNMPReadCommunity, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSNMPReadCommunity", err,
				"Failure at UpdateSNMPReadCommunity, unexpected response", ""))
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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateSNMPReadCommunity", err1))
				return diags
			}
		}
	}

	return resourceGlobalCredentialSNMPv2ReadCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2ReadCommunityDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete GlobalCredentialSNMPv2ReadCommunity on Catalyst Center
	//       Returning empty diags to delete it on Terraform
	// DeleteGlobalCredentialsByID
	client := m.(*dnacentersdkgo.Client)

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "SNMPV2_READ_COMMUNITY"
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Write(m, queryParams1, vID)
	if item == nil && err != nil {
		return resourceGlobalCredentialSNMPv2ReadCommunityRead(ctx, d, m)
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
func expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateSNMPReadCommunity {
	request := dnacentersdkgo.RequestDiscoveryCreateSNMPReadCommunity{}
	if v := expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunityItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunityItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity {
	request := []dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity{}
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
		i := expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunityItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2ReadCommunityCreateSNMPReadCommunityItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity {
	request := dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2ReadCommunityUpdateSNMPReadCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateSNMPReadCommunity {
	request := dnacentersdkgo.RequestDiscoveryUpdateSNMPReadCommunity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscoveryGetGlobalCredentialsSmpv2Read(m interface{}, queryParams dnacentersdkgo.GetGlobalCredentialsQueryParams, vID string) (*dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
	var ite *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentials
	queryParams.CredentialSubType = "SNMPV2_READ_COMMUNITY"
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
		if item.ID == vID {
			var getItem *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
