
data "dnacenter_event_series_audit_logs_summary" "example" {
  provider           = dnacenter
  category           = "string"
  context            = "string"
  description        = "string"
  device_id          = "string"
  domain             = "string"
  end_time           = "hh:mm"
  event_hierarchy    = "string"
  event_id           = "string"
  instance_id        = "string"
  is_parent_only     = "false"
  is_system_events   = "false"
  name               = "string"
  parent_instance_id = "string"
  severity           = "string"
  site_id            = "string"
  source             = "string"
  start_time         = "hh:mm"
  sub_domain         = "string"
  user_id            = "string"
}

output "dnacenter_event_series_audit_logs_summary_example" {
  value = data.dnacenter_event_series_audit_logs_summary.example.items
}
