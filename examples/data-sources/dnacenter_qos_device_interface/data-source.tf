
data "dnacenter_qos_device_interface" "example" {
  provider          = dnacenter
  network_device_id = "string"
}

output "dnacenter_qos_device_interface_example" {
  value = data.dnacenter_qos_device_interface.example.items
}
