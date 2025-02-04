package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHealthScoreDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get all health score defintions.
Supported filters are id, name and overall health include status. A health score definition can be different across
device type. So, deviceType in the query param is important and default is all device types.
By default all supported attributes are listed in response. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml

- Get health score defintion for the given id. Definition includes all properties from HealthScoreDefinition schema by
default. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceHealthScoreDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. These are the attributes supported in health score definitions response. By default, all properties are sent in response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Health score definition id.
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
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"definition_status": &schema.Schema{
							Description: `Definition Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"include_for_overall_health": &schema.Schema{
							Description: `Include For Overall Health`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"synchronize_to_issue_threshold": &schema.Schema{
							Description: `Synchronize To Issue Threshold`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"threshold_value": &schema.Schema{
							Description: `Threshold Value`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHealthScoreDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vID, okID := d.GetOk("id")
	vIncludeForOverallHealth, okIncludeForOverallHealth := d.GetOk("include_for_overall_health")
	vAttribute, okAttribute := d.GetOk("attribute")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")

	method1 := []bool{okDeviceType, okID, okIncludeForOverallHealth, okAttribute, okOffset, okLimit, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllHealthScoreDefinitionsForGivenFilters")

		headerParams1 := dnacentersdkgo.GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams{}
		queryParams1 := dnacentersdkgo.GetAllHealthScoreDefinitionsForGivenFiltersQueryParams{}

		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okIncludeForOverallHealth {
			queryParams1.IncludeForOverallHealth = vIncludeForOverallHealth.(bool)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Devices.GetAllHealthScoreDefinitionsForGivenFilters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllHealthScoreDefinitionsForGivenFilters", err,
				"Failure at GetAllHealthScoreDefinitionsForGivenFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetHealthScoreDefinitionForTheGivenID")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.GetHealthScoreDefinitionForTheGivenIDHeaderParams{}

		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Devices.GetHealthScoreDefinitionForTheGivenID(vvID, &headerParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetHealthScoreDefinitionForTheGivenID", err,
				"Failure at GetHealthScoreDefinitionForTheGivenID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItems2 := flattenDevicesGetHealthScoreDefinitionForTheGivenIDItems(response2.Response)
		if err := d.Set("items", vItems2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetHealthScoreDefinitionForTheGivenID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetHealthScoreDefinitionForTheGivenIDItems(items *[]dnacentersdkgo.ResponseDevicesGetHealthScoreDefinitionForTheGivenIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["display_name"] = item.DisplayName
		respItem["device_family"] = item.DeviceFamily
		respItem["description"] = item.Description
		respItem["include_for_overall_health"] = boolPtrToString(item.IncludeForOverallHealth)
		respItem["definition_status"] = item.DefinitionStatus
		respItem["threshold_value"] = item.ThresholdValue
		respItem["synchronize_to_issue_threshold"] = boolPtrToString(item.SynchronizeToIssueThreshold)
		respItem["last_modified"] = item.LastModified
		respItems = append(respItems, respItem)
	}
	return respItems
}
