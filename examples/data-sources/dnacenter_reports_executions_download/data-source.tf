
data "dnacenter_reports_executions_download" "example" {
  provider     = dnacenter
  dirpath      = "string"
  execution_id = "string"
  report_id    = "string"
}

output "dnacenter_reports_executions_download_example" {
  value = data.dnacenter_reports_executions_download.example.item
}
