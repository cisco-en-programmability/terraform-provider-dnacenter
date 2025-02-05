
data "dnacenter_fabric_site_health_summaries_count" "example" {
  provider    = dnacenter
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  xca_lle_rid = "string"
}

output "dnacenter_fabric_site_health_summaries_count_example" {
  value = data.dnacenter_fabric_site_health_summaries_count.example.item
}
