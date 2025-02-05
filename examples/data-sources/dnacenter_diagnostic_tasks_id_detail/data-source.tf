
data "dnacenter_diagnostic_tasks_id_detail" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_diagnostic_tasks_id_detail_example" {
  value = data.dnacenter_diagnostic_tasks_id_detail.example.item
}
