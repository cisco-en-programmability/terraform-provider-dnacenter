
resource "dnacenter_wireless_enterprise_ssid" "example" {
  provider = dnacenter
  parameters {

    basic_service_set_client_idle_timeout = 1
    client_exclusion_timeout              = 1
    enable_basic_service_set_max_idle     = "false"
    enable_broadcast_ssi_d                = "false"
    enable_client_exclusion               = "false"
    enable_directed_multicast_service     = "false"
    enable_fast_lane                      = "false"
    enable_mac_filtering                  = "false"
    enable_neighbor_list                  = "false"
    enable_session_time_out               = "false"
    fast_transition                       = "string"
    mfp_client_protection                 = "string"
    name                                  = "string"
    passphrase                            = "string"
    radio_policy                          = "string"
    security_level                        = "string"
    session_time_out                      = 1
    ssid_name                             = "string"
    traffic_type                          = "string"
  }
}

output "dnacenter_wireless_enterprise_ssid_example" {
  value = dnacenter_wireless_enterprise_ssid.example
}