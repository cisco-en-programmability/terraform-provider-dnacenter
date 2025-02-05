
resource "dnacenter_wireless_settings_ap_profiles_id" "example" {
  provider = dnacenter

  parameters {

    ap_power_profile_name  = "string"
    ap_profile_name        = "string"
    awips_enabled          = "false"
    awips_forensic_enabled = "false"
    calendar_power_profiles {

      duration {

        scheduler_date       = "string"
        scheduler_day        = "string"
        scheduler_end_time   = "string"
        scheduler_start_time = "string"
      }
      power_profile_name = "string"
      scheduler_type     = "string"
    }
    client_limit = 1
    country_code = "string"
    description  = "string"
    id           = "string"
    management_setting {

      auth_type                  = "string"
      cdp_state                  = "false"
      dot1x_password             = "string"
      dot1x_username             = "string"
      management_enable_password = "string"
      management_password        = "string"
      management_user_name       = "string"
      ssh_enabled                = "false"
      telnet_enabled             = "false"
    }
    mesh_enabled = "false"
    mesh_setting {

      backhaul_client_access    = "false"
      bridge_group_name         = "string"
      ghz24_backhaul_data_rates = "string"
      ghz5_backhaul_data_rates  = "string"
      range                     = 1
      rap_downlink_backhaul     = "string"
    }
    pmf_denial_enabled    = "false"
    remote_worker_enabled = "false"
    rogue_detection_setting {

      rogue_detection                    = "false"
      rogue_detection_min_rssi           = 1
      rogue_detection_report_interval    = 1
      rogue_detection_transient_interval = 1
    }
    time_zone                = "string"
    time_zone_offset_hour    = 1
    time_zone_offset_minutes = 1
  }
}

output "dnacenter_wireless_settings_ap_profiles_id_example" {
  value = dnacenter_wireless_settings_ap_profiles_id.example
}
