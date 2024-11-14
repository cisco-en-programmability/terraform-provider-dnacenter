
data "dnacenter_sda_fabric_devices_count" "example" {
  provider          = dnacenter
  device_roles      = "string"
  fabric_id         = "string"
  network_device_id = "string"
}

output "dnacenter_sda_fabric_devices_count_example" {
  value = data.dnacenter_sda_fabric_devices_count.example.item
}
