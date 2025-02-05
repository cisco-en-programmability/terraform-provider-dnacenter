
resource "dnacenter_dna_network_devices_query_count" "example" {
  provider = dnacenter
  parameters = [{

    end_time = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    start_time = 1
  }]
}

output "dnacenter_dna_network_devices_query_count_example" {
  value = dnacenter_dna_network_devices_query_count.example
}
