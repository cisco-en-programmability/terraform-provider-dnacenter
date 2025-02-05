
resource "dnacenter_application_visibility_network_devices_enable_cbar" "example" {
  provider = dnacenter
  parameters = [{

    network_devices = [{

      exclude_interface_ids = ["string"]
      exclude_wlan_modes    = ["string"]
      id                    = "string"
    }]
  }]
}

output "dnacenter_application_visibility_network_devices_enable_cbar_example" {
  value = dnacenter_application_visibility_network_devices_enable_cbar.example
}
