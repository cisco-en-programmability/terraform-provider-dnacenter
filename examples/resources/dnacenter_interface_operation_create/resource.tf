
resource "dnacenter_interface_operation_create" "example" {
  provider        = dnacenter
  deployment_mode = "string"
  interface_uuid  = "string"
  parameters {

    operation = "string"
    payload   = "string"
  }
}

output "dnacenter_interface_operation_create_example" {
  value = dnacenter_interface_operation_create.example
}