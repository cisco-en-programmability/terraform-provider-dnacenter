
data "dnacenter_wireless_enterprise_ssid" "example" {
  provider  = dnacenter
  ssid_name = "string"
}

output "dnacenter_wireless_enterprise_ssid_example" {
  value = data.dnacenter_wireless_enterprise_ssid.example.items
}
