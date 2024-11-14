
data "dnacenter_lan_automation_sessions" "example" {
  provider = dnacenter
}

output "dnacenter_lan_automation_sessions_example" {
  value = data.dnacenter_lan_automation_sessions.example.item
}
