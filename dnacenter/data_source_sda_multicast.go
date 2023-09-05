package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

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

						"description": &schema.Schema{
							Description: `multicast configuration info retrieved successfully from sda fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multicast_method": &schema.Schema{
							Description: `Multicast Method
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multicast_type": &schema.Schema{
							Description: `Multicast Type
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
										Description: `ExternalRpIpAddress
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"internal_rp_ip_address": &schema.Schema{
										Description: `InternalRpIpAddress
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to Fabric Site
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssm_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ssm_group_range": &schema.Schema{
													Description: `SSM group range
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ssm_wildcard_mask": &schema.Schema{
													Description: `SSM Wildcard Mask 
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to Fabric Site
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"status": &schema.Schema{
							Description: `Status
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
		log.Printf("[DEBUG] Selected method: GetMulticastDetailsFromSdaFabric")
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
	respItem["multicast_method"] = item.MulticastMethod
	respItem["multicast_type"] = item.MulticastType
	respItem["multicast_vn_info"] = flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfo(item.MulticastVnInfo)
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfo(items *[]dnacentersdkgo.ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["internal_rp_ip_address"] = item.InternalRpIPAddress
		respItem["external_rp_ip_address"] = item.ExternalRpIPAddress
		respItem["ssm_info"] = flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfoSsmInfo(item.SsmInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetMulticastDetailsFromSdaFabricItemMulticastVnInfoSsmInfo(items *[]dnacentersdkgo.ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssm_group_range"] = item.SsmGroupRange
		respItem["ssm_wildcard_mask"] = item.SsmWildcardMask
		respItems = append(respItems, respItem)
	}
	return respItems
}
