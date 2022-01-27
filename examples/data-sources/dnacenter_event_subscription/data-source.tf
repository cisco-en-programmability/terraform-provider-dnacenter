
data "dnacenter_event_subscription" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = "#"
  offset    = "#"
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_example" {
  value = data.dnacenter_event_subscription.example.items
}
