
data "dnacenter_wireless_controllers_secondary_managed_ap_locations" "example" {
  provider          = dnacenter
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "dnacenter_wireless_controllers_secondary_managed_ap_locations_example" {
  value = data.dnacenter_wireless_controllers_secondary_managed_ap_locations.example.items
}
