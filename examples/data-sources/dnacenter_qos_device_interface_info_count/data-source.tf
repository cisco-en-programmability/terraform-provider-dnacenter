
data "dnacenter_qos_device_interface_info_count" "example" {
  provider = dnacenter
}

output "dnacenter_qos_device_interface_info_count_example" {
  value = data.dnacenter_qos_device_interface_info_count.example.item
}
