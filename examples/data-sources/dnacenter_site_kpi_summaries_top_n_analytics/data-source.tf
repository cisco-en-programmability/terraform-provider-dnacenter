
data "dnacenter_site_kpi_summaries_top_n_analytics" "example" {
  provider    = dnacenter
  task_id     = "string"
  xca_lle_rid = "string"
}

output "dnacenter_site_kpi_summaries_top_n_analytics_example" {
  value = data.dnacenter_site_kpi_summaries_top_n_analytics.example.items
}
