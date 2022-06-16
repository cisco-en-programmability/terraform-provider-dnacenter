
data "dnacenter_reports_view_group" "example" {
  provider = dnacenter
}

output "dnacenter_reports_view_group_example" {
  value = data.dnacenter_reports_view_group.example.items
}

data "dnacenter_reports_view_group" "example" {
  provider      = dnacenter
  view_group_id = "string"
}

output "dnacenter_reports_view_group_example" {
  value = data.dnacenter_reports_view_group.example.item
}
