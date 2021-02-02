
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.0.3"]
      source   = "hashicorp.com/edu/dnacenter"
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

