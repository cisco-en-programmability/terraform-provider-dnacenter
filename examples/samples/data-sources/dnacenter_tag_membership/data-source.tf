
data "dnacenter_tag_membership" "example" {
  provider                = dnacenter
  id                      = "string"
  level                   = "string"
  limit                   = 1
  member_association_type = "string"
  member_type             = "string"
  offset                  = 1
}

output "dnacenter_tag_membership_example" {
  value = data.dnacenter_tag_membership.example.items
}
