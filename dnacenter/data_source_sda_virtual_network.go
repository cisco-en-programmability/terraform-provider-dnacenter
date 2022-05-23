package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get virtual network (VN) from SDA Fabric
`,

		ReadContext: dataSourceSdaVirtualNetworkRead,
		Schema: map[string]*schema.Schema{
			"site_name_hierarchy": &schema.Schema{
				Description: `siteNameHierarchy query parameter.`,
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

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_management_ip_address": &schema.Schema{
							Description: `Device Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"roles": &schema.Schema{
							Description: `Roles`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVirtualNetworkName := d.Get("virtual_network_name")
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVnFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVnFromSdaFabric", err,
				"Failure at GetVnFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetVnFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVnFromSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetVnFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetVnFromSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["name"] = item.Name
	respItem["roles"] = item.Roles
	respItem["device_management_ip_address"] = item.DeviceManagementIPAddress
	respItem["site_hierarchy"] = item.SiteHierarchy
	return []map[string]interface{}{
		respItem,
	}
}
