
data "dnacenter_wireless_controllers_managed_ap_locations_count" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_wireless_controllers_managed_ap_locations_count_example" {
  value = data.dnacenter_wireless_controllers_managed_ap_locations_count.example.item
}
