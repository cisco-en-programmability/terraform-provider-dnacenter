
data "dnacenter_network_device_replacements" "example" {
  provider                         = dnacenter
  family                           = "string"
  faulty_device_name               = "string"
  faulty_device_platform           = "string"
  faulty_device_serial_number      = "string"
  limit                            = 1
  offset                           = 1
  replacement_device_platform      = "string"
  replacement_device_serial_number = "string"
  replacement_status               = "string"
  sort_by                          = "string"
  sort_order                       = "string"
}

output "dnacenter_network_device_replacements_example" {
  value = data.dnacenter_network_device_replacements.example.items
}
