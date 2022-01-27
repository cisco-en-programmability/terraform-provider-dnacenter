
data "dnacenter_compliance_device_by_id_detail" "example" {
  provider        = dnacenter
  category        = "string"
  compliance_type = "string"
  device_uuid     = "string"
  diff_list       = "false"
  key             = "string"
  value           = "string"
}

output "dnacenter_compliance_device_by_id_detail_example" {
  value = data.dnacenter_compliance_device_by_id_detail.example.items
}
