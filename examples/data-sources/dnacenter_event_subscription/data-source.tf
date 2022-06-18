
data "dnacenter_event_subscription" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_example" {
  value = data.dnacenter_event_subscription.example.items
}
