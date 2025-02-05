
data "dnacenter_virtual_network_health_summaries_id" "example" {
  provider    = dnacenter
  attribute   = "string"
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  view        = "string"
  xca_lle_rid = "string"
}

output "dnacenter_virtual_network_health_summaries_id_example" {
  value = data.dnacenter_virtual_network_health_summaries_id.example.item
}
