
data "dnacenter_sda_multicast_virtual_networks" "example" {
  provider             = dnacenter
  fabric_id            = "string"
  limit                = 1
  offset               = 1
  virtual_network_name = "string"
}

output "dnacenter_sda_multicast_virtual_networks_example" {
  value = data.dnacenter_sda_multicast_virtual_networks.example.items
}
