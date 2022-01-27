
resource "dnacenter_network_device" "example" {
  provider = dnacenter
  parameters {

    id = "string"
  }
}

output "dnacenter_network_device_example" {
  value = dnacenter_network_device.example
}