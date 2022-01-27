package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabric() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get SDA Fabric Info
`,

		ReadContext: dataSourceSdaFabricRead,
		Schema: map[string]*schema.Schema{
			"fabric_name": &schema.Schema{
				Description: `fabricName query parameter. Fabric Name
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fabric_domain_type": &schema.Schema{
							Description: `Fabric Domain type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_name": &schema.Schema{
							Description: `Fabric name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_type": &schema.Schema{
							Description: `Fabric type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaFabricRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricName := d.Get("fabric_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSdaFabricInfo")
		queryParams1 := dnacentersdkgo.GetSdaFabricInfoQueryParams{}

		queryParams1.FabricName = vFabricName.(string)

		response1, restyResp1, err := client.Sda.GetSdaFabricInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSdaFabricInfo", err,
				"Failure at GetSdaFabricInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetSdaFabricInfoItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaFabricInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetSdaFabricInfoItem(item *dnacentersdkgo.ResponseSdaGetSdaFabricInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["execution_id"] = item.ExecutionID
	respItem["fabric_name"] = item.FabricName
	respItem["fabric_type"] = item.FabricType
	respItem["fabric_domain_type"] = item.FabricDomainType
	return []map[string]interface{}{
		respItem,
	}
}
