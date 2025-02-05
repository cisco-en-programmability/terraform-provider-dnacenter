
data "dnacenter_field_notices_results_network_devices_network_device_id_notices_count" "example" {
  provider          = dnacenter
  id                = "string"
  network_device_id = "string"
  type              = "string"
}

output "dnacenter_field_notices_results_network_devices_network_device_id_notices_count_example" {
  value = data.dnacenter_field_notices_results_network_devices_network_device_id_notices_count.example.item
}
