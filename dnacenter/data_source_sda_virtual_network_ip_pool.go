package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaVirtualNetworkIPPool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get IP Pool from SDA Virtual Network
`,

		ReadContext: dataSourceSdaVirtualNetworkIPPoolRead,
		Schema: map[string]*schema.Schema{
			"ip_pool_name": &schema.Schema{
				Description: `ipPoolName query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"virtual_network_name": &schema.Schema{
				Description: `virtualNetworkName query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authentication_policy_name": &schema.Schema{
							Description: `Authentication Policy Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_l2_flooding_enabled": &schema.Schema{
							Description: `Is L2 Flooding Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_this_critical_pool": &schema.Schema{
							Description: `Is This Critical Pool`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"traffic_type": &schema.Schema{
							Description: `Traffic Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaVirtualNetworkIPPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vIPPoolName := d.Get("ip_pool_name")
	vVirtualNetworkName := d.Get("virtual_network_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetIPPoolFromSdaVirtualNetwork")
		queryParams1 := dnacentersdkgo.GetIPPoolFromSdaVirtualNetworkQueryParams{}

		queryParams1.IPPoolName = vIPPoolName.(string)

		queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)

		response1, restyResp1, err := client.Sda.GetIPPoolFromSdaVirtualNetwork(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIPPoolFromSdaVirtualNetwork", err,
				"Failure at GetIPPoolFromSdaVirtualNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetIPPoolFromSdaVirtualNetworkItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPPoolFromSdaVirtualNetwork response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetIPPoolFromSdaVirtualNetworkItem(item *dnacentersdkgo.ResponseSdaGetIPPoolFromSdaVirtualNetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["virtual_network_name"] = item.VirtualNetworkName
	respItem["ip_pool_name"] = item.IPPoolName
	respItem["authentication_policy_name"] = item.AuthenticationPolicyName
	respItem["traffic_type"] = item.TrafficType
	respItem["scalable_group_name"] = item.ScalableGroupName
	respItem["is_l2_flooding_enabled"] = boolPtrToString(item.IsL2FloodingEnabled)
	respItem["is_this_critical_pool"] = boolPtrToString(item.IsThisCriticalPool)
	return []map[string]interface{}{
		respItem,
	}
}
