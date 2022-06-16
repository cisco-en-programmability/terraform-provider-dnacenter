
data "dnacenter_event_subscription_details_email" "example" {
  provider    = dnacenter
  instance_id = "string"
  limit       = 1
  name        = "string"
  offset      = 1
  order       = "string"
  sort_by     = "string"
}

output "dnacenter_event_subscription_details_email_example" {
  value = data.dnacenter_event_subscription_details_email.example.items
}
