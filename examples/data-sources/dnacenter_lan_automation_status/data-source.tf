
data "dnacenter_lan_automation_status" "example" {
  provider = dnacenter
  limit    = "string"
  offset   = "string"
}

output "dnacenter_lan_automation_status_example" {
  value = data.dnacenter_lan_automation_status.example.items
}

data "dnacenter_lan_automation_status" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_lan_automation_status_example" {
  value = data.dnacenter_lan_automation_status.example.item
}
