
data "dnacenter_field_notices_results_notices_id_network_devices_network_device_id" "example" {
  provider          = dnacenter
  id                = "string"
  network_device_id = "string"
}

output "dnacenter_field_notices_results_notices_id_network_devices_network_device_id_example" {
  value = data.dnacenter_field_notices_results_notices_id_network_devices_network_device_id.example.item
}
