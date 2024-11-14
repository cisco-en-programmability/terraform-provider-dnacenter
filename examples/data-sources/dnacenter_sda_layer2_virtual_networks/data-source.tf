
data "dnacenter_sda_layer2_virtual_networks" "example" {
  provider                               = dnacenter
  associated_layer3_virtual_network_name = "string"
  fabric_id                              = "string"
  id                                     = "string"
  limit                                  = 1
  offset                                 = 1
  traffic_type                           = "string"
  vlan_id                                = 1.0
  vlan_name                              = "string"
}

output "dnacenter_sda_layer2_virtual_networks_example" {
  value = data.dnacenter_sda_layer2_virtual_networks.example.items
}
