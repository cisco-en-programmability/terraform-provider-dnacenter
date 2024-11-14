
resource "dnacenter_discovery_range_delete" "example" {
  provider          = dnacenter
  records_to_delete = 1
  start_index       = 1
  parameters {

  }
}

output "dnacenter_discovery_range_delete_example" {
  value = dnacenter_discovery_range_delete.example
}