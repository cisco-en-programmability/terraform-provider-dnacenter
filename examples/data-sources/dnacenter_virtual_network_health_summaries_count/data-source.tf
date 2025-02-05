
data "dnacenter_virtual_network_health_summaries_count" "example" {
  provider    = dnacenter
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  vn_layer    = "string"
  xca_lle_rid = "string"
}

output "dnacenter_virtual_network_health_summaries_count_example" {
  value = data.dnacenter_virtual_network_health_summaries_count.example.item
}
