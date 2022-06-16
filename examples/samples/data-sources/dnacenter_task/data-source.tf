
data "dnacenter_task" "example" {
  provider       = dnacenter
  data           = "string"
  end_time       = "string"
  error_code     = "string"
  failure_reason = "string"
  is_error       = "string"
  limit          = "string"
  offset         = "string"
  order          = "string"
  parent_id      = "string"
  progress       = "string"
  service_type   = "string"
  sort_by        = "string"
  start_time     = "string"
  username       = "string"
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.items
}

data "dnacenter_task" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.item
}
