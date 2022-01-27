
data "dnacenter_network_device_config_count" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_config_count_example" {
  value = data.dnacenter_network_device_config_count.example.item
}
