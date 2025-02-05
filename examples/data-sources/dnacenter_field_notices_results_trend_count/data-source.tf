
data "dnacenter_field_notices_results_trend_count" "example" {
  provider  = dnacenter
  scan_time = 1.0
}

output "dnacenter_field_notices_results_trend_count_example" {
  value = data.dnacenter_field_notices_results_trend_count.example.item
}
