
data "dnacenter_tag_member_type" "example" {
  provider = dnacenter
}

output "dnacenter_tag_member_type_example" {
  value = data.dnacenter_tag_member_type.example.items
}
