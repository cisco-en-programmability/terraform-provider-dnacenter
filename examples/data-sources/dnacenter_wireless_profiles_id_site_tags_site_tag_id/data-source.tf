
data "dnacenter_wireless_profiles_id_site_tags_site_tag_id" "example" {
  provider    = dnacenter
  id          = "string"
  site_tag_id = "string"
}

output "dnacenter_wireless_profiles_id_site_tags_site_tag_id_example" {
  value = data.dnacenter_wireless_profiles_id_site_tags_site_tag_id.example.item
}
