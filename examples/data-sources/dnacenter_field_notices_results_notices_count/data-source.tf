
data "dnacenter_field_notices_results_notices_count" "example" {
  provider     = dnacenter
  device_count = 1.0
  id           = "string"
  type         = "string"
}

output "dnacenter_field_notices_results_notices_count_example" {
  value = data.dnacenter_field_notices_results_notices_count.example.item
}
