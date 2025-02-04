package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
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
							Description: `Virtual Network info retrieved successfully from SDA Fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_name": &schema.Schema{
							Description: `Fabric Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_default_vn": &schema.Schema{
							Description: `Default VN
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_infra_vn": &schema.Schema{
							Description: `Infra VN
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_context_id": &schema.Schema{
							Description: `Virtual Network Context Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_id": &schema.Schema{
							Description: `Virtual Network Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name
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

func dataSourceSdaVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVirtualNetworkName := d.Get("virtual_network_name")
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetVnFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetVnFromSdaFabric", err,
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
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["virtual_network_name"] = item.VirtualNetworkName
	respItem["fabric_name"] = item.FabricName
	respItem["is_infra_vn"] = boolPtrToString(item.IsInfraVn)
	respItem["is_default_vn"] = boolPtrToString(item.IsDefaultVn)
	respItem["virtual_network_context_id"] = item.VirtualNetworkContextID
	respItem["virtual_network_id"] = item.VirtualNetworkID
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["execution_id"] = item.ExecutionID
	return []map[string]interface{}{
		respItem,
	}
}
