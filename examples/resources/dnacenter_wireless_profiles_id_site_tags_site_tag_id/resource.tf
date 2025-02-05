
resource "dnacenter_wireless_profiles_id_site_tags_site_tag_id" "example" {
  provider = dnacenter

  parameters {

    ap_profile_name   = "string"
    flex_profile_name = "string"
    id                = "string"
    site_ids          = ["string"]
    site_tag_id       = "string"
    site_tag_name     = "string"
  }
}

output "dnacenter_wireless_profiles_id_site_tags_site_tag_id_example" {
  value = dnacenter_wireless_profiles_id_site_tags_site_tag_id.example
}
