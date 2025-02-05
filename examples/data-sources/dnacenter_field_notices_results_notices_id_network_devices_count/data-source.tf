
data "dnacenter_field_notices_results_notices_id_network_devices_count" "example" {
  provider          = dnacenter
  id                = "string"
  network_device_id = "string"
  scan_status       = "string"
}

output "dnacenter_field_notices_results_notices_id_network_devices_count_example" {
  value = data.dnacenter_field_notices_results_notices_id_network_devices_count.example.item
}
