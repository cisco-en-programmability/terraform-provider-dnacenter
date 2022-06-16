
resource "dnacenter_reserve_ip_subpool" "example" {
  provider = dnacenter
  parameters {

    id                 = "string"
    ipv4_dhcp_servers  = ["string"]
    ipv4_dns_servers   = ["string"]
    ipv4_gate_way      = "string"
    ipv4_global_pool   = "string"
    ipv4_prefix        = "false"
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
    ipv6_total_host    = 1
    name               = "string"
    site_id            = "string"
    slaac_support      = "false"
    type               = "string"
  }
}

output "dnacenter_reserve_ip_subpool_example" {
  value = dnacenter_reserve_ip_subpool.example
}