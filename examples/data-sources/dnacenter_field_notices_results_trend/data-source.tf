
data "dnacenter_field_notices_results_trend" "example" {
  provider  = dnacenter
  limit     = 1
  offset    = 1
  scan_time = 1.0
}

output "dnacenter_field_notices_results_trend_example" {
  value = data.dnacenter_field_notices_results_trend.example.items
}
