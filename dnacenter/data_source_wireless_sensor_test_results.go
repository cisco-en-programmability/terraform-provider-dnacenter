package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSensorTestResults() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Intent API to get SENSOR test result summary
`,

		ReadContext: dataSourceWirelessSensorTestResultsRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. The epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Assurance site UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. The epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"test_failure_by": &schema.Schema{
				Description: `testFailureBy query parameter. Obtain failure statistics group by "area", "building", or "floor"
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failure_stats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"error_code": &schema.Schema{
										Description: `Error Code`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"error_title": &schema.Schema{
										Description: `Error Title`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"test_category": &schema.Schema{
										Description: `Test Category`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"test_type": &schema.Schema{
										Description: `Test Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_con_nec_tiv_ity": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fil_etr_ans_fer": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
														},
													},
												},

												"hos_tre_ach_abi_lit_y": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"web_ser_ver": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"ema_il": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"mai_lse_rve_r": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"net_wor_kse_rvi_ces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dns": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"onb_oar_din_g": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ass_oc": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"aut_h": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"dhc_p": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"per_for_man_ce": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ips_las_end_er": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"rfa_sse_ssm_ent": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dat_ara_te": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"snr": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Fail Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},

															"pass_count": &schema.Schema{
																Description: `Pass Count`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},

									"total_test_count": &schema.Schema{
										Description: `Total Test Count`,
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

func dataSourceWirelessSensorTestResultsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vTestFailureBy, okTestFailureBy := d.GetOk("test_failure_by")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SensorTestResults")
		queryParams1 := dnacentersdkgo.SensorTestResultsQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okTestFailureBy {
			queryParams1.TestFailureBy = vTestFailureBy.(string)
		}

		response1, restyResp1, err := client.Wireless.SensorTestResults(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SensorTestResults", err,
				"Failure at SensorTestResults, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessSensorTestResultsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SensorTestResults response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessSensorTestResultsItem(item *dnacentersdkgo.ResponseWirelessSensorTestResults) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["summary"] = flattenWirelessSensorTestResultsItemSummary(item.Summary)
	respItem["failure_stats"] = flattenWirelessSensorTestResultsItemFailureStats(item.FailureStats)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessSensorTestResultsItemSummary(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_test_count"] = item.TotalTestCount
	respItem["onb_oar_din_g"] = flattenWirelessSensorTestResultsItemSummaryOnBoarding(item.OnBoarding)
	respItem["per_for_man_ce"] = flattenWirelessSensorTestResultsItemSummaryPERfORMAncE(item.PERfORMAncE)
	respItem["net_wor_kse_rvi_ces"] = flattenWirelessSensorTestResultsItemSummaryNETWORKSERVICES(item.NETWORKSERVICES)
	respItem["app_con_nec_tiv_ity"] = flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITY(item.ApPCONNECTIVITY)
	respItem["rfa_sse_ssm_ent"] = flattenWirelessSensorTestResultsItemSummaryRfASSESSMENT(item.RfASSESSMENT)
	respItem["ema_il"] = flattenWirelessSensorTestResultsItemSummaryEmail(item.Email)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryOnBoarding(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryOnBoarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aut_h"] = flattenWirelessSensorTestResultsItemSummaryOnBoardingAuth(item.Auth)
	respItem["dhc_p"] = flattenWirelessSensorTestResultsItemSummaryOnBoardingDHCP(item.DHCP)
	respItem["ass_oc"] = flattenWirelessSensorTestResultsItemSummaryOnBoardingAssoc(item.Assoc)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryOnBoardingAuth(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryOnBoardingAuth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryOnBoardingDHCP(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryOnBoardingDHCP) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryOnBoardingAssoc(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryOnBoardingAssoc) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryPERfORMAncE(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryPERfORMAncE) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ips_las_end_er"] = flattenWirelessSensorTestResultsItemSummaryPERfORMAncEIPSLASENDER(item.IPSLASENDER)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryPERfORMAncEIPSLASENDER(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryPERfORMAncEIPSLASENDER) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryNETWORKSERVICES(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryNETWORKSERVICES) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dns"] = flattenWirelessSensorTestResultsItemSummaryNETWORKSERVICESDNS(item.DNS)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryNETWORKSERVICESDNS(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryNETWORKSERVICESDNS) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITY(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITY) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hos_tre_ach_abi_lit_y"] = flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYHOSTREACHABILITY(item.HOSTREACHABILITY)
	respItem["web_ser_ver"] = flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYWebServer(item.WebServer)
	respItem["fil_etr_ans_fer"] = flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYFileTransfer(item.FileTransfer)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYHOSTREACHABILITY(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYHOSTREACHABILITY) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYWebServer(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYWebServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryApPCONNECTIVITYFileTransfer(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYFileTransfer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryRfASSESSMENT(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryRfASSESSMENT) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dat_ara_te"] = flattenWirelessSensorTestResultsItemSummaryRfASSESSMENTDATARATE(item.DATARATE)
	respItem["snr"] = flattenWirelessSensorTestResultsItemSummaryRfASSESSMENTSNR(item.SNR)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryRfASSESSMENTDATARATE(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryRfASSESSMENTDATARATE) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryRfASSESSMENTSNR(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryRfASSESSMENTSNR) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryEmail(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryEmail) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mai_lse_rve_r"] = flattenWirelessSensorTestResultsItemSummaryEmailMailServer(item.MailServer)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemSummaryEmailMailServer(item *dnacentersdkgo.ResponseWirelessSensorTestResultsSummaryEmailMailServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsItemFailureStats(items *[]dnacentersdkgo.ResponseWirelessSensorTestResultsFailureStats) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["error_code"] = item.ErrorCode
		respItem["error_title"] = item.ErrorTitle
		respItem["test_type"] = item.TestType
		respItem["test_category"] = item.TestCategory
		respItems = append(respItems, respItem)
	}
	return respItems
}
