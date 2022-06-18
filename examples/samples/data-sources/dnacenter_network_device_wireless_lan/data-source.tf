
data "dnacenter_network_device_wireless_lan" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_wireless_lan_example" {
  value = data.dnacenter_network_device_wireless_lan.example.item
}
