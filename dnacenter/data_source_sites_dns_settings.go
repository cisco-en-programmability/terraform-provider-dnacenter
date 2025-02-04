package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesDNSSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieve DNS settings for a site; *null* values indicate that the setting will be inherited from the parent site;
empty objects (*{}*) indicate that the setting is unset at a site.
`,

		ReadContext: dataSourceSitesDNSSettingsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id
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

						"dns": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_servers": &schema.Schema{
										Description: `DNS servers for hostname resolution.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"domain_name": &schema.Schema{
										Description: `Network's domain name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name.
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

func dataSourceSitesDNSSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveDNSSettingsForASite")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.RetrieveDNSSettingsForASiteQueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrieveDNSSettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveDNSSettingsForASite", err,
				"Failure at RetrieveDNSSettingsForASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveDNSSettingsForASiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveDNSSettingsForASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveDNSSettingsForASiteItem(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dns"] = flattenNetworkSettingsRetrieveDNSSettingsForASiteItemDNS(item.DNS)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrieveDNSSettingsForASiteItemDNS(item *dnacentersdkgo.ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponseDNS) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["domain_name"] = item.DomainName
	respItem["dns_servers"] = item.DNSServers
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}
