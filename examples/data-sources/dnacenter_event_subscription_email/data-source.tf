
data "dnacenter_event_subscription_email" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = "#"
  offset    = "#"
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_email_example" {
  value = data.dnacenter_event_subscription_email.example.items
}
