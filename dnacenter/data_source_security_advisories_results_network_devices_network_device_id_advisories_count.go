package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisoriesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of security advisories affecting the network device
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisoriesCountRead,
		Schema: map[string]*schema.Schema{
			"cvss_base_score": &schema.Schema{
				Description: `cvssBaseScore query parameter. Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of the security advisory
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"security_impact_rating": &schema.Schema{
				Description: `securityImpactRating query parameter. Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
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

func dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisoriesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vID, okID := d.GetOk("id")
	vCvssBaseScore, okCvssBaseScore := d.GetOk("cvss_base_score")
	vSecurityImpactRating, okSecurityImpactRating := d.GetOk("security_impact_rating")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okCvssBaseScore {
			queryParams1.CvssBaseScore = vCvssBaseScore.(string)
		}
		if okSecurityImpactRating {
			queryParams1.SecurityImpactRating = vSecurityImpactRating.(string)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice", err,
				"Failure at GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceItem(item *dnacentersdkgo.ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
