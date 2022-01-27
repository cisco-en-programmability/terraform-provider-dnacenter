package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemPerformance() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source gives the aggregated performance indicators. The data can be retrieved for the last 3 months.
`,

		ReadContext: dataSourceSystemPerformanceRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. This is the epoch end time in milliseconds upto which performance indicator need to be fetched
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"function": &schema.Schema{
				Description: `function query parameter. Valid values: sum,average,max
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"kpi": &schema.Schema{
				Description: `kpi query parameter. Valid values: cpu,memory,network
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

												"utilization": &schema.Schema{
													Description: `cpu usage in units
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

												"utilization": &schema.Schema{
													Description: `Memory usage in units
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

												"utilization": &schema.Schema{
													Description: `Network rx_rate in units
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

												"utilization": &schema.Schema{
													Description: `Network tx_rate in units
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

func dataSourceSystemPerformanceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vKpi, okKpi := d.GetOk("kpi")
	vFunction, okFunction := d.GetOk("function")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SystemPerformanceApI")
		queryParams1 := dnacentersdkgo.SystemPerformanceApIQueryParams{}

		if okKpi {
			queryParams1.Kpi = vKpi.(string)
		}
		if okFunction {
			queryParams1.Function = vFunction.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}

		response1, restyResp1, err := client.HealthAndPerformance.SystemPerformanceApI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SystemPerformanceApI", err,
				"Failure at SystemPerformanceApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemPerformanceApIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemPerformanceApI response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemPerformanceApIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceApI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["version"] = item.Version
	respItem["kpis"] = flattenHealthAndPerformanceSystemPerformanceApIItemKpis(item.Kpis)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemPerformanceApIItemKpis(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu"] = flattenHealthAndPerformanceSystemPerformanceApIItemKpisCPU(item.CPU)
	respItem["memory"] = flattenHealthAndPerformanceSystemPerformanceApIItemKpisMemory(item.Memory)
	respItem["network_tx_rate"] = flattenHealthAndPerformanceSystemPerformanceApIItemKpisNetworktxRate(item.NetworktxRate)
	respItem["network_rx_rate"] = flattenHealthAndPerformanceSystemPerformanceApIItemKpisNetworkrxRate(item.NetworkrxRate)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceApIItemKpisCPU(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisCPU) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units
	respItem["utilization"] = item.Utilization

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceApIItemKpisMemory(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisMemory) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units
	respItem["utilization"] = item.Utilization

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceApIItemKpisNetworktxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworktxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units
	respItem["utilization"] = item.Utilization

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceApIItemKpisNetworkrxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworkrxRate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["units"] = item.Units
	respItem["utilization"] = item.Utilization

	return []map[string]interface{}{
		respItem,
	}

}
