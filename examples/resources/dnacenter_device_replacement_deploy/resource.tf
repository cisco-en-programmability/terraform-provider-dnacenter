
resource "dnacenter_device_replacement_deploy" "example" {
  provider = dnacenter
  parameters {

    faulty_device_serial_number      = "string"
    replacement_device_serial_number = "string"
  }
}

output "dnacenter_device_replacement_deploy_example" {
  value = dnacenter_device_replacement_deploy.example
}