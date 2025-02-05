
resource "dnacenter_dhcp_services_id_trend_analytics" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
  parameters = [{

    aggregate_attributes = [{

      function = "string"
      name     = "string"
    }]
    attributes = ["string"]
    end_time   = 1
    filters = [{

      filters          = ["string"]
      key              = "string"
      logical_operator = "string"
      operator         = "string"

    }]
    group_by = ["string"]
    page = [{

      limit           = 1
      offset          = 1
      timestamp_order = "string"
    }]
    start_time     = 1
    trend_interval = "string"
  }]
}

output "dnacenter_dhcp_services_id_trend_analytics_example" {
  value = dnacenter_dhcp_services_id_trend_analytics.example
}
