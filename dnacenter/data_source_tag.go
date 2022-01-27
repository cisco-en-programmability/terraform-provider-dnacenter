package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTag() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Returns the tags for given filter criteria

- Returns tag specified by Id
`,

		ReadContext: dataSourceTagRead,
		Schema: map[string]*schema.Schema{
			"additional_info_attributes": &schema.Schema{
				Description: `additionalInfo.attributes query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"additional_info_name_space": &schema.Schema{
				Description: `additionalInfo.nameSpace query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field": &schema.Schema{
				Description: `field query parameter. Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Description: `level query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Tag name is mandatory when filter operation is used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Available values are asc and des
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. size in kilobytes(KB)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Only supported attribute is name. SortyBy is mandatory when order is used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_tag": &schema.Schema{
				Description: `systemTag query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vAdditionalInfonameSpace, okAdditionalInfonameSpace := d.GetOk("additional_info_name_space")
	vAdditionalInfoattributes, okAdditionalInfoattributes := d.GetOk("additional_info_attributes")
	vLevel, okLevel := d.GetOk("level")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSize, okSize := d.GetOk("size")
	vField, okField := d.GetOk("field")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSystemTag, okSystemTag := d.GetOk("system_tag")
	vID, okID := d.GetOk("id")

	method1 := []bool{okName, okAdditionalInfonameSpace, okAdditionalInfoattributes, okLevel, okOffset, okLimit, okSize, okField, okSortBy, okOrder, okSystemTag}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTag")
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okAdditionalInfonameSpace {
			queryParams1.AdditionalInfonameSpace = vAdditionalInfonameSpace.(string)
		}
		if okAdditionalInfoattributes {
			queryParams1.AdditionalInfoattributes = vAdditionalInfoattributes.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okSize {
			queryParams1.Size = vSize.(string)
		}
		if okField {
			queryParams1.Field = vField.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSystemTag {
			queryParams1.SystemTag = vSystemTag.(string)
		}

		response1, restyResp1, err := client.Tag.GetTag(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTag", err,
				"Failure at GetTag, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTagGetTagItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTag response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTagByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Tag.GetTagByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTagGetTagByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagGetTagItems(items *[]dnacentersdkgo.ResponseTagGetTagResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["system_tag"] = boolPtrToString(item.SystemTag)
		respItem["description"] = item.Description
		respItem["dynamic_rules"] = flattenTagGetTagItemsDynamicRules(item.DynamicRules)
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagItemsDynamicRules(items *[]dnacentersdkgo.ResponseTagGetTagResponseDynamicRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["member_type"] = item.MemberType
		respItem["rules"] = flattenTagGetTagItemsDynamicRulesRules(item.Rules)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagItemsDynamicRulesRules(item *dnacentersdkgo.ResponseTagGetTagResponseDynamicRulesRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["values"] = item.Values
	respItem["items"] = flattenTagGetTagItemsDynamicRulesRulesItems(item.Items)
	respItem["operation"] = item.Operation
	respItem["name"] = item.Name
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTagGetTagItemsDynamicRulesRulesItems(items *[]dnacentersdkgo.ResponseTagGetTagResponseDynamicRulesRulesItems) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenTagGetTagByIDItem(item *dnacentersdkgo.ResponseTagGetTagByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["system_tag"] = boolPtrToString(item.SystemTag)
	respItem["description"] = item.Description
	respItem["dynamic_rules"] = flattenTagGetTagByIDItemDynamicRules(item.DynamicRules)
	respItem["name"] = item.Name
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTagGetTagByIDItemDynamicRules(items *[]dnacentersdkgo.ResponseTagGetTagByIDResponseDynamicRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["member_type"] = item.MemberType
		respItem["rules"] = flattenTagGetTagByIDItemDynamicRulesRules(item.Rules)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagByIDItemDynamicRulesRules(item *dnacentersdkgo.ResponseTagGetTagByIDResponseDynamicRulesRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["values"] = item.Values
	respItem["items"] = flattenTagGetTagByIDItemDynamicRulesRulesItems(item.Items)
	respItem["operation"] = item.Operation
	respItem["name"] = item.Name
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTagGetTagByIDItemDynamicRulesRulesItems(items *[]dnacentersdkgo.ResponseTagGetTagByIDResponseDynamicRulesRulesItems) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
