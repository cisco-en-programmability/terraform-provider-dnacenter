
data "dnacenter_sda_anycast_gateways_count" "example" {
  provider             = dnacenter
  fabric_id            = "string"
  ip_pool_name         = "string"
  virtual_network_name = "string"
  vlan_id              = 1.0
  vlan_name            = "string"
}

output "dnacenter_sda_anycast_gateways_count_example" {
  value = data.dnacenter_sda_anycast_gateways_count.example.item
}
