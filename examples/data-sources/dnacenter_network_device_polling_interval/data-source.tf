
data "dnacenter_network_device_polling_interval" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_polling_interval_example" {
  value = data.dnacenter_network_device_polling_interval.example.item
}
