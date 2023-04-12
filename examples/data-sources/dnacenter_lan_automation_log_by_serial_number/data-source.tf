
data "dnacenter_lan_automation_log_by_serial_number" "example" {
  provider      = dnacenter
  id            = "string"
  log_level     = "string"
  serial_number = "string"
}

output "dnacenter_lan_automation_log_by_serial_number_example" {
  value = data.dnacenter_lan_automation_log_by_serial_number.example.items
}
