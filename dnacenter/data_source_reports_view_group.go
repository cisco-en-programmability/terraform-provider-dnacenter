package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsViewGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Gives a list of summary of all view groups.

- Gives a list of summary of all views in a viewgroup. Use "Get all view groups" API to get the viewGroupIds (required
as a query param for this API) for available viewgroups.
`,

		ReadContext: dataSourceReportsViewGroupRead,
		Schema: map[string]*schema.Schema{
			"view_group_id": &schema.Schema{
				Description: `viewGroupId path parameter. viewGroupId of viewgroup.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"view_group_id": &schema.Schema{
							Description: `viewgroup Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"views": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_id": &schema.Schema{
										Description: `Unique id for a view within viewgroup
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
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category": &schema.Schema{
							Description: `category of the view group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `view group description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `name of view group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_group_id": &schema.Schema{
							Description: `id of viewgroup
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

func dataSourceReportsViewGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vViewGroupID, okViewGroupID := d.GetOk("view_group_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okViewGroupID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAllViewGroups")

		response1, restyResp1, err := client.Reports.GetAllViewGroups()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllViewGroups", err,
				"Failure at GetAllViewGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenReportsGetAllViewGroupsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllViewGroups response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetViewsForAGivenViewGroup")
		vvViewGroupID := vViewGroupID.(string)

		response2, restyResp2, err := client.Reports.GetViewsForAGivenViewGroup(vvViewGroupID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetViewsForAGivenViewGroup", err,
				"Failure at GetViewsForAGivenViewGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenReportsGetViewsForAGivenViewGroupItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetViewsForAGivenViewGroup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetAllViewGroupsItems(items *dnacentersdkgo.ResponseReportsGetAllViewGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["category"] = item.Category
		respItem["description"] = item.Description
		respItem["name"] = item.Name
		respItem["view_group_id"] = item.ViewGroupID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewsForAGivenViewGroupItem(item *dnacentersdkgo.ResponseReportsGetViewsForAGivenViewGroup) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_group_id"] = item.ViewGroupID
	respItem["views"] = flattenReportsGetViewsForAGivenViewGroupItemViews(item.Views)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetViewsForAGivenViewGroupItemViews(items *[]dnacentersdkgo.ResponseReportsGetViewsForAGivenViewGroupViews) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["description"] = item.Description
		respItem["view_id"] = item.ViewID
		respItem["view_name"] = item.ViewName
		respItems = append(respItems, respItem)
	}
	return respItems
}
