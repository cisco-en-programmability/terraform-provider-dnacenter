package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricEdgeDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get edge device from SDA Fabric
`,

		ReadContext: dataSourceSdaFabricEdgeDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter.`,
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

func dataSourceSdaFabricEdgeDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEdgeDeviceFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetEdgeDeviceFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetEdgeDeviceFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEdgeDeviceFromSdaFabric", err,
				"Failure at GetEdgeDeviceFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetEdgeDeviceFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEdgeDeviceFromSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetEdgeDeviceFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetEdgeDeviceFromSdaFabric) []map[string]interface{} {
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
