
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_interface_operation_create" "example" {
  provider = dnacenter
  parameters {

    interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
    operation      = "string"
    payload        = "string"
  }
}

output "dnacenter_interface_operation_create_example" {
  value = dnacenter_interface_operation_create.example
}