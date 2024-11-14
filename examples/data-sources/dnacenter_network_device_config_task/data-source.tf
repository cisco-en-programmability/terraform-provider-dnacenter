
data "dnacenter_network_device_config_task" "example" {
  provider       = dnacenter
  parent_task_id = "string"
}

output "dnacenter_network_device_config_task_example" {
  value = data.dnacenter_network_device_config_task.example.items
}
