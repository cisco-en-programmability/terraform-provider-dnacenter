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

resource "dna_snmpv2_write_community_credential" "response" {
  provider = dnacenter
  item {
    description     = "SNMP WO test 1"
    write_community = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
  }
}
output "dna_snmpv2_write_community_credential_response" {
  sensitive = true
  value = dna_snmpv2_write_community_credential.response
}
