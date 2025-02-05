
data "dnacenter_field_notices_results_network_devices_network_device_id_notices" "example" {
  provider          = dnacenter
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  order             = "string"
  sort_by           = "string"
  type              = "string"
}

output "dnacenter_field_notices_results_network_devices_network_device_id_notices_example" {
  value = data.dnacenter_field_notices_results_network_devices_network_device_id_notices.example.items
}
