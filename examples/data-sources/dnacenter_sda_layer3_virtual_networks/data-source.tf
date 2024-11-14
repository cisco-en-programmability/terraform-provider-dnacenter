
data "dnacenter_sda_layer3_virtual_networks" "example" {
  provider             = dnacenter
  anchored_site_id     = "string"
  fabric_id            = "string"
  limit                = 1
  offset               = 1
  virtual_network_name = "string"
}

output "dnacenter_sda_layer3_virtual_networks_example" {
  value = data.dnacenter_sda_layer3_virtual_networks.example.items
}
