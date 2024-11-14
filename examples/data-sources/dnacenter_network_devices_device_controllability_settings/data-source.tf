
data "dnacenter_network_devices_device_controllability_settings" "example" {
  provider = dnacenter
}

output "dnacenter_network_devices_device_controllability_settings_example" {
  value = data.dnacenter_network_devices_device_controllability_settings.example.item
}
