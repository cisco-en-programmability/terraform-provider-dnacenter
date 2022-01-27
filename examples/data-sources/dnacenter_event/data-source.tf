
data "dnacenter_event" "example" {
  provider = dnacenter
  event_id = "string"
  limit    = "#"
  offset   = "#"
  order    = "string"
  sort_by  = "string"
  tags     = "string"
}

output "dnacenter_event_example" {
  value = data.dnacenter_event.example.items
}
