
data "dnacenter_field_notices_results_network_devices" "example" {
  provider          = dnacenter
  limit             = 1
  network_device_id = "string"
  notice_count      = 1.0
  offset            = 1
  order             = "string"
  scan_status       = "string"
  sort_by           = "string"
}

output "dnacenter_field_notices_results_network_devices_example" {
  value = data.dnacenter_field_notices_results_network_devices.example.items
}
