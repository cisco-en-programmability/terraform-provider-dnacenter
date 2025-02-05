
data "dnacenter_network_bugs_results_trend" "example" {
  provider  = dnacenter
  limit     = 1
  offset    = 1
  scan_time = 1.0
}

output "dnacenter_network_bugs_results_trend_example" {
  value = data.dnacenter_network_bugs_results_trend.example.items
}
