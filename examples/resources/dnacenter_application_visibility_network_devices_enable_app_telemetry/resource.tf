
resource "dnacenter_application_visibility_network_devices_enable_app_telemetry" "example" {
  provider = dnacenter
  parameters = [{

    network_devices = [{

      id                  = "string"
      include_guest_ssids = "false"
      include_wlan_modes  = ["string"]
    }]
  }]
}

output "dnacenter_application_visibility_network_devices_enable_app_telemetry_example" {
  value = dnacenter_application_visibility_network_devices_enable_app_telemetry.example
}
