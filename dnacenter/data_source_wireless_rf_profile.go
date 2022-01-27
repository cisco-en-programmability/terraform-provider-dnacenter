package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessRfProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieve all RF profiles
`,

		ReadContext: dataSourceWirelessRfProfileRead,
		Schema: map[string]*schema.Schema{
			"rf_profile_name": &schema.Schema{
				Description: `rf-profile-name query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"channel_width": &schema.Schema{
							Description: `Channel Width
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_rf_profile": &schema.Schema{
							Description: `is default radio profile
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_brown_field": &schema.Schema{
							Description: `is brownfield enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_custom": &schema.Schema{
							Description: `is Custom Enable
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_a": &schema.Schema{
							Description: `Enable Radio Type A
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_b": &schema.Schema{
							Description: `Enable Radio Type B
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `radio profile name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Max Power Level
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Min Power Level
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold V1
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `Radio Channels
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Max Power Level
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Min Power Level
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold V1
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `Radio Channels
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold
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
	}
}

func dataSourceWirelessRfProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vRfProfileName, okRfProfileName := d.GetOk("rf_profile_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RetrieveRfProfiles")
		queryParams1 := dnacentersdkgo.RetrieveRfProfilesQueryParams{}

		if okRfProfileName {
			queryParams1.RfProfileName = vRfProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.RetrieveRfProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetrieveRfProfiles", err,
				"Failure at RetrieveRfProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessRetrieveRfProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveRfProfiles response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveRfProfilesItems(items *[]dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
		respItem["channel_width"] = item.ChannelWidth
		respItem["enable_brown_field"] = boolPtrToString(item.EnableBrownField)
		respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
		respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
		respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
		respItem["radio_type_a_properties"] = flattenWirelessRetrieveRfProfilesItemsRadioTypeAProperties(item.RadioTypeAProperties)
		respItem["radio_type_b_properties"] = flattenWirelessRetrieveRfProfilesItemsRadioTypeBProperties(item.RadioTypeBProperties)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessRetrieveRfProfilesItemsRadioTypeAProperties(item *dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponseRadioTypeAProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessRetrieveRfProfilesItemsRadioTypeBProperties(item *dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponseRadioTypeBProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel

	return []map[string]interface{}{
		respItem,
	}

}
