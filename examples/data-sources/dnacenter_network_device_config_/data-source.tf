
data "dnacenter_network_device_config_" "example" {
  provider     = dnacenter
  created_by   = "string"
  created_time = "string"
  device_id    = "string"
  file_type    = "string"
  limit        = 1
  offset       = 1
}

output "dnacenter_network_device_config__example" {
  value = data.dnacenter_network_device_config_.example.items
}
