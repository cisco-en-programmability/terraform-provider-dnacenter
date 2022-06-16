
data "dnacenter_compliance_device_details" "example" {
  provider          = dnacenter
  compliance_status = "string"
  compliance_type   = "string"
  device_uuid       = "string"
  limit             = "string"
  offset            = "string"
}

output "dnacenter_compliance_device_details_example" {
  value = data.dnacenter_compliance_device_details.example.items
}
