
resource "dnacenter_network_devices_summary_analytics" "example" {
  provider = dnacenter
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

      limit  = 1
      offset = 1
      sort_by {

        name  = "string"
        order = "string"
      }
    }
    start_time = 1
  }
}

output "dnacenter_network_devices_summary_analytics_example" {
  value = dnacenter_network_devices_summary_analytics.example
}