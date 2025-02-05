
data "dnacenter_site_kpi_summaries_count" "example" {
  provider          = dnacenter
  end_time          = 1609459200
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_id           = "string"
  site_type         = "string"
  start_time        = 1609459200
  xca_lle_rid       = "string"
}

output "dnacenter_site_kpi_summaries_count_example" {
  value = data.dnacenter_site_kpi_summaries_count.example.item
}
