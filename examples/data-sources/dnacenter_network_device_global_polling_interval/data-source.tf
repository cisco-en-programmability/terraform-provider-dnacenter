
data "dnacenter_network_device_global_polling_interval" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_global_polling_interval_example" {
  value = data.dnacenter_network_device_global_polling_interval.example.item
}
