
resource "dnacenter_network_devices_query" "example" {
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
    page {

      count   = 1
      limit   = 1
      offset  = 1
      sort_by = "string"
    }
    start_time = 1
    views      = ["string"]
  }
}

output "dnacenter_network_devices_query_example" {
  value = dnacenter_network_devices_query.example
}