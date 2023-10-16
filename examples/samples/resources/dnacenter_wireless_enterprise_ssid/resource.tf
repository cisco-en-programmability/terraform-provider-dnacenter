
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.23-beta"
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
    name           = "TestPersonal2"
    security_level = "WPA2_ENTERPRISE"
  }
}

output "dnacenter_wireless_enterprise_ssid_example" {
  value = dnacenter_wireless_enterprise_ssid.example
}
