
resource "dnacenter_site_kpi_summaries_summary_analytics" "example" {
  provider = dnacenter

  parameters {

    attributes = ["string"]
    end_time   = 1
    filters {

      key      = "string"
      operator = "string"
      value    = "string"
    }
    start_time = 1
  }
}

output "dnacenter_site_kpi_summaries_summary_analytics_example" {
  value = dnacenter_site_kpi_summaries_summary_analytics.example
}
