
data "dnacenter_topology_physical" "example" {
  provider  = dnacenter
  node_type = "string"
}

output "dnacenter_topology_physical_example" {
  value = data.dnacenter_topology_physical.example.item
}
