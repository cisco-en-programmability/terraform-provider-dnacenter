
data "dnacenter_network_device_maintenance_schedules_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_maintenance_schedules_id_example" {
  value = data.dnacenter_network_device_maintenance_schedules_id.example.item
}
