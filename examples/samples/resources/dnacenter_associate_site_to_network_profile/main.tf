
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.9-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_associate_site_to_network_profile" "example" {
  provider = dnacenter

  parameters {
    network_profile_id = "string"
    site_id            = "string"
  }
}