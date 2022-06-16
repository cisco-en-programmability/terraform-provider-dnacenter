
data "dnacenter_network_device_config" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_config_example" {
  value = data.dnacenter_network_device_config.example.items
}

data "dnacenter_network_device_config" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_network_device_config_example" {
  value = data.dnacenter_network_device_config.example.item
}
