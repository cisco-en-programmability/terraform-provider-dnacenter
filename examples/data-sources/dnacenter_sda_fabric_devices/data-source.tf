
data "dnacenter_sda_fabric_devices" "example" {
  provider          = dnacenter
  device_roles      = "string"
  fabric_id         = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "dnacenter_sda_fabric_devices_example" {
  value = data.dnacenter_sda_fabric_devices.example.items
}
