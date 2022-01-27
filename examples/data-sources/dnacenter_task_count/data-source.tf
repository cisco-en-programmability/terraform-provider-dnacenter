
data "dnacenter_task_count" "example" {
  provider       = dnacenter
  data           = "string"
  end_time       = "string"
  error_code     = "string"
  failure_reason = "string"
  is_error       = "string"
  parent_id      = "string"
  progress       = "string"
  service_type   = "string"
  start_time     = "string"
  username       = "string"
}

output "dnacenter_task_count_example" {
  value = data.dnacenter_task_count.example.item
}
