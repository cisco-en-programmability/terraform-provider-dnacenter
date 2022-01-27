
data "dnacenter_network_device_equipment" "example" {
  provider    = dnacenter
  device_uuid = "string"
  type        = "string"
}

output "dnacenter_network_device_equipment_example" {
  value = data.dnacenter_network_device_equipment.example.items
}
