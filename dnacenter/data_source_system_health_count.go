package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemHealthCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source gives the count of the latest system events
`,

		ReadContext: dataSourceSystemHealthCountRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: `domain query parameter. Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdomain": &schema.Schema{
				Description: `subdomain query parameter. Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count of the events
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

func dataSourceSystemHealthCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDomain, okDomain := d.GetOk("domain")
	vSubdomain, okSubdomain := d.GetOk("subdomain")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SystemHealthCountApI")
		queryParams1 := dnacentersdkgo.SystemHealthCountApIQueryParams{}

		if okDomain {
			queryParams1.Domain = vDomain.(string)
		}
		if okSubdomain {
			queryParams1.Subdomain = vSubdomain.(string)
		}

		response1, restyResp1, err := client.HealthAndPerformance.SystemHealthCountApI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SystemHealthCountApI", err,
				"Failure at SystemHealthCountApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemHealthCountApIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemHealthCountApI response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemHealthCountApIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemHealthCountApI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
