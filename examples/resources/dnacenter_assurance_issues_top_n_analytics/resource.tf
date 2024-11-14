
resource "dnacenter_assurance_issues_top_n_analytics" "example" {
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

      filters {

        key      = "string"
        operator = "string"
        value    = "string"
      }
      key              = "string"
      logical_operator = "string"
      operator         = "string"
      value            = "string"
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
    top_n      = 1
  }
}

output "dnacenter_assurance_issues_top_n_analytics_example" {
  value = dnacenter_assurance_issues_top_n_analytics.example
}