
data "dnacenter_flexible_report_executions" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_flexible_report_executions_example" {
  value = data.dnacenter_flexible_report_executions.example.item
}
