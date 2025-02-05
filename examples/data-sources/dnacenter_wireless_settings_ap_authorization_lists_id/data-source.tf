
data "dnacenter_wireless_settings_ap_authorization_lists_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_ap_authorization_lists_id_example" {
  value = data.dnacenter_wireless_settings_ap_authorization_lists_id.example.item
}
