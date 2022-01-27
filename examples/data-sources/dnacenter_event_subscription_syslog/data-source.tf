
data "dnacenter_event_subscription_syslog" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = "#"
  offset    = "#"
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_syslog_example" {
  value = data.dnacenter_event_subscription_syslog.example.items
}
