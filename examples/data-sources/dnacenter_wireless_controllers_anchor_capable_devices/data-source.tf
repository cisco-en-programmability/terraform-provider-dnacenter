
data "dnacenter_wireless_controllers_anchor_capable_devices" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_controllers_anchor_capable_devices_example" {
  value = data.dnacenter_wireless_controllers_anchor_capable_devices.example.item
}
