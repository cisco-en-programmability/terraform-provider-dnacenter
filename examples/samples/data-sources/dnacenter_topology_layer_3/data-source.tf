
data "dnacenter_topology_layer_3" "example" {
  provider      = dnacenter
  topology_type = "string"
}

output "dnacenter_topology_layer_3_example" {
  value = data.dnacenter_topology_layer_3.example.item
}
