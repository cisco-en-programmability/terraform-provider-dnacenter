package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTopologySite() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Topology.

- Returns site topology
`,

		ReadContext: dataSourceTopologySiteRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sites": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"group_name_hierarchy": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"latitude": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"location_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"location_country": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"location_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"longitude": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"parent_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTopologySiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSiteTopology")

		response1, restyResp1, err := client.Topology.GetSiteTopology()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSiteTopology", err,
				"Failure at GetSiteTopology, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTopologyGetSiteTopologyItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteTopology response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTopologyGetSiteTopologyItem(item *dnacentersdkgo.ResponseTopologyGetSiteTopologyResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sites"] = flattenTopologyGetSiteTopologyItemSites(item.Sites)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTopologyGetSiteTopologyItemSites(items *[]dnacentersdkgo.ResponseTopologyGetSiteTopologyResponseSites) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["group_name_hierarchy"] = item.GroupNameHierarchy
		respItem["id"] = item.ID
		respItem["latitude"] = item.Latitude
		respItem["location_address"] = item.LocationAddress
		respItem["location_country"] = item.LocationCountry
		respItem["location_type"] = item.LocationType
		respItem["longitude"] = item.Longitude
		respItem["name"] = item.Name
		respItem["parent_id"] = item.ParentID
		respItems = append(respItems, respItem)
	}
	return respItems
}
