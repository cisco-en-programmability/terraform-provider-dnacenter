
data "dnacenter_event_subscription_details_email" "example" {
  provider       = dnacenter
  connector_type = "string"
  instance_id    = "string"
  name           = "string"
}

output "dnacenter_event_subscription_details_email_example" {
  value = data.dnacenter_event_subscription_details_email.example.items
}
