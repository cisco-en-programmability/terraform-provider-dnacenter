package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugs() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get bugs affecting the network device
`,

		ReadContext: dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. Id of the network bug
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
			"severity": &schema.Schema{
				Description: `severity query parameter. Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE.
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

						"affected_versions": &schema.Schema{
							Description: `Versions that are affected by the network bug
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"device_count": &schema.Schema{
							Description: `Number of devices which are vulnerable to this network bug
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"has_workaround": &schema.Schema{
							Description: `Indicates if the network bug has a workaround
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"headline": &schema.Schema{
							Description: `Title of the network bug
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the network bug
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"integrated_releases": &schema.Schema{
							Description: `Versions that have the fix for the network bug
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"publication_url": &schema.Schema{
							Description: `Url for getting network bug details on cisco website
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"severity": &schema.Schema{
							Description: `'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"workaround": &schema.Schema{
							Description: `Workaround if any that exists for the network bug
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

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vID, okID := d.GetOk("id")
	vSeverity, okSeverity := d.GetOk("severity")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetBugsAffectingTheNetworkDevice")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetBugsAffectingTheNetworkDeviceQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
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

		response1, restyResp1, err := client.Compliance.GetBugsAffectingTheNetworkDevice(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetBugsAffectingTheNetworkDevice", err,
				"Failure at GetBugsAffectingTheNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetBugsAffectingTheNetworkDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetBugsAffectingTheNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetBugsAffectingTheNetworkDeviceItems(items *[]dnacentersdkgo.ResponseComplianceGetBugsAffectingTheNetworkDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["headline"] = item.Headline
		respItem["publication_url"] = item.PublicationURL
		respItem["device_count"] = item.DeviceCount
		respItem["severity"] = item.Severity
		respItem["has_workaround"] = boolPtrToString(item.HasWorkaround)
		respItem["workaround"] = item.Workaround
		respItem["affected_versions"] = item.AffectedVersions
		respItem["integrated_releases"] = item.IntegratedReleases
		respItems = append(respItems, respItem)
	}
	return respItems
}
