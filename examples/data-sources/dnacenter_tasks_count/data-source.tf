
data "dnacenter_tasks_count" "example" {
  provider   = dnacenter
  end_time   = 1
  parent_id  = "string"
  root_id    = "string"
  start_time = 1
  status     = "string"
}

output "dnacenter_tasks_count_example" {
  value = data.dnacenter_tasks_count.example.item
}
