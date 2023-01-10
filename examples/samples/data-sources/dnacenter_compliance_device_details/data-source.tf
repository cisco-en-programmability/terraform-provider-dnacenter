
data "dnacenter_compliance_device_details" "example" {
  provider          = dnacenter
  compliance_status = "string"
  compliance_type   = "string"
  device_uuid       = "string"
  limit             = 1
  offset            = 1
}

output "dnacenter_compliance_device_details_example" {
  value = data.dnacenter_compliance_device_details.example.items
}
