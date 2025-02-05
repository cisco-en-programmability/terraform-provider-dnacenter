
data "dnacenter_field_notices_results_network_devices_count" "example" {
  provider          = dnacenter
  network_device_id = "string"
  notice_count      = 1.0
  scan_status       = "string"
}

output "dnacenter_field_notices_results_network_devices_count_example" {
  value = data.dnacenter_field_notices_results_network_devices_count.example.item
}
