terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_wireless_rf_profile" "example" {
  provider        = dnacenter
  rf_profile_name = "TYPICAL"
}

output "dnacenter_wireless_rf_profile_example" {
  value = data.dnacenter_wireless_rf_profile.example.items
}
