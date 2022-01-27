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

data "dnacenter_sda_fabric_authentication_profile" "example" {
  provider                   = dnacenter
  authenticate_template_name = "string"
  site_name_hierarchy        = "string"
}

output "dnacenter_sda_fabric_authentication_profile_example" {
  value = data.dnacenter_sda_fabric_authentication_profile.example.item
}
