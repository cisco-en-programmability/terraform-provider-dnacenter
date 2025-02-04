package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalCredentialSNMPv2WriteCommunity() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Adds global SNMP write community

- Updates global SNMP write community

- Deletes global credential for the given ID
`,

		CreateContext: resourceGlobalCredentialSNMPv2WriteCommunityCreate,
		ReadContext:   resourceGlobalCredentialSNMPv2WriteCommunityRead,
		UpdateContext: resourceGlobalCredentialSNMPv2WriteCommunityUpdate,
		DeleteContext: resourceGlobalCredentialSNMPv2WriteCommunityDelete,
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
				Description: `Array of RequestDiscoveryCreateSNMPWriteCommunity`,
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
						"write_community": &schema.Schema{
							Description: `SNMP write community
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

func resourceGlobalCredentialSNMPv2WriteCommunityCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunity(ctx, "parameters", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vWriteCommunity := resourceItem["write_community"]
	vvWriteCommunity := interfaceToString(vWriteCommunity)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "SNMPV2_WRITE_COMMUNITY"
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Write(m, queryParams1, vvID)
	if item != nil || err != nil {
		resourceMap := make(map[string]string)
		resourceMap["write_community"] = vvWriteCommunity
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Discovery.CreateSNMPWriteCommunity(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSNMPWriteCommunity", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSNMPWriteCommunity", err))
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
				"Failure when executing CreateSNMPWriteCommunity", err1))
			return diags
		}
		vvID = response2.Response.Progress
	}

	resourceMap := make(map[string]string)
	resourceMap["write_community"] = vvWriteCommunity
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "SNMPV2_WRITE_COMMUNITY"
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalCredentials")
		queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

		queryParams1.CredentialSubType = vCredentialSubType

		response1, err := searchDiscoveryGetGlobalCredentialsSmpv2Write(m, queryParams1, vID)
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

func resourceGlobalCredentialSNMPv2WriteCommunityUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "SNMPV2_WRITE_COMMUNITY"
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}
	queryParams1.CredentialSubType = vCredentialSubType
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Write(m, queryParams1, vID)
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
		request1 := expandRequestGlobalCredentialSNMPv2WriteCommunityUpdateSNMPWriteCommunity(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateSNMPWriteCommunity(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSNMPWriteCommunity", err, restyResp1.String(),
					"Failure at UpdateSNMPWriteCommunity, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSNMPWriteCommunity", err,
				"Failure at UpdateSNMPWriteCommunity, unexpected response", ""))
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
					"Failure when executing UpdateSNMPWriteCommunity", err1))
				return diags
			}
		}
	}

	return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2WriteCommunityDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// NOTE: Unable to delete GlobalCredentialSNMPv2WriteCommunity on Catalyst Center
	//       Returning empty diags to delete it on Terraform
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "SNMPV2_WRITE_COMMUNITY"
	item, err := searchDiscoveryGetGlobalCredentialsSmpv2Write(m, queryParams1, vID)
	if item == nil && err != nil {
		return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
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
func expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateSNMPWriteCommunity {
	request := dnacentersdkgo.RequestDiscoveryCreateSNMPWriteCommunity{}
	if v := expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunityItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunityItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity {
	request := []dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity{}
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
		i := expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunityItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunityItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity {
	request := dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".write_community")))) {
		request.WriteCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv2WriteCommunityUpdateSNMPWriteCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateSNMPWriteCommunity {
	request := dnacentersdkgo.RequestDiscoveryUpdateSNMPWriteCommunity{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".write_community")))) {
		request.WriteCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscoveryGetGlobalCredentialsSmpv2Write(m interface{}, queryParams dnacentersdkgo.GetGlobalCredentialsQueryParams, vID string) (*dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
	var ite *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentials
	queryParams.CredentialSubType = "SNMPV2_WRITE_COMMUNITY"
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
