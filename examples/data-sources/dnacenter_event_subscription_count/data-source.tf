
data "dnacenter_event_subscription_count" "example" {
  provider  = dnacenter
  event_ids = "string"
}

output "dnacenter_event_subscription_count_example" {
  value = data.dnacenter_event_subscription_count.example.item
}
