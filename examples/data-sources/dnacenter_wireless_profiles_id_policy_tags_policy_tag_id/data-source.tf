
data "dnacenter_wireless_profiles_id_policy_tags_policy_tag_id" "example" {
  provider      = dnacenter
  id            = "string"
  policy_tag_id = "string"
}

output "dnacenter_wireless_profiles_id_policy_tags_policy_tag_id_example" {
  value = data.dnacenter_wireless_profiles_id_policy_tags_policy_tag_id.example.item
}
