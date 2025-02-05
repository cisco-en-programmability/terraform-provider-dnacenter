
data "dnacenter_network_bugs_results_network_devices" "example" {
  provider          = dnacenter
  bug_count         = 1.0
  limit             = 1
  network_device_id = "string"
  offset            = 1
  order             = "string"
  scan_mode         = "string"
  scan_status       = "string"
  sort_by           = "string"
}

output "dnacenter_network_bugs_results_network_devices_example" {
  value = data.dnacenter_network_bugs_results_network_devices.example.items
}
