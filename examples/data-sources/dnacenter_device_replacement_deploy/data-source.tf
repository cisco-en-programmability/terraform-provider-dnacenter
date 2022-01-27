
data "dnacenter_device_replacement_deploy" "example" {
  provider                         = dnacenter
  faulty_device_serial_number      = "string"
  replacement_device_serial_number = "string"
}