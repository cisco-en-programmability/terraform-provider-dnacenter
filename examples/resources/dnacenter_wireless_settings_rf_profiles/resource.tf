
resource "dnacenter_wireless_settings_rf_profiles" "example" {
  provider = dnacenter

  parameters {

    default_rf_profile      = "false"
    enable_radio_type6_g_hz = "false"
    enable_radio_type_a     = "false"
    enable_radio_type_b     = "false"
    id                      = "string"
    radio_type6_g_hz_properties {

      data_rates                    = "string"
      enable_standard_power_service = "false"
      mandatory_data_rates          = "string"
      max_dbs_width                 = 1
      max_power_level               = 1
      min_dbs_width                 = 1
      min_power_level               = 1
      multi_bssid_properties {

        dot11ax_parameters {

          mu_mimo_down_link = "false"
          mu_mimo_up_link   = "false"
          ofdma_down_link   = "false"
          ofdma_up_link     = "false"
        }
        dot11be_parameters {

          mu_mimo_down_link = "false"
          mu_mimo_up_link   = "false"
          ofdma_down_link   = "false"
          ofdma_multi_ru    = "false"
          ofdma_up_link     = "false"
        }
        target_wake_time      = "false"
        twt_broadcast_support = "false"
      }
      parent_profile     = "string"
      power_threshold_v1 = 1
      preamble_puncture  = "false"
      radio_channels     = "string"
      rx_sop_threshold   = "string"
    }
    radio_type_a_properties {

      channel_width        = "string"
      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1
      min_power_level      = 1
      parent_profile       = "string"
      power_threshold_v1   = 1
      preamble_puncture    = "false"
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    radio_type_b_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1
      min_power_level      = 1
      parent_profile       = "string"
      power_threshold_v1   = 1
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    rf_profile_name = "string"
  }
}

output "dnacenter_wireless_settings_rf_profiles_example" {
  value = dnacenter_wireless_settings_rf_profiles.example
}