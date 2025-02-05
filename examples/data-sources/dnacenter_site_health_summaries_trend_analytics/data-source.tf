
data "dnacenter_site_health_summaries_trend_analytics" "example" {
  provider          = dnacenter
  attribute         = "string"
  end_time          = 1609459200
  id                = "string"
  limit             = 1
  offset            = 1
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_type         = "string"
  start_time        = 1609459200
  task_id           = "string"
  time_sort_order   = "string"
  trend_interval    = "string"
  xca_lle_rid       = "string"
}

output "dnacenter_site_health_summaries_trend_analytics_example" {
  value = data.dnacenter_site_health_summaries_trend_analytics.example.items
}
