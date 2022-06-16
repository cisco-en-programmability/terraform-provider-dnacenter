
data "dnacenter_topology_vlan_details" "example" {
  provider = dnacenter
}

output "dnacenter_topology_vlan_details_example" {
  value = data.dnacenter_topology_vlan_details.example.items
}
