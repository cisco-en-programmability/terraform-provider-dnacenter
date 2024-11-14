
data "dnacenter_wireless_controllers_ssid_details_count" "example" {
  provider          = dnacenter
  admin_status      = "false"
  managed           = "false"
  network_device_id = "string"
}

output "dnacenter_wireless_controllers_ssid_details_count_example" {
  value = data.dnacenter_wireless_controllers_ssid_details_count.example.item
}
