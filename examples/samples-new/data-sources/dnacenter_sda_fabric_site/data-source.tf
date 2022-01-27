terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}


data "dnacenter_sda_fabric_site" "example" {
  provider            = dnacenter
  site_name_hierarchy = "Global/San Francisco"
}

output "dnacenter_sda_fabric_site_example" {
  value = data.dnacenter_sda_fabric_site.example.item
}
