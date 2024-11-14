
data "dnacenter_assurance_events" "example" {
  provider            = dnacenter
  ap_mac              = "string"
  attribute           = "string"
  client_mac          = "string"
  device_family       = "string"
  end_time            = 1609459200
  limit               = 1
  message_type        = "string"
  network_device_id   = "string"
  network_device_name = "string"
  offset              = 1
  order               = "string"
  severity            = 1.0
  site_hierarchy_id   = "string"
  site_id             = "string"
  sort_by             = "string"
  start_time          = 1609459200
  view                = "string"
  xca_lle_rid         = "string"
}

output "dnacenter_assurance_events_example" {
  value = data.dnacenter_assurance_events.example.items
}

data "dnacenter_assurance_events" "example" {
  provider    = dnacenter
  attribute   = "string"
  id          = "string"
  view        = "string"
  xca_lle_rid = "string"
}

output "dnacenter_assurance_events_example" {
  value = data.dnacenter_assurance_events.example.item
}
