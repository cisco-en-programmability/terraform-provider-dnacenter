
data "dnacenter_wireless_controllers_network_device_id_ap_authorization_lists" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_wireless_controllers_network_device_id_ap_authorization_lists_example" {
  value = data.dnacenter_wireless_controllers_network_device_id_ap_authorization_lists.example.item
}
