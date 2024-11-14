
resource "dnacenter_network_devices_interfaces_query" "example" {
  provider  = dnacenter
  device_id = "string"
  parameters {

    end_time = 1
    query {

      fields = ["string"]
      filters {

        key      = "string"
        operator = "string"
        value    = "string"
      }
      page {

        limit  = 1
        offset = 1
        order_by {

          name  = "string"
          order = "string"
        }
      }
    }
    start_time = 1
  }
}

output "dnacenter_network_devices_interfaces_query_example" {
  value = dnacenter_network_devices_interfaces_query.example
}