
resource "dnacenter_flexible_report_schedule" "example" {
  provider = dnacenter
  parameters {

    report_id = "string"
    schedule  = "string"
  }
}

output "dnacenter_flexible_report_schedule_example" {
  value = dnacenter_flexible_report_schedule.example
}