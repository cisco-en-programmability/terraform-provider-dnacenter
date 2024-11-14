
data "dnacenter_tasks" "example" {
  provider   = dnacenter
  end_time   = 1
  limit      = 1
  offset     = 1
  order      = "string"
  parent_id  = "string"
  root_id    = "string"
  sort_by    = "string"
  start_time = 1
  status     = "string"
}

output "dnacenter_tasks_example" {
  value = data.dnacenter_tasks.example.items
}

data "dnacenter_tasks" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_tasks_example" {
  value = data.dnacenter_tasks.example.item
}
