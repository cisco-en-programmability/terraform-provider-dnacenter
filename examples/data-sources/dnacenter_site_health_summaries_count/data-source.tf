
data "dnacenter_site_health_summaries_count" "example" {
  provider          = dnacenter
  end_time          = 1609459200
  id                = "string"
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_type         = "string"
  xca_lle_rid       = "string"
}

output "dnacenter_site_health_summaries_count_example" {
  value = data.dnacenter_site_health_summaries_count.example.item
}
