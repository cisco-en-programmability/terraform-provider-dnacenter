package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloorsFloorIDPlannedAccessPointPositionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieve all Planned Access Points Positions count designated for a specific floor
`,

		ReadContext: dataSourceFloorsFloorIDPlannedAccessPointPositionsCountRead,
		Schema: map[string]*schema.Schema{
			"floor_id": &schema.Schema{
				Description: `floorId path parameter. Floor Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Planned Access Point mac address.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Planned Access Point name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Planned Access Point type.
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
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFloorsFloorIDPlannedAccessPointPositionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFloorID := d.Get("floor_id")
	vName, okName := d.GetOk("name")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vType, okType := d.GetOk("type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPlannedAccessPointsPositionsCountV2")
		vvFloorID := vFloorID.(string)
		queryParams1 := dnacentersdkgo.GetPlannedAccessPointsPositionsCountV2QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}

		response1, restyResp1, err := client.SiteDesign.GetPlannedAccessPointsPositionsCountV2(vvFloorID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPlannedAccessPointsPositionsCountV2", err,
				"Failure at GetPlannedAccessPointsPositionsCountV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetPlannedAccessPointsPositionsCountV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPlannedAccessPointsPositionsCountV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetPlannedAccessPointsPositionsCountV2Item(item *dnacentersdkgo.ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
