package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

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
				Description: `rf-profile-name query parameter. RF Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
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
							Description: `is Default Rf Profile
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_brown_field": &schema.Schema{
							Description: `Enable Brown Field
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_custom": &schema.Schema{
							Description: `Enable Custom
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

						"enable_radio_type_c": &schema.Schema{
							Description: `Enable Radio Type C (6GHz)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `RF Profile Name
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
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Rx Sop Threshold
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent Profile
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold V1
`,
										Type:     schema.TypeFloat,
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
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Min Power Level
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent Profile
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold V1
`,
										Type:     schema.TypeFloat,
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

						"radio_type_c_properties": &schema.Schema{
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
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Min Power Level
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent Profile
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold V1
`,
										Type:     schema.TypeFloat,
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
		log.Printf("[DEBUG] Selected method: RetrieveRfProfiles")
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

		vItem1 := flattenWirelessRetrieveRfProfilesItem(response1.Response)
		if err := d.Set("items", vItem1); err != nil {
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

func flattenWirelessRetrieveRfProfilesItem(items *[]dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
		respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
		respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
		respItem["channel_width"] = item.ChannelWidth
		respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
		respItem["enable_brown_field"] = boolPtrToString(item.EnableBrownField)
		respItem["radio_type_a_properties"] = flattenWirelessRetrieveRfProfilesItemRadioTypeAProperties(item.RadioTypeAProperties)
		respItem["radio_type_b_properties"] = flattenWirelessRetrieveRfProfilesItemRadioTypeBProperties(item.RadioTypeBProperties)
		respItem["radio_type_c_properties"] = flattenWirelessRetrieveRfProfilesItemRadioTypeCProperties(item.RadioTypeCProperties)
		respItem["enable_radio_type_c"] = boolPtrToString(item.EnableRadioTypeC)
		respItems = append(respItems, respItem)
	}

	return respItems
}

func flattenWirelessRetrieveRfProfilesItemRadioTypeAProperties(item *dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponseRadioTypeAProperties) []map[string]interface{} {
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

func flattenWirelessRetrieveRfProfilesItemRadioTypeBProperties(item *dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponseRadioTypeBProperties) []map[string]interface{} {
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

func flattenWirelessRetrieveRfProfilesItemRadioTypeCProperties(item *dnacentersdkgo.ResponseWirelessRetrieveRfProfilesResponseRadioTypeCProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["power_threshold_v1"] = item.PowerThresholdV1

	return []map[string]interface{}{
		respItem,
	}

}
