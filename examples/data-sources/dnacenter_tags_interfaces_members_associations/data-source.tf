
data "dnacenter_tags_interfaces_members_associations" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_tags_interfaces_members_associations_example" {
  value = data.dnacenter_tags_interfaces_members_associations.example.items
}
