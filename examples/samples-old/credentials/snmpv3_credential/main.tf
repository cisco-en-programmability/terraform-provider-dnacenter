terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
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
