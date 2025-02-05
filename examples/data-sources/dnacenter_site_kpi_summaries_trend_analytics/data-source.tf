
data "dnacenter_site_kpi_summaries_trend_analytics" "example" {
  provider    = dnacenter
  task_id     = "string"
  xca_lle_rid = "string"
}

output "dnacenter_site_kpi_summaries_trend_analytics_example" {
  value = data.dnacenter_site_kpi_summaries_trend_analytics.example.items
}
