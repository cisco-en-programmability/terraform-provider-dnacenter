
data "dnacenter_wireless_settings_anchor_groups" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_anchor_groups_example" {
  value = data.dnacenter_wireless_settings_anchor_groups.example.item
}
