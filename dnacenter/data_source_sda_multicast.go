package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaMulticast() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get multicast details from SDA fabric
`,

		ReadContext: dataSourceSdaMulticastRead,
		Schema: map[string]*schema.Schema{
			"site_name_hierarchy": &schema.Schema{
				Description: `siteNameHierarchy query parameter. fabric site name hierarchy
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"multicast_method": &schema.Schema{
							Description: `Multicast Methods
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multicast_vn_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_rp_ip_address": &schema.Schema{
										Description: `External Rp Ip Address, required for muticastType=asm_with_external_rp
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_group_range": &schema.Schema{
										Description: `Valid SSM group range ip address(e.g., 230.0.0.0)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_info": &schema.Schema{
										Description: `Source-specific multicast information, required if muticastType=ssm
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_wildcard_mask": &schema.Schema{
										Description: `Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"muticast_type": &schema.Schema{
							Description: `Muticast Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Full path of sda fabric siteNameHierarchy
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

func dataSourceSdaMulticastRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetMulticastDetailsFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMulticastDetailsFromSdaFabric", err,
				"Failure at GetMulticastDetailsFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetMulticastDetailsFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMulticastDetailsFromSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetMulticastDetailsFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetMulticastDetailsFromSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["multicast_method"] = item.MulticastMethod
	respItem["muticast_type"] = item.MuticastType
	respItem["multicast_vn_info"] = flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfo(item.MulticastVnInfo)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfo(item *dnacentersdkgo.ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["virtual_network_name"] = item.VirtualNetworkName
	respItem["ip_pool_name"] = item.IPPoolName
	respItem["external_rp_ip_address"] = item.ExternalRpIPAddress
	respItem["ssm_info"] = flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfoSsmInfo(item.SsmInfo)
	respItem["ssm_group_range"] = item.SsmGroupRange
	respItem["ssm_wildcard_mask"] = item.SsmWildcardMask

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfoSsmInfo(item *dnacentersdkgo.ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
