package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalCredentialSNMPv3() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Updates global SNMPv3 credential

- Adds global SNMPv3 credentials
`,

		CreateContext: resourceGlobalCredentialSNMPv3Create,
		ReadContext:   resourceGlobalCredentialSNMPv3Read,
		UpdateContext: resourceGlobalCredentialSNMPv3Update,
		DeleteContext: resourceGlobalCredentialSNMPv3Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateSNMPv3Credentials`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
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
						"privacy_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"privacy_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGlobalCredentialSNMPv3Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalCredentialSNMPv3CreateSNMPv3Credentials(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Discovery.CreateSNMPv3Credentials(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSNMPv3Credentials", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSNMPv3Credentials", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialSNMPv3Read(ctx, d, m)
}

func resourceGlobalCredentialSNMPv3Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := resourceMap["credential_sub_type"]
	vSortBy := resourceMap["sort_by"]
	vOrder := resourceMap["order"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalCredentials")
		queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

		queryParams1.CredentialSubType = vCredentialSubType

		if okSortBy {
			queryParams1.SortBy = vSortBy
		}
		if okOrder {
			queryParams1.Order = vOrder
		}

		response1, restyResp1, err := client.Discovery.GetGlobalCredentials(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGlobalCredentials", err,
				"Failure at GetGlobalCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenDiscoveryGetGlobalCredentialsItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalCredentialSNMPv3Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := resourceMap["credential_sub_type"]
	vSortBy := resourceMap["sort_by"]
	vOrder := resourceMap["order"]

	queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams
	queryParams1.CredentialSubType = vCredentialSubType
	queryParams1.SortBy = vSortBy
	queryParams1.Order = vOrder
	item, err := searchDiscoveryGetGlobalCredentials(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalCredentials", err,
			"Failure at GetGlobalCredentials, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestGlobalCredentialSNMPv3UpdateSNMPv3Credentials(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateSNMPv3Credentials(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSNMPv3Credentials", err, restyResp1.String(),
					"Failure at UpdateSNMPv3Credentials, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSNMPv3Credentials", err,
				"Failure at UpdateSNMPv3Credentials, unexpected response", ""))
			return diags
		}
	}

	return resourceGlobalCredentialSNMPv3Read(ctx, d, m)
}

func resourceGlobalCredentialSNMPv3Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete GlobalCredentialSNMPv3 on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestGlobalCredentialSNMPv3CreateSNMPv3Credentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateSNMPv3Credentials {
	request := dnacentersdkgo.RequestDiscoveryCreateSNMPv3Credentials{}
	if v := expandRequestGlobalCredentialSNMPv3CreateSNMPv3CredentialsItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv3CreateSNMPv3CredentialsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateSNMPv3Credentials {
	request := []dnacentersdkgo.RequestItemDiscoveryCreateSNMPv3Credentials{}
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
		i := expandRequestGlobalCredentialSNMPv3CreateSNMPv3CredentialsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv3CreateSNMPv3CredentialsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateSNMPv3Credentials {
	request := dnacentersdkgo.RequestItemDiscoveryCreateSNMPv3Credentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_password")))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_type")))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialSNMPv3UpdateSNMPv3Credentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateSNMPv3Credentials {
	request := dnacentersdkgo.RequestDiscoveryUpdateSNMPv3Credentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_password")))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_type")))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscoveryGetGlobalCredentials(m interface{}, queryParams dnacentersdkgo.GetGlobalCredentialsQueryParams) (*dnacentersdkgo.ResponseItemDiscoveryGetGlobalCredentials, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDiscoveryGetGlobalCredentials
	var ite *dnacentersdkgo.ResponseDiscoveryGetGlobalCredentials
	ite, _, err = client.Discovery.GetGlobalCredentials(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemDiscoveryGetGlobalCredentials
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
