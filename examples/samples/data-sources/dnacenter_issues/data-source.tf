
data "dnacenter_issues" "example" {
  provider     = dnacenter
  ai_driven    = "string"
  device_id    = "string"
  end_time     = 1609459200
  issue_status = "string"
  mac_address  = "string"
  priority     = "string"
  site_id      = "string"
  start_time   = 1609459200
}

output "dnacenter_issues_example" {
  value = data.dnacenter_issues.example.items
}
