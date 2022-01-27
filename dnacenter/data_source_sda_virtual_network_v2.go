package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaVirtualNetworkV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get virtual network with scalable groups
`,

		ReadContext: dataSourceSdaVirtualNetworkV2Read,
		Schema: map[string]*schema.Schema{
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

						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_guest_virtual_network": &schema.Schema{
							Description: `Is Guest Virtual Network`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_names": &schema.Schema{
							Description: `Scalable Group Names`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"task_id": &schema.Schema{
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"task_status_url": &schema.Schema{
							Description: `Task Status Url`,
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

func dataSourceSdaVirtualNetworkV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVirtualNetworkName := d.Get("virtual_network_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVirtualNetworkWithScalableGroups")
		queryParams1 := dnacentersdkgo.GetVirtualNetworkWithScalableGroupsQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)

		response1, restyResp1, err := client.Sda.GetVirtualNetworkWithScalableGroups(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualNetworkWithScalableGroups", err,
				"Failure at GetVirtualNetworkWithScalableGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetVirtualNetworkWithScalableGroupsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkWithScalableGroups response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetVirtualNetworkWithScalableGroupsItem(item *dnacentersdkgo.ResponseSdaGetVirtualNetworkWithScalableGroups) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["virtual_network_name"] = item.VirtualNetworkName
	respItem["is_guest_virtual_network"] = boolPtrToString(item.IsGuestVirtualNetwork)
	respItem["scalable_group_names"] = item.ScalableGroupNames
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["task_id"] = item.TaskID
	respItem["task_status_url"] = item.TaskStatusURL
	respItem["execution_status_url"] = item.ExecutionStatusURL
	return []map[string]interface{}{
		respItem,
	}
}
