
data "dnacenter_wireless_profiles_id_policy_tags_count" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_profiles_id_policy_tags_count_example" {
  value = data.dnacenter_wireless_profiles_id_policy_tags_count.example.item
}
