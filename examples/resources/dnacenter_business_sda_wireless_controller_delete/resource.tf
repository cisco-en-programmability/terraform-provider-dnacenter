
resource "dnacenter_business_sda_wireless_controller_delete" "example" {
  provider         = dnacenter
  device_ipaddress = "string"
}

output "dnacenter_business_sda_wireless_controller_delete_example" {
  value = dnacenter_business_sda_wireless_controller_delete.example
}