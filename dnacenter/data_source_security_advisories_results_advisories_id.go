package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsAdvisoriesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get security advisory affecting the network devices by Id
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsAdvisoriesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the security advisory
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cve_ids": &schema.Schema{
							Description: `CVE (Common Vulnerabilities and Exposures) ID of the advisory
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"cvss_base_score": &schema.Schema{
							Description: `Common Vulnerability Scoring System(CVSS) base score
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"device_count": &schema.Schema{
							Description: `Number of devices which are vulnerable to this advisory
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"first_fixed_versions_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fixed_versions": &schema.Schema{
										Description: `First versions that have the fix for the advisory
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"vulnerable_version": &schema.Schema{
										Description: `Version that is vulnerable to the advisory
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id of the security advisory
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"publication_url": &schema.Schema{
							Description: `Url for getting advisory details on cisco website
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"security_impact_rating": &schema.Schema{
							Description: `'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
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

func dataSourceSecurityAdvisoriesResultsAdvisoriesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityAdvisoryAffectingTheNetworkDevicesByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Compliance.GetSecurityAdvisoryAffectingTheNetworkDevicesByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSecurityAdvisoryAffectingTheNetworkDevicesByID", err,
				"Failure at GetSecurityAdvisoryAffectingTheNetworkDevicesByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityAdvisoryAffectingTheNetworkDevicesByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDItem(item *dnacentersdkgo.ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["device_count"] = item.DeviceCount
	respItem["cve_ids"] = item.CveIDs
	respItem["publication_url"] = item.PublicationURL
	respItem["cvss_base_score"] = item.CvssBaseScore
	respItem["security_impact_rating"] = item.SecurityImpactRating
	respItem["first_fixed_versions_list"] = flattenComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDItemFirstFixedVersionsList(item.FirstFixedVersionsList)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDItemFirstFixedVersionsList(items *[]dnacentersdkgo.ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponseFirstFixedVersionsList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["vulnerable_version"] = item.VulnerableVersion
		respItem["fixed_versions"] = item.FixedVersions
		respItems = append(respItems, respItem)
	}
	return respItems
}
