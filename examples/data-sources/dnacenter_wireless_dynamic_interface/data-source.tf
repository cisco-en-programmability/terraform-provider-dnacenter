
data "dnacenter_wireless_dynamic_interface" "example" {
  provider       = dnacenter
  interface_name = "string"
}

output "dnacenter_wireless_dynamic_interface_example" {
  value = data.dnacenter_wireless_dynamic_interface.example.items
}
