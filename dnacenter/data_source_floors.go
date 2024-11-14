package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloors() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Gets a floor in the network hierarchy.
`,

		ReadContext: dataSourceFloorsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Floor Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"units_of_measure": &schema.Schema{
				Description: `_unitsOfMeasure query parameter. Floor units of measure
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"floor_number": &schema.Schema{
							Description: `Floor number
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"height": &schema.Schema{
							Description: `Floor height. Example : 10.1
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Floor Id. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"length": &schema.Schema{
							Description: `Floor length. Example : 110.3
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Floor name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name_hierarchy": &schema.Schema{
							Description: `Floor hierarchical name. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent Id.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rf_model": &schema.Schema{
							Description: `RF Model
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Example : floor
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"units_of_measure": &schema.Schema{
							Description: `Units Of Measure`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"width": &schema.Schema{
							Description: `Floor width. Example : 100.5
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFloorsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vUnitsOfMeasure, okUnitsOfMeasure := d.GetOk("units_of_measure")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAFloorV2")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetsAFloorV2QueryParams{}

		if okUnitsOfMeasure {
			queryParams1.UnitsOfMeasure = vUnitsOfMeasure.(string)
		}

		response1, restyResp1, err := client.SiteDesign.GetsAFloorV2(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAFloorV2", err,
				"Failure at GetsAFloorV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsAFloorV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAFloorV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetsAFloorV2Item(item *dnacentersdkgo.ResponseSiteDesignGetsAFloorV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_id"] = item.ParentID
	respItem["name"] = item.Name
	respItem["floor_number"] = item.FloorNumber
	respItem["rf_model"] = item.RfModel
	respItem["width"] = item.Width
	respItem["length"] = item.Length
	respItem["height"] = item.Height
	respItem["units_of_measure"] = item.UnitsOfMeasure
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["name_hierarchy"] = item.NameHierarchy
	return []map[string]interface{}{
		respItem,
	}
}
