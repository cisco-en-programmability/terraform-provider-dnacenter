
data "dnacenter_field_notices_results_notices" "example" {
  provider     = dnacenter
  device_count = 1.0
  id           = "string"
  limit        = 1
  offset       = 1
  order        = "string"
  sort_by      = "string"
  type         = "string"
}

output "dnacenter_field_notices_results_notices_example" {
  value = data.dnacenter_field_notices_results_notices.example.items
}
