package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplications() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get applications by offset/limit or by name
`,

		ReadContext: dataSourceApplicationsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The maximum number of applications to be returned
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Application's name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset of the first application to be returned
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_set": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id_ref": &schema.Schema{
										Description: `Id Ref`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"indicative_network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `displayName`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"lower_port": &schema.Schema{
										Description: `lowerPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ports": &schema.Schema{
										Description: `ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upper_port": &schema.Schema{
										Description: `upperPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_protocol": &schema.Schema{
										Description: `App Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"application_sub_type": &schema.Schema{
										Description: `Application Sub Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"application_type": &schema.Schema{
										Description: `Application Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"category_id": &schema.Schema{
										Description: `Category Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dscp": &schema.Schema{
										Description: `Dscp`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"engine_id": &schema.Schema{
										Description: `Engine Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"help_string": &schema.Schema{
										Description: `Help String`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ignore_conflict": &schema.Schema{
										Description: `Ignore Conflict`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"long_description": &schema.Schema{
										Description: `Long Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"popularity": &schema.Schema{
										Description: `Popularity`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rank": &schema.Schema{
										Description: `Rank`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"server_name": &schema.Schema{
										Description: `Server Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"traffic_class": &schema.Schema{
										Description: `Traffic Class`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"url": &schema.Schema{
										Description: `Url`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"lower_port": &schema.Schema{
										Description: `Lower Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ports": &schema.Schema{
										Description: `Ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upper_port": &schema.Schema{
										Description: `Upper Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplications")
		queryParams1 := dnacentersdkgo.GetApplicationsQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplications(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplications", err,
				"Failure at GetApplications, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplications response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationsItem(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["indicative_network_identity"] = flattenApplicationPolicyGetApplicationsItemsIndicativeNetworkIDentity(item.IndicativeNetworkIDentity)
	respItem["network_applications"] = flattenApplicationPolicyGetApplicationsItemsNetworkApplications(item.NetworkApplications)
	respItem["network_identity"] = flattenApplicationPolicyGetApplicationsItemsNetworkIDentity(item.NetworkIDentity)
	respItem["application_set"] = flattenApplicationPolicyGetApplicationsItemsApplicationSet(item.ApplicationSet)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenApplicationPolicyGetApplicationsItems(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["indicative_network_identity"] = flattenApplicationPolicyGetApplicationsItemsIndicativeNetworkIDentity(item.IndicativeNetworkIDentity)
		respItem["network_applications"] = flattenApplicationPolicyGetApplicationsItemsNetworkApplications(item.NetworkApplications)
		respItem["network_identity"] = flattenApplicationPolicyGetApplicationsItemsNetworkIDentity(item.NetworkIDentity)
		respItem["application_set"] = flattenApplicationPolicyGetApplicationsItemsApplicationSet(item.ApplicationSet)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsItemsIndicativeNetworkIDentity(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponseIndicativeNetworkIDentity) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["display_name"] = item.DisplayName
		respItem["lower_port"] = item.LowerPort
		respItem["ports"] = item.Ports
		respItem["protocol"] = item.Protocol
		respItem["upper_port"] = item.UpperPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsItemsNetworkApplications(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponseNetworkApplications) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["app_protocol"] = item.AppProtocol
		respItem["application_sub_type"] = item.ApplicationSubType
		respItem["application_type"] = item.ApplicationType
		respItem["category_id"] = item.CategoryID
		respItem["display_name"] = item.DisplayName
		respItem["engine_id"] = item.EngineID
		respItem["help_string"] = item.HelpString
		respItem["long_description"] = item.LongDescription
		respItem["name"] = item.Name
		respItem["popularity"] = item.Popularity
		respItem["rank"] = item.Rank
		respItem["traffic_class"] = item.TrafficClass
		respItem["server_name"] = item.ServerName
		respItem["url"] = item.URL
		respItem["dscp"] = item.Dscp
		respItem["ignore_conflict"] = item.IgnoreConflict
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsItemsNetworkIDentity(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponseNetworkIDentity) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["display_name"] = item.DisplayName
		respItem["lower_port"] = item.LowerPort
		respItem["ports"] = item.Ports
		respItem["protocol"] = item.Protocol
		respItem["upper_port"] = item.UpperPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsItemsApplicationSet(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsResponseApplicationSet) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id_ref"] = item.IDRef

	return []map[string]interface{}{
		respItem,
	}

}
