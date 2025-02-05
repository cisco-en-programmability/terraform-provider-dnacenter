
data "dnacenter_network_device_config_files" "example" {
  provider          = dnacenter
  file_type         = "string"
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "dnacenter_network_device_config_files_example" {
  value = data.dnacenter_network_device_config_files.example.items
}
