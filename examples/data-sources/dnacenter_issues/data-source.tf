
data "dnacenter_issues" "example" {
  provider     = dnacenter
  ai_driven    = "string"
  device_id    = "string"
  end_time     = "hh:mm"
  issue_status = "string"
  mac_address  = "string"
  priority     = "string"
  site_id      = "string"
  start_time   = "hh:mm"
}

output "dnacenter_issues_example" {
  value = data.dnacenter_issues.example.items
}
