terraform {
  required_providers {
    dnacenter = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_reserve_ip_subpool" "example" {
  provider = dnacenter
  parameters {
    ipv4_prefix      = "false"
    ipv4_global_pool = "10.43.192.0"

    //id                 = "string"
    # ipv4_dhcp_servers  = []
    # ipv4_dns_servers   = []
    ipv4_gate_way      = "10.43.192.1"
    ipv4_prefix_length = 1
    ipv4_subnet        = "17"
    # ipv4_total_host    = 2
    # ipv6_address_space = "false"
    # ipv6_dhcp_servers  = []
    # ipv6_dns_servers   = []
    # ipv6_gate_way      = "string2"
    # ipv6_global_pool   = "string"
    # ipv6_prefix        = "false"
    # ipv6_prefix_length = 1
    # ipv6_subnet        = "2"
    # ipv6_total_host    = 2
    name    = "Test"
    site_id = "4e1fb72d-4896-49d7-aa98-57b64d93ce0a"
    /*slaac_support      = "false"
    type               = "string"*/
  }
}

output "dnacenter_reserve_ip_subpool_example" {
  value = dnacenter_reserve_ip_subpool.example
}
