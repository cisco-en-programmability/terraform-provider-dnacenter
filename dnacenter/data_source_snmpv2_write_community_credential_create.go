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

// dataSourceAction
func dataSourceSNMPv2WriteCommunityCredentialCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Discovery.

- Adds global SNMP write community
`,

		ReadContext: dataSourceSNMPv2WriteCommunityCredentialCreateRead,
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"write_community": &schema.Schema{
				Description: `SNMP write community
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSNMPv2WriteCommunityCredentialCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateSNMPWriteCommunity")
		request1 := expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunity(ctx, "", d)

		response1, restyResp1, err := client.Discovery.CreateSNMPWriteCommunity(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSNMPWriteCommunity", err,
				"Failure at CreateSNMPWriteCommunity, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryCreateSNMPWriteCommunityItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSNMPWriteCommunity response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateSNMPWriteCommunity {
	request := dnacentersdkgo.RequestDiscoveryCreateSNMPWriteCommunity{}
	if v := expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunityItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunityItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity {
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
		i := expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunityItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPv2WriteCommunityCredentialCreateCreateSNMPWriteCommunityItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateSNMPWriteCommunity {
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

func flattenDiscoveryCreateSNMPWriteCommunityItem(item *dnacentersdkgo.ResponseDiscoveryCreateSNMPWriteCommunityResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
