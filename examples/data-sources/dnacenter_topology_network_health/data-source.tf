
data "dnacenter_topology_network_health" "example" {
  provider  = dnacenter
  timestamp = "string"
}

output "dnacenter_topology_network_health_example" {
  value = data.dnacenter_topology_network_health.example.items
}
