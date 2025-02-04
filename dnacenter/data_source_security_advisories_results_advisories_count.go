package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsAdvisoriesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of security advisories affecting the network devices
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsAdvisoriesCountRead,
		Schema: map[string]*schema.Schema{
			"cvss_base_score": &schema.Schema{
				Description: `cvssBaseScore query parameter. Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_count": &schema.Schema{
				Description: `deviceCount query parameter. Return advisories with deviceCount greater than this deviceCount
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of the security advisory
`,
				Type:     schema.TypeString,
				Optional: true,
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

func dataSourceSecurityAdvisoriesResultsAdvisoriesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vDeviceCount, okDeviceCount := d.GetOk("device_count")
	vCvssBaseScore, okCvssBaseScore := d.GetOk("cvss_base_score")
	vSecurityImpactRating, okSecurityImpactRating := d.GetOk("security_impact_rating")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices")
		queryParams1 := dnacentersdkgo.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okDeviceCount {
			queryParams1.DeviceCount = vDeviceCount.(float64)
		}
		if okCvssBaseScore {
			queryParams1.CvssBaseScore = vCvssBaseScore.(string)
		}
		if okSecurityImpactRating {
			queryParams1.SecurityImpactRating = vSecurityImpactRating.(string)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices", err,
				"Failure at GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesItem(item *dnacentersdkgo.ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
