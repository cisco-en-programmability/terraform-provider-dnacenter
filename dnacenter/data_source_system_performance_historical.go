package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemPerformanceHistorical() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- Retrieves the average values of cluster key performance indicators (KPIs), like CPU utilization, memory utilization or
network rates grouped by time intervals within a specified time range. The data will be available from the past 24
hours.
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

									"cpu_avg": &schema.Schema{
										Description: `CPU average utilization
`,
										Type:     schema.TypeString,
										Computed: true,
									},

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

									"memory_avg": &schema.Schema{
										Description: `Memory average utilization
`,
										Type:     schema.TypeString,
										Computed: true,
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
		log.Printf("[DEBUG] Selected method: SystemPerformanceHistoricalAPI")
		queryParams1 := dnacentersdkgo.SystemPerformanceHistoricalAPIQueryParams{}

		if okKpi {
			queryParams1.Kpi = vKpi.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}

		response1, restyResp1, err := client.HealthAndPerformance.SystemPerformanceHistoricalAPI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemPerformanceHistoricalAPI", err,
				"Failure at SystemPerformanceHistoricalAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemPerformanceHistoricalAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["version"] = item.Version
	respItem["kpis"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpis(item.Kpis)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpis(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["legends"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegends(item.Legends)
	respItem["data"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisData(item.Data)
	respItem["cpu_avg"] = item.CPUAvg
	respItem["memory_avg"] = item.MemoryAvg

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegends(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegends) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsCPU(item.CPU)
	respItem["memory"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsMemory(item.Memory)
	respItem["network_tx_rate"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsNetworktxRate(item.NetworktxRate)
	respItem["network_rx_rate"] = flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsNetworkrxRate(item.NetworkrxRate)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsCPU(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsCPU) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsMemory(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsMemory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsNetworktxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworktxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisLegendsNetworkrxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworkrxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceHistoricalAPIItemKpisData(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["t1"] = item

	return []map[string]interface{}{
		respItem,
	}

}
