
data "dnacenter_compliance_device_status_count" "example" {
  provider          = dnacenter
  compliance_status = "string"
}

output "dnacenter_compliance_device_status_count_example" {
  value = data.dnacenter_compliance_device_status_count.example.item
}
