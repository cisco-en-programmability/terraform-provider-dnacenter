terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_sda_fabric_sites" "example" {
  provider = dnacenter

  parameters {
    payload {
      authentication_profile_name = "string"
      id                          = "string"
      is_pub_sub_enabled          = "false"
      site_id                     = "string"
    }
  }
}

output "dnacenter_sda_fabric_sites_example" {
  value = dnacenter_sda_fabric_sites.example
}
