terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}

resource "dnacenter_reserve_ip_subpool" "example" {
  provider = dnacenter
  parameters {
    ipv4_prefix        = "false"
    ipv4_global_pool   = "175.175.0.0/16"
    /*
    id                 = "string"
    ipv4_dhcp_servers  = ["string"]
    ipv4_dns_servers   = ["string"]
    ipv4_gate_way      = "string"
    ipv4_global_pool   = "string"
    
    ipv4_prefix_length = 1
    ipv4_subnet        = "string"
    ipv4_total_host    = 1
    ipv6_address_space = "false"
    ipv6_dhcp_servers  = ["string"]
    ipv6_dns_servers   = ["string"]
    ipv6_gate_way      = "string"
    ipv6_global_pool   = "string"
    ipv6_prefix        = "false"
    ipv6_prefix_length = 1
    ipv6_subnet        = "string"
    ipv6_total_host    = 1*/
    name               = "IPv4v62002Block"
    site_id            = "9e860d9e-6499-40d1-9645-4b45bd684219"
    /*slaac_support      = "false"
    type               = "string"*/
  }
}

output "dnacenter_reserve_ip_subpool_example" {
  value = dnacenter_reserve_ip_subpool.example
}