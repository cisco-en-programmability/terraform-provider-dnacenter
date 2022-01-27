
data "dnacenter_tag_membership" "example" {
  provider = dnacenter
  member_to_tags {

    key = ["string"]
  }
  member_type = "string"
}