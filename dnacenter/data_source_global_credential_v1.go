package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGlobalCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns global credential for the given credential sub type

- Returns the credential sub type for the given Id
`,

		ReadContext: dataSourceGlobalCredentialRead,
		Schema: map[string]*schema.Schema{
			"credential_sub_type": &schema.Schema{
				Description: `credentialSubType query parameter. Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Global Credential ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order of sorting. 'asc' or 'des'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Field to sort the results by. Sorts by 'instanceId' if no value is provided
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Credential type as 'CLICredential', 'HTTPReadCredential', 'HTTPWriteCredential', 'NetconfCredential', 'SNMPv2ReadCommunity', 'SNMPv2WriteCommunity', 'SNMPv3Credential'
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_password": &schema.Schema{
							Description: `SNMPV3 Auth Password
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_type": &schema.Schema{
							Description: `SNMPV3 Auth Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"comments": &schema.Schema{
							Description: `Comments to identify the Global Credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"credential_type": &schema.Schema{
							Description: `Credential type to identify the application that uses the Global credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description for Global Credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_password": &schema.Schema{
							Description: `CLI Enable Password
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the Global Credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id of the Global Credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the Global Credential
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_port": &schema.Schema{
							Description: `Netconf Port
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"password": &schema.Schema{
							Description: `CLI Password
`,
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},

						"port": &schema.Schema{
							Description: `HTTP(S) port
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"privacy_password": &schema.Schema{
							Description: `SNMPV3 Privacy Password
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"privacy_type": &schema.Schema{
							Description: `SNMPV3 Privacy Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"read_community": &schema.Schema{
							Description: `SNMP Read Community
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"secure": &schema.Schema{
							Description: `Flag for HTTP(S)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_mode": &schema.Schema{
							Description: `SNMP Mode
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Description: `CLI Username
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"write_community": &schema.Schema{
							Description: `SNMP Write Community
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vCredentialSubType, okCredentialSubType := d.GetOk("credential_sub_type")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vID, okID := d.GetOk("id")

	method1 := []bool{okCredentialSubType, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGlobalCredentials")
		queryParams1 := dnacentersdkgo.GetGlobalCredentialsQueryParams{}

		if okCredentialSubType {
			queryParams1.CredentialSubType = vCredentialSubType.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Discovery.GetGlobalCredentials(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetGlobalCredentials", err,
				"Failure at GetGlobalCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetGlobalCredentialsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetCredentialSubTypeByCredentialID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Discovery.GetCredentialSubTypeByCredentialID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCredentialSubTypeByCredentialID", err,
				"Failure at GetCredentialSubTypeByCredentialID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDiscoveryGetCredentialSubTypeByCredentialIDItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCredentialSubTypeByCredentialID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetGlobalCredentialsItems(items *[]dnacentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["username"] = item.Username
		respItem["enable_password"] = item.EnablePassword
		respItem["password"] = item.Password
		respItem["netconf_port"] = item.NetconfPort
		respItem["read_community"] = item.ReadCommunity
		respItem["write_community"] = item.WriteCommunity
		respItem["auth_password"] = item.AuthPassword
		respItem["auth_type"] = item.AuthType
		respItem["privacy_password"] = item.PrivacyPassword
		respItem["privacy_type"] = item.PrivacyType
		respItem["snmp_mode"] = item.SNMPMode
		respItem["secure"] = item.Secure
		respItem["port"] = item.Port
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetCredentialSubTypeByCredentialIDItem(item *dnacentersdkgo.ResponseDiscoveryGetCredentialSubTypeByCredentialID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
