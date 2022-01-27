package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get  DHCP and DNS center server details.
`,

		ReadContext: dataSourceNetworkRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site id to get the network settings associated with the site.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"group_uuid": &schema.Schema{
							Description: `Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"key": &schema.Schema{
							Description: `Key`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_dnac_ip": &schema.Schema{
										Description: `Configure Dnac I P`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_addresses": &schema.Schema{
										Description: `Ip Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetwork")
		queryParams1 := dnacentersdkgo.GetNetworkQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetNetwork(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetwork", err,
				"Failure at GetNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetNetworkItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetwork response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetNetworkItems(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_type"] = item.InstanceType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["namespace"] = item.Namespace
		respItem["type"] = item.Type
		respItem["key"] = item.Key
		respItem["version"] = item.Version
		respItem["value"] = flattenNetworkSettingsGetNetworkItemsValue(item.Value)
		respItem["group_uuid"] = item.GroupUUID
		respItem["inherited_group_uuid"] = item.InheritedGroupUUID
		respItem["inherited_group_name"] = item.InheritedGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetNetworkItemsValue(items *[]dnacentersdkgo.ResponseNetworkSettingsGetNetworkResponseValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_addresses"] = item.IPAddresses
		respItem["configure_dnac_ip"] = boolPtrToString(item.ConfigureDnacIP)
		respItems = append(respItems, respItem)
	}
	return respItems
}
