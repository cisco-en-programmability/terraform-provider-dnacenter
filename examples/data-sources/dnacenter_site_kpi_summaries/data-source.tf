
data "dnacenter_site_kpi_summaries" "example" {
  provider          = dnacenter
  attribute         = "string"
  band              = "string"
  end_time          = 1609459200
  failure_category  = "string"
  failure_reason    = "string"
  limit             = 1
  offset            = 1
  order             = "string"
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_id           = "string"
  site_type         = "string"
  sort_by           = "string"
  ssid              = "string"
  start_time        = 1609459200
  task_id           = "string"
  view              = "string"
  xca_lle_rid       = "string"
}

output "dnacenter_site_kpi_summaries_example" {
  value = data.dnacenter_site_kpi_summaries.example.items
}
