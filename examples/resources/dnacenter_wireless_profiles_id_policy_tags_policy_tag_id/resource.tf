
resource "dnacenter_wireless_profiles_id_policy_tags_policy_tag_id" "example" {
  provider = dnacenter

  parameters {

    ap_zones        = ["string"]
    id              = "string"
    policy_tag_id   = "string"
    policy_tag_name = "string"
    site_ids        = ["string"]
  }
}

output "dnacenter_wireless_profiles_id_policy_tags_policy_tag_id_example" {
  value = dnacenter_wireless_profiles_id_policy_tags_policy_tag_id.example
}
