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
