
data "dnacenter_flexible_report_schedules" "example" {
  provider = dnacenter
}

output "dnacenter_flexible_report_schedules_example" {
  value = data.dnacenter_flexible_report_schedules.example.items
}
