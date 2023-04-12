
data "dnacenter_event_snmp_config" "example" {
  provider  = dnacenter
  config_id = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_snmp_config_example" {
  value = data.dnacenter_event_snmp_config.example.item
}
