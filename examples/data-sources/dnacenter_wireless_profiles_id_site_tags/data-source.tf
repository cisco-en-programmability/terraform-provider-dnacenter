
data "dnacenter_wireless_profiles_id_site_tags" "example" {
  provider      = dnacenter
  id            = "string"
  limit         = 1
  offset        = 1
  site_tag_name = "string"
}

output "dnacenter_wireless_profiles_id_site_tags_example" {
  value = data.dnacenter_wireless_profiles_id_site_tags.example.items
}
