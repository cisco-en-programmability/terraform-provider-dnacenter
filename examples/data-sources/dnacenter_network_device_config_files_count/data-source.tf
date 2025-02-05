
data "dnacenter_network_device_config_files_count" "example" {
  provider          = dnacenter
  file_type         = "string"
  id                = "string"
  network_device_id = "string"
}

output "dnacenter_network_device_config_files_count_example" {
  value = data.dnacenter_network_device_config_files_count.example.item
}
