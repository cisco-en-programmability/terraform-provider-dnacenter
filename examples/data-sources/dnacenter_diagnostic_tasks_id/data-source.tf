
data "dnacenter_diagnostic_tasks_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_diagnostic_tasks_id_example" {
  value = data.dnacenter_diagnostic_tasks_id.example.item
}
