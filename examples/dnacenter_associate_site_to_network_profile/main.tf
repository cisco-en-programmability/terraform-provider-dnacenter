terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_associate_site_to_network_profile" "associate" {
  provider = dnacenter
  network_profile_id=1
  site_id=1
}
output "dnacenter_associate_site_to_network_profile_associate" {
  value = data.dnacenter_associate_site_to_network_profile.associate
}


