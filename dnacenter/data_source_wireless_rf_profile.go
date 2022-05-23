package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

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

						"a_radio_channels": &schema.Schema{
							Description: `A Radio Channels`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"b_radio_channels": &schema.Schema{
							Description: `B Radio Channels`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"c_radio_channels": &schema.Schema{
							Description: `C Radio Channels`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"channel_width": &schema.Schema{
							Description: `Channel Width`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data_rates_a": &schema.Schema{
							Description: `Data Rates A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data_rates_b": &schema.Schema{
							Description: `Data Rates B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data_rates_c": &schema.Schema{
							Description: `Data Rates C`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_rf_profile": &schema.Schema{
							Description: `Default Rf Profile`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_a_radio_type": &schema.Schema{
							Description: `Enable ARadio Type`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_b_radio_type": &schema.Schema{
							Description: `Enable BRadio Type`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_brown_field": &schema.Schema{
							Description: `Enable Brown Field`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_c_radio_type": &schema.Schema{
							Description: `Enable CRadio Type`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_custom": &schema.Schema{
							Description: `Enable Custom`,

							Type:     schema.TypeString,
							Computed: true,
						},

						"mandatory_data_rates_a": &schema.Schema{
							Description: `Mandatory Data Rates A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mandatory_data_rates_b": &schema.Schema{
							Description: `Mandatory Data Rates B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mandatory_data_rates_c": &schema.Schema{
							Description: `Mandatory Data Rates C`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"max_power_level_a": &schema.Schema{
							Description: `Max Power Level A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"max_power_level_b": &schema.Schema{
							Description: `Max Power Level B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"min_power_level_a": &schema.Schema{
							Description: `Min Power Level A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"min_power_level_b": &schema.Schema{
							Description: `Min Power Level B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"min_power_level_c": &schema.Schema{
							Description: `Min Power Level C`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_profile_a": &schema.Schema{
							Description: `Parent Profile A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_profile_b": &schema.Schema{
							Description: `Parent Profile B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"power_threshold_v1_a": &schema.Schema{
							Description: `Power Threshold V1 A`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"power_threshold_v1_b": &schema.Schema{
							Description: `Power Threshold V1 B`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"power_threshold_v1_c": &schema.Schema{
							Description: `Power Threshold V1 C`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"rx_sop_threshold_a": &schema.Schema{
							Description: `Rx Sop Threshold A`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rx_sop_threshold_b": &schema.Schema{
							Description: `Rx Sop Threshold B`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rx_sop_threshold_c": &schema.Schema{
							Description: `Rx Sop Threshold C`,
							Type:        schema.TypeString,
							Computed:    true,
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
		respItem["parent_profile_a"] = item.ParentProfileA
		respItem["parent_profile_b"] = item.ParentProfileB
		respItem["enable_a_radio_type"] = boolPtrToString(item.EnableARadioType)
		respItem["enable_b_radio_type"] = boolPtrToString(item.EnableBRadioType)
		respItem["enable_c_radio_type"] = boolPtrToString(item.EnableCRadioType)
		respItem["channel_width"] = item.ChannelWidth
		respItem["a_radio_channels"] = item.ARadioChannels
		respItem["b_radio_channels"] = item.BRadioChannels
		respItem["c_radio_channels"] = item.CRadioChannels
		respItem["data_rates_a"] = item.DataRatesA
		respItem["data_rates_b"] = item.DataRatesB
		respItem["data_rates_c"] = item.DataRatesC
		respItem["mandatory_data_rates_a"] = item.MandatoryDataRatesA
		respItem["mandatory_data_rates_b"] = item.MandatoryDataRatesB
		respItem["mandatory_data_rates_c"] = item.MandatoryDataRatesC
		respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
		respItem["min_power_level_a"] = item.MinPowerLevelA
		respItem["min_power_level_b"] = item.MinPowerLevelB
		respItem["min_power_level_c"] = item.MinPowerLevelC
		respItem["max_power_level_a"] = item.MaxPowerLevelA
		respItem["max_power_level_b"] = item.MaxPowerLevelB
		respItem["power_threshold_v1_a"] = item.PowerThresholdV1A
		respItem["power_threshold_v1_b"] = item.PowerThresholdV1B
		respItem["power_threshold_v1_c"] = item.PowerThresholdV1C
		respItem["rx_sop_threshold_a"] = item.RxSopThresholdA
		respItem["rx_sop_threshold_b"] = item.RxSopThresholdB
		respItem["rx_sop_threshold_c"] = item.RxSopThresholdC
		respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
		respItem["enable_brown_field"] = boolPtrToString(item.EnableBrownField)
		respItems = append(respItems, respItem)
	}
	return respItems
}
