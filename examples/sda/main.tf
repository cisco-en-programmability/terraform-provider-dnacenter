terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_sda_fabric" "response" {
  provider    = dnacenter
  fabric_name = "MyFabricName2"
}
output "dna_sda_fabric_response" {
  value = dna_sda_fabric.response
}
