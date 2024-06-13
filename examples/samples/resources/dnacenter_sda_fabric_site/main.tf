terraform {
  required_providers {
    dnacenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_sda_fabric_site" "example" {
  provider = dnacenter
  parameters {
    fabric_name         = "Default LAN Fabric 2"
    site_name_hierarchy = "Global/San Francisco"
  }
}

output "dnacenter_sda_fabric_site_example" {
  value = dnacenter_sda_fabric_site.example
}
