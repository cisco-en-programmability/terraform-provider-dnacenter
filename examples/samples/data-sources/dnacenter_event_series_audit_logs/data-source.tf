
data "dnacenter_event_series_audit_logs" "example" {
  provider           = dnacenter
  category           = "string"
  context            = "string"
  description        = "string"
  device_id          = "string"
  domain             = "string"
  end_time           = 1609459200
  event_hierarchy    = "string"
  event_id           = "string"
  instance_id        = "string"
  is_system_events   = "false"
  limit              = 1
  name               = "string"
  offset             = 1
  order              = "string"
  parent_instance_id = "string"
  severity           = "string"
  site_id            = "string"
  sort_by            = "string"
  source             = "string"
  start_time         = 1609459200
  sub_domain         = "string"
  user_id            = "string"
}

output "dnacenter_event_series_audit_logs_example" {
  value = data.dnacenter_event_series_audit_logs.example.items
}
