
data "dnacenter_flexible_report_schedule" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_flexible_report_schedule_example" {
  value = data.dnacenter_flexible_report_schedule.example.item
}
