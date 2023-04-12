
data "dnacenter_network_device_user_defined_field" "example" {
  provider = dnacenter
  id       = "string"
  name     = "string"
}

output "dnacenter_network_device_user_defined_field_example" {
  value = data.dnacenter_network_device_user_defined_field.example.items
}
