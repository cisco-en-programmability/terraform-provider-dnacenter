
data "dnacenter_compliance_device_by_id" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_compliance_device_by_id_example" {
  value = data.dnacenter_compliance_device_by_id.example.item
}
