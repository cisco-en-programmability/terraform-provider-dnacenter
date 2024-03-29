
resource "dnacenter_tag_member" "example" {
  provider = dnacenter
  parameters {

    id        = "string"
    member_id = "string"
    object    = "string"
  }
}

output "dnacenter_tag_member_example" {
  value = dnacenter_tag_member.example
}