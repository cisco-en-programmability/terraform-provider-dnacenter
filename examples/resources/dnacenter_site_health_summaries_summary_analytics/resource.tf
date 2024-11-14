
resource "dnacenter_site_health_summaries_summary_analytics" "example" {
  provider = dnacenter
  parameters {

    attributes = ["string"]
    end_time   = 1
    start_time = 1
    views      = ["string"]
  }
}

output "dnacenter_site_health_summaries_summary_analytics_example" {
  value = dnacenter_site_health_summaries_summary_analytics.example
}