package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationSets() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get appllication-sets by offset/limit or by name
`,

		ReadContext: dataSourceApplicationSetsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApplicationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationSets")
		queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationSets(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationSets", err,
				"Failure at GetApplicationSets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationSetsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationSets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationSetsItems(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["identity_source"] = flattenApplicationPolicyGetApplicationSetsItemsIDentitySource(item.IDentitySource)
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationSetsItem(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}

	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["identity_source"] = flattenApplicationPolicyGetApplicationSetsItemsIDentitySource(item.IDentitySource)
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}
}

func flattenApplicationPolicyGetApplicationSetsItemsIDentitySource(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponseIDentitySource) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
