package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaAnycastGatewaysCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of anycast gateways that match the provided query parameters.
`,

		ReadContext: dataSourceSdaAnycastGatewaysCountRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the anycast gateway is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_pool_name": &schema.Schema{
				Description: `ipPoolName query parameter. Name of the IP pool associated with the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_network_name": &schema.Schema{
				Description: `virtualNetworkName query parameter. Name of the virtual network associated with the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Description: `vlanId query parameter. VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Description: `vlanName query parameter. VLAN name of the anycast gateways.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The number of anycast gateways.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaAnycastGatewaysCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vVirtualNetworkName, okVirtualNetworkName := d.GetOk("virtual_network_name")
	vIPPoolName, okIPPoolName := d.GetOk("ip_pool_name")
	vVLANName, okVLANName := d.GetOk("vlan_name")
	vVLANID, okVLANID := d.GetOk("vlan_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnycastGatewayCount")
		queryParams1 := dnacentersdkgo.GetAnycastGatewayCountQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okVirtualNetworkName {
			queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)
		}
		if okIPPoolName {
			queryParams1.IPPoolName = vIPPoolName.(string)
		}
		if okVLANName {
			queryParams1.VLANName = vVLANName.(string)
		}
		if okVLANID {
			queryParams1.VLANID = vVLANID.(float64)
		}

		response1, restyResp1, err := client.Sda.GetAnycastGatewayCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAnycastGatewayCount", err,
				"Failure at GetAnycastGatewayCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetAnycastGatewayCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnycastGatewayCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetAnycastGatewayCountItem(item *dnacentersdkgo.ResponseSdaGetAnycastGatewayCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
