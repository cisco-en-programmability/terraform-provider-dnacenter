provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_tag_membership" "example" {
  provider = dnacenter
  parameters {
    tag_id      = "string"
    member_type = "string"
    member_id   = "string"

  }
}

output "dnacenter_tag_membership_example" {
  value = dnacenter_tag_membership.example
}