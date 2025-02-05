
data "dnacenter_sda_pending_fabric_events" "example" {
  provider  = dnacenter
  fabric_id = "string"
  limit     = 1
  offset    = 1
}

output "dnacenter_sda_pending_fabric_events_example" {
  value = data.dnacenter_sda_pending_fabric_events.example.items
}
