
data "dnacenter_discovery_device_count" "example" {
  provider = dnacenter
  id       = "string"
  task_id  = "string"
}

output "dnacenter_discovery_device_count_example" {
  value = data.dnacenter_discovery_device_count.example.item
}
