
resource "dnacenter_network_devices_resync_interval_settings_id" "example" {
  provider = dnacenter

  parameters {

    id       = "string"
    interval = 1
  }
}

output "dnacenter_network_devices_resync_interval_settings_id_example" {
  value = dnacenter_network_devices_resync_interval_settings_id.example
}