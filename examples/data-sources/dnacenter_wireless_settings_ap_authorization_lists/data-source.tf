
data "dnacenter_wireless_settings_ap_authorization_lists" "example" {
  provider                   = dnacenter
  ap_authorization_list_name = "string"
  limit                      = "string"
  offset                     = "string"
}

output "dnacenter_wireless_settings_ap_authorization_lists_example" {
  value = data.dnacenter_wireless_settings_ap_authorization_lists.example.item
}
