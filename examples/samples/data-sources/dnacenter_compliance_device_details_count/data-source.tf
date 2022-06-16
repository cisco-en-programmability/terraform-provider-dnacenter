
data "dnacenter_compliance_device_details_count" "example" {
  provider          = dnacenter
  compliance_status = "string"
  compliance_type   = "string"
}

output "dnacenter_compliance_device_details_count_example" {
  value = data.dnacenter_compliance_device_details_count.example.item
}
