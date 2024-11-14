
data "dnacenter_interfaces_members_associations_count" "example" {
  provider = dnacenter
}

output "dnacenter_interfaces_members_associations_count_example" {
  value = data.dnacenter_interfaces_members_associations_count.example.item
}
