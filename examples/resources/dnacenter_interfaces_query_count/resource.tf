
resource "dnacenter_interfaces_query_count" "example" {
  provider = dnacenter
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
    page {

      limit  = 1
      offset = 1
      sort_by {

        name  = "string"
        order = "string"
      }
    }
    start_time = 1
    views      = ["string"]
  }
}

output "dnacenter_interfaces_query_count_example" {
  value = dnacenter_interfaces_query_count.example
}