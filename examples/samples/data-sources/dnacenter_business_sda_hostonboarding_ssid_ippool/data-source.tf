terraform {
  required_providers {
    dnacenter = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
  vlan_name           = "string"
}

output "dnacenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = data.dnacenter_business_sda_hostonboarding_ssid_ippool.example.item
}
