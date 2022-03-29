
resource "dnacenter_device_replacement_deploy" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    faulty_device_serial_number      = "string"
    replacement_device_serial_number = "string"
  }
}