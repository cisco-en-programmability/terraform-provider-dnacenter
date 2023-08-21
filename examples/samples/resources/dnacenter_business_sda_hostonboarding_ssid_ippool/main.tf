
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.11-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider = dnacenter
  parameters {

    # scalable_group_name = "string"
    site_name_hierarchy = "Global/San Francisco/Default LAN Fabric 2"
    ssid_names          = ["lab"]
    vlan_name           = "wlan1"
  }
}

output "dnacenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = dnacenter_business_sda_hostonboarding_ssid_ippool.example
}
