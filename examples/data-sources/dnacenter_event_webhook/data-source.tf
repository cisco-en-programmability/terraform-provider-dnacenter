
data "dnacenter_event_webhook" "example" {
  provider    = dnacenter
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  webhook_ids = "string"
}

output "dnacenter_event_webhook_example" {
  value = data.dnacenter_event_webhook.example.item
}
