
data "dnacenter_task_operation" "example" {
  provider     = dnacenter
  limit        = 1
  offset       = 1
  operation_id = "string"
}

output "dnacenter_task_operation_example" {
  value = data.dnacenter_task_operation.example.items
}
