
data "dnacenter_lan_automation_count" "example" {
  provider = dnacenter
}

output "dnacenter_lan_automation_count_example" {
  value = data.dnacenter_lan_automation_count.example.item
}
