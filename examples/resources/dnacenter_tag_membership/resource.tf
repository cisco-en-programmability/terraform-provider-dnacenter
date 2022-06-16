
resource "dnacenter_tag_membership" "example" {
  provider = dnacenter
  parameters {

    id        = "string"
    member_id = "string"
    object    = "string"
  }
}

output "dnacenter_tag_membership_example" {
  value = dnacenter_tag_membership.example
}