
data "dnacenter_network_device_with_snmp_v3_des" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  order    = "string"
  site_id  = "string"
  sort_by  = "string"
}

output "dnacenter_network_device_with_snmp_v3_des_example" {
  value = data.dnacenter_network_device_with_snmp_v3_des.example.items
}
