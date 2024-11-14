
data "dnacenter_wireless_settings_interfaces" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_wireless_settings_interfaces_example" {
  value = data.dnacenter_wireless_settings_interfaces.example.items
}

data "dnacenter_wireless_settings_interfaces" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_interfaces_example" {
  value = data.dnacenter_wireless_settings_interfaces.example.item
}
