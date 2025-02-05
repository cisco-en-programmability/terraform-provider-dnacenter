
data "dnacenter_endpoint_analytics_tasks" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_endpoint_analytics_tasks_example" {
  value = data.dnacenter_endpoint_analytics_tasks.example.item
}
