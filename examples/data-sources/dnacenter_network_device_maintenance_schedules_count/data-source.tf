
data "dnacenter_network_device_maintenance_schedules_count" "example" {
  provider           = dnacenter
  network_device_ids = "string"
  status             = "string"
}

output "dnacenter_network_device_maintenance_schedules_count_example" {
  value = data.dnacenter_network_device_maintenance_schedules_count.example.item
}
