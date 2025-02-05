package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloorsFloorIDAccessPointPositionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieve Access Points positions count assigned for a specific floor
`,

		ReadContext: dataSourceFloorsFloorIDAccessPointPositionsCountRead,
		Schema: map[string]*schema.Schema{
			"floor_id": &schema.Schema{
				Description: `floorId path parameter. Floor Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Access Point mac address.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": &schema.Schema{
				Description: `model query parameter. Access Point model.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Access Point name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Access Point type.
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

func dataSourceFloorsFloorIDAccessPointPositionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFloorID := d.Get("floor_id")
	vName, okName := d.GetOk("name")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vType, okType := d.GetOk("type")
	vModel, okModel := d.GetOk("model")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointsPositionsCountV2")
		vvFloorID := vFloorID.(string)
		queryParams1 := dnacentersdkgo.GetAccessPointsPositionsCountV2QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okModel {
			queryParams1.Model = vModel.(string)
		}

		response1, restyResp1, err := client.SiteDesign.GetAccessPointsPositionsCountV2(vvFloorID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointsPositionsCountV2", err,
				"Failure at GetAccessPointsPositionsCountV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetAccessPointsPositionsCountV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointsPositionsCountV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetAccessPointsPositionsCountV2Item(item *dnacentersdkgo.ResponseSiteDesignGetAccessPointsPositionsCountV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
