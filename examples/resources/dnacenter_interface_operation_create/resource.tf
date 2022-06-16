
resource "dnacenter_interface_operation_create" "example" {
  provider = dnacenter
  parameters {

    interface_uuid = "string"
    operation      = "string"
    payload        = "string"
  }
}

output "dnacenter_interface_operation_create_example" {
  value = dnacenter_interface_operation_create.example
}