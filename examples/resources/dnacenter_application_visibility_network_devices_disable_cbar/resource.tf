
resource "dnacenter_application_visibility_network_devices_disable_cbar" "example" {
  provider = dnacenter
  parameters = [{

    network_device_ids = ["string"]
  }]
}

output "dnacenter_application_visibility_network_devices_disable_cbar_example" {
  value = dnacenter_application_visibility_network_devices_disable_cbar.example
}
