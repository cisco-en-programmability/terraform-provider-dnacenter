package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceGetSensorTestResults() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Intent API to get SENSOR test result summary
`,

		ReadContext: dataSourceAssuranceGetSensorTestResultsRead,
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
				Description: `testFailureBy query parameter. Obtain failure statistics group by "area", "building", or "floor" (case insensitive)
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
										Description: `The error code
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"error_title": &schema.Schema{
										Description: `The error title
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"test_category": &schema.Schema{
										Description: `The test category
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"test_type": &schema.Schema{
										Description: `The test type
`,
										Type:     schema.TypeString,
										Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
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
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"total_test_count": &schema.Schema{
										Description: `Total test count
`,
										Type:     schema.TypeInt,
										Computed: true,
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

func dataSourceAssuranceGetSensorTestResultsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vTestFailureBy, okTestFailureBy := d.GetOk("test_failure_by")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SensorTestResults")
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
				"Failure when executing 2 SensorTestResults", err,
				"Failure at SensorTestResults, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessSensorTestResultsItem(response1.Response)
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
