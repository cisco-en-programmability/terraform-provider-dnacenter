
data "dnacenter_network_bugs_results_trend_count" "example" {
  provider  = dnacenter
  scan_time = 1.0
}

output "dnacenter_network_bugs_results_trend_count_example" {
  value = data.dnacenter_network_bugs_results_trend_count.example.item
}
