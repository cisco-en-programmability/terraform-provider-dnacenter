
resource "dnacenter_flexible_report_execute" "example" {
  provider  = dnacenter
  report_id = "string"
  parameters {

  }
}

output "dnacenter_flexible_report_execute_example" {
  value = dnacenter_flexible_report_execute.example
}