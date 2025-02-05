
data "dnacenter_wireless_settings_ap_authorization_lists_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_ap_authorization_lists_count_example" {
  value = data.dnacenter_wireless_settings_ap_authorization_lists_count.example.item
}
