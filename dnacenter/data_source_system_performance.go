package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemPerformance() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- Retrieves the aggregated metrics (total, average or maximum) of cluster key performance indicators (KPIs), such as CPU
utilization, memory utilization or network rates recorded within a specified time period. The data will be available
from the past 24 hours.
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
		log.Printf("[DEBUG] Selected method: SystemPerformanceAPI")
		queryParams1 := dnacentersdkgo.SystemPerformanceAPIQueryParams{}

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

		response1, restyResp1, err := client.HealthAndPerformance.SystemPerformanceAPI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemPerformanceAPI", err,
				"Failure at SystemPerformanceAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemPerformanceAPIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemPerformanceAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemPerformanceAPIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["version"] = item.Version
	respItem["kpis"] = flattenHealthAndPerformanceSystemPerformanceAPIItemKpis(item.Kpis)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemPerformanceAPIItemKpis(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu"] = flattenHealthAndPerformanceSystemPerformanceAPIItemKpisCPU(item.CPU)
	respItem["memory"] = flattenHealthAndPerformanceSystemPerformanceAPIItemKpisMemory(item.Memory)
	respItem["network_tx_rate"] = flattenHealthAndPerformanceSystemPerformanceAPIItemKpisNetworktxRate(item.NetworktxRate)
	respItem["network_rx_rate"] = flattenHealthAndPerformanceSystemPerformanceAPIItemKpisNetworkrxRate(item.NetworkrxRate)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenHealthAndPerformanceSystemPerformanceAPIItemKpisCPU(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisCPU) []map[string]interface{} {
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

func flattenHealthAndPerformanceSystemPerformanceAPIItemKpisMemory(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisMemory) []map[string]interface{} {
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

func flattenHealthAndPerformanceSystemPerformanceAPIItemKpisNetworktxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworktxRate) []map[string]interface{} {
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

func flattenHealthAndPerformanceSystemPerformanceAPIItemKpisNetworkrxRate(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworkrxRate) []map[string]interface{} {
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
