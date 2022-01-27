
data "dnacenter_event_subscription_details_syslog" "example" {
  provider       = dnacenter
  connector_type = "string"
  instance_id    = "string"
  name           = "string"
}

output "dnacenter_event_subscription_details_syslog_example" {
  value = data.dnacenter_event_subscription_details_syslog.example.items
}
