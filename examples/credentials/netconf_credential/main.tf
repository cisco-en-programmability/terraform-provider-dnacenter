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

resource "dna_netconf_credential" "response" {
  provider = dnacenter
  item {
    netconf_port = 23
    description  = "netconf 23"
  }
}
output "dna_netconf_credential_response" {
  value = dna_netconf_credential.response
}
