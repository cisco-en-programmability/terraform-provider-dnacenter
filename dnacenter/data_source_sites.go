package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSites() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Get sites.
`,

		ReadContext: dataSourceSitesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Site name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_hierarchy": &schema.Schema{
				Description: `nameHierarchy query parameter. Site name hierarchy.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Site type.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"units_of_measure": &schema.Schema{
				Description: `_unitsOfMeasure query parameter. Floor units of measure
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"address": &schema.Schema{
							Description: `Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"country": &schema.Schema{
							Description: `Country name for the building.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

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
							Description: `Site Id. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"latitude": &schema.Schema{
							Description: `Building Latitude. Example: 37.403712
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"length": &schema.Schema{
							Description: `Floor length. Example : 110.3
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"longitude": &schema.Schema{
							Description: `Building Longitude. Example: -121.971063
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Site name.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name_hierarchy": &schema.Schema{
							Description: `Site hierarchical name. Read only. Example: Global/USA/San Jose/Building1
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent Id. Read only
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rf_model": &schema.Schema{
							Description: `Floor RF Model
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"units_of_measure": &schema.Schema{
							Description: `Floor unit of measure
`,
							Type:     schema.TypeString,
							Computed: true,
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

func dataSourceSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vNameHierarchy, okNameHierarchy := d.GetOk("name_hierarchy")
	vType, okType := d.GetOk("type")
	vUnitsOfMeasure, okUnitsOfMeasure := d.GetOk("units_of_measure")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSites")
		queryParams1 := dnacentersdkgo.GetSitesQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okNameHierarchy {
			queryParams1.NameHierarchy = vNameHierarchy.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okUnitsOfMeasure {
			queryParams1.UnitsOfMeasure = vUnitsOfMeasure.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		response1, restyResp1, err := client.SiteDesign.GetSites(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSites", err,
				"Failure at GetSites, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignGetSitesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSites response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetSitesItems(items *[]dnacentersdkgo.ResponseSiteDesignGetSitesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name_hierarchy"] = item.NameHierarchy
		respItem["name"] = item.Name
		respItem["latitude"] = item.Latitude
		respItem["longitude"] = item.Longitude
		respItem["address"] = item.Address
		respItem["country"] = item.Country
		respItem["floor_number"] = item.FloorNumber
		respItem["rf_model"] = item.RfModel
		respItem["width"] = item.Width
		respItem["length"] = item.Length
		respItem["height"] = item.Height
		respItem["units_of_measure"] = item.UnitsOfMeasure
		respItem["type"] = item.Type
		respItem["id"] = item.ID
		respItem["parent_id"] = item.ParentID
		respItems = append(respItems, respItem)
	}
	return respItems
}
