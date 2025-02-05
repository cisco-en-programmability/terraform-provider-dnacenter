
data "dnacenter_tags_interfaces_members_associations_count" "example" {
  provider = dnacenter
}

output "dnacenter_tags_interfaces_members_associations_count_example" {
  value = data.dnacenter_tags_interfaces_members_associations_count.example.item
}
