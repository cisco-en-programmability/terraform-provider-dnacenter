
data "dnacenter_snmp_properties" "example" {
  provider = dnacenter
}

output "dnacenter_snmp_properties_example" {
  value = data.dnacenter_snmp_properties.example.items
}
