
resource "dnacenter_sites_wireless_settings_ssids" "example" {
  provider = dnacenter

  parameters {

    aaa_override                                       = "false"
    acct_servers                                       = ["string"]
    acl_name                                           = "string"
    auth_server                                        = "string"
    auth_servers                                       = ["string"]
    auth_type                                          = "string"
    basic_service_set_client_idle_timeout              = 1
    basic_service_set_max_idle_enable                  = "false"
    cckm_tsf_tolerance                                 = 1
    client_exclusion_enable                            = "false"
    client_exclusion_timeout                           = 1
    client_rate_limit                                  = 1
    coverage_hole_detection_enable                     = "false"
    directed_multicast_service_enable                  = "false"
    egress_qos                                         = "string"
    external_auth_ip_address                           = "string"
    fast_transition                                    = "string"
    fast_transition_over_the_distributed_system_enable = "false"
    ghz24_policy                                       = "string"
    ghz6_policy_client_steering                        = "false"
    id                                                 = "string"
    ingress_qos                                        = "string"
    is_ap_beacon_protection_enabled                    = "false"
    is_auth_key8021x                                   = "false"
    is_auth_key8021x_plus_ft                           = "false"
    is_auth_key8021x_sha256                            = "false"
    is_auth_key_easy_psk                               = "false"
    is_auth_key_owe                                    = "false"
    is_auth_key_psk                                    = "false"
    is_auth_key_psk_plus_ft                            = "false"
    is_auth_key_psk_sha256                             = "false"
    is_auth_key_sae                                    = "false"
    is_auth_key_sae_ext                                = "false"
    is_auth_key_sae_ext_plus_ft                        = "false"
    is_auth_key_sae_plus_ft                            = "false"
    is_auth_key_suite_b1921x                           = "false"
    is_auth_key_suite_b1x                              = "false"
    is_broadcast_ssi_d                                 = "false"
    is_cckm_enabled                                    = "false"
    is_enabled                                         = "false"
    is_fast_lane_enabled                               = "false"
    is_hex                                             = "false"
    is_mac_filtering_enabled                           = "false"
    is_posturing_enabled                               = "false"
    is_random_mac_filter_enabled                       = "false"
    l3_auth_type                                       = "string"
    management_frame_protection_clientprotection       = "string"
    multi_psk_settings {

      passphrase      = "string"
      passphrase_type = "string"
      priority        = 1
    }
    nas_options                = ["string"]
    neighbor_list_enable       = "false"
    open_ssid                  = "string"
    passphrase                 = "string"
    profile_name               = "string"
    protected_management_frame = "string"
    rsn_cipher_suite_ccmp128   = "false"
    rsn_cipher_suite_ccmp256   = "false"
    rsn_cipher_suite_gcmp128   = "false"
    rsn_cipher_suite_gcmp256   = "false"
    session_time_out           = 1
    session_time_out_enable    = "false"
    site_id                    = "string"
    sleeping_client_enable     = "false"
    sleeping_client_timeout    = 1
    ssid                       = "string"
    ssid_radio_type            = "string"
    web_passthrough            = "false"
    wlan_band_select_enable    = "false"
    wlan_type                  = "string"
  }
}

output "dnacenter_sites_wireless_settings_ssids_example" {
  value = dnacenter_sites_wireless_settings_ssids.example
}