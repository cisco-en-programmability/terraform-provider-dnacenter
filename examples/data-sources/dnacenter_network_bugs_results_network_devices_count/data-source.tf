
data "dnacenter_network_bugs_results_network_devices_count" "example" {
  provider          = dnacenter
  bug_count         = 1.0
  network_device_id = "string"
  scan_mode         = "string"
  scan_status       = "string"
}

output "dnacenter_network_bugs_results_network_devices_count_example" {
  value = data.dnacenter_network_bugs_results_network_devices_count.example.item
}
