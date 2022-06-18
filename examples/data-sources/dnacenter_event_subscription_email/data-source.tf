
data "dnacenter_event_subscription_email" "example" {
  provider   = dnacenter
  category   = "string"
  domain     = "string"
  event_ids  = "string"
  limit      = 1
  name       = "string"
  offset     = 1
  order      = "string"
  sort_by    = "string"
  sub_domain = "string"
  type       = "string"
}

output "dnacenter_event_subscription_email_example" {
  value = data.dnacenter_event_subscription_email.example.items
}
