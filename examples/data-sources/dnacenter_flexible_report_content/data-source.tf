
data "dnacenter_flexible_report_content" "example" {
  provider     = dnacenter
  execution_id = "string"
  report_id    = "string"
}

output "dnacenter_flexible_report_content_example" {
  value = data.dnacenter_flexible_report_content.example.item
}
