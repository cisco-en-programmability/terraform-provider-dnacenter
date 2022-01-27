
data "dnacenter_compliance_device" "example" {
  provider          = dnacenter
  compliance_status = "string"
  device_uuid       = "string"
  limit             = "#"
  offset            = "#"
}

output "dnacenter_compliance_device_example" {
  value = data.dnacenter_compliance_device.example.items
}
