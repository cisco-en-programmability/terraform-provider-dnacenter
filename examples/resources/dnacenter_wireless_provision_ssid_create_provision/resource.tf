
resource "dnacenter_wireless_provision_ssid_create_provision" "example" {
  provider = dnacenter
  parameters {

    enable_fabric = "false"
    flex_connect {

      enable_flex_connect = "false"
      local_to_vlan       = 1
    }
    managed_aplocations = ["string"]
    ssid_details {

      auth_key_mgmt               = ["string"]
      enable_broadcast_ssi_d      = "false"
      enable_fast_lane            = "false"
      enable_mac_filtering        = "false"
      fast_transition             = "string"
      ghz24_policy                = "string"
      ghz6_policy_client_steering = "false"
      name                        = "string"
      passphrase                  = "string"
      radio_policy                = "string"
      rsn_cipher_suite_ccmp256    = "false"
      rsn_cipher_suite_gcmp128    = "false"
      rsn_cipher_suite_gcmp256    = "false"
      security_level              = "string"
      traffic_type                = "string"
      web_auth_url                = "string"
    }
    ssid_type = "string"
  }
}

output "dnacenter_wireless_provision_ssid_create_provision_example" {
  value = dnacenter_wireless_provision_ssid_create_provision.example
}