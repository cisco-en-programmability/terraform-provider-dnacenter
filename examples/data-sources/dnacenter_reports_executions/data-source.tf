
data "dnacenter_reports_executions" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_reports_executions_example" {
  value = data.dnacenter_reports_executions.example.item
}
