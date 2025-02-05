package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGlobalCredentialV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- API to get device credentials' details. It fetches all global credentials of all types at once, without the need to
pass any input parameters.
`,

		ReadContext: dataSourceGlobalCredentialV2Read,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the CLI credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the CLI credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the CLI credential
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
										Description: `Id of the CLI Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of CLI Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of CLI Credential
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

									"username": &schema.Schema{
										Description: `CLI Username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"https_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the HTTP(S) Read credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the HTTP(S) Read credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description for HTTP(S) Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the HTTP(S) Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of HTTP(S) Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of HTTP(S) Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Description: `HTTP(S) Read Password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Description: `HTTP(S) Port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										Description: `Flag for HTTP(S) Read
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Description: `HTTP(S) Read Username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"https_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the HTTP(S) Write credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the HTTP(S) Write credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description for HTTP(S) Write Credetntials
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the HTTP(S) Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of HTTP(S) Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of HTTP(S) Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Description: `HTTP(S) Write Password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"port": &schema.Schema{
										Description: `HTTP(S) Port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"secure": &schema.Schema{
										Description: `Flag for HTTP(S) Write
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Description: `HTTP(S) Write Username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v2c_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the SNMP Read credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the SNMP Read credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description for Snmp RO community
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the SNMP Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of SNMP Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of SNMP Read Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"read_community": &schema.Schema{
										Description: `Snmp RO community
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v2c_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments to identify the SNMP Write credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the SNMP Write credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description for Snmp RW community
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of SNMP Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of SNMP Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of SNMP Write Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"write_community": &schema.Schema{
										Description: `Snmp RW community
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_password": &schema.Schema{
										Description: `Auth Password for SNMP V3
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_type": &schema.Schema{
										Description: `SNMP auth protocol. SHA' or 'MD5'
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"comments": &schema.Schema{
										Description: `Comments to identify the SNMP V3 credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential type to identify the application that uses the SNMP V3 credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description for Snmp V3 Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the SNMP V3 Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id of SNMP V3 Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Tenant Id of SNMP V3 Credential
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"privacy_password": &schema.Schema{
										Description: `Privacy Password for SNMP privacy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"privacy_type": &schema.Schema{
										Description: `SNMP privacy protocol. 'AES128','AES192','AES256'
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"snmp_mode": &schema.Schema{
										Description: `Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"username": &schema.Schema{
										Description: `SNMP V3 Username
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalCredentialV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllGlobalCredentialsV2")

		response1, restyResp1, err := client.Discovery.GetAllGlobalCredentialsV2()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllGlobalCredentialsV2", err,
				"Failure at GetAllGlobalCredentialsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryGetAllGlobalCredentialsV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllGlobalCredentialsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetAllGlobalCredentialsV2Item(item *dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli_credential"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemCliCredential(item.CliCredential)
	respItem["snmp_v2c_read"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV2CRead(item.SNMPV2CRead)
	respItem["snmp_v2c_write"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV2CWrite(item.SNMPV2CWrite)
	respItem["https_read"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemHTTPSRead(item.HTTPSRead)
	respItem["https_write"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemHTTPSWrite(item.HTTPSWrite)
	respItem["snmp_v3"] = flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV3(item.SNMPV3)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemCliCredential(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseCliCredential) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["password"] = item.Password
		respItem["username"] = item.Username
		respItem["enable_password"] = item.EnablePassword
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV2CRead(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CRead) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["read_community"] = item.ReadCommunity
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV2CWrite(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CWrite) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["write_community"] = item.WriteCommunity
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemHTTPSRead(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSRead) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["password"] = item.Password
		respItem["port"] = item.Port
		respItem["username"] = item.Username
		respItem["secure"] = boolPtrToString(item.Secure)
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemHTTPSWrite(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSWrite) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["password"] = item.Password
		respItem["port"] = item.Port
		respItem["username"] = item.Username
		respItem["secure"] = boolPtrToString(item.Secure)
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetAllGlobalCredentialsV2ItemSNMPV3(items *[]dnacentersdkgo.ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV3) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["username"] = item.Username
		respItem["auth_password"] = item.AuthPassword
		respItem["auth_type"] = item.AuthType
		respItem["privacy_password"] = item.PrivacyPassword
		respItem["privacy_type"] = item.PrivacyType
		respItem["snmp_mode"] = item.SNMPMode
		respItem["description"] = item.Description
		respItem["comments"] = item.Comments
		respItem["credential_type"] = item.CredentialType
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
