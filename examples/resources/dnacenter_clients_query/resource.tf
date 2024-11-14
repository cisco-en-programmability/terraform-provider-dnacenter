
resource "dnacenter_clients_query" "example" {
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

output "dnacenter_clients_query_example" {
  value = dnacenter_clients_query.example
}