
data "dnacenter_network_device_maintenance_schedules" "example" {
  provider           = dnacenter
  limit              = "string"
  network_device_ids = "string"
  offset             = "string"
  order              = "string"
  sort_by            = "string"
  status             = "string"
}

output "dnacenter_network_device_maintenance_schedules_example" {
  value = data.dnacenter_network_device_maintenance_schedules.example.items
}
