
data "dnacenter_security_advisories_results_trend" "example" {
  provider  = dnacenter
  limit     = 1
  offset    = 1
  scan_time = 1.0
}

output "dnacenter_security_advisories_results_trend_example" {
  value = data.dnacenter_security_advisories_results_trend.example.items
}
