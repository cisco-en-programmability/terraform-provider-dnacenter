
resource "dnacenter_clients_summary_analytics" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
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
      value    = 1
    }
    group_by = ["string"]
    page {

      cursor = "string"
      limit  = 1
      sort_by {

        name  = "string"
        order = "string"
      }
    }
    start_time = 1
  }
}

output "dnacenter_clients_summary_analytics_example" {
  value = dnacenter_clients_summary_analytics.example
}