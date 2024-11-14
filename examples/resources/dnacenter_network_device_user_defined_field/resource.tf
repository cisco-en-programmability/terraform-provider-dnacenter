
resource "dnacenter_network_device_user_defined_field" "example" {
  provider = dnacenter

  parameters {

    description = "string"
    id          = "string"
    name        = "string"
  }
}

output "dnacenter_network_device_user_defined_field_example" {
  value = dnacenter_network_device_user_defined_field.example
}