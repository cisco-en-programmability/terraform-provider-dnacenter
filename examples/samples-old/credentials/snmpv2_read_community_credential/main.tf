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

resource "dna_snmpv2_read_community_credential" "response" {
  provider = dnacenter
  item {
    description     = "SNMP RO test 1"
    read_community  = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    id              = "e566053f-d5cd-4a81-841e-cb91a712af20"
  }
}
output "dna_snmpv2_read_community_credential_response" {
  value = dna_snmpv2_read_community_credential.response
}
