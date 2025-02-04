package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHealthScoreDefinitionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health
include status. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceHealthScoreDefinitionsCountRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
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
			"include_for_overall_health": &schema.Schema{
				Description: `includeForOverallHealth query parameter. The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
`,
				Type:     schema.TypeBool,
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

func dataSourceHealthScoreDefinitionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vID, okID := d.GetOk("id")
	vIncludeForOverallHealth, okIncludeForOverallHealth := d.GetOk("include_for_overall_health")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters")

		headerParams1 := dnacentersdkgo.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams{}
		queryParams1 := dnacentersdkgo.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams{}

		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okIncludeForOverallHealth {
			queryParams1.IncludeForOverallHealth = vIncludeForOverallHealth.(bool)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters", err,
				"Failure at GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersItem(item *dnacentersdkgo.ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
