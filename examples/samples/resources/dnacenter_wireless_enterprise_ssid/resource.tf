
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}


resource "dnacenter_wireless_enterprise_ssid" "example" {
  provider = dnacenter
  parameters {
    name                       = "TestPersonal2"
    ssid_name                  = "TestPersonal2"
    security_level             = "WPA2_ENTERPRISE"
    traffic_type               = "voicedata"
    radio_policy               = "Triple band operation (2.6GHz, 5GHz and 6GHz)"
    fast_transition            = "Adaptive"
    mfp_client_protection      = "Optional"
    protected_management_frame = "Optional"
    # multi_psk_settings {
    #   passphraseType = "ASCII"
    # }

  }
}



output "dnacenter_wireless_enterprise_ssid_example" {
  value = dnacenter_wireless_enterprise_ssid.example
}
