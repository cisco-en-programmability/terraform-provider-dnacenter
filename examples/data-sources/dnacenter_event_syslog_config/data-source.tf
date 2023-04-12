
data "dnacenter_event_syslog_config" "example" {
  provider  = dnacenter
  config_id = "string"
  limit     = 1
  name      = "string"
  offset    = 1
  order     = "string"
  protocol  = "string"
  sort_by   = "string"
}

output "dnacenter_event_syslog_config_example" {
  value = data.dnacenter_event_syslog_config.example.item
}
