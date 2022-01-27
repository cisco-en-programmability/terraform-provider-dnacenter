
data "dnacenter_discovery_range" "example" {
  provider          = dnacenter
  records_to_return = 1
  start_index       = 1
}

output "dnacenter_discovery_range_example" {
  value = data.dnacenter_discovery_range.example.items
}
