
data "dnacenter_network_bugs_results_network_devices_network_device_id_bugs" "example" {
  provider          = dnacenter
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  order             = "string"
  severity          = "string"
  sort_by           = "string"
}

output "dnacenter_network_bugs_results_network_devices_network_device_id_bugs_example" {
  value = data.dnacenter_network_bugs_results_network_devices_network_device_id_bugs.example.items
}
