
data "dnacenter_sda_site_member_member" "example" {
  provider    = dnacenter
  id          = "string"
  level       = "string"
  limit       = "string"
  member_type = "string"
  offset      = "string"
}

output "dnacenter_sda_site_member_member_example" {
  value = data.dnacenter_sda_site_member_member.example.items
}
