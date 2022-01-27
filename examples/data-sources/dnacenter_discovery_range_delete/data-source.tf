
data "dnacenter_discovery_range_delete" "example" {
  provider          = dnacenter
  records_to_delete = 1
  start_index       = 1
}