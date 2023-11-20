terraform {
  required_providers {
    dnacenter = {
      version = "1.1.26-beta"
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
    ipv4_global_pool = "10.0.0.0"

    //id                 = "string"
    ipv4_dhcp_servers  = ["abc.def.ghi.jkl"]
    ipv4_dns_servers   = ["abc.def.ghi.jkl"]
    ipv4_gate_way      = "abc.def.ghi.jkl"
    ipv4_prefix_length = 27
    ipv4_subnet        = "17"
    # ipv4_total_host    = 2
    # ipv6_address_space = "false"
    ipv6_dhcp_servers = ["abc.def.ghi.jkl2"]
    ipv6_dns_servers   = ["abc.def.ghi.jkl"]
    ipv6_gate_way      = "abc.def.ghi.jkl"
    ipv6_global_pool   = "10.0.0.1"
    # ipv6_prefix        = "false"
    ipv6_prefix_length = 10
    ipv6_subnet        = "2"
    # ipv6_total_host    = 2
    name    = "devhf_SVG_medtek_vlan855"
    site_id = "4e1fb72d-4896-49d7-aa98-57b64d93ce0a"
    /*slaac_support      = "false"
    type               = "string"*/
  }
}

output "dnacenter_reserve_ip_subpool_example" {
  value = dnacenter_reserve_ip_subpool.example
}
