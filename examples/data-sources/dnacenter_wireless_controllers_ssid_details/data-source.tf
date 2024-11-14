
data "dnacenter_wireless_controllers_ssid_details" "example" {
  provider          = dnacenter
  admin_status      = "false"
  limit             = 1
  managed           = "false"
  network_device_id = "string"
  offset            = 1
  ssid_name         = "string"
}

output "dnacenter_wireless_controllers_ssid_details_example" {
  value = data.dnacenter_wireless_controllers_ssid_details.example.items
}
