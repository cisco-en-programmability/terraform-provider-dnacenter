
data "dnacenter_event_subscription_rest" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = "#"
  offset    = "#"
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_rest_example" {
  value = data.dnacenter_event_subscription_rest.example.items
}
