
resource "dnacenter_snmp_properties" "example" {
  provider = dnacenter
  parameters {

    id                   = "string"
    instance_tenant_id   = "string"
    instance_uuid        = "string"
    int_value            = 1
    system_property_name = "string"
  }
}

output "dnacenter_snmp_properties_example" {
  value = dnacenter_snmp_properties.example
}