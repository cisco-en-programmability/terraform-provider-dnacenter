
resource "dnacenter_wireless_accespoint_configuration" "example" {
  provider = dnacenter
  parameters {

    admin_status = "false"
    ap_list {

      ap_name     = "string"
      ap_name_new = "string"
      mac_address = "string"
    }
    ap_mode                        = 1
    configure_admin_status         = "false"
    configure_ap_mode              = "false"
    configure_failover_priority    = "false"
    configure_hacontroller         = "false"
    configure_led_brightness_level = "false"
    configure_led_status           = "false"
    configure_location             = "false"
    failover_priority              = 1
    is_assigned_site_as_location   = "false"
    led_brightness_level           = 1
    led_status                     = "false"
    location                       = "string"
    primary_controller_name        = "string"
    primary_ip_address {

      address = "string"
    }
    radio_configurations {

      admin_status                    = "false"
      antenna_cable_name              = "string"
      antenna_gain                    = 1
      antenna_pattern_name            = "string"
      cable_loss                      = 1.0
      channel_assignment_mode         = 1
      channel_number                  = 1
      channel_width                   = 1
      clean_air_si                    = 1
      configure_admin_status          = "false"
      configure_antenna_cable         = "false"
      configure_antenna_pattern_name  = "false"
      configure_channel               = "false"
      configure_channel_width         = "false"
      configure_clean_air_si          = "false"
      configure_power                 = "false"
      configure_radio_role_assignment = "false"
      power_assignment_mode           = 1
      powerlevel                      = 1
      radio_band                      = "string"
      radio_role_assignment           = "string"
      radio_type                      = 1
    }
    secondary_controller_name = "string"
    secondary_ip_address {

      address = "string"
    }
    tertiary_controller_name = "string"
    tertiary_ip_address {

      address = "string"
    }
  }
}

output "dnacenter_wireless_accespoint_configuration_example" {
  value = dnacenter_wireless_accespoint_configuration.example
}