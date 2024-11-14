
resource "dnacenter_clients_query_count" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
  parameters {

    end_time = 1
    filters {

      key      = "string"
      operator = "string"
      value    = 1
    }
    start_time = 1
  }
}

output "dnacenter_clients_query_count_example" {
  value = dnacenter_clients_query_count.example
}