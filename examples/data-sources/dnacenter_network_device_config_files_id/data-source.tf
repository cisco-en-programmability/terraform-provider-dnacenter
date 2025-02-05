
data "dnacenter_network_device_config_files_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_config_files_id_example" {
  value = data.dnacenter_network_device_config_files_id.example.item
}
