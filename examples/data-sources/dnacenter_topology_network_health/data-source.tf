
data "dnacenter_topology_network_health" "example" {
  provider  = dnacenter
  timestamp = 1.0
}

output "dnacenter_topology_network_health_example" {
  value = data.dnacenter_topology_network_health.example.items
}
