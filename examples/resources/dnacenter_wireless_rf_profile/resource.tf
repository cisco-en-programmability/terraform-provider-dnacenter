
resource "dnacenter_wireless_rf_profile" "example" {
  provider = dnacenter

  parameters {

    channel_width       = "string"
    default_rf_profile  = "false"
    enable_brown_field  = "false"
    enable_custom       = "false"
    enable_radio_type_a = "false"
    enable_radio_type_b = "false"
    enable_radio_type_c = "false"
    name                = "string"
    radio_type_a_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    radio_type_b_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    radio_type_c_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    rf_profile_name = "string"
  }
}

output "dnacenter_wireless_rf_profile_example" {
  value = dnacenter_wireless_rf_profile.example
}