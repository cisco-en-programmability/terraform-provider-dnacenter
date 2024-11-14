package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time
`,

		ReadContext: dataSourceClientHealthRead,
		Schema: map[string]*schema.Schema{
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the Client health data is required
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"score_detail": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_count": &schema.Schema{
										Description: `Total client count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"client_unique_count": &schema.Schema{
										Description: `Total unique client count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"connected_to_udn_count": &schema.Schema{
										Description: `Total connected to UDN count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"duid_count": &schema.Schema{
										Description: `Device UUID count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"endtime": &schema.Schema{
										Description: `UTC timestamp of data end time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"maintenance_affected_client_count": &schema.Schema{
										Description: `Total client count affected by maintenance
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"random_mac_count": &schema.Schema{
										Description: `Total client count with random MAC count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"score_category": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"score_category": &schema.Schema{
													Description: `Health score category
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Description: `Health score category value
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"score_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_count": &schema.Schema{
													Description: `Total client count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"client_unique_count": &schema.Schema{
													Description: `Total unique client count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"connected_to_udn_count": &schema.Schema{
													Description: `Total connected to UDN count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"duid_count": &schema.Schema{
													Description: `Device UUID count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"endtime": &schema.Schema{
													Description: `UTC timestamp of data end time
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"maintenance_affected_client_count": &schema.Schema{
													Description: `Total client count affected by maintenance
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"random_mac_count": &schema.Schema{
													Description: `Total client count with random MAC count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"score_category": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"score_category": &schema.Schema{
																Description: `Category of the overall health score
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"value": &schema.Schema{
																Description: `Health score category value
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"score_value": &schema.Schema{
													Description: `Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"starttime": &schema.Schema{
													Description: `UTC timestamp of data start time
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"unconnected_to_udn_count": &schema.Schema{
													Description: `Total unconnected to UDN count
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"score_value": &schema.Schema{
										Description: `Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"starttime": &schema.Schema{
										Description: `UTC timestamp of data start time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"unconnected_to_udn_count": &schema.Schema{
										Description: `Total unconnected to UDN count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"site_id": &schema.Schema{
							Description: `Site UUID or 'global'
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

func dataSourceClientHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTimestamp, okTimestamp := d.GetOk("timestamp")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOverallClientHealth")
		queryParams1 := dnacentersdkgo.GetOverallClientHealthQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(float64)
		}

		response1, restyResp1, err := client.Clients.GetOverallClientHealth(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetOverallClientHealth", err,
				"Failure at GetOverallClientHealth, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenClientsGetOverallClientHealthItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetOverallClientHealth response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsGetOverallClientHealthItems(items *[]dnacentersdkgo.ResponseClientsGetOverallClientHealthResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_id"] = item.SiteID
		respItem["score_detail"] = flattenClientsGetOverallClientHealthItemsScoreDetail(item.ScoreDetail)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetOverallClientHealthItemsScoreDetail(items *[]dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetail) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["score_category"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreCategory(item.ScoreCategory)
		respItem["score_value"] = item.ScoreValue
		respItem["client_count"] = item.ClientCount
		respItem["client_unique_count"] = item.ClientUniqueCount
		respItem["maintenance_affected_client_count"] = item.MaintenanceAffectedClientCount
		respItem["random_mac_count"] = item.RandomMacCount
		respItem["duid_count"] = item.DuidCount
		respItem["starttime"] = item.Starttime
		respItem["endtime"] = item.Endtime
		respItem["connected_to_udn_count"] = item.ConnectedToUdnCount
		respItem["unconnected_to_udn_count"] = item.UnconnectedToUdnCount
		respItem["score_list"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreList(item.ScoreList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreCategory(item *dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreCategory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["score_category"] = item.ScoreCategory
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreList(items *[]dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["score_category"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreCategory(item.ScoreCategory)
		respItem["score_value"] = item.ScoreValue
		respItem["client_count"] = item.ClientCount
		respItem["client_unique_count"] = item.ClientUniqueCount
		respItem["maintenance_affected_client_count"] = item.MaintenanceAffectedClientCount
		respItem["random_mac_count"] = item.RandomMacCount
		respItem["duid_count"] = item.DuidCount
		respItem["starttime"] = item.Starttime
		respItem["endtime"] = item.Endtime
		respItem["connected_to_udn_count"] = item.ConnectedToUdnCount
		respItem["unconnected_to_udn_count"] = item.UnconnectedToUdnCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreCategory(item *dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreCategory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["score_category"] = item.ScoreCategory
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}
