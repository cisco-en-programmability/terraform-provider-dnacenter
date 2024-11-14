
resource "dnacenter_assurance_events_query" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
  parameters {

    attributes    = ["string"]
    device_family = ["string"]
    end_time      = 1
    filters {

      key      = "string"
      operator = "string"
      value    = "string"
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

output "dnacenter_assurance_events_query_example" {
  value = dnacenter_assurance_events_query.example
}