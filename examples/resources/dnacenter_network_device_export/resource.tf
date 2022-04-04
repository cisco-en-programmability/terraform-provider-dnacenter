
resource "dnacenter_network_device_export" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_uuids   = ["string"]
    id             = "string"
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}