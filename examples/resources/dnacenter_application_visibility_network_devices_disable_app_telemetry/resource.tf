
resource "dnacenter_application_visibility_network_devices_disable_app_telemetry" "example" {
  provider = dnacenter
  parameters = [{

    network_device_ids = ["string"]
  }]
}

output "dnacenter_application_visibility_network_devices_disable_app_telemetry_example" {
  value = dnacenter_application_visibility_network_devices_disable_app_telemetry.example
}
