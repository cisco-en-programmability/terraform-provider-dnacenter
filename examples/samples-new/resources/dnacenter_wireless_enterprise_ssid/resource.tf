
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_wireless_enterprise_ssid" "example" {
  provider = dnacenter
  parameters {

      basic_service_set_client_idle_timeout= 0
      client_exclusion_timeout= 0
      enable_basic_service_set_max_idle= "true"
      enable_broadcast_ssi_d= "true"
      enable_client_exclusion= "true"
      enable_directed_multicast_service= "true"
      enable_fast_lane= "true"
      enable_mac_filtering= "true"
      enable_neighbor_list= "true"
      enable_session_time_out= "true"
      fast_transition= "Adaptive"
      mfp_client_protection= "Optional"
      name= "Test"
      passphrase= ""
      radio_policy= "Dual band operation (2.4GHz and 5GHz)"
      security_level= "WPA2_ENTERPRISE"
      session_time_out= 0
      traffic_type= "voicedata"

  }
}

output "dnacenter_wireless_enterprise_ssid_example" {
  value = dnacenter_wireless_enterprise_ssid.example
}