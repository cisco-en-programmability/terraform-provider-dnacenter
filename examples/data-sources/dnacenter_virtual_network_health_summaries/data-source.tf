
data "dnacenter_virtual_network_health_summaries" "example" {
  provider    = dnacenter
  attribute   = "string"
  end_time    = 1609459200
  id          = "string"
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  start_time  = 1609459200
  view        = "string"
  vn_layer    = "string"
  xca_lle_rid = "string"
}

output "dnacenter_virtual_network_health_summaries_example" {
  value = data.dnacenter_virtual_network_health_summaries.example.items
}
