
data "dnacenter_lan_automation_log" "example" {
  provider = dnacenter
  limit    = "string"
  offset   = "string"
}

output "dnacenter_lan_automation_log_example" {
  value = data.dnacenter_lan_automation_log.example.items
}

data "dnacenter_lan_automation_log" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_lan_automation_log_example" {
  value = data.dnacenter_lan_automation_log.example.item
}
