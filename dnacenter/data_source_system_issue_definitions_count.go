package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemIssueDefinitionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Get the count of system defined issue definitions based on provided filters. Supported filters are id, name, profileId
and definition enable status. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSystemIssueDefinitionsCountRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. The definition identifier.
Examples:
id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request)
id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"issue_enabled": &schema.Schema{
				Description: `issueEnabled query parameter. The enablement status of the issue definition, either true or false.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The list of system defined issue names. (Ex."BGP_Down")
Examples:
name=BGP_Down (single entity uuid requested)
name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Description: `priority query parameter. Issue priority, possible values are P1, P2, P3, P4.
*P1*: A critical issue that needs immediate attention and can have a wide impact on network operations.
*P2*: A major issue that can potentially impact multiple devices or clients.
*P3*: A minor issue that has a localized or minimal impact.
*P4*: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Description: `profileId query parameter. The profile identier to fetch the profile associated issue defintions. The default is *global*. Please refer Network design profiles documentation for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceSystemIssueDefinitionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vProfileID, okProfileID := d.GetOk("profile_id")
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vPriority, okPriority := d.GetOk("priority")
	vIssueEnabled, okIssueEnabled := d.GetOk("issue_enabled")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters")

		headerParams1 := dnacentersdkgo.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams{}
		queryParams1 := dnacentersdkgo.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams{}

		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okProfileID {
			queryParams1.ProfileID = vProfileID.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okIssueEnabled {
			queryParams1.IssueEnabled = vIssueEnabled.(bool)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Issues.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters", err,
				"Failure at GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersItem(item *dnacentersdkgo.ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
