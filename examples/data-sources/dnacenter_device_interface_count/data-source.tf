
data "dnacenter_device_interface_count" "example" {
  provider = dnacenter
}

output "dnacenter_device_interface_count_example" {
  value = data.dnacenter_device_interface_count.example.item
}
