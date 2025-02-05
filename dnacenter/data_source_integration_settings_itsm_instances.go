package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationSettingsItsmInstances() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM Integration.

- Fetches all ITSM Integration settings
`,

		ReadContext: dataSourceIntegrationSettingsItsmInstancesRead,
		Schema: map[string]*schema.Schema{
			"order": &schema.Schema{
				Description: `order query parameter. Specify the sorting order asc for ascending or desc for descending
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Indicates the current page number to display.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"page_size": &schema.Schema{
				Description: `page_size query parameter. Specifies the number of records to display per page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. The field name used to sort the records.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"type_id": &schema.Schema{
										Description: `Deprecated: Should not be used and will be removed in an upcoming release
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"created_by": &schema.Schema{
										Description: `Created By`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"created_date": &schema.Schema{
										Description: `Created Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dyp_id": &schema.Schema{
										Description: `Dyp Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dyp_major_version": &schema.Schema{
										Description: `Dyp Major Version`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dyp_name": &schema.Schema{
										Description: `Dyp Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"schema_version": &schema.Schema{
										Description: `Schema Version`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"software_version_log": &schema.Schema{
										Description: `Software Version Log`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"unique_key": &schema.Schema{
										Description: `Unique Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"updated_by": &schema.Schema{
										Description: `Updated By`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"updated_date": &schema.Schema{
										Description: `Updated Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"page": &schema.Schema{
							Description: `Page`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"page_size": &schema.Schema{
							Description: `Page Size`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_pages": &schema.Schema{
							Description: `Total Pages`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_records": &schema.Schema{
							Description: `Total Records`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationSettingsItsmInstancesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllItsmIntegrationSettings")

		response1, restyResp1, err := client.ItsmIntegration.GetAllItsmIntegrationSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllItsmIntegrationSettings", err,
				"Failure at GetAllItsmIntegrationSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenItsmIntegrationGetAllItsmIntegrationSettingsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllItsmIntegrationSettings response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmIntegrationGetAllItsmIntegrationSettingsItem(item *dnacentersdkgo.ResponseItsmIntegrationGetAllItsmIntegrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["page"] = item.Page
	respItem["page_size"] = item.PageSize
	respItem["total_pages"] = item.TotalPages
	respItem["data"] = flattenItsmIntegrationGetAllItsmIntegrationSettingsItemData(&item.Data)
	respItem["total_records"] = item.TotalRecords
	return []map[string]interface{}{
		respItem,
	}
}

func flattenItsmIntegrationGetAllItsmIntegrationSettingsItemData(items *[]dnacentersdkgo.ResponseItsmIntegrationGetAllItsmIntegrationSettingsData) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type_id"] = item.TypeID
		respItem["id"] = item.ID
		respItem["created_by"] = item.CreatedBy
		respItem["description"] = item.Description
		respItem["dyp_id"] = item.DypID
		respItem["dyp_major_version"] = item.DypMajorVersion
		respItem["dyp_name"] = item.DypName
		respItem["name"] = item.Name
		respItem["schema_version"] = item.SchemaVersion
		respItem["software_version_log"] = flattenItsmIntegrationGetAllItsmIntegrationSettingsItemDataSoftwareVersionLog(item.SoftwareVersionLog)
		respItem["unique_key"] = item.UniqueKey
		respItem["updated_by"] = item.UpdatedBy
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenItsmIntegrationGetAllItsmIntegrationSettingsItemDataSoftwareVersionLog(items *[]dnacentersdkgo.ResponseItsmIntegrationGetAllItsmIntegrationSettingsDataSoftwareVersionLog) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
