terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_sda_fabric" "example" {
  provider = dnacenter
  parameters {
    fabric_name = "b"
  }
}

output "dnacenter_sda_fabric_example" {
  value = dnacenter_sda_fabric.example
}