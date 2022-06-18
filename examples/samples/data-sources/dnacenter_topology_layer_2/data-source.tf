
data "dnacenter_topology_layer_2" "example" {
  provider = dnacenter
  vlan_id  = "string"
}

output "dnacenter_topology_layer_2_example" {
  value = data.dnacenter_topology_layer_2.example.item
}
