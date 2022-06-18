
data "dnacenter_tag_membership" "example" {
  provider                = dnacenter
  id                      = "string"
  level                   = "string"
  limit                   = "string"
  member_association_type = "string"
  member_type             = "string"
  offset                  = "string"
}

output "dnacenter_tag_membership_example" {
  value = data.dnacenter_tag_membership.example.items
}
