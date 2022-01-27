package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemPerformanceHistorical() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source retrieves the historical performance indicators . The data can be retrieved for the last 3 months.
`,

		ReadContext: dataSourceSystemPerformanceHistoricalRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. This is the epoch end time in milliseconds upto which performance indicator need to be fetched
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"kpi": &schema.Schema{
				Description: `kpi query parameter. Fetch historical data for this kpi. Valid values: cpu,memory,network
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. This is the epoch start time in milliseconds from which performance indicator need to be fetched
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"host_name": &schema.Schema{
							Description: `Hostname`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"kpis": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"t1": &schema.Schema{
													Description: `Time in  'YYYY-MM-DDT00:00:00Z' format with values for legends
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"legends": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"units": &schema.Schema{
																Description: `Units for cpu usage
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"memory": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"units": &schema.Schema{
																Description: `Units for memory usage
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"network_rx_rate": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"units": &schema.Schema{
																Description: `Units for network rx_rate
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"network_tx_rate": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"units": &schema.Schema{
																Description: `Units for network tx_rate
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version of the API
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

func dataSourceSystemPerformanceHistoricalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vKpi, okKpi := d.GetOk("kpi")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SystemPerformanceHistoricalApI")
		queryParams1 := dnacentersdkgo.SystemPerformanceHistoricalApIQueryParams{}

		if okKpi {
			queryParams1.Kpi = vKpi.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}

		response1, restyResp1, err := client.HealthAndPerformance.SystemPerformanceHistoricalApI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SystemPerformanceHistoricalApI", err,
				"Failure at SystemPerformanceHistoricalApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemPerformanceHistoricalApIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemPerformanceHistoricalApI response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalApI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["version"] = item.Version
	respItem["kpis"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpis(item.Kpis)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpis(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["legends"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegends(item.Legends)
	respItem["data"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegends(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegends) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsCPU(item.CPU)
	respItem["memory"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsMemory(item.Memory)
	respItem["network_tx_rate"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsNetworktxRate(item.NetworktxRate)
	respItem["network_rx_rate"] = flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsNetworkrxRate(item.NetworkrxRate)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsCPU(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsCPU) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsMemory(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsMemory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsNetworktxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworktxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisLegendsNetworkrxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworkrxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalApIItemKpisData(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["t1"] = item.T1

	return []map[string]interface{}{
		respItem,
	}

}
