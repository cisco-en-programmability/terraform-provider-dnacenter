
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

resource "dna_pnp_device" "response" {
  provider = dnacenter
  item {
    device_info {
      serial_number = "FOCTEST2"
    }
  }
}
output "dna_pnp_device_response" {
  value = dna_pnp_device.response
}

