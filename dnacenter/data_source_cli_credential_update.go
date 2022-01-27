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
func dataSourceCliCredentialUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Discovery.

- Updates global CLI credentials
`,

		ReadContext: dataSourceCliCredentialUpdateRead,
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
			"enable_password": &schema.Schema{
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
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceCliCredentialUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateCliCredentials")
		request1 := expandRequestCliCredentialUpdateUpdateCliCredentials(ctx, "", d)

		response1, restyResp1, err := client.Discovery.UpdateCliCredentials(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateCliCredentials", err,
				"Failure at UpdateCliCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryUpdateCliCredentialsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateCliCredentials response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestCliCredentialUpdateUpdateCliCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateCliCredentials {
	request := dnacentersdkgo.RequestDiscoveryUpdateCliCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDiscoveryUpdateCliCredentialsItem(item *dnacentersdkgo.ResponseDiscoveryUpdateCliCredentialsResponse) []map[string]interface{} {
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
