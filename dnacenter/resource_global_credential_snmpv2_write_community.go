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

func resourceGlobalCredentialSNMPv2WriteCommunity() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Adds global SNMP write community

- Updates global SNMP write community
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
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateSNMPWriteCommunity`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
	request1 := expandRequestGlobalCredentialSNMPv2WriteCommunityCreateSNMPWriteCommunity(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

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
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceGlobalCredentialSNMPv2WriteCommunityUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	}

	return resourceGlobalCredentialSNMPv2WriteCommunityRead(ctx, d, m)
}

func resourceGlobalCredentialSNMPv2WriteCommunityDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete GlobalCredentialSNMPv2WriteCommunity on Dna Center
	//       Returning empty diags to delete it on Terraform
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
