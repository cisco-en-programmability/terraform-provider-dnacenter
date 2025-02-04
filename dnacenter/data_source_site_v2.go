package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- API to get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters  are not given as an
input.
`,

		ReadContext: dataSourceSiteV2Read,
		Schema: map[string]*schema.Schema{
			"group_name_hierarchy": &schema.Schema{
				Description: `groupNameHierarchy query parameter. Site name hierarchy (E.g. Global/USA/CA)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Site Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of sites to be listed. Default and max supported value is 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset/starting index for pagination
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Site type (Acceptable values: area, building, floor)
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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `Site address (e.g. 269 East Tasman Drive, San Jose, California 95134, United States)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"address_inherited_from": &schema.Schema{
													Description: `Site instance UUID from where address inherited (e.g. 576c7859-e485-4073-a46f-305f475de4c5)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"country": &schema.Schema{
													Description: `Site Country (e.g. United States)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"latitude": &schema.Schema{
													Description: `Site latitude (e.g. 37.413082)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"longitude": &schema.Schema{
													Description: `Site longitude (e.g. -121.933886)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `Site type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"name_space": &schema.Schema{
										Description: `Site name space. Default value is 'Location'
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"group_hierarchy": &schema.Schema{
							Description: `Site hierarchy by instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c/576c7859-e485-4073-a46f-305f475de4c5)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"group_name_hierarchy": &schema.Schema{
							Description: `Site hierarchy by name (e.g. Global/USA/CA/San Jose/Building4)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"group_type_list": &schema.Schema{
							Description: `There are different group types like 'RBAC', 'POLICY', 'SITE', 'TAG', 'PORT', 'DEVICE_TYPE'. This API is for site, so list contains 'SITE' only
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Site instance UUID (e.g. bb5122ce-4527-4af5-8718-44b746a3a3d8)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Tenant instance Id where site created (e.g. 63bf047b64ec9c1c45f9019c)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Site name (e.g. Building4)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent site Instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c)
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

func dataSourceSiteV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vGroupNameHierarchy, okGroupNameHierarchy := d.GetOk("group_name_hierarchy")
	vID, okID := d.GetOk("id")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteV2")
		queryParams1 := dnacentersdkgo.GetSiteV2QueryParams{}

		if okGroupNameHierarchy {
			queryParams1.GroupNameHierarchy = vGroupNameHierarchy.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
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

		response1, restyResp1, err := client.Sites.GetSiteV2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteV2", err,
				"Failure at GetSiteV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetSiteV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteV2Items(items *[]dnacentersdkgo.ResponseSitesGetSiteV2Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["parent_id"] = item.ParentID
		respItem["group_type_list"] = item.GroupTypeList
		respItem["group_hierarchy"] = item.GroupHierarchy
		respItem["additional_info"] = flattenSitesGetSiteV2ItemsAdditionalInfo(item.AdditionalInfo)
		respItem["group_name_hierarchy"] = item.GroupNameHierarchy
		respItem["name"] = item.Name
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteV2ItemsAdditionalInfo(items *[]dnacentersdkgo.ResponseSitesGetSiteV2ResponseAdditionalInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name_space"] = item.NameSpace
		respItem["attributes"] = flattenSitesGetSiteV2ItemsAdditionalInfoAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteV2ItemsAdditionalInfoAttributes(item *dnacentersdkgo.ResponseSitesGetSiteV2ResponseAdditionalInfoAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["address_inherited_from"] = item.AddressInheritedFrom
	respItem["type"] = item.Type
	respItem["country"] = item.Country
	respItem["address"] = item.Address
	respItem["latitude"] = item.Latitude
	respItem["longitude"] = item.Longitude

	return []map[string]interface{}{
		respItem,
	}

}
