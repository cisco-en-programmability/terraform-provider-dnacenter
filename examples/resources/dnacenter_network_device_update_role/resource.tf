
resource "dnacenter_network_device_update_role" "example" {
  provider = dnacenter
  parameters {

    id          = "string"
    role        = "string"
    role_source = "string"
  }
}

output "dnacenter_network_device_update_role_example" {
  value = dnacenter_network_device_update_role.example
}