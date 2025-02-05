
data "dnacenter_network_bugs_results_network_devices_network_device_id_bugs_id" "example" {
  provider          = dnacenter
  id                = "string"
  network_device_id = "string"
}

output "dnacenter_network_bugs_results_network_devices_network_device_id_bugs_id_example" {
  value = data.dnacenter_network_bugs_results_network_devices_network_device_id_bugs_id.example.item
}
