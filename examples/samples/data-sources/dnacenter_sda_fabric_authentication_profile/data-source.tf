terraform {
  required_providers {
    dnacenter = {
      version = "1.1.12-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_sda_fabric_authentication_profile" "example" {
  provider = dnacenter
  # authenticate_template_name = "Test"
  site_name_hierarchy = "Global/Pennsylvania"
}

output "dnacenter_sda_fabric_authentication_profile_example" {
  value = data.dnacenter_sda_fabric_authentication_profile.example.item
}
