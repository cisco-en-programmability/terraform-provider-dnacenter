
data "dnacenter_wireless_settings_anchor_groups_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_anchor_groups_id_example" {
  value = data.dnacenter_wireless_settings_anchor_groups_id.example.item
}
