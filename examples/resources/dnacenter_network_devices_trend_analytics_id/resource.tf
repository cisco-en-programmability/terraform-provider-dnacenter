
resource "dnacenter_network_devices_trend_analytics_id" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

    aggregate_attributes {

      function = "string"
      name     = "string"
    }
    attributes = ["string"]
    end_time   = 1
    filters {

      filters          = ["string"]
      key              = "string"
      logical_operator = "string"
      operator         = "string"
      value            = "string"
    }
    group_by = ["string"]
    page {

      limit           = 1
      offset          = 1
      timestamp_order = "string"
    }
    start_time                = 1
    trend_interval_in_minutes = 1
  }
}

output "dnacenter_network_devices_trend_analytics_id_example" {
  value = dnacenter_network_devices_trend_analytics_id.example
}