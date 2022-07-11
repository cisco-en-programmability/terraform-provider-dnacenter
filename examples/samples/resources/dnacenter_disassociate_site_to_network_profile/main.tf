terraform {
  required_providers {
    dnacenter = {
      version = "0.3.1"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_disassociate_site_to_network_profile" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    network_profile_id = "string"
    site_id            = "string"
  }
}