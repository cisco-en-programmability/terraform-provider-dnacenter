
data "dnacenter_tag_member_count" "example" {
  provider                = dnacenter
  id                      = "string"
  level                   = "string"
  member_association_type = "string"
  member_type             = "string"
}

output "dnacenter_tag_member_count_example" {
  value = data.dnacenter_tag_member_count.example.item
}
