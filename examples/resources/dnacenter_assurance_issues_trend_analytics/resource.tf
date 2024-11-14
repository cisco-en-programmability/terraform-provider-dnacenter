
resource "dnacenter_assurance_issues_trend_analytics" "example" {
  provider        = dnacenter
  accept_language = "string"
  xca_lle_rid     = "string"
  parameters {

    aggregate_attributes {

      function = "string"
      name     = "string"
    }
    attributes = ["string"]
    end_time   = 1
    filters {

      key      = "string"
      operator = "string"
      value    = "string"
    }
    group_by = ["string"]
    page {

      limit           = 1
      offset          = 1
      timestamp_order = "string"
    }
    start_time     = 1
    trend_interval = "string"
  }
}

output "dnacenter_assurance_issues_trend_analytics_example" {
  value = dnacenter_assurance_issues_trend_analytics.example
}