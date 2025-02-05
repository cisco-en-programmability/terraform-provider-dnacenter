
data "dnacenter_wireless_settings_anchor_groups_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_anchor_groups_count_example" {
  value = data.dnacenter_wireless_settings_anchor_groups_count.example.item
}
