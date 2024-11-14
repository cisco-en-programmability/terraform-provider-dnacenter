
resource "dnacenter_wireless_enterprise_ssid" "example" {
  provider = dnacenter

  parameters {

    aaa_override                          = "false"
    auth_key_mgmt                         = ["string"]
    basic_service_set_client_idle_timeout = 1
    client_exclusion_timeout              = 1
    client_rate_limit                     = 1.0
    coverage_hole_detection_enable        = "false"
    enable_basic_service_set_max_idle     = "false"
    enable_broadcast_ssi_d                = "false"
    enable_client_exclusion               = "false"
    enable_directed_multicast_service     = "false"
    enable_fast_lane                      = "false"
    enable_mac_filtering                  = "false"
    enable_neighbor_list                  = "false"
    enable_session_time_out               = "false"
    fast_transition                       = "string"
    ghz24_policy                          = "string"
    ghz6_policy_client_steering           = "false"
    mfp_client_protection                 = "string"
    multi_psk_settings {

      passphrase      = "string"
      passphrase_type = "string"
      priority        = 1
    }
    name                       = "string"
    nas_options                = ["string"]
    passphrase                 = "string"
    policy_profile_name        = "string"
    profile_name               = "string"
    protected_management_frame = "string"
    radio_policy               = "string"
    rsn_cipher_suite_ccmp256   = "false"
    rsn_cipher_suite_gcmp128   = "false"
    rsn_cipher_suite_gcmp256   = "false"
    security_level             = "string"
    session_time_out           = 1
    ssid_name                  = "string"
    traffic_type               = "string"
  }
}

output "dnacenter_wireless_enterprise_ssid_example" {
  value = dnacenter_wireless_enterprise_ssid.example
}