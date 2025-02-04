package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationsCountV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get the number of all existing applications
`,

		ReadContext: dataSourceApplicationsCountV2Read,
		Schema: map[string]*schema.Schema{
			"scalable_group_type": &schema.Schema{
				Description: `scalableGroupType query parameter. scalable group type to retrieve, valid value APPLICATION
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Total number of Application
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version
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

func dataSourceApplicationsCountV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vScalableGroupType := d.Get("scalable_group_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationCountV2")
		queryParams1 := dnacentersdkgo.GetApplicationCountV2QueryParams{}

		queryParams1.ScalableGroupType = vScalableGroupType.(string)

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationCountV2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApplicationCountV2", err,
				"Failure at GetApplicationCountV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyGetApplicationCountV2Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationCountV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationCountV2Item(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationCountV2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
