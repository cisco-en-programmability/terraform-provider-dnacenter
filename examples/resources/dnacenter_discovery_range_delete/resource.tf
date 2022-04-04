
resource "dnacenter_discovery_range_delete" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    records_to_delete = 1
    start_index       = 1
  }
}