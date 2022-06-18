
resource "dnacenter_discovery_range_delete" "example" {
  provider = dnacenter
  parameters {

    records_to_delete = []
    start_index       = []
  }
}

output "dnacenter_discovery_range_delete_example" {
  value = dnacenter_discovery_range_delete.example
}