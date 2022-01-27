package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsViewGroupView() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Gives complete information of the view that is required to configure a report. Use "Get views for a given view group"
API to get the viewIds  (required as a query param for this API) for available views.
`,

		ReadContext: dataSourceReportsViewGroupViewRead,
		Schema: map[string]*schema.Schema{
			"view_group_id": &schema.Schema{
				Description: `viewGroupId path parameter. viewGroupId of viewgroup
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"view_id": &schema.Schema{
				Description: `viewId path parameter. view id of view
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deliveries": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default": &schema.Schema{
										Description: `true, if the delivery type is considered default
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `delivery type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"description": &schema.Schema{
							Description: `view description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"field_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"field_group_display_name": &schema.Schema{
										Description: `Field group label/displayname for user
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"field_group_name": &schema.Schema{
										Description: `Field group name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fields": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `field label/displayname
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `field name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"table_id": &schema.Schema{
										Description: `Table Id of the corresponding table mapped to fieldgroup
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_info": &schema.Schema{
										Description: `Additional info for managing filter options
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"cache_filter": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"data_type": &schema.Schema{
										Description: `data type of filter value
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Description: `filter label/displayname
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"filter_source": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data_source": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_value_path": &schema.Schema{
													Description: `JSONPath of the label of filter option from the filter option as root
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"root_path": &schema.Schema{
													Description: `JSONPath of the filter options array in the API response
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"value_path": &schema.Schema{
													Description: `JSONPath of the value of filter option from the filter option as root
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"name": &schema.Schema{
										Description: `filter name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"required": &schema.Schema{
										Description: `true if the filter is required
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"time_options": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"info": &schema.Schema{
													Description: `Time range option description
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_value": &schema.Schema{
													Description: `Maximum number of hours allowed for the time range option. (Client Validation)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"min_value": &schema.Schema{
													Description: `Minimum number of hours allowed for the time range option. (Client Validation)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Time range option label
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Description: `Time range option value
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"type": &schema.Schema{
										Description: `filter type. Used to handle filter value selection by the client for report configuration.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"formats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default": &schema.Schema{
										Description: `true, if the format type is considered default
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"format": &schema.Schema{
										Description: `format type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `format name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"template": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"js_template_id": &schema.Schema{
													Description: `TemplateId of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"schedules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default": &schema.Schema{
										Description: `true, if the schedule type is default
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `schedule type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"view_id": &schema.Schema{
							Description: `Unique view Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_info": &schema.Schema{
							Description: `view filters info
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_name": &schema.Schema{
							Description: `view name
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

func dataSourceReportsViewGroupViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vViewGroupID := d.Get("view_group_id")
	vViewID := d.Get("view_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetViewDetailsForAGivenViewGroupView")
		vvViewGroupID := vViewGroupID.(string)
		vvViewID := vViewID.(string)

		response1, restyResp1, err := client.Reports.GetViewDetailsForAGivenViewGroupView(vvViewGroupID, vvViewID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetViewDetailsForAGivenViewGroupView", err,
				"Failure at GetViewDetailsForAGivenViewGroupView, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenReportsGetViewDetailsForAGivenViewGroupViewItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetViewDetailsForAGivenViewGroupView response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItem(item *dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupView) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deliveries"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemDeliveries(item.Deliveries)
	respItem["description"] = item.Description
	respItem["field_groups"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFieldGroups(item.FieldGroups)
	respItem["filters"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFilters(item.Filters)
	respItem["formats"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFormats(item.Formats)
	respItem["schedules"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemSchedules(item.Schedules)
	respItem["view_id"] = item.ViewID
	respItem["view_info"] = item.ViewInfo
	respItem["view_name"] = item.ViewName
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemDeliveries(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewDeliveries) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["default"] = boolPtrToString(item.Default)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFieldGroups(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["field_group_display_name"] = item.FieldGroupDisplayName
		respItem["field_group_name"] = item.FieldGroupName
		respItem["fields"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFieldGroupsFields(item.Fields)
		respItem["table_id"] = item.TableID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFieldGroupsFields(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroupsFields) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFilters(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFilters) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["additional_info"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersAdditionalInfo(item.AdditionalInfo)
		respItem["cache_filter"] = boolPtrToString(item.CacheFilter)
		respItem["data_type"] = item.DataType
		respItem["display_name"] = item.DisplayName
		respItem["filter_source"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersFilterSource(item.FilterSource)
		respItem["name"] = item.Name
		respItem["required"] = boolPtrToString(item.Required)
		respItem["time_options"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersTimeOptions(item.TimeOptions)
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersAdditionalInfo(item *dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersFilterSource(item *dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSource) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data_source"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersFilterSourceDataSource(item.DataSource)
	respItem["display_value_path"] = item.DisplayValuePath
	respItem["root_path"] = item.RootPath
	respItem["value_path"] = item.ValuePath

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersFilterSourceDataSource(item *dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSourceDataSource) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFiltersTimeOptions(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersTimeOptions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["info"] = item.Info
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFormats(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFormats) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["format"] = item.Format
		respItem["name"] = item.Name
		respItem["default"] = boolPtrToString(item.Default)
		respItem["template"] = flattenReportsGetViewDetailsForAGivenViewGroupViewItemFormatsTemplate(item.Template)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemFormatsTemplate(item *dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewFormatsTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["js_template_id"] = item.JsTemplateID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetViewDetailsForAGivenViewGroupViewItemSchedules(items *[]dnacentersdkgo.ResponseReportsGetViewDetailsForAGivenViewGroupViewSchedules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["default"] = boolPtrToString(item.Default)
		respItems = append(respItems, respItem)
	}
	return respItems
}
