
resource "dnacenter_interfaces_id_trend_analytics" "example" {
  provider = dnacenter
  id       = "string"
  parameters = [{

    aggregate_attributes = [{

      function = "string"
      name     = "string"
    }]
    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    start_time                = 1
    timestamp_order           = "string"
    trend_interval_in_minutes = 1
  }]
}

output "dnacenter_interfaces_id_trend_analytics_example" {
  value = dnacenter_interfaces_id_trend_analytics.example
}
