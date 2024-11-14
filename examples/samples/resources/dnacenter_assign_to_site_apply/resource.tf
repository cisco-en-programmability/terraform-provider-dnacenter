terraform {
  required_providers {
    dnacenter = {
      version = "1.1.33-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_assign_to_site_apply" "example" {
  provider = dnacenter
  parameters {

    device_ids = ["string"]
    site_id    = "string"
  }
}

output "dnacenter_assign_to_site_apply_example" {
  value = dnacenter_assign_to_site_apply.example
}