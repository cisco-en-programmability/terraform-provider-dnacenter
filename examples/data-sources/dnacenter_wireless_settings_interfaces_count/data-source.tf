
data "dnacenter_wireless_settings_interfaces_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_interfaces_count_example" {
  value = data.dnacenter_wireless_settings_interfaces_count.example.item
}
