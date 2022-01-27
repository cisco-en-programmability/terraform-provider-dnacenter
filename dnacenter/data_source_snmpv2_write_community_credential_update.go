package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSNMPv2WriteCommunityCredentialUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Discovery.

- Updates global SNMP write community
`,

		ReadContext: dataSourceSNMPv2WriteCommunityCredentialUpdateRead,
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
				Description: `SNMP write community. NO!$DATA!$ for no value change
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSNMPv2WriteCommunityCredentialUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateSNMPWriteCommunity")
		request1 := expandRequestSNMPv2WriteCommunityCredentialUpdateUpdateSNMPWriteCommunity(ctx, "", d)

		response1, restyResp1, err := client.Discovery.UpdateSNMPWriteCommunity(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSNMPWriteCommunity", err,
				"Failure at UpdateSNMPWriteCommunity, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryUpdateSNMPWriteCommunityItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateSNMPWriteCommunity response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSNMPv2WriteCommunityCredentialUpdateUpdateSNMPWriteCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateSNMPWriteCommunity {
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

func flattenDiscoveryUpdateSNMPWriteCommunityItem(item *dnacentersdkgo.ResponseDiscoveryUpdateSNMPWriteCommunityResponse) []map[string]interface{} {
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
