terraform {
  required_providers {
    dnacenter = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_assign_device_to_site" "example" {
  provider = dnacenter
  parameters {

    device {

      ip = "10.122.1.1"
    }
    site_id = "1ba55132-0fb8-4986-a31e-674d30e8b8ee"
  }
}

output "dnacenter_assign_device_to_site_example" {
  value = dnacenter_assign_device_to_site.example
}