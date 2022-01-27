
data "dnacenter_reports" "example" {
  provider      = dnacenter
  view_group_id = "string"
  view_id       = "string"
}

output "dnacenter_reports_example" {
  value = data.dnacenter_reports.example.items
}

data "dnacenter_reports" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_reports_example" {
  value = data.dnacenter_reports.example.item
}
