
data "dnacenter_network_devices_members_associations_count" "example" {
  provider = dnacenter
}

output "dnacenter_network_devices_members_associations_count_example" {
  value = data.dnacenter_network_devices_members_associations_count.example.item
}
