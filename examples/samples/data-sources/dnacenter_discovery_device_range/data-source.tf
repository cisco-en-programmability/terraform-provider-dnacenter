
data "dnacenter_discovery_device_range" "example" {
  provider          = dnacenter
  id                = "string"
  records_to_return = 1
  start_index       = 1
  task_id           = "string"
}

output "dnacenter_discovery_device_range_example" {
  value = data.dnacenter_discovery_device_range.example.items
}
