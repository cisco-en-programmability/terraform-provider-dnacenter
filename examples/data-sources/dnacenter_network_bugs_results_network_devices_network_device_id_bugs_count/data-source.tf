
data "dnacenter_network_bugs_results_network_devices_network_device_id_bugs_count" "example" {
  provider          = dnacenter
  id                = "string"
  network_device_id = "string"
  severity          = "string"
}

output "dnacenter_network_bugs_results_network_devices_network_device_id_bugs_count_example" {
  value = data.dnacenter_network_bugs_results_network_devices_network_device_id_bugs_count.example.item
}
