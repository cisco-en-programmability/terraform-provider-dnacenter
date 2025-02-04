package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Security Advisories.

- Retrieves summary of advisories on the network.
`,

		ReadContext: dataSourceSecurityAdvisoriesSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cri_tic_al": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"hig_h": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"inf_orm_ati_ona_l": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"low": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"med_ium": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"na": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
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

func dataSourceSecurityAdvisoriesSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdvisoriesSummary")

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoriesSummary()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoriesSummary", err,
				"Failure at GetAdvisoriesSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSecurityAdvisoriesGetAdvisoriesSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoriesSummary response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItem(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["inf_orm_ati_ona_l"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemINFORMATIONAL(item.INFORMATIONAL)
	respItem["low"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemLOW(item.LOW)
	respItem["med_ium"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemMEDIUM(item.MEDIUM)
	respItem["hig_h"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemHIGH(item.HIGH)
	respItem["cri_tic_al"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemCRITICaL(item.CRITICaL)
	respItem["na"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryItemNA(item.NA)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemINFORMATIONAL(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseINFORMATIONAL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemLOW(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseLOW) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemMEDIUM(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseMEDIUM) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemHIGH(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseHIGH) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemCRITICaL(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseCRITICaL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryItemNA(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseNA) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}
