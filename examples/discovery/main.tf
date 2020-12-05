
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

resource "dna_discovery" "response" {
  provider = dnacenter
  item {
    cdp_level                 = 16
    discovery_type            = "CDP"
    global_credential_id_list = ["90acbab8-03d5-4726-9c19-e1e51a40b3cd", "f979d842-f6fd-456a-8137-2cb5113cd2e8"]
    ip_address_list           = "10.10.22.22"
    name                      = "start_discovery_test"
    netconf_port              = "65535"
    protocol_order            = "ssh"
  }
}
output "dna_discovery_response" {
  value = dna_discovery.response
}

