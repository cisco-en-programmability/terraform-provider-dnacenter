package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisories() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get security advisories affecting the network device
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisoriesRead,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Default value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
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
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
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

func dataSourceSecurityAdvisoriesResultsNetworkDevicesNetworkDeviceIDAdvisoriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vID, okID := d.GetOk("id")
	vCvssBaseScore, okCvssBaseScore := d.GetOk("cvss_base_score")
	vSecurityImpactRating, okSecurityImpactRating := d.GetOk("security_impact_rating")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityAdvisoriesAffectingTheNetworkDevice")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okCvssBaseScore {
			queryParams1.CvssBaseScore = vCvssBaseScore.(string)
		}
		if okSecurityImpactRating {
			queryParams1.SecurityImpactRating = vSecurityImpactRating.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Compliance.GetSecurityAdvisoriesAffectingTheNetworkDevice(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSecurityAdvisoriesAffectingTheNetworkDevice", err,
				"Failure at GetSecurityAdvisoriesAffectingTheNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityAdvisoriesAffectingTheNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceItems(items *[]dnacentersdkgo.ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["device_count"] = item.DeviceCount
		respItem["cve_ids"] = item.CveIDs
		respItem["publication_url"] = item.PublicationURL
		respItem["cvss_base_score"] = item.CvssBaseScore
		respItem["security_impact_rating"] = item.SecurityImpactRating
		respItem["first_fixed_versions_list"] = flattenComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceItemsFirstFixedVersionsList(item.FirstFixedVersionsList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceItemsFirstFixedVersionsList(items *[]dnacentersdkgo.ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponseFirstFixedVersionsList) []map[string]interface{} {
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
