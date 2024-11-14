
resource "dnacenter_network_device_user_defined_field_delete" "example" {
  provider  = dnacenter
  device_id = "string"
  name      = "string"
  parameters {

  }
}

output "dnacenter_network_device_user_defined_field_delete_example" {
  value = dnacenter_network_device_user_defined_field_delete.example
}