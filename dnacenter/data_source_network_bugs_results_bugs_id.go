package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsBugsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get network bug by Id
`,

		ReadContext: dataSourceNetworkBugsResultsBugsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the network bug
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceNetworkBugsResultsBugsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkBugByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Compliance.GetNetworkBugByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkBugByID", err,
				"Failure at GetNetworkBugByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetNetworkBugByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkBugByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetNetworkBugByIDItem(item *dnacentersdkgo.ResponseComplianceGetNetworkBugByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	return []map[string]interface{}{
		respItem,
	}
}
