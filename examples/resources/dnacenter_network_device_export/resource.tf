
resource "dnacenter_network_device_export" "example" {
  provider = dnacenter
  parameters {

    device_uuids   = ["string"]
    id             = "string"
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}

output "dnacenter_network_device_export_example" {
  value = dnacenter_network_device_export.example
}