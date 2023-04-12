
data "dnacenter_accesspoint_configuration_details_by_task_id" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_accesspoint_configuration_details_by_task_id_example" {
  value = data.dnacenter_accesspoint_configuration_details_by_task_id.example.items
}
