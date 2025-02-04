package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesDeviceCredentials() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Gets device credential settings for a site; *null* values indicate that the setting will be inherited from the parent
site; empty objects (*{}*) indicate that the credential is unset, and that no credential of that type will be used for
the site.
`,

		ReadContext: dataSourceSitesDeviceCredentialsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id, retrievable from the *id* attribute in */dna/intent/api/v1/sites*
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"inherited": &schema.Schema{
				Description: `_inherited query parameter. Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when *false*, *null* values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv2c_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv2c_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv3_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceSitesDeviceCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCredentialSettingsForASite")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetDeviceCredentialSettingsForASiteQueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.GetDeviceCredentialSettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceCredentialSettingsForASite", err,
				"Failure at GetDeviceCredentialSettingsForASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCredentialSettingsForASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItem(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemCliCredentialsID(item.CliCredentialsID)
	respItem["snmpv2c_read_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv2CReadCredentialsID(item.SNMPv2CReadCredentialsID)
	respItem["snmpv2c_write_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv2CWriteCredentialsID(item.SNMPv2CWriteCredentialsID)
	respItem["snmpv3_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv3CredentialsID(item.SNMPv3CredentialsID)
	respItem["http_read_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemHTTPReadCredentialsID(item.HTTPReadCredentialsID)
	respItem["http_write_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemHTTPWriteCredentialsID(item.HTTPWriteCredentialsID)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemCliCredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseCliCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv2CReadCredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CReadCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv2CWriteCredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CWriteCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemSNMPv3CredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv3CredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemHTTPReadCredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPReadCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteItemHTTPWriteCredentialsID(item *dnacentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPWriteCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}
