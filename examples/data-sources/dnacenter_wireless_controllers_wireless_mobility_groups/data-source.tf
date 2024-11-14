
data "dnacenter_wireless_controllers_wireless_mobility_groups" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_wireless_controllers_wireless_mobility_groups_example" {
  value = data.dnacenter_wireless_controllers_wireless_mobility_groups.example.items
}
