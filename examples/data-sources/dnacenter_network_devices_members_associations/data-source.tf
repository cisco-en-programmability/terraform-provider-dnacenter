
data "dnacenter_network_devices_members_associations" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_network_devices_members_associations_example" {
  value = data.dnacenter_network_devices_members_associations.example.items
}
