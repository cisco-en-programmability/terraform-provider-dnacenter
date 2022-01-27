package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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
				Type:     schema.TypeString,
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
										Description: `Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"client_unique_count": &schema.Schema{
										Description: `Client Unique Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"endtime": &schema.Schema{
										Description: `Endtime`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"score_category": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"score_category": &schema.Schema{
													Description: `Score Category`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
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
													Description: `Client Count`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"client_unique_count": &schema.Schema{
													Description: `Client Unique Count`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"endtime": &schema.Schema{
													Description: `Endtime`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"score_category": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"score_category": &schema.Schema{
																Description: `Score Category`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Description: `Client Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"client_unique_count": &schema.Schema{
																Description: `Client Unique Count`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"endtime": &schema.Schema{
																Description: `Endtime`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"score_category": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"score_category": &schema.Schema{
																			Description: `Score Category`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"value": &schema.Schema{
																			Description: `Value`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},

															"score_value": &schema.Schema{
																Description: `Score Value`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"starttime": &schema.Schema{
																Description: `Starttime`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"score_value": &schema.Schema{
													Description: `Score Value`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"starttime": &schema.Schema{
													Description: `Starttime`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},

									"score_value": &schema.Schema{
										Description: `Score Value`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"starttime": &schema.Schema{
										Description: `Starttime`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
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
		log.Printf("[DEBUG] Selected method 1: GetOverallClientHealth")
		queryParams1 := dnacentersdkgo.GetOverallClientHealthQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(string)
		}

		response1, restyResp1, err := client.Clients.GetOverallClientHealth(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetOverallClientHealth", err,
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
		respItem["starttime"] = item.Starttime
		respItem["endtime"] = item.Endtime
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
		respItem["starttime"] = item.Starttime
		respItem["endtime"] = item.Endtime
		respItem["score_list"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreList(item.ScoreList)
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

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreList(items *[]dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["score_category"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreListScoreCategory(item.ScoreCategory)
		respItem["score_value"] = item.ScoreValue
		respItem["client_count"] = item.ClientCount
		respItem["client_unique_count"] = flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreListClientUniqueCount(item.ClientUniqueCount)
		respItem["starttime"] = item.Starttime
		respItem["endtime"] = item.Endtime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreListScoreCategory(item *dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListScoreCategory) []map[string]interface{} {
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

func flattenClientsGetOverallClientHealthItemsScoreDetailScoreListScoreListClientUniqueCount(item *dnacentersdkgo.ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListClientUniqueCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
