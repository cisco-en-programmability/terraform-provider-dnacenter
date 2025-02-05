package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsBugsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of network bugs
`,

		ReadContext: dataSourceNetworkBugsResultsBugsCountRead,
		Schema: map[string]*schema.Schema{
			"device_count": &schema.Schema{
				Description: `deviceCount query parameter. Return network bugs with deviceCount greater than this deviceCount
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of the network bug
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
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
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkBugsResultsBugsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vDeviceCount, okDeviceCount := d.GetOk("device_count")
	vSeverity, okSeverity := d.GetOk("severity")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfNetworkBugs")
		queryParams1 := dnacentersdkgo.GetCountOfNetworkBugsQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okDeviceCount {
			queryParams1.DeviceCount = vDeviceCount.(float64)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfNetworkBugs(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfNetworkBugs", err,
				"Failure at GetCountOfNetworkBugs, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfNetworkBugsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfNetworkBugs response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfNetworkBugsItem(item *dnacentersdkgo.ResponseComplianceGetCountOfNetworkBugsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
