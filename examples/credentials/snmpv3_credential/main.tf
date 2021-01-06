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

resource "dna_snmpv3_credential" "response" {
  provider = dnacenter
  item {
    snmp_mode = "NOAUTHNOPRIV"
    username  = "user3"
  }
}
output "dna_snmpv3_credential_response" {
  value = dna_snmpv3_credential.response
}
