
data "dnacenter_transit_peer_network" "example" {
  provider                  = dnacenter
  transit_peer_network_name = "string"
}

output "dnacenter_transit_peer_network_example" {
  value = data.dnacenter_transit_peer_network.example.item
}
