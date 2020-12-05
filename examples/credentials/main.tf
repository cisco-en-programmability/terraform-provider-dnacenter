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

resource "dna_cli_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "123456"
    credential_type = "APP"
    id              = "34091dbf-8d55-48b8-aec4-5a572c265370"
  }
}
output "dna_cli_credential_response" {
  value = dna_cli_credential.response
}

resource "dna_http_read_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 23
  }
}
output "dna_http_read_credential_response" {
  value = dna_http_read_credential.response
}

resource "dna_http_write_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 24
    id              = "5abaa9c9-4470-46c4-90d6-107594164845"
  }
}
output "dna_http_write_credential_response" {
  value = dna_http_write_credential.response
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

resource "dna_snmpv2_write_community_credential" "response" {
  provider = dnacenter
  item {
    description     = "SNMP WO test 1"
    read_community  = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
  }
}
output "dna_snmpv2_write_community_credential_response" {
  value = dna_snmpv2_write_community_credential.response
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
