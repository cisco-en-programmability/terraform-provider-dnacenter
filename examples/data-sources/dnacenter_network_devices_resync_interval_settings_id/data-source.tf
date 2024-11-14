
data "dnacenter_network_devices_resync_interval_settings_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_devices_resync_interval_settings_id_example" {
  value = data.dnacenter_network_devices_resync_interval_settings_id.example.item
}
