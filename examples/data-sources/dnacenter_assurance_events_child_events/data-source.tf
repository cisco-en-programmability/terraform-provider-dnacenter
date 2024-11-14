
data "dnacenter_assurance_events_child_events" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_assurance_events_child_events_example" {
  value = data.dnacenter_assurance_events_child_events.example.items
}
