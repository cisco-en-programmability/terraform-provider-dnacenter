
data "dnacenter_discovery_device" "example" {
  provider = dnacenter
  id       = "string"
  task_id  = "string"
}

output "dnacenter_discovery_device_example" {
  value = data.dnacenter_discovery_device.example.items
}
