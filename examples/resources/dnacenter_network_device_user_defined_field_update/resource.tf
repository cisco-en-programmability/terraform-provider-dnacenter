
resource "dnacenter_network_device_user_defined_field_update" "example" {
  provider  = dnacenter
  device_id = "string"
  parameters {

    name  = "string"
    value = "string"
  }
}

output "dnacenter_network_device_user_defined_field_update_example" {
  value = dnacenter_network_device_user_defined_field_update.example
}