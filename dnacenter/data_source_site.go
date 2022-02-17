package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSite() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Get site using siteNameHierarchy/siteId/type ,return all sites if these parameters are not given as input.
`,

		ReadContext: dataSourceSiteRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of sites to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. siteNameHierarchy (ex: global/groupName)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site id to which site details to retrieve.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. type (ex: area, building, floor)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_info": &schema.Schema{
							Description: `Additional Info`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"address_inherited_from": &schema.Schema{
													Description: `addressInheritedFrom`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"country": &schema.Schema{
													Description: `country`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"floor_index": &schema.Schema{
													Description: `floorIndex`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"height": &schema.Schema{
													Description: `height`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"latitude": &schema.Schema{
													Description: `latitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"length": &schema.Schema{
													Description: `length`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"longitude": &schema.Schema{
													Description: `longitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"offset_x": &schema.Schema{
													Description: `offsetX`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"offset_y": &schema.Schema{
													Description: `offsetY`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"rf_model": &schema.Schema{
													Description: `rfModel`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"type": &schema.Schema{
													Description: `type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"width": &schema.Schema{
													Description: `width`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"namespace": &schema.Schema{
										Description: `namespace`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vSiteID, okSiteID := d.GetOk("site_id")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSite")
		queryParams1 := dnacentersdkgo.GetSiteQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.Sites.GetSite(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSite", err,
				"Failure at GetSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetSiteItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSite response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteItems(items *[]dnacentersdkgo.ResponseSitesGetSiteResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["parent_id"] = item.ParentID
		respItem["name"] = item.Name
		respItem["additional_info"] = flattenSitesGetSiteItemsAdditionalInfo(&item.AdditionalInfo)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteItemsAdditionalInfo(items *[]dnacentersdkgo.ResponseSitesGetSiteResponseAdditionalInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["namespace"] = item.Namespace
		respItem["attributes"] = flattenSitesGetSiteItemsAdditionalInfoAttributes(&item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteItemsAdditionalInfoAttributes(item *dnacentersdkgo.ResponseSitesGetSiteResponseAdditionalInfoAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["country"] = item.Country
	respItem["address"] = item.Address
	respItem["latitude"] = item.Latitude
	respItem["address_inherited_from"] = item.AddressInheritedFrom
	respItem["type"] = item.Type
	respItem["longitude"] = item.Longitude
	respItem["offset_x"] = item.OffsetX
	respItem["offset_y"] = item.OffsetY
	respItem["length"] = item.Length
	respItem["width"] = item.Width
	respItem["height"] = item.Height
	respItem["rf_model"] = item.RfModel
	respItem["floor_index"] = item.FloorIndex

	return []map[string]interface{}{
		respItem,
	}

}
