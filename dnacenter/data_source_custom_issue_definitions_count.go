package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomIssueDefinitionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Get the total number of Custom issue definitions count based on the provided filters. The supported filters are id,
name, profileId and definition enable status, severity, facility and mnemonic. For detailed information about the usage
of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceCustomIssueDefinitionsCountRead,
		Schema: map[string]*schema.Schema{
			"facility": &schema.Schema{
				Description: `facility query parameter. The syslog facility name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. The custom issue definition identifier and unique identifier across the profile. Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_enabled": &schema.Schema{
				Description: `isEnabled query parameter. The enable status of the custom issue definition, either true or false.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mnemonic": &schema.Schema{
				Description: `mnemonic query parameter. The syslog mnemonic name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The list of UDI issue names. (Ex."TestUdiIssues")
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Description: `priority query parameter. The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Description: `profileId query parameter. The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
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

func dataSourceCustomIssueDefinitionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vProfileID, okProfileID := d.GetOk("profile_id")
	vName, okName := d.GetOk("name")
	vPriority, okPriority := d.GetOk("priority")
	vIsEnabled, okIsEnabled := d.GetOk("is_enabled")
	vSeverity, okSeverity := d.GetOk("severity")
	vFacility, okFacility := d.GetOk("facility")
	vMnemonic, okMnemonic := d.GetOk("mnemonic")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters")

		headerParams1 := dnacentersdkgo.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams{}
		queryParams1 := dnacentersdkgo.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okProfileID {
			queryParams1.ProfileID = vProfileID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okIsEnabled {
			queryParams1.IsEnabled = vIsEnabled.(bool)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(float64)
		}
		if okFacility {
			queryParams1.Facility = vFacility.(string)
		}
		if okMnemonic {
			queryParams1.Mnemonic = vMnemonic.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Issues.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters", err,
				"Failure at GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersItem(item *dnacentersdkgo.ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
