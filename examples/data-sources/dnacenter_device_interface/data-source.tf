
data "dnacenter_device_interface" "example" {
  provider         = dnacenter
  last_input_time  = "string"
  last_output_time = "string"
  limit            = 1
  offset           = 1
}

output "dnacenter_device_interface_example" {
  value = data.dnacenter_device_interface.example.items
}

data "dnacenter_device_interface" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_device_interface_example" {
  value = data.dnacenter_device_interface.example.item
}
